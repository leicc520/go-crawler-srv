package models

import (
	"github.com/leicc520/go-orm"
	"reflect"
)

type SpAmazonAsin struct {
	*orm.ModelSt
}

//结构体实例的结构说明
type SpAmazonAsinSt struct {
	Id		int64		`json:"id"`		
	Asin		string		`json:"asin"`		
	Version		string		`json:"version"`		
	Json		string		`json:"json"`		
	UpdatedAt		string		`json:"updated_at"`		
	CreatedAt		string		`json:"created_at"`		
}

//这里默认引用全局的连接池句柄
func NewSpAmazonAsin() *SpAmazonAsin {
	fields := map[string]reflect.Kind{
		"id":		reflect.Int64,		//记录ID
		"asin":		reflect.String,		//唯一编号
		"version":		reflect.String,		//解析器版本
		"json":		reflect.String,		//采集数据
		"updated_at":		reflect.String,		//更新时间
		"created_at":		reflect.String,		//创建时间
	}
	
	args  := map[string]interface{}{
		"table":		"sp_amazon_asin",
		"orgtable":		"sp_amazon_asin",
		"prikey":		"id",
		"dbmaster":		"spiderdbmaster",
		"dbslaver":		"spiderdbslaver",
		"slot":			0,
	}

	data := &SpAmazonAsin{&orm.ModelSt{}}
	data.Init(&orm.GdbPoolSt, args, fields)
	return data
}
