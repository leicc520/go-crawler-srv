package parse

import (
	"github.com/antchfx/htmlquery"
	"github.com/leicc520/go-orm/log"
	"golang.org/x/net/html"
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
	Temp 		string  `json:"temp"`
	CrawlType 	int8 	`json:"crawl_type"`
	NodeType	int8 	`json:"node_type"`
}

//执行业务逻辑解析处理逻辑
func (s TempNodeSt) RunParse(node *html.Node) {
	switch s.NodeType {
	case NODE_TYPE_XPATH:
		if s.CrawlType == CRAWL_TYPE_NODE {
			htmlquery.InnerText()
		} else {

		}
	case NODE_TYPE_REGREP:

	default:
		log.Write(log.ERROR, "spider node parse node not supportd")
	}
}
