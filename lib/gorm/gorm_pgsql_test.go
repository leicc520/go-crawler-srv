package gorm

import (
	"fmt"
	"github.com/leicc520/go-crawler-srv/lib"
	"github.com/leicc520/go-crawler-srv/service/dbsrv/spider/model"
	"github.com/leicc520/go-orm"
	"testing"
	"time"
)

//type SysActivationCode struct {
//	Id           uint `gorm:"primaryKey;autoIncrement"`
//	Status       bool
//	Code         string
//	ActivateTime time.Time
//	CreateTime   time.Time
//	ExpireTime   time.Time
//}

func Test(t *testing.T) {
	getData()
}

func getData() {
	DbConfig := orm.DbConfig{
		Driver:       "postgres",
		Host:         "host=10.100.72.102 port=5432 user=postgres password=postgres dbname=spider sslmode=disable",
		SKey:         "",
		MaxOpenConns: 128,
		MaxIdleConns: 24,
	}

	sysConfig := lib.ConfigSt{
		SpiderDbMaster: DbConfig,
	}
	InitPostgresDbPool(sysConfig)
	db := GetDB()
	var ac *model.SysActivationCode
	resutl := db.Where(model.SysActivationCode{
		Code: "454545-7890--7890",
	}).First(&ac)
	fmt.Println(ac)
	fmt.Println(resutl.RowsAffected)
	expireTime, err := time.Parse("2006-01-02", "2022-12-01")
	if err != nil {

	}
	sysActivationCode := model.SysActivationCode{Code: "4545-7890--7890-7899", Status: false, CreateTime: time.Now(), ExpireTime: &expireTime}
	result := db.Create(&sysActivationCode)
	fmt.Println(sysActivationCode.Id)
	fmt.Println(result.RowsAffected)

	//user := User{Name: "Jinzhu", Age: 18, Birthday: time.Now()}
	//
	//result := db.Create(&user) // 通过数据的指针来创建
}
