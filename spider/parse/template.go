package parse

type TemplateSt struct {
	Templates []struct{
		TaskName string 	 `json:"task_name"`  //通过任务名称定位模板
		Elements []ElementSt `json:"elements"`   //使用的模板提取元素
	}  `json:"templates"`
}

//获取模板配置数据资料信息
type CompilerSt struct {
	doc  string
	XPathParser *XPathParseSt `json:"-"`
	QueryParser *QueryParseSt `json:"-"`
	JsonParser  *JsonParseSt `json:"-"`
}

//解析模块数据资料信息
func (s *CompilerSt) SetDoc(doc string)  {
	s.doc = doc
}

//获取解析匹配的模板
func (s *CompilerSt) GetDoc() string {
	return s.doc
}

//获取解析器模板引擎
func (s *CompilerSt) getParser(nodeType int8) InNodeParser {
	switch nodeType {
	case NODE_TYPE_XPATH:
		return InNodeParser(s.getXPathParser())
	case NODE_TYPE_QUERY:
		return InNodeParser(s.getQueryParser())
	case NODE_TYPE_JSON:
		return InNodeParser(s.getJsonParser())
	case NODE_TYPE_REGREP:
		return InNodeParser(RegExpParseSt(s.doc))
	}
	panic("get Parser nodeType Not Support")
}

//获取Json的解析器
func (s *CompilerSt) getJsonParser() *JsonParseSt {
	var err error = nil
	if s.JsonParser == nil {
		s.JsonParser, err = NewJsonParser(s.doc)
		if err != nil {//异常退出 非json的情况
			panic(err)
		}
	}
	return s.JsonParser
}

//获取go-query的解析器
func (s *CompilerSt) getQueryParser() *QueryParseSt {
	var err error = nil
	if s.QueryParser == nil {
		s.QueryParser, err = NewQueryParse(s.doc)
		if err != nil {//异常退出 非json的情况
			panic(err)
		}
	}
	return s.QueryParser
}

//获取go-xpath的解析器
func (s *CompilerSt) getXPathParser() *XPathParseSt {
	var err error = nil
	if s.XPathParser == nil {
		s.XPathParser, err = NewXPathParser(s.doc)
		if err != nil {//异常退出 非json的情况
			panic(err)
		}
	}
	return s.XPathParser
}