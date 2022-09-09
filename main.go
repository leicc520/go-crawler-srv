package main

import (
	"github.com/leicc520/go-crawler-srv/api"
	"github.com/leicc520/go-crawler-srv/lib"
	"github.com/leicc520/go-gin-http"
	"github.com/leicc520/go-gin-http/micro"
)

func main() {
	micro.CmdInit(func() {
		lib.InitConfig()
	}) //初始化配置
	defer api.Release() //资源释放处理逻辑
	core.NewApp(&lib.Config.App).RegHandler(api.Router).Start()
}
