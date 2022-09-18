package api

import (
	"github.com/gin-gonic/gin"
	"github.com/leicc520/go-crawler-srv/lib"
	"github.com/leicc520/go-crawler-srv/lib/proxy"
	"net/http"
)

func PreLoader() {
	//完成代理的初始化业务逻辑
	proxy.Init(lib.Config.HttpProxy, lib.Redis)
}

// 释放数据资料处理逻辑
func Release() {
	if lib.Redis != nil {
		lib.Redis.Close()
		lib.Redis = nil
	}
}

// 请求提交一个url地址数据资料信息
func Router(c *gin.Engine) {
	c.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Simple Life For Spider！")
	})
	state := c.Group("_state")
	state.GET("proxy", proxyState)
	api := c.Group("api")
	api.POST("/task/create", taskCreate)
	api.POST("/activation/active", activationCode)
}
