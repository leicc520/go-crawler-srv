package parse

import (
	"errors"
	"regexp"
	"strings"

	"github.com/leicc520/go-crawler-srv/lib"
	"github.com/leicc520/go-orm"
	"github.com/leicc520/go-orm/log"
)

const (
	NODE_TYPE_XPATH = 1  //通过xpath定位元素
	NODE_TYPE_QUERY = 2  //通过go-query查询
	NODE_TYPE_JSON  = 3  //通过json-query解析
	NODE_TYPE_REGREP= 4  //通过正则提取元素

	ELEMENT_TYPE_TEXT = "TEXT"  //采集内部元素
	ELEMENT_TYPE_IMAGE= "IMG"   //采集内部图片元素
	ELEMENT_TYPE_ATTR = "ATTR"  //采集元素属性
	ELEMENT_TYPE_URL  = "URL"   //采集元素地址
	ELEMENT_TYPE_LIST = "LIST"  //采集元素列表
)

//元素节点提取配置，便利模板节点直到知道数据才结束
type ElementSt struct {
	Tag   string 	  `json:"tag" yaml:"tag"` 	   //提取之后放到这个名字的map当中
	Name  string  	  `json:"name" yaml:"name"`	   //元素节点名称
	XPath 	[]string  `json:"xPath" yaml:"xPath"`
	CssPath []string  `json:"cssPath" yaml:"cssPath"`
	Json    []string  `json:"json"  yaml:"json"`
	Regexp  []string  `json:"regexp" yaml:"regexp"`
	MatchReg string   `json:"matchReg" yaml:"matchReg"`
	Type 	string 	  `json:"type" yaml:"type"`
	Elements []ElementSt `json:"elements" yaml:"elements"`  //允许递归的获取元素，在当前解析节点继续解析
}

//格式化成字符串输出数据
func (s ElementSt) String() string {
	arrStr := []string{s.Name, s.Tag, s.MatchReg}
	if s.XPath != nil && len(s.XPath) > 0 {
		arrStr  = append(arrStr, "xpath:" + strings.Join(s.XPath, "|"))
	}
	if s.CssPath != nil && len(s.CssPath) > 0 {
		arrStr  = append(arrStr, "csspath:" + strings.Join(s.CssPath, "|"))
	}
	if s.Json != nil && len(s.Json) > 0 {
		arrStr  = append(arrStr, "json:" + strings.Join(s.Json, "|"))
	}
	if s.Regexp != nil && len(s.Regexp) > 0 {
		arrStr  = append(arrStr, "regexp:" + strings.Join(s.Regexp, "|"))
	}
	return strings.Join(arrStr, ";")
}

//执行业务逻辑解析处理逻辑
func (s ElementSt) RunParse(t IFCompiler, result orm.SqlMap) error {
	value, err := s.getValue(t)
	if err != nil {
		return err
	}
	//判断是否继续匹配逻辑
	if s.Elements != nil && len(s.Elements) > 0 {
		aStr := convertSlice(value)
		list := make([]orm.SqlMap, 0)
		for _, doc := range aStr {
			//在每个匹配节点下接续查找数据
			newCP:= t.Clone(doc)
			item := orm.SqlMap{}
			for _, el := range s.Elements {
				err = el.RunParse(newCP, item)
				if err != nil {
					return err
				}
			}
			if len(item) > 0 {
				list = append(list, item)
			}
		}
		value = list
	}
	//记录匹配结果到map当中 为空的字段忽略不返回
	if len(s.Tag) > 0 && s.Tag != "-" {
		result[s.Tag] = value
	}
	return nil
}

//执行获取匹配的结果数据处理逻辑
func (s ElementSt) getValue(t IFCompiler) (value interface{}, err error) {
	value, err = s.nodeParse(t)
	if err != nil {//节点取值匹配失败的情况
		log.Write(-1, s.Name, s.Tag, err)
		log.Write(log.INFO, t.GetDoc())
		return
	}
	result  := convertString(value)
	if len(s.MatchReg) > 0 {//正则过滤提取
		if list, err := s.regFilter(result); err == nil {
			value = list
		}
	}
	summary := lib.CutStr(result, 64, "...")
	log.Write(log.INFO, s.String(), summary)
	return
}

//正则提取逻辑
func (s ElementSt) regFilter(result string) (string, error)  {
	if reg, err := regexp.Compile(s.MatchReg); err == nil {
		str := reg.FindString(result)
		if len(str) > 0 {
			return str, nil
		}
	}
	return "", errors.New("过滤器未生效:"+s.MatchReg)
}

//执行业务逻辑解析处理逻辑
func (s ElementSt) nodeParse(t IFCompiler) (result interface{}, err error) {
	if len(s.XPath) > 0 {
		p := t.GetParser(NODE_TYPE_XPATH)
		result, err = s.runParse(s.XPath, p)
		if err == nil {//模板解析提取成功
			return
		}
	}
	if len(s.CssPath) > 0 {
		p := t.GetParser(NODE_TYPE_QUERY)
		result, err = s.runParse(s.CssPath, p)
		if err == nil {//模板解析提取成功
			return
		}
	}
	if len(s.Regexp) > 0 {
		p := t.GetParser(NODE_TYPE_REGREP)
		result, err = s.runParse(s.Regexp, p)
		if err == nil {//模板解析提取成功
			return
		}
	}
	if len(s.Json) > 0 {
		p := t.GetParser(NODE_TYPE_JSON)
		result, err = s.runParse(s.Json, p)
		if err == nil {//模板解析提取成功
			return
		}
	}
	if err == nil { //节点配置异常的情况
		err = ErrNode
	}
	return
}

//执行业务逻辑处理情况
func (s ElementSt) runParse(arrTemps []string, p IFNodeParser) (result interface{}, err error) {
	arrStr := make([]string, 0) //设置存储组合结果
	for _, tempStr := range arrTemps {
		switch s.Type {
		case ELEMENT_TYPE_TEXT:
			result, err = p.InnerText(tempStr)
		case ELEMENT_TYPE_LIST:
			result, err = p.InnerTexts(tempStr)
			if err == nil {//数组类型的不允许再做组合
				return
			}
		default:
			err = ErrType
			return
		}
		if err == nil && result != nil {//得到匹配结果的情况
			arrStr = append(arrStr, convertString(result))
		}
	}
	if len(arrStr) > 0 { //如果找到的组合的情况逻辑
		result = arrStr
	}
	return
}



