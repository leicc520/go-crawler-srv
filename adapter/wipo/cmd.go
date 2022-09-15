package wipo

import (
	"github.com/leicc520/go-crawler-srv/adapter"
	"github.com/leicc520/go-crawler-srv/lib"
	"github.com/leicc520/go-crawler-srv/lib/proxy"
	"github.com/leicc520/go-crawler-srv/plugins"
	"github.com/leicc520/go-orm"
	"github.com/leicc520/go-orm/log"
)

const (
	Name = "go-wipo-cmd"
	CMD  = "wipo"
)

type goWipoCmd struct {
}

//定义cmd配置管理处理逻辑
type configSt struct {
	State struct{
		StartDate string `yaml:"start_date"`
		EndDate   string `yaml:"end_date"`
		Country   string `yaml:"country"`
	} `yaml:"state"`
	Country  string 			`yaml:"country"`
	ChromeDp plugins.ChromeDpSt	`yaml:"chrome_dp"`
	HttpProxy []proxy.ProxySt   `yaml:"http_proxy"`
	DbMaster orm.DbConfig 		`yaml:"wipo_db_master"`
	DbSlaver orm.DbConfig 		`yaml:"wipo_db_slaver"`
}

//定义变量处理逻辑
var wState *configSt = nil

func init() {
	adapter.CommandRegister(CMD, &goWipoCmd{})
}

//执行命令行业务逻辑
func (s *goWipoCmd) Run() {
	if err := recover(); err != nil {
		log.Write(log.ERROR, err)
	}
	log.Write(-1, "开启执行命令行:"+CMD)
	wState = &configSt{}
	lib.LoadConfigByName(Name, wState)
	proxy.Init(wState.HttpProxy, lib.Redis)
	ss := &WipoSt{dpc:&wState.ChromeDp, Country: wState.State.Country}
	orm.InitDBPoolSt().Set("wipo_db_master", &wState.DbMaster)
	orm.InitDBPoolSt().Set("wipo_db_slaver", &wState.DbSlaver)
	ss.Run(wState.State.StartDate, wState.State.EndDate)
	log.Write(-1, "结束执行命令行:"+CMD)
}