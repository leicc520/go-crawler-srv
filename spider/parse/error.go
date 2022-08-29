package parse

import "strings"

type ParseError []string

//获取错误数据信息收集
func (e ParseError) Error() string {
	return strings.Join(e, "\r\n")
}

//判定是否为空的情况
func (e ParseError) IsEmpty() bool {
	return len(e) < 1
}

//包裹错误数据信息
func (e *ParseError) Wrapped(field string, err error) {
	*e = append(*e, field +"="+ err.Error())
}



