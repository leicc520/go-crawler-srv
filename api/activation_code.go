package api

import (
	"github.com/gin-gonic/gin"
	"github.com/leicc520/go-crawler-srv/service/dbsrv/spider/activationcode"
	"net/http"
)

type Activation struct {
	Code string `json:"code"`
}

func activationCode(c *gin.Context) {
	var activation Activation
	// 将request的body中的数据，自动按照json格式解析到结构体
	if err := c.ShouldBindJSON(&activation); err != nil {
		// 返回错误信息
		// gin.H封装了生成json数据的工具
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if activation.Code == "" {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "序列码不能为空"})
		return
	}
	result := activationcode.UpdateStatusByCode(activation.Code)
	if !result.Status {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": result.Msg})
	}
	c.Writer.WriteHeader(http.StatusOK)
}
