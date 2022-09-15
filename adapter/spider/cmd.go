package spider

import (
	"github.com/leicc520/go-crawler-srv/adapter"
	"github.com/leicc520/go-orm/log"
)

const (
	Name = "go-spider-cmd"
	CMD  = "spider"
)

type goSpiderCmd struct {
}

func init() {
	adapter.CommandRegister(CMD, &goSpiderCmd{})
}

//执行命令行业务逻辑
func (s *goSpiderCmd) Run() {
	if err := recover(); err != nil {
		log.Write(log.ERROR, err)
	}
	log.Write(-1, "开启执行命令行:"+CMD)
	//todo 更新数据信息
	log.Write(-1, "结束执行命令行:"+CMD)
}