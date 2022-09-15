package adapter

import (
	"flag"
	"strings"
	"sync"
)

type IFCommand interface {
	Run()
}

var (
	commandOnce  = sync.Once{}
	commandStr   = "all"
	command map[string]IFCommand = nil
)

//注册命令行执行业务处理逻辑
func CommandRegister(name string, cmd IFCommand)  {
	if command == nil {
		commandOnce.Do(func() {
			command = make(map[string]IFCommand)
		})
	}
	command[name] = cmd
}

//注册要获取的命令行逻辑
func CommandCmd() {
	arrStr := make([]string, 0)
	for key, _ := range command {
		arrStr  = append(arrStr, key)
	}
	usage := "请输入要执行的命令行["+strings.Join(arrStr, ",")+"],all-开启所有任务..."
	flag.StringVar(&commandStr,"cmd", "all", usage)
}

//执行业务逻辑
func Run() bool {
	if cmd, ok := command[commandStr]; ok {
		cmd.Run() //执行业务处理逻辑
		return false
	}
	if strings.ToLower(commandStr) != "all" {
		panic("未找到要执行的命令行："+commandStr)
	}
	//如果没有设置的话开启做个业务
	for _, cmd := range command {
		go cmd.Run()
	}
	return true
}


