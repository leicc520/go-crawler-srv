package parse

//元素节点提取配置，便利模板节点直到知道数据才结束
type ElementSt struct {
	Tag   string 		 `json:"tag"` 		//提取之后放到这个名字的map当中
	Nodes []TempNodeSt 	 `json:"nodes"`     //节点解析的模板配置
	Elements []ElementSt `json:"elements"`  //允许递归的获取元素，在当前解析节点继续解析
}
