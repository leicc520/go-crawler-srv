package models

import (
	"github.com/leicc520/go-orm"
	"reflect"
	"time"
)

type SysActivationCode struct {
	*orm.ModelSt
}

//结构体实例的结构说明
type SysActivationCodeSt struct {
	Code		string		`json:"code"`		
	ActivateTime		time.Time		`json:"activate_time"`		
	CreateTime		time.Time		`json:"create_time"`		
	ExpireTime		time.Time		`json:"expire_time"`		
	Id		int64		`json:"id"`		
	Status		int		`json:"status"`		
}

//这里默认引用全局的连接池句柄
func NewSysActivationCode() *SysActivationCode {
	fields := map[string]reflect.Kind{
		"code":		reflect.String,		//激活码
		"activate_time":		orm.DT_TIMESTAMP,		//激活时间
		"create_time":		orm.DT_TIMESTAMP,		//激活码生成时间
		"expire_time":		orm.DT_TIMESTAMP,		//过期时间
		"id":		reflect.Int64,		//记录ID
		"status":		reflect.Int,		//状态 0-待激活 1-已激活
	}
	
	args  := map[string]interface{}{
		"table":		"sys_activation_code",
		"orgtable":		"sys_activation_code",
		"prikey":		"id",
		"dbmaster":		"spiderdbmaster",
		"dbslaver":		"spiderdbslaver",
		"slot":			0,
	}

	data := &SysActivationCode{&orm.ModelSt{}}
	data.Init(&orm.GdbPoolSt, args, fields)
	return data
}
