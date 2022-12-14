package proxy

import (
	"net/http"
	"sync"
	"sync/atomic"
	"time"

	"github.com/go-redis/redis"
	"github.com/leicc520/go-orm/log"
)

/************************************************************
	代理请求数据业务统计处理逻辑
 */
//汇总数据处理逻辑
type Monitor struct {
	proxy []ProxySt
	len    int
	Request uint64 //请求的统计数值
}

//获取数据资料信息
var (
	monitorState map[string]*Monitor = nil
	statRedis *RedisStateSt = nil
	onceInit = sync.Once{}
)

//初始化对象数据资料信息
func NewMonitor(proxy []ProxySt) *Monitor {
	return &Monitor{proxy: proxy, len: len(proxy)}
}

//初始化数据资料信息 放置默认代理
func Init(proxy []ProxySt, rds *redis.Client) {
	onceInit.Do(func() {
		statRedis = NewRedisStateSt(rds)
		if len(proxy) > 0 {
			monitorState[PROXY_DEFUALT_NAME] = NewMonitor(proxy)
		}
	})
}

//设置注册代理监控
func SetMonitor(name string, monitor *Monitor) {
	monitorState[name] = monitor
}

//返回统计数据资料信息
func GetMonitor(name string) *Monitor {
	if len(name) < 1 {
		name = PROXY_DEFUALT_NAME
	}
	if ss, ok := monitorState[name]; ok {
		return ss
	}
	panic("代理监控["+name+"]不存在,无法获取")
}

//获取代理地址
func (s *Monitor) GetProxy(isTcp, isCut bool) string {
	nlen := len(s.proxy)
	for i := 0; i < nlen; i++ {
		item := &s.proxy[i]
		if len(item.Url) < 1 || isCut {
			item.CutProxy() //切换代理
		}
		if isTcp == item.IsTcp {
			return item.Url
		}
	}
	return ""
}

//获取数据资料信息
func (s *Monitor) ItemNotify(proxy string) string {
	if statRedis == nil {
		return "未开启代理监控统计..."
	}
	_, state := statRedis.state(proxy)
	if state != nil && state["proxy"] != proxy {
		return proxy+"404不存在"
	}
	return formatNotify(state)
}

//上报统计数据资料信息往队列写，然后异步协程同步更新到redis当中
func (s *Monitor) Report(idx int, host string, statusCode int)  {
	if idx < 0 || idx > len(s.proxy) || statRedis == nil {//如果没有定位到代理的情况
		return
	}
	log.Write(log.INFO, "代理监控状态通知....")
	logState := logStateSt{ProxyIdx: idx, Host: host, Status: statusCode, Proxy: s.proxy[idx].Proxy}
	statRedis.Chan() <- logState.toString()
	if statusCode != http.StatusOK {//请求失败的情况
		n := atomic.AddUint64(&s.proxy[idx].Error, 1)
		if n > PROXY_ERROR_LIMIT {
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
func (s *Monitor) Proxy() (int, string) {
	n := atomic.AddUint64(&s.Request, 1)
	for i := 0; i < s.len; i++ {
		idx  := int((n+uint64(i))%uint64(s.len))
		item := &s.proxy[idx]
		if len(item.Url) < 1 {
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
	log.Write(-1, "代理全军覆膜，使用本机IP尝试.")
	return -1, ""
}