package activationcode

import (
	"fmt"
	"github.com/leicc520/go-crawler-srv/lib/gorm"
	"github.com/leicc520/go-crawler-srv/service/dbsrv/spider/model"
	"github.com/leicc520/go-orm/log"
	"time"
)

type Response struct {
	Status bool
	Msg    string
}

func GetActivationCodeByCode(code string) *model.SysActivationCode {
	db := gorm.GetDB()
	var ac *model.SysActivationCode
	result := db.Where(model.SysActivationCode{
		Code:   code,
		Status: false,
	}).Find(&ac)
	if result.RowsAffected == 0 {
		return nil
	}
	return ac
}

// UpdateStatusByCode 更新激活状态
func UpdateStatusByCode(code string) Response {
	sysActivationCode := GetActivationCodeByCode(code)
	if sysActivationCode == nil {
		msg := fmt.Sprintf("未找到记录, code: %s", code)
		log.Write(log.ERROR, msg)
		return Response{
			Status: false,
			Msg:    msg,
		}
	}
	if sysActivationCode.Status {
		msg := fmt.Sprintf("code: %s, 该激活码已经被激活", code)
		log.Write(log.ERROR, msg)
		return Response{
			Status: false,
			Msg:    msg,
		}
	}
	db := gorm.GetDB()
	now := time.Now()
	result := db.Model(sysActivationCode).Updates(model.SysActivationCode{ActivateTime: &now, Status: true})
	return Response{
		Status: result.RowsAffected > 0,
	}
}
