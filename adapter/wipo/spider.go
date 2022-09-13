package wipo

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"

	LZString "github.com/Lazarus/lz-string-go"
	"github.com/leicc520/go-crawler-srv/lib/proxy"
	"github.com/leicc520/go-orm"
	"github.com/leicc520/go-orm/log"
)

const (
	wipoBaseURL = "https://branddb.wipo.int"
	wipoEndDate = "2022-09-12"
	wipo7Days   = time.Hour*24*7
	wipoPageSize= 60
)


type WipoSt struct {
	TotalPage int           `json:"end_page"`
	IndexPage int           `json:"index_page"`
	EndDate   time.Time     `json:"end_date"`
	StartDate time.Time     `json:"start_date"`
	RangeDate []time.Time   `json:"range_date"`
	Qi        string        `json:"qi"`
	Qk        string        `json:"qk"`
	OpRequest int           `json:"op_request"`
	qz        Request       `json:"-"`
	client   *proxy.HttpSt  `json:"-"`
}

//初始化处理流程
func (s *WipoSt) init()  {
	s.client = proxy.NewHttpRequest().IsProxy(true)
	agent := s.client.ResetAgent() //浏览器头信息
	s.qz = Request{Type: "brand", La: "en", Queue: 1, Field6: "11932", P: PSt{Rows: wipoPageSize, Start: 0}}
	s.stepInitCookie(agent)
	s.stepInitSelect() //设置操作
}

//循环请求，直到成功为止，失败的话休眠最多1分钟
func (s *WipoSt) request(link string, body []byte, method string) (result string, err error) {
	nTry, sleepTime := 0, time.Duration(0)
	for {
		result, err = s.client.Request(link, body, method)
		if err == nil {
			return
		}
		nTry ++
		sleepTime = time.Millisecond*100*time.Duration(nTry)
		if sleepTime > time.Second*60 {//最多休眠1分钟
			sleepTime = time.Second*60
			log.Write(-1, "请求出现异常，重新加入处理逻辑")
		}
		time.Sleep(sleepTime)
		if nTry > 1000 {//失败太多次了也是走重新初始化的逻辑
			return
		}
	}
}

//去掉空格，然后正则提取
func (s *WipoSt) parseQk2QiString(str string) string {
	str     = strings.ReplaceAll(str, " ", "")
	reg, _ := regexp.Compile("qk=\"([^\"]+)\"")
	arrStr := reg.FindAllStringSubmatch(str, -1)
	if len(arrStr) < 1 {
		return ""
	}
	//取最后一个QK为有效的QK
	return arrStr[len(arrStr)-1][1]
}

//第一步初始化要做的事件
func (s *WipoSt) stepInitCookie(agent string) error  {
	s.client.Reset() //初始化处理逻辑，包括jar的初始化
	link      := wipoBaseURL+"/branddb/en/#"
	result, _ := s.request(link, nil, http.MethodGet)
	visitorId, err := wipoVisitorUunId(agent)
	if err != nil {
		log.Write(-1, "获取wipo访客ID失败.", err)
		return err
	}
	s.client.SetCookie(link+"/", "wipo-visitor-uunid="+visitorId)
	s.Qk = s.parseQk2QiString(result)
	if len(s.Qk) < 1 {
		log.Write(-1, "获取首页QK失败")
		return errors.New("获取首页QK失败")
	}
	s.Qi = "0-"+s.Qk
	return nil
}

//初始化设置操作处理逻辑
func (s *WipoSt) stepInitSelect() {
	link   := wipoBaseURL + "/branddb/jsp/select.jsp"
	query  := url.Values{}
	query.Set("type", "brand")
	query.Set("q", "ID:TIME")
	s.OpRequest += 1
	query.Set("qi", fmt.Sprintf("%d-%s", s.OpRequest, s.Qk))
	header := map[string]string{"x-requested-with":"XMLHttpRequest", "referer":wipoBaseURL+"/branddb/en/"}
	s.client.SetHeader(header)
	link   += "?"+query.Encode()
	result, err := s.client.Request(link, nil, http.MethodGet)
	s.client.DelHeader(header)
	if err != nil { //请求失败的情况
		log.Write(-1, err)
	}
	log.Write(log.INFO, "初始化设置操作", result)
}

//请求获取数据资料信息
func (s *WipoSt) selectData() (*Response, error) {
	qzStr  := s.genSelectQZParams()
	link   := wipoBaseURL + "/branddb/jsp/select.jsp"
	query  := url.Values{}
	query.Set("qz", qzStr)
	body   := []byte(query.Encode())
	header := map[string]string{"x-requested-with":"XMLHttpRequest", "accept":"application/json, text/javascript, */*; q=0.01",
		"content-type":"application/x-www-form-urlencoded; charset=UTF-8", "referer":wipoBaseURL+"/branddb/en/",
		"content-length":strconv.FormatInt(int64(len(body)), 10)}
	s.client.SetHeader(header)
	time.Sleep(time.Second*5) //休眠时间，不休眠的话会500
	result, err := s.client.Request(link, body, http.MethodPost)
	s.client.DelHeader(header)
	if err != nil { //请求失败的情况
		return nil, err
	}
	sp  := Response{} //结构到结构体当中
	if err = json.Unmarshal([]byte(result), &sp); err != nil {
		log.Write(-1, "返回结果结构异常", err)
		return nil, err
	}
	//解析请求，复制到下一个请求的处理逻辑
	s.Qi = sp.Qi //向上取整，返回页数
	if len(sp.Qk) > 0 && len(sp.Qi) > 0 {
		s.Qk = sp.Qk
		arrStr := strings.Split(sp.Qi, "=")
		if len(arrStr) > 1 {//获取qi的另外一种方式
			idx, err := strconv.ParseInt(arrStr[0], 10, 64)
			if err == nil && idx >= 0 {
				s.Qi= fmt.Sprintf("%d-%s", idx, sp.Qk)
			}
		}
	}
	s.TotalPage = int(math.Ceil(float64(sp.Response.NumFound) / float64(wipoPageSize)))
	s.IndexPage+= 1
	return &sp, nil
}

//生成请求参数QZ的处理逻辑
func (s *WipoSt) genSelectQZParams() string {
	startDate := s.RangeDate[0].Format(orm.DATEYMDSTRFormat)
	endDate   := s.RangeDate[1].Format(orm.DATEYMDSTRFormat)
	dtStr     := "["+startDate+" TO "+endDate+"]"
	sq1 := SqSt{Dt: dtStr, Fi: "AD", Te: "["+startDate+"T00:00:00Z TO "+endDate+"T23:59:59Z]"}
	sq2 := SqSt{Dt: "", Fi:"OO", Te: "US"} //限制只要美国
	s.qz.P.Search = SearchSt{Sq: []SqSt{sq1, sq2}}
	s.qz.Qi       = s.Qi
	s.qz.P.Start  = s.IndexPage * wipoPageSize //计算已经到了第几页
	body, _ := json.Marshal(s.qz)
	qzStr   := LZString.Compress(string(body), "")
	log.Write(log.INFO, string(body))
	log.Write(log.INFO, qzStr)
	return qzStr
}

//解析数据处理逻辑
func (s *WipoSt) handle(indexPage int, sp *Response) {
	sorm := newApWipoBrand()
	if (sp.Response.Docs != nil && len(sp.Response.Docs) > 0) {
		for _, doc := range sp.Response.Docs {
			//分拆字段插入数据库当中
			docStr, _   := json.Marshal(doc.HOL)
			ncStr, _    := json.Marshal(doc.NC)
			brandStr, _ := json.Marshal(doc.BRAND)
			holcStr,  _ := json.Marshal(doc.HOLC)
			originStr, _:= json.Marshal(doc)
			sorm.NewOneFromHandler(func(st *orm.QuerySt) *orm.QuerySt {
				st.Value("oo", doc.OO).Value("score", doc.Score).Value( "status",doc.STATUS)
				st.Value("ad", doc.AD.Unix()).Value("hol", string(docStr)).Value("doc", doc.DOC)
				st.Value("nc", string(ncStr)).Value("img", doc.IMG).Value("source", doc.SOURCE)
				st.Value("number", doc.ID).Value("brand", string(brandStr)).Value("holc", string(holcStr))
				st.Value("origin", string(originStr)).Value("stime", time.Now().Unix())
				return st
			}, func(st *orm.QuerySt) interface{} {
				st.ConflictField("number") //冲突的话执行更新操作
				st.Duplicate("oo", doc.OO).Duplicate("score", doc.Score).Duplicate( "status",doc.STATUS)
				st.Duplicate("ad", doc.AD.Unix()).Duplicate("hol", string(docStr)).Duplicate("doc", doc.DOC)
				st.Duplicate("nc", string(ncStr)).Duplicate("img", doc.IMG).Duplicate("source", doc.SOURCE)
				st.Duplicate("brand", string(brandStr)).Duplicate("holc", string(holcStr))
				st.Duplicate("origin", string(originStr)).Duplicate("stime", time.Now().Unix())
				return nil
			})
		}
	}
	log.Write(log.INFO, "抓取数据", indexPage, s.RangeDate, sp)
}

//格式化日期数据资料信息
func (s *WipoSt) formatRange() string {
	startDate := s.RangeDate[0].Format(orm.DATEYMDSTRFormat)
	endDate   := s.RangeDate[0].Format(orm.DATEYMDSTRFormat)
	dtStr     := "["+startDate+" TO "+endDate+"]"
	return dtStr
}

//初始化处理逻辑
func (s *WipoSt) Run(startDate, endDate string)  {
	var err1 error = nil
	var err2 error = nil
	s.StartDate, err1 = time.Parse(orm.DATEYMDSTRFormat, startDate)
	s.EndDate, err2   = time.Parse(orm.DATEYMDSTRFormat, endDate)
	if err1 != nil || err2 != nil {
		log.Write(-1, "起止日期解析异常", err2, err1)
		panic("起止日期解析异常")
	}
	//这个时间前后抓数据 获取最近一周的数据
	s.RangeDate = []time.Time{s.StartDate, s.StartDate.Add(wipo7Days)}
	s.init() //初始化完成第一个请求
	getCache(s) //初始化逻辑
	for {
		log.Write(log.INFO, "遍历开始抓取："+s.formatRange()+"的数据")
		for { //遍历抓取该日期的数据列表
			if sp, err := s.selectData(); err != nil {
				s.init() //请求失败的情况 重新初始化
				continue
			} else {//请求正常的情况
				s.handle(s.IndexPage, sp)
				setCache(s) //设置缓存处理逻辑
				if s.IndexPage >= s.TotalPage  {//页数已经抓取完成了
					break
				}
			}
		}
		log.Write(log.INFO, "遍历结束抓取："+s.formatRange()+"的数据")
		//日期往后挪7天的处理逻辑
		startDate  := s.RangeDate[1].Add(time.Hour*24)
		endDate    := startDate.Add(wipo7Days)
		s.RangeDate = []time.Time{startDate, endDate}
		if startDate.After(s.EndDate) {
			log.Write(log.INFO, s.formatRange(), "数据已经抓取完毕...")
			break
		}
		s.IndexPage = 0 //重新从0开始抓
	}
}


