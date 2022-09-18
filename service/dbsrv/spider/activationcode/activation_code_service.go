package activationcode

import (
	"github.com/leicc520/go-crawler-srv/lib/gorm"
	"github.com/leicc520/go-crawler-srv/service/dbsrv/spider/model"
	"github.com/leicc520/go-orm/log"
	"time"
)

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
func UpdateStatusByCode(code string) bool {
	sysActivationCode := GetActivationCodeByCode(code)
	if sysActivationCode == nil {
		log.Write(log.ERROR, "未找到记录, code:"+code)
		return false
	}
	db := gorm.GetDB()
	now := time.Now()
	result := db.Model(sysActivationCode).Updates(model.SysActivationCode{ActivateTime: &now, Status: true})
	return result.RowsAffected > 1
}
