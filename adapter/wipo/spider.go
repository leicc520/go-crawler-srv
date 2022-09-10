package wipo

import (
	"fmt"
	"github.com/leicc520/go-crawler-srv/lib/proxy"
	"github.com/leicc520/go-orm/log"
	"net/http"
	"regexp"
	"strings"
	"time"
)

const (
	wipoBaseURL = "https://branddb.wipo.int"
)

//定义请求变量处理逻辑
type Request struct {
	P struct {
		Search struct {
			Sq []struct {
				Te string `json:"te"`
				Fi string `json:"fi"`
				Dt string `json:"dt"`
			} `json:"sq"`
		} `json:"search"`
		Rows  int `json:"rows"`
		Start int `json:"start"`
	} `json:"p"`
	Type   string `json:"type"`
	La     string `json:"la"`
	Qi     string `json:"qi"`
	Queue  int    `json:"queue"`
	Field6 string `json:"_"`
}

//返回结果的回参数据信息
type Response struct {
	LastUpdated int64  `json:"lastUpdated"`
	Sv          string `json:"sv"`
	Response    struct {
		Docs     []interface{} `json:"docs"`
		NumFound int           `json:"numFound"`
		Start    int           `json:"start"`
		MaxScore int           `json:"maxScore"`
	} `json:"response"`
	Qi           string `json:"qi"`
	Highlighting struct {
	} `json:"highlighting"`
	FacetCounts struct {
		FacetIntervals struct {
		} `json:"facet_intervals"`
		FacetQueries struct {
			EDNOWDAY1DAYTONOWDAY6MONTHS int `json:"ED:[NOW/DAY+1DAY TO NOW/DAY+6MONTHS]"`
			ITYCOMBINED                 int `json:"ITY:COMBINED"`
			EDNOWDAY1DAYTONOWDAY1MONTH  int `json:"ED:[NOW/DAY+1DAY TO NOW/DAY+1MONTH]"`
			EDNOWDAY1DAYTO              int `json:"ED:[NOW/DAY+1DAY TO *]"`
			EDTONOWDAY                  int `json:"ED:[* TO NOW/DAY]"`
			EDNOWDAY1DAYTONOWDAY1YEAR   int `json:"ED:[NOW/DAY+1DAY TO NOW/DAY+1YEAR]"`
			ITYNONVERBAL                int `json:"ITY:NONVERBAL"`
			EDNOWDAY1MONTHTONOWDAY      int `json:"ED:[NOW/DAY-1MONTH TO NOW/DAY]"`
			ITYVERBAL                   int `json:"ITY:VERBAL"`
			ITYUNKNOWN                  int `json:"ITY:UNKNOWN"`
		} `json:"facet_queries"`
		FacetFields struct {
			OO struct {
			} `json:"OO"`
			STATUS struct {
			} `json:"STATUS"`
			MTY struct {
			} `json:"MTY"`
			SOURCE struct {
			} `json:"SOURCE"`
			HOLC struct {
			} `json:"HOLC"`
		} `json:"facet_fields"`
		FacetHeatmaps struct {
		} `json:"facet_heatmaps"`
		FacetRanges struct {
		} `json:"facet_ranges"`
	} `json:"facet_counts"`
}

type WipoSt struct {
	StartPage int
	EndPage   int
	Qi        string
	PageSize  int
	IndexPage int
	client   *proxy.HttpSt
}

//初始化处理流程
func (s *WipoSt) Init()  {
	s.client = proxy.NewHttpRequest()
	agent := s.client.ResetAgent() //浏览器头信息
	s.firstStepInitCookie(agent)
}

//循环请求，直到成功为止，失败的话休眠最多1分钟
func (s *WipoSt) Request(link string, body []byte, method string) (result string, err error) {
	nTry, sleepTime := 0, time.Duration(0)
	for {
		result, err = s.client.Request(link, body, method)
		if err == nil {
			return
		}
		sleepTime = time.Millisecond*100*time.Duration(nTry)
		if sleepTime > time.Second*60 {//最多休眠1分钟
			sleepTime = time.Second*60
		}
		time.Sleep(sleepTime)
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
func (s *WipoSt) firstStepInitCookie(agent string)  {
	link := wipoBaseURL+"/branddb/en/#"
	result, _ := s.Request(link, nil, http.MethodGet)
	visitorId, err := wipoVisitorUunId(agent)
	if err != nil {
		log.Write(-1, "获取wipo访客ID失败.", err)
	}
	s.client.SetCookie(link+"/", "wipo-visitor-uunid="+visitorId)
	s.Qi = "1-"+s.parseQk2QiString(result)
	fmt.Println(result)
	//qk = "rx0aodmPMMlXiR/ABFsCBmatZsCMvsgv7XRUXKfGaQI";
	/*
	{"type":"brand","la":"en","qi":"0-Q5f6M8X2Sz6h5v3xcvS0Pk0d6TFgZFZWrmbkMw3YD60=","queue":1,"_":"11932"} <nil>
	{"type":"brand","la":"en","qi":"0-EUDiJKCxsjZjpAR1fo1TYxWLr5cZa23epH+D+4oTCx8=","queue":1,"_":"11932"} <nil>
		
		qk = "EUDiJKCxsjZjpAR1fo1TYxWLr5cZa23epH+D+4oTCx8=";*/

}

//初始化处理逻辑
func (s *WipoSt) Run()  {

}


