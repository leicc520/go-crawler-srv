package wipo

import (
	"github.com/leicc520/go-orm"
	"reflect"
)

const (
	DBSRV = "host=10.100.72.102 port=5432 user=postgres password=postgres dbname=zbox sslmode=disable"
)

//初始化数据库信息
func init() {
	master := orm.DbConfig{"postgres", DBSRV, "dbmaster", 128, 64}
	slaver := orm.DbConfig{"postgres", DBSRV, "dbslaver", 128, 64}
	orm.InitDBPoolSt().Set(master.SKey, &master)
	orm.InitDBPoolSt().Set(slaver.SKey, &slaver)
}

type apWipoBrand struct {
	*orm.ModelSt
}

//这里的dbPool
func newApWipoBrand() *apWipoBrand {
	fields := map[string]reflect.Kind{
		"id": reflect.Int64,	//账号id
	}
	args  := map[string]interface{}{
		"table":		"ap_wipo_brand",
		"orgtable":		"ap_wipo_brand",
		"prikey":		"id",
		"dbmaster":		"dbmaster",
		"dbslaver":		"dbslaver",
		"slot":			0,
	}
	data := &apWipoBrand{&orm.ModelSt{}}
	data.Init(&orm.GdbPoolSt, args, fields)
	return data
}



