package gorm

import (
	"github.com/leicc520/go-crawler-srv/lib"
	"github.com/leicc520/go-orm/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// db连接
var db *gorm.DB
var Config *lib.ConfigSt = nil

func InitPostgresDbPool(config lib.ConfigSt) {
	dsn := config.SpiderDbMaster.Host
	dbConn, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		log.Write(log.ERROR, "数据库:"+config.SpiderDbMaster.Host+" 连接 失败:"+err.Error())
	}
	sqlDB, _ := dbConn.DB()

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(config.SpiderDbMaster.MaxIdleConns)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(config.SpiderDbMaster.MaxOpenConns)
	db = dbConn
}

func GetDB() *gorm.DB {
	sqlDB, err := db.DB()
	if err != nil {
		log.Write(log.ERROR, "数据库连接失败:"+err.Error())
		reConnectionDB()
	}
	if err := sqlDB.Ping(); err != nil {
		sqlDB.Close()
		reConnectionDB()
	}
	return db
}

func reConnectionDB() {
	Config = &lib.ConfigSt{}
	lib.LoadConfigByName(lib.NAME, Config)
	InitPostgresDbPool(*Config)
}
