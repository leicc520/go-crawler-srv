package main

import (
	"github.com/leicc520/go-crawler-srv/adapter"
	_ "github.com/leicc520/go-crawler-srv/adapter/wipo"
	"github.com/leicc520/go-crawler-srv/api"
	"github.com/leicc520/go-crawler-srv/lib"
	"github.com/leicc520/go-crawler-srv/lib/gorm"
	core "github.com/leicc520/go-gin-http"
	"github.com/leicc520/go-gin-http/micro"
	_ "github.com/lib/pq"
)

func main() {
	micro.PushCmd(adapter.CommandCmd)
	micro.CmdInit(func() {
		Config := lib.InitConfig()
		gorm.InitPostgresDbPool(*Config)
	}) //初始化配置
	defer api.Release() //资源释放处理逻辑
	if adapter.Run() {  //执行业务逻辑
		core.NewApp(&lib.Config.App).RegHandler(api.Router).Start()
	}
}
