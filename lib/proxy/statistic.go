package proxy

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/go-redis/redis"
	"github.com/leicc520/go-crawler-srv/lib/proxy/channal"
	"github.com/leicc520/go-orm/log"
)

/************************************************************
	代理请求数据业务统计处理逻辑
 */

const (
	PROXY_ERROR_LOCK_TIME = time.Millisecond*100
	PROXY_SYNC_REDIS_TIME = time.Second*30
	PROXY_SYNC_NOTIFY_TIME= time.Hour*24
	PROXY_ERROR_LIMIT     = 30 //连续错误30次切换ip
)

type ProxySt struct {
	Proxy  string  `yaml:"proxy"  json:"proxy"`  //代理名称
	Url    string  `yaml:"url"    json:"url"`       //代理请求地址
	Status int8    `yaml:"status" json:"status"` //状态0-禁用 1-正常 2-锁定
	IsTcp  bool    `yaml:"-" json:"-"`			 //是否tcp代理
	Error  uint64  `yaml:"-"      json:"-"`      //请求失败的统计
	IFGet  channal.IFProxy   `yaml:"-"  json:"-"` //代理请求获取地址
	Expire int64   `yaml:"-" json:"-"` //锁定时间
}

//自动切换代理处理逻辑
func (s *ProxySt) CutProxy()  {
	ipAddress := s.IFGet.GetProxy(channal.PROXY_SOCK5)
	if len(ipAddress) > 0 {//请求失败很多更好地址
		s.Url = "http://"+ipAddress
	}
	s.IsTcp = true
	log.Write(-1, "自动切换IP", ipAddress)
}

//记录日志的状态 统计记录
type LogStateSt struct {
	ProxyIdx 	int
	Host 		string
	Status 		int
}

//汇总数据处理逻辑
type Statistic struct {
	proxy []ProxySt
	len    int
	Request uint64 //请求的统计数值
	logChan chan string
	rds *redis.Client
}

//获取数据资料信息
var (
	statPtr *Statistic = nil
	onceInit = sync.Once{}
)

//初始化数据资料信息
func Init(proxy []ProxySt, rds *redis.Client) {
	onceInit.Do(func() {
		statPtr = NewStatistic(proxy, rds)
	})
}

//返回统计数据资料信息
func GetStatistic() *Statistic {
	return statPtr
}

//获取代理地址
func (s *Statistic) GetProxy(isTcp, isCut bool) string {
	nlen := len(s.proxy)
	for i := 0; i < nlen; i++ {
		item := &s.proxy[i]
		if (len(item.Url) < 1 || isCut) && item.IFGet != nil {
			item.CutProxy() //切换代理
		}
		if isTcp == item.IsTcp {
			return item.Url
		}
	}
	return ""
}

//格式化成字符串
func (s LogStateSt) String() string {
	return fmt.Sprintf("%d;%s;%d", s.ProxyIdx, s.Host, s.Status)
}

//格式化状态数据资料信息
func LogStateBuilder(logStr string) *LogStateSt {
	arrStr := strings.Split(logStr, ";")
	if len(arrStr) != 3 {
		log.Write(log.ERROR, logStr, "代理监控数据异常...")
		return nil
	}
	proxyIdx, _ := strconv.ParseInt(arrStr[0], 10, 64)
	status, _   := strconv.ParseInt(arrStr[2], 10, 64)
	return &LogStateSt{ProxyIdx:int(proxyIdx), Host: arrStr[1], Status: int(status)}
}

//初始化对象数据资料信息
func NewStatistic(proxy []ProxySt, rds *redis.Client) *Statistic {
	logChan := make(chan string, 1024*10)
	ss := &Statistic{proxy: proxy, rds: rds, len: len(proxy), logChan: logChan}
	go ss.goAsyncNotify() //开启异步执行队列 持久化数据到redis当中
	return ss
}

//当天剩余的时间处理逻辑
func (s *Statistic) dayDuration() time.Duration {
	n := time.Now()
	l := time.Date(n.Year(), n.Month(), n.Day(), 23, 59, 59, 0, n.Location())
	t := l.Sub(n)
	return t
}

//异步任务通知，格式化存储到数据库
func (s *Statistic) goAsyncNotify() {
	state := make(map[string]int)
	syncChan   := time.After(PROXY_SYNC_REDIS_TIME)
	notifyChan := time.After(s.dayDuration())
	for {
		//接收请求处理逻辑
		select {
			case logStr, ok := <-s.logChan:
				if !ok {//句柄广告异常关闭了退出
					log.Write(-1, "async proxy monitor exit!")
					return
				}
				if _, ok = state[logStr]; !ok {
					state[logStr]  = 1
				} else {
					state[logStr] += 1
				}
				//数据存储的比较多 也做一次同步
				if len(state) > 256 {
					s.syncRedis(state)
				}
				log.Write(log.INFO, "完成代理状态收集...")
			case <-syncChan:
				s.syncRedis(state) //做一个定期同步处理逻辑
				syncChan = time.After(PROXY_SYNC_REDIS_TIME)
			case <-notifyChan:
				s.syncReset() //将redis数据清理并生产汇总报表
				notifyChan = time.After(PROXY_SYNC_NOTIFY_TIME)
		}
	}
}

//每日做一个重置处理逻辑
func (s *Statistic) syncReset() {
	if s.rds == nil {
		return
	}
	for _, proxyItem := range s.proxy {
		ckey  := s.redisStatisticKey(proxyItem.Proxy)
		cmd   := s.rds.HGetAll(ckey)
		state := cmd.Val()
		s.rds.Del(ckey) //删除key信息
		if state != nil && state["proxy"] != proxyItem.Proxy {
			continue
		}
		str := s.formatNotify(state)
		//todo 发送钉钉通知处理逻辑
		log.Write(log.DEBUG, proxyItem.Proxy, str)
		fmt.Println(str)
	}
}

//获取数据资料信息
func (s *Statistic) ItemNotify(proxyHost string) string {
	ckey  := s.redisStatisticKey(proxyHost)
	cmd   := s.rds.HGetAll(ckey)
	state := cmd.Val()
	if state != nil && state["proxy"] != proxyHost {
		return proxyHost+"404不存在"
	}
	return s.formatNotify(state)
}

//统计格式化统计数据资料信息
func (s *Statistic) formatNotify(state map[string]string) string {
	success, _ := strconv.ParseInt(state["success"], 10, 64)
	if success < 1 {
		success += 1
	}
	regCmp, _  := regexp.Compile(":[\\d]+$")
	failed, _  := strconv.ParseInt(state["failed"], 10, 64)
	ratio  := fmt.Sprintf("%.6f", float64(success) / float64(success + failed) * 100.00)
	strBuf := strings.Builder{}
	strBuf.WriteString("代理服务:"+state["proxy"]+"\r\n")
	strBuf.WriteString("状态200请求数:"+state["success"]+"\r\n")
	strBuf.WriteString("状态非200请求数:"+state["failed"]+"\r\n")
	strBuf.WriteString("计算成功率:"+ratio+"\r\n")
	strBuf.WriteString("请求失败明细:\r\n")
	for keyStr, val := range state {
		if ok := regCmp.MatchString(keyStr); ok {
			strBuf.WriteString("\t-"+keyStr+" 累计数:"+val+"\r\n")
		}
	}
	return strBuf.String()
}

//获取redis数据资料信息
func (s Statistic) redisStatisticKey(proxy string) string {
	return "proxy@"+proxy
}

//将数据迁移到redis当中的处理逻辑
func (s *Statistic) syncRedis(state map[string]int) {
	if s.rds == nil {//数据为空的情况
		return
	}
	for logStr, nSize := range state {
		logState := LogStateBuilder(logStr)
		//丢弃异常的数据 数据处理逻辑 失败的情况
		if logState == nil || logState.ProxyIdx > len(s.proxy) {
			continue
		}
		//统计代理异常情况数据资料信息
		proxyItem := &s.proxy[logState.ProxyIdx]
		ckey, field := s.redisStatisticKey(proxyItem.Proxy), "success"
		if logState.Status != http.StatusOK {
			field    = "failed"
		}
		s.rds.HSetNX(ckey, "proxy", proxyItem.Proxy)
		s.rds.HIncrBy(ckey, field, int64(nSize))
		if logState.Status != http.StatusOK {//记录失败的域名明细
			field = fmt.Sprintf("%s:%d", logState.Host, logState.Status)
			s.rds.HIncrBy(ckey, field, int64(nSize))
		}
		delete(state, logStr)
	}
}

//上报统计数据资料信息往队列写，然后异步协程同步更新到redis当中
func (s *Statistic) Report(idx int, host string, statusCode int)  {
	if idx < 0 || idx > len(s.proxy) {//如果没有定位到代理的情况
		return
	}
	log.Write(log.INFO, "代理监控状态通知....")
	logState := LogStateSt{ProxyIdx: idx, Host: host, Status: statusCode}
	s.logChan <- logState.String()
	if statusCode != http.StatusOK {//请求失败的情况
		n := atomic.AddUint64(&s.proxy[idx].Error, 1)
		if n > PROXY_ERROR_LIMIT && s.proxy[idx].IFGet != nil {
			(&s.proxy[idx]).CutProxy() //自动切换ip
			return
		}
		maxLockTime := time.Second*300
		if PROXY_ERROR_LOCK_TIME * time.Duration(n) > maxLockTime {
			s.proxy[idx].Expire = time.Now().UnixNano() + int64(maxLockTime)
			return //设置最多锁定上线5分钟
		}
		if s.proxy[idx].Status == 1 {
			s.proxy[idx].Expire = time.Now().UnixNano()
		}
		s.proxy[idx].Status  = 2
		s.proxy[idx].Expire += int64(PROXY_ERROR_LOCK_TIME)
	} else {//只要成功就重置
		s.proxy[idx].Expire, s.proxy[idx].Status = 0, 1
		atomic.StoreUint64(&s.proxy[idx].Error, 0)
	}
}

//代理调度处理逻辑
func (s *Statistic) Proxy() (int, string) {
	n := atomic.AddUint64(&s.Request, 1)
	for i := 0; i < s.len; i++ {
		idx  := int((n+uint64(i))%uint64(s.len))
		item := &s.proxy[idx]
		if item.IFGet != nil && len(item.Url) < 1 {
			item.CutProxy() //切代理
		}
		//状态正常 且解锁的状态 直接处理逻辑即可
		if item.Status == 1 {
			return idx, item.Url
		} else if item.Status == 2 && item.Expire < time.Now().UnixNano() {
			item.Status = 1
			return idx, item.Url
		}
	}
	return -1, ""
}