package models

import (
	"reflect"
	"time"
	"github.com/leicc520/go-orm"
)

type SpAmazonSeller struct {
	*orm.ModelSt
}

//结构体实例的结构说明
type SpAmazonSellerSt struct {
	Id		int64		`json:"id"`		
	SellerId		string		`json:"seller_id"`		
	Json		string		`json:"json"`		
	UpdatedAt		time.Time		`json:"updated_at"`		
	CreatedAt		time.Time		`json:"created_at"`		
	Version		string		`json:"version"`		
}

//这里默认引用全局的连接池句柄
func NewSpAmazonSeller() *SpAmazonSeller {
	fields := map[string]reflect.Kind{
		"id":		reflect.Int64,		//记录ID
		"seller_id":		reflect.String,		//卖家ID
		"json":		reflect.String,		//抓取的json
		"updated_at":		orm.DT_TIMESTAMP,		//更新时间
		"created_at":		orm.DT_TIMESTAMP,		//创建时间
		"version":		reflect.String,		//解析器版本
	}
	
	args  := map[string]interface{}{
		"table":		"sp_amazon_seller",
		"orgtable":		"sp_amazon_seller",
		"prikey":		"id",
		"dbmaster":		"spiderdbmaster",
		"dbslaver":		"spiderdbslaver",
		"slot":			0,
	}

	data := &SpAmazonSeller{&orm.ModelSt{}}
	data.Init(&orm.GdbPoolSt, args, fields)
	return data
}
