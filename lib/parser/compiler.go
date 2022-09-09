package parser

import (
	"errors"

	"github.com/leicc520/go-crawler-srv/lib/parser/parse"
	"github.com/leicc520/go-orm"
)

//获取模板配置数据资料信息
type CompilerSt struct {
	DocHtml  string
	XPathParser *parse.XPathParseSt `json:"-"`
	QueryParser *parse.QueryParseSt `json:"-"`
	JsonParser  *parse.JsonParseSt  `json:"-"`
}

//获取生成一个模板编译器
func NewCompiler(doc string) *CompilerSt {
	return &CompilerSt{DocHtml: doc}
}

//模板解析器处理逻辑
func (s *CompilerSt) Parse(elements []parse.ElementSt) (result orm.SqlMap, err error) {
	result = make(orm.SqlMap)
	if elements == nil || len(elements) < 1 {
		err = errors.New("解析器模块元素获取失败")
		return
	}
	parseErr := parse.ParseError{}
	for _, element := range elements {
		err = element.RunParse(s, result)
		if err != nil {
			parseErr.Wrapped(element.Tag, err)
		}
	}
	//如果数据不为空的情况 直接返回空数据信息
	if !parseErr.IsEmpty() {
		err = parseErr
	}
	return
}

//解析模块数据资料信息
func (s *CompilerSt) SetDoc(doc string) {
	s.DocHtml = doc
}

//获取解析匹配的模板
func (s *CompilerSt) GetDoc() string {
	return s.DocHtml
}

//克隆一个对象返回接口对象
func (s *CompilerSt) Clone(doc string) parse.IFCompiler {
	return parse.IFCompiler(&CompilerSt{DocHtml: doc})
}

//获取解析器模板引擎
func (s *CompilerSt) GetParser(nodeType int8) parse.IFNodeParser {
	switch nodeType {
	case parse.NODE_TYPE_XPATH:
		return parse.IFNodeParser(s.getXPathParser())
	case parse.NODE_TYPE_QUERY:
		return parse.IFNodeParser(s.getQueryParser())
	case parse.NODE_TYPE_JSON:
		return parse.IFNodeParser(s.getJsonParser())
	case parse.NODE_TYPE_REGREP:
		return parse.IFNodeParser(parse.RegExpParseSt(s.DocHtml))
	}
	panic("get Parser nodeType Not Support")
}

//获取Json的解析器
func (s *CompilerSt) getJsonParser() *parse.JsonParseSt {
	var err error = nil
	if s.JsonParser == nil {
		s.JsonParser, err = parse.NewJsonParser(s.DocHtml)
		if err != nil {//异常退出 非json的情况
			panic(err)
		}
	}
	return s.JsonParser
}

//获取go-query的解析器
func (s *CompilerSt) getQueryParser() *parse.QueryParseSt {
	var err error = nil
	if s.QueryParser == nil {
		s.QueryParser, err = parse.NewQueryParse(s.DocHtml)
		if err != nil {//异常退出 非json的情况
			panic(err)
		}
	}
	return s.QueryParser
}

//获取go-xpath的解析器
func (s *CompilerSt) getXPathParser() *parse.XPathParseSt {
	var err error = nil
	if s.XPathParser == nil {
		s.XPathParser, err = parse.NewXPathParser(s.DocHtml)
		if err != nil {//异常退出 非json的情况
			panic(err)
		}
	}
	return s.XPathParser
}