package wipo

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
	Qk           string `json:"qk"`
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
