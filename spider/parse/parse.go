package parse

import (
	"errors"
	"regexp"
)

var ErrEmpty = errors.New("Spider Node Not Exists!")

type RegExpParseSt string

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