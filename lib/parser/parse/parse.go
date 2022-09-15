package parse

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

var ErrEmpty = errors.New("要解析的节点数据不存在!")
var ErrType  = errors.New("类型不支持,无法操作!")
var ErrNode  = errors.New("节点配置异常,无法解析")

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


//将结果转换成slice
func convertSlice(result interface{}) []string {
	if aStr, ok := result.([]string); ok {
		return aStr
	}
	aStr := []string{fmt.Sprintf("%v", result)}
	return aStr
}

//转换成字符串
func convertString(result interface{}) string {
	if aStr, ok := result.([]string); ok {
		return strings.Join(aStr, "\r\n")
	}
	return fmt.Sprintf("%v", result)
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