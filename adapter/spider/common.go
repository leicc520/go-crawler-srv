package spider


//定义队列中，接收的任务数据信息
type TaskSt struct {
	req  	 BaseRequest `json:"req"`
	taskId   string  	 `json:"task_id"`
	template string  	 `json:"template"`
}

func (s TaskSt) Run() error {
	return nil
}



