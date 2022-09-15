package lib

import (
	"fmt"
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
	App      core.AppConfigSt	   	`yaml:"app"`
	Logger   log.LogFileSt	   		`yaml:"logger"`
	Redis    string                 `yaml:"redis"`
	AQMQueue queue.RabbitMqSt       `yaml:"aqm_queue"`
	ChromeDp struct {
		ReTry     int `json:"retry" yaml:"retry"`
		HeadLess bool `json:"head_less" yaml:"head_less"`
		ProxyUrl string `json:"proxy_url" yaml:"proxy_url"`
		DevtoolsWs []string `json:"devtools_ws" yaml:"devtools_ws"`
	}	  	 						`yaml:"chrome_dp"`
	HttpProxy []proxy.ProxySt   	`yaml:"http_proxy"`
	SpiderService struct{
		JobConCurrency      	int     `yaml:"job_con_currency"`
		ChromeDpConCurrency 	int     `yaml:"chrome_dp_con_currency"`
		SpiderConCurrency   	int     `yaml:"spider_con_currency"`
	}  `yaml:"spider_service"`
	ParserService struct{
		JobConCurrency      	int     `yaml:"job_con_currency"`
		ParserConCurrency   	int     `yaml:"parser_con_currency"`
	}  `yaml:"parser_service"`
	DispatcherService struct{
		JobConCurrency      	int     `yaml:"job_con_currency"`
		DispatcherConCurrency   int     `yaml:"dispatcher_con_currency"`
	}  `yaml:"dispatcher_service"`
}

var Config *ConfigSt = nil

//初始化加载配置信息
func InitConfig() *ConfigSt {
	Config    = &ConfigSt{}
	LoadConfigByName(NAME, Config)
	Config.App.Name    = NAME
	Config.App.Version = VERSION
	InitRedis(Config.Redis) //初始化redis连接池
	amazonConfigLoad("config/amazon.yml")
	log.SetLogger(Config.Logger.Init())
	return Config
}

//通过名称加载配置数据资料信息
func LoadConfigByName(name string, config interface{}) {
	workDir, _ := os.Getwd()
	filePath := fmt.Sprintf("config/%s-%s.yml", name, os.Getenv(core.DCENV))
	filePath  = filepath.Join(workDir, filePath)
	if _, err := micro.LoadFile(filePath, config); err != nil {
		panic("配置加载异常:"+filePath)
	}
}