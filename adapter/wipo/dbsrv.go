package wipo

import (
	"github.com/leicc520/go-orm"
	"reflect"
)

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
		"dbmaster":		"wipo_db_master",
		"dbslaver":		"wipo_db_slaver",
		"slot":			0,
	}
	data := &apWipoBrand{&orm.ModelSt{}}
	data.Init(&orm.GdbPoolSt, args, fields)
	return data
}

type apWipoResult struct {
	*orm.ModelSt
}

//这里的dbPool
func newApWipoResult() *apWipoResult {
	fields := map[string]reflect.Kind{
		"id": reflect.Int64,	//账号id
	}
	args  := map[string]interface{}{
		"table":		"ap_wipo_result",
		"orgtable":		"ap_wipo_result",
		"prikey":		"id",
		"dbmaster":		"wipo_db_master",
		"dbslaver":		"wipo_db_slaver",
		"slot":			0,
	}
	data := &apWipoResult{&orm.ModelSt{}}
	data.Init(&orm.GdbPoolSt, args, fields)
	return data
}



