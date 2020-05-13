package database

import (
	"github.com/jinzhu/gorm"
	"github.com/king-tu/mallweb/common/conf"
	"go.uber.org/zap"
	"time"
)

func InitDB() *gorm.DB {
	//	连接mysql
	connStr := conf.MySQL_USER + ":" + conf.MySQL_PASSWD
	connStr = connStr + "@tcp(" + conf.MySQL_ADDR + ")/" + conf.MySQL_DBNAME
	connStr = connStr + "?charset=utf8&parseTime=True&loc=Local"

	zap.L().Debug("Connect to the database", zap.String("connStr", connStr))
	db, err := gorm.Open("mysql", connStr)
	if err != nil {
		panic(err)
	}

	//db.DB()
	db.LogMode(true)
	err = db.DB().Ping()
	if err != nil {
		panic(err)
	}

	db.DB().SetMaxIdleConns(conf.MySQL_MaxIdleConns)
	db.DB().SetMaxOpenConns(conf.MySQL_MaxOpenConns)
	db.DB().SetConnMaxLifetime(time.Second * conf.MySQL_ConnMaxLifetime)
	// 表名使用单数（默认复数）
	db.SingularTable(true)

	return db
}
