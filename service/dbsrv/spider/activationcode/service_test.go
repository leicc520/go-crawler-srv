package activationcode

import (
	"github.com/leicc520/go-crawler-srv/lib"
	"github.com/leicc520/go-crawler-srv/lib/gorm"
	"github.com/leicc520/go-orm"
	"testing"
)

func Test(t *testing.T) {
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
	gorm.InitPostgresDbPool(sysConfig)

	UpdateStatusByCode("4545-7890--7890-7899")
}
