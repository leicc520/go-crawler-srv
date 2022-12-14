package parse

import (
	"fmt"
	"strings"

	"github.com/antchfx/htmlquery"
	"github.com/leicc520/go-crawler-srv/lib"
	"golang.org/x/net/html"
)


//Xpath解析器的使用情况逻辑
type XPathParseSt struct {
	node *html.Node
}

//解析数据资料信息
func NewXPathParser(htmlStr string) (*XPathParseSt, error) {
	topNode, err := htmlquery.Parse(strings.NewReader(htmlStr))
	if err != nil {
		return nil, err
	}
	return &XPathParseSt{node: topNode}, nil
}

//通过文件获取解析器的逻辑
func NewXPathParserFromFile(file string) (*XPathParseSt, error) {
	htmlStr := lib.ReadFile(file)
	return NewXPathParser(htmlStr)
}

//验证是否取到xpath节点数据
func (s *XPathParseSt) HasNode(expr string) (has bool, err error) {
	node, err := htmlquery.Query(s.node, expr)
	if err == nil && node != nil {
		has = true
	}
	return
}

//获取节点的内部内容数据信息
func (s *XPathParseSt) InnerText(expr string) (text string, err error) {
	node, err := htmlquery.Query(s.node, expr)
	if err != nil {
		return
	} else if node == nil {
		err = ErrEmpty
		return
	} else {
		text = strings.TrimSpace(htmlquery.InnerText(node))
		return
	}
}

//获取节点数据 数组切片列表
func (s *XPathParseSt) InnerTexts(expr string) (texts []string, err error) {
	texts = make([]string, 0)
	nodes, err := htmlquery.QueryAll(s.node, expr)
	if err != nil {
		return
	}
	for _, node := range nodes {
		texts = append(texts, strings.TrimSpace(htmlquery.InnerText(node)))
	}
	return
}

//解析表格数据信息提取
func (s *XPathParseSt) ParseTable(expr string) (table [][]string, err error) {
	lines, err := htmlquery.QueryAll(s.node, fmt.Sprintf("%s//tr", expr))
	if err != nil {
		return
	}
	var cells []*html.Node
	for _, line := range lines {
		var rows []string
		cells, err = htmlquery.QueryAll(line, "./th|./td")
		if err != nil {
			return
		}
		for _, cell := range cells {
			rows = append(rows, strings.TrimSpace(lib.NormalizeSpace(htmlquery.InnerText(cell))))
		}
		table = append(table, rows)
	}
	if table == nil {
		err = ErrEmpty
	}
	return
}

//查询所有的节点数据
func (s *XPathParseSt) QueryAll(expr string) (nodes []*html.Node, err error) {
	nodes, err = htmlquery.QueryAll(s.node, expr)
	return
}

//查询返回一个节点数据信息即可
func (s *XPathParseSt) QueryNode(expr string) (node *html.Node, err error) {
	node, err = htmlquery.Query(s.node, expr)
	if err != nil {
		return
	}
	if node == nil {
		err = ErrEmpty
	}
	return
}

//获取节点执行属性的值信息
func (s *XPathParseSt) QueryNodeAttr(expr string, attrName string) (v string, err error) {
	node, err := s.QueryNode(expr)
	if err != nil {
		return
	}
	for _, attribute := range node.Attr {
		if attribute.Key == attrName {
			return attribute.Val, nil
		}
	}
	return "", ErrEmpty
}