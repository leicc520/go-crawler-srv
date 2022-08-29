package parse

/*********************************************************************
 	配置模板参数数据资料信息，只有user-agent是随机的，其他走配置
 */
type HeaderSt struct{
	TaskName string   `json:"task_name"`
	Headers  []struct {
		Key string    `json:"key"`
		Value string  `json:"value"`
	} `json:"headers"`
}

//转换成map数据信息存在起来
func (s HeaderSt) ASMap() map[string]string {
	data := map[string]string{}
	for _, item := range s.Headers {
		data[item.Key] = item.Value
	}
	return data
}

//配置请求头数据资料信息
type HttpHeaderSt struct {
	Headers []HeaderSt `json:"headers"`
}

//配置模板数据资料信息
type TemplateSt struct {
	Templates []struct{
		TaskName string 	  `json:"task_name"`  //通过任务名称定位模板
		Elements []ElementSt  `json:"elements"`   //使用的模板提取元素
	}  `json:"templates"`
}


//获取模板名称数据资料信息
func (s TemplateSt) getTemplate(taskName string) []ElementSt {
	if s.Templates == nil || len(s.Templates) < 1 {
		return nil
	}
	for _, temp := range s.Templates {
		if temp.TaskName == taskName {
			return temp.Elements
		}
	}
	return nil
}