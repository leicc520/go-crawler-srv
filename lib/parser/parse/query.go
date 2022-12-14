package parse

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/leicc520/go-crawler-srv/lib"
	"strings"
)

//Xpath解析器的使用情况逻辑
type QueryParseSt struct {
	node *goquery.Document
}

//解析数据资料信息
func NewQueryParse(htmlStr string) (*QueryParseSt, error) {
	topNode, err := goquery.NewDocumentFromReader(strings.NewReader(htmlStr))
	if err != nil {
		return nil, err
	}
	return &QueryParseSt{node: topNode}, nil
}

//通过文件获取解析器的逻辑
func NewQueryParseFromFile(file string) (*QueryParseSt, error) {
	htmlStr := lib.ReadFile(file)
	return NewQueryParse(htmlStr)
}

//获取节点的数据资料信息
func (s *QueryParseSt) TextHTML(expr string) (text string, err error) {
	sel := s.node.Find(expr)
	if sel.Length() < 1 {
		err = ErrEmpty
		return
	}
	text, err = sel.Html()
	return
}

//获取节点的数据资料信息
func (s *QueryParseSt) InnerText(expr string) (text string, err error) {
	sel := s.node.Find(expr)
	if sel.Length() < 1 {
		err = ErrEmpty
		return
	}
	sel.Find("*").RemoveFiltered("style,noscript,script")
	text = lib.NormalizeSpace(strings.TrimSpace(sel.Text()))
	return
}

//获取节点的数据资料信息
func (s *QueryParseSt) InnerTexts(expr string) (texts []string, err error) {
	sel := s.node.Find(expr)
	if sel.Length() < 1 {
		err = ErrEmpty
		return
	}
	texts  = make([]string, 0)
	sel.Find("*").RemoveFiltered("style,noscript,script")
	f   := func(_ int, tmpSel *goquery.Selection) string {
		return lib.NormalizeSpace(strings.TrimSpace(tmpSel.Text()))
	}
	texts = sel.Map(f)
	return
}

