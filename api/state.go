package api

import (
	"github.com/gin-gonic/gin"
	"github.com/leicc520/go-crawler-srv/lib"
	"github.com/leicc520/go-crawler-srv/lib/proxy"
	"github.com/leicc520/go-gin-http"
	"os"
)

//获取代理id统计数据信息格式化输出处理逻辑
func proxyState(c *gin.Context)  {
	if c.Query("_s") == os.Getenv(core.DCJWT) {
		c.AbortWithStatus(403)
		return
	}
	//为开启代理监控
	ss := proxy.GetStatistic()
	if ss == nil || len(lib.Config.HttpProxy) < 1 {
		core.PanicHttpError(401, "未开启代理监控.")
	}
	arrState := make([]string, 0)
	for _, item := range lib.Config.HttpProxy {
		arrState = append(arrState, ss.ItemNotify(item.Proxy))
	}
	//格式化输出处理逻辑
	c.Writer.WriteHeader(200)
	c.Writer.WriteString(lib.PrettyJson(arrState))
}
