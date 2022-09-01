package parse

import (
	"errors"
	"regexp"
)

var ErrEmpty = errors.New("Spider Node Not Exists!")

type RegExpParseSt string

//根据提取规格获取数据
type IFNodeParser interface {
	InnerText(expr string) (text string, err error)
	InnerTexts(expr string) (texts []string, err error)
}

//编译器的实现接口类型 解决依赖问题
type IFCompiler interface {
	SetDoc(doc string)
	GetDoc() string
	Clone(doc string) IFCompiler
	GetParser(nodeType int8) IFNodeParser
}


//获取节点的数据资料信息
func (s RegExpParseSt) InnerText(expr string) (text string, err error) {
	var reg *regexp.Regexp
	reg, err = regexp.Compile(expr)
	if err != nil  {
		return
	}
	text = reg.FindString(string(s))
	if len(text) < 1 {
		err = ErrEmpty
	}
	return
}

//获取节点的数据资料信息
func (s RegExpParseSt) InnerTexts(expr string) (texts []string, err error) {
	var reg *regexp.Regexp
	reg, err = regexp.Compile(expr)
	if err != nil  {
		return
	}
	texts = reg.FindAllString(string(s), -1)
	if len(texts) < 1 {
		err = ErrEmpty
	}
	return
}