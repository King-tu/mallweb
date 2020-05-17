package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/king-tu/mallweb/common/conf"
	"go.uber.org/zap"
	"sync"
	"time"
)

var db *gorm.DB
var once sync.Once

func DB() *gorm.DB {
	var err error
	once.Do(func() {
		//	连接mysql
		connStr := conf.MySQL_USER + ":" + conf.MySQL_PASSWD
		connStr = connStr + "@tcp(" + conf.MySQL_ADDR + ")/" + conf.MySQL_DBNAME
		connStr = connStr + "?charset=utf8&parseTime=True&loc=Local"

		zap.L().Debug("Connect to the database", zap.String("connStr", connStr))
		db, err = gorm.Open("mysql", connStr)
		if err != nil {
			zap.L().Sugar().Fatalf("models.Setup err: %v", err)
		}

		db.DB().SetMaxIdleConns(conf.MySQL_MaxIdleConns)
		db.DB().SetMaxOpenConns(conf.MySQL_MaxOpenConns)
		db.DB().SetConnMaxLifetime(time.Second * conf.MySQL_ConnMaxLifetime)
		// 表名使用单数（默认复数）
		db.SingularTable(true)
	})

	return db
}
