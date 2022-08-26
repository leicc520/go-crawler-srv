package parse

import (
	"errors"
	"fmt"
)

const (
	NODE_TYPE_XPATH = 1  //通过xpath定位元素
	NODE_TYPE_QUERY = 2  //通过go-query查询
	NODE_TYPE_JSON  = 3  //通过json-query解析
	NODE_TYPE_REGREP= 4  //通过正则提取元素

	CRAWL_TYPE_NODE = 1  //采集单个元素
	CRAWL_TYPE_LIST = 2  //采集一个列表元素
)

//模板节点配置，每个字符模板有类型 允许递归查询节点
type TempNodeSt struct {
	Name        string  `json:"name"`
	Temp 		string  `json:"temp"`
	CrawlType 	int8 	`json:"crawl_type"`
	NodeType	int8 	`json:"node_type"`
}

//执行业务逻辑解析处理逻辑
func (s TempNodeSt) RunParse(p InNodeParser) (result interface{}, err error) {
	switch s.CrawlType {
	case CRAWL_TYPE_NODE:
		return  p.InnerText(s.Temp)
	case CRAWL_TYPE_LIST:
		return p.InnerTexts(s.Temp)
	}
	return nil, errors.New("Crawler Type Not Support")
}

//将结果转换成slice
func convertSlice(result interface{}) []string {
	if aStr, ok := result.([]string); ok {
		return aStr
	}
	aStr := []string{fmt.Sprintf("%v", result)}
	return aStr
}
