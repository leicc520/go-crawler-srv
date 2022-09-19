package lib

import (
	"fmt"
	"github.com/leicc520/go-orm"
	"os"
	"path/filepath"

	"github.com/leicc520/go-crawler-srv/lib/proxy"
	"github.com/leicc520/go-crawler-srv/lib/queue"
	"github.com/leicc520/go-gin-http"
	"github.com/leicc520/go-gin-http/micro"
	"github.com/leicc520/go-orm/log"
)

const (
	NAME    = "go-crawler-srv"
	VERSION = "v1.0.0"
)

type ConfigSt struct {
	App            core.AppConfigSt `yaml:"app"`
	Logger         log.LogFileSt    `yaml:"logger"`
	Redis          string           `yaml:"redis"`
	AQMQueue       queue.RabbitMqSt `yaml:"aqm_queue"`
	SpiderDbMaster orm.DbConfig     `yaml:"spider_db_master"`
	SpiderDbSlaver orm.DbConfig     `yaml:"spider_db_slaver"`
	ChromeDp       struct {
		ReTry      int      `yaml:"retry"`
		HeadLess   bool     `yaml:"head_less"`
		ProxyUrl   string   `yaml:"proxy_url"`
		DevtoolsWs []string `yaml:"devtools_ws"`
	} `yaml:"chrome_dp"`
	HttpProxy     []proxy.ProxySt `yaml:"http_proxy"`
	SpiderService struct {
		JobConCurrency      int `yaml:"job_con_currency"`
		ChromeDpConCurrency int `yaml:"chrome_dp_con_currency"`
		SpiderConCurrency   int `yaml:"spider_con_currency"`
	} `yaml:"spider_service"`
}

var Config *ConfigSt = nil

// 初始化加载配置信息
func InitConfig() *ConfigSt {
	Config = &ConfigSt{}
	LoadConfigByName(NAME, Config)
	Config.App.Name = NAME
	Config.App.Version = VERSION
	InitRedis(Config.Redis) //初始化redis连接池
	amazonConfigLoad("config/amazon.yml")
	log.SetLogger(Config.Logger.Init())
	orm.InitDBPoolSt().LoadDbConfig(Config) //配置数据库结构注册到数据库调用配置当中
	return Config
}

// 通过名称加载配置数据资料信息
func LoadConfigByName(name string, config interface{}) {
	workDir, _ := os.Getwd()
	filePath := fmt.Sprintf("config/%s-%s.yml", name, os.Getenv(core.DCENV))
	filePath = filepath.Join(workDir, filePath)
	if _, err := micro.LoadFile(filePath, config); err != nil {
		panic("配置加载异常:" + filePath)
	}
}
