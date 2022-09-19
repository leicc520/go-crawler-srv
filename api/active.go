package api

import (
	"github.com/gin-gonic/gin"
	"github.com/leicc520/go-crawler-srv/service/dbsrv/spider/models"
	core "github.com/leicc520/go-gin-http"
	"github.com/leicc520/go-orm"
	"time"
)

//获取激活码的处理逻辑
func activeCode(c *gin.Context) {
	args := struct {
		Code string `form:"code" json:"code" binding:"required,min=1"`
	}{}
	if err := c.ShouldBind(&args); err != nil {
		core.PanicValidateHttpError(1001, err)
	}
	sorm   := models.NewSysActivationCode()
	state  := models.SysActivationCodeSt{}
	if err := sorm.GetItem(func(st *orm.QuerySt) string {
		st.Where("code", args.Code)
		return st.GetWheres()
	}, "*").ToStruct(&state); err != nil {
		core.PanicHttpError(1002, "激活码不存在,无法激活")
	}
	if state.Status != 0 || time.Now().Before(state.ExpireTime) {
		core.PanicHttpError(1003, "激活码已被激活或过期,无法使用")
	}
	sorm.Save(state.Id, orm.SqlMap{"status":1, "activate_time":time.Now()})
	core.NewHttpView(c).JsonDisplay(nil)
}
