package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)


//请求提交一个url地址数据资料信息
func Router(c *gin.Engine) {
	c.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Simple Life For Spider！")
	})
	state := c.Group("_state")
	state.GET("proxy", proxyState)
	api := c.Group("api")
	api.POST("/task/create", taskCreate)
}
