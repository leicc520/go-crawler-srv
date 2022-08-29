package parse

import (
	"github.com/leicc520/go-orm"
	"github.com/leicc520/go-orm/log"
)

//元素节点提取配置，便利模板节点直到知道数据才结束
type ElementSt struct {
	Tag   string 		 `json:"tag"` 		//提取之后放到这个名字的map当中
	Nodes []TempNodeSt 	 `json:"nodes"`     //节点解析的模板配置
	Elements []ElementSt `json:"elements"`  //允许递归的获取元素，在当前解析节点继续解析
}

//执行获取匹配的结果数据处理逻辑
func (s ElementSt) getValue(t *CompilerSt) (value interface{}, err error) {
	var p InNodeParser
	for _, node := range s.Nodes {
		p = t.getParser(node.NodeType)
		value, err = node.RunParse(p)
		if err != nil {
			continue
		}
		//匹配结果打印日志
		log.Write(log.DEBUG, node.Temp, value)
	}
	if err != nil {//节点取值匹配失败的情况
		log.Write(log.ERROR, "get element value error ", err, t.GetDoc())
	}
	return
}

//执行业务逻辑解析处理逻辑
func (s ElementSt) RunParse(t *CompilerSt, result orm.SqlMap) error {
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
			t     = &CompilerSt{doc:doc}
			item := orm.SqlMap{}
			for _, el := range s.Elements {
				err = el.RunParse(t, item)
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
