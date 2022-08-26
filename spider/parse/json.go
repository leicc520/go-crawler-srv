package parse

import (
	"github.com/antchfx/jsonquery"
	"github.com/leicc520/go-crawler-srv/lib"
	"strings"
)

type JsonParseSt struct {
	node *jsonquery.Node
}

//解析数据资料信息
func NewJsonParser(jsonStr string) (*JsonParseSt, error) {
	topNode, err := jsonquery.Parse(strings.NewReader(jsonStr))
	if err != nil {
		return nil, err
	}
	return &JsonParseSt{node: topNode}, nil
}

//通过文件获取解析器的逻辑
func NewJsonParserFromFile(file string) (*JsonParseSt, error) {
	jsonStr := lib.ReadFile(file)
	return NewJsonParser(jsonStr)
}

//获取节点的数据资料信息
func (s *JsonParseSt) InnerText(expr string) (text string, err error) {
	node, err := jsonquery.Query(s.node, expr)
	if err != nil {
		return
	}
	if node == nil {
		err = ErrEmpty
		return
	}
	text = node.InnerText()
	return
}

//获取节点的数据资料信息
func (s *JsonParseSt) InnerTexts(expr string) (texts []string, err error) {
	nodes, err := jsonquery.QueryAll(s.node, expr)
	if err != nil {
		return
	}
	texts  = make([]string, 0)
	for _, node := range nodes {
		texts = append(texts, node.InnerText())
	}
	return
}

//验证是否取到 节点数据
func (s *JsonParseSt) HasNode(expr string) (has bool, err error) {
	node, err := jsonquery.Query(s.node, expr)
	if err == nil && node != nil {
		has = true
	}
	return
}

//验证节点取值是否true
func (s *JsonParseSt) NodeValueIsTrue(expr string) (r bool, err error) {
	node, err := jsonquery.Query(s.node, expr)
	if err != nil {
		return
	}
	if node == nil {
		err = ErrEmpty
		return
	}
	if node.InnerText() == "true" {
		r = true
	} else {
		r = false
	}
	return
}




