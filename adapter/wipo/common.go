package wipo

import (
	"encoding/json"
	"github.com/leicc520/go-crawler-srv/lib"
	"github.com/leicc520/go-orm"
	"github.com/leicc520/go-orm/log"
	"time"
)

type SqSt struct {
	Te string `json:"te"`
	Fi string `json:"fi"`
	Dt string `json:"dt"`
}

type SearchSt struct {
	Sq []SqSt `json:"sq"`
}

type PSt struct {
	Search SearchSt `json:"search"`
	Rows  int `json:"rows"`
	Start int `json:"start"`
}

//定义请求变量处理逻辑
type Request struct {
	P      PSt    `json:"p"`
	Type   string `json:"type"`
	La     string `json:"la"`
	Qi     string `json:"qi"`
	Queue  int    `json:"queue"`
	Field6 string `json:"_"`
}

type DocItem struct {
	OO     string    `json:"OO"`
	Score  float64   `json:"score"`
	STATUS string    `json:"STATUS"`
	AD     time.Time `json:"AD"`
	HOL    []string  `json:"HOL"`
	NC     []int     `json:"NC"`
	IMG    string    `json:"IMG"`
	SOURCE string    `json:"SOURCE"`
	DOC    string    `json:"DOC"`
	ID     string    `json:"ID"`
	BRAND  []string  `json:"BRAND"`
	HOLC   []string  `json:"HOLC"`
}

//返回结果的回参数据信息
type Response struct {
	LastUpdated int64  `json:"lastUpdated"`
	Sv          string `json:"sv"`
	Response    struct {
		Docs []DocItem `json:"docs"`
		NumFound int     `json:"numFound"`
		Start    int     `json:"start"`
		MaxScore float64 `json:"maxScore"`
	} `json:"response"`
	Qi           string `json:"qi"`
	Qk           string `json:"qk"`
}

//设置缓存的情况处理逻辑
func setCache(state *WipoSt) {
	if lib.Redis == nil {//未作初始化的情况
		return
	}
	skey := "wipo@"+state.StartDate.Format(orm.DATEYMDFormat)
	skey += "-"+state.EndDate.Format(orm.DATEYMDFormat)
	str, _ := json.Marshal(state)
	lib.Redis.Set(skey, str, time.Hour*72) //缓存三天
}

//获取缓存数据资料信息
func getCache(state *WipoSt) {
	if lib.Redis == nil {//未作初始化的情况
		return
	}
	skey := "wipo@"+state.StartDate.Format(orm.DATEYMDFormat)
	skey += "-"+state.EndDate.Format(orm.DATEYMDFormat)
	str, err := lib.Redis.Get(skey).Result()
	if len(str) < 1 || err != nil {
		return
	}
	cState := WipoSt{}
	if err = json.Unmarshal([]byte(str), &cState); err == nil {
		if cState.RangeDate != nil && len(cState.RangeDate) == 2 {
			state.RangeDate = cState.RangeDate
		}
		state.IndexPage = cState.IndexPage
		log.Write(-1, "完成初始化逻辑", cState.IndexPage, cState.RangeDate)
	}
}