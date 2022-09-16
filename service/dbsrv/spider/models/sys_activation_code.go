package models

import (
	"reflect"
	"github.com/leicc520/go-orm"
)

type SysActivationCode struct {
	*orm.ModelSt
}

//结构体实例的结构说明
type SysActivationCodeSt struct {
	Id		int		`json:"id"`		
	Status		string		`json:"status"`		
	Code		string		`json:"code"`		
	ActivateTime		string		`json:"activate_time"`		
	CreateTime		string		`json:"create_time"`		
	ExpireTime		string		`json:"expire_time"`		
}

//这里默认引用全局的连接池句柄
func NewSysActivationCode() *SysActivationCode {
	fields := map[string]reflect.Kind{
		"id":		reflect.Int,		//
		"status":		reflect.String,		//是否激活 默认未激活
		"code":		reflect.String,		//激活码
		"activate_time":		reflect.String,		//激活时间
		"create_time":		reflect.String,		//激活码生成时间
		"expire_time":		reflect.String,		//过期时间
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
