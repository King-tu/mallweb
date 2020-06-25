package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/king-tu/mallweb/app/global"
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
		connStr := global.Config.Mysql.Username + ":" + global.Config.Mysql.Password
		connStr = connStr + "@tcp(" + global.Config.Mysql.Addr + ")/" + global.Config.Mysql.Dbname
		connStr = connStr + "?charset=utf8&parseTime=True&loc=Local"

		global.Logger.Bg().Debug("Connect to the database", zap.String("connStr", connStr))
		db, err = gorm.Open("mysql", connStr)
		if err != nil {
			global.Logger.Bg().Fatal("models.Setup err", zap.Error(err))
		}

		db.DB().SetMaxIdleConns(global.Config.Mysql.MaxIdleConns)
		db.DB().SetMaxOpenConns(global.Config.Mysql.MaxOpenConns)
		db.DB().SetConnMaxLifetime(time.Second * time.Duration(global.Config.Mysql.MaxConnLifetime))
		// 表名使用单数（默认复数）
		db.SingularTable(true)

		// 自动迁移(创建)表
		MigrateTables()
	})

	return db
}

func MigrateTables() {
	// 自动迁移(创建)表
	db.AutoMigrate(new(User), new(Address))
	db.AutoMigrate(new(Goods), new(GoodsType), new(GoodsSKU), new(GoodsImage), new(IndexGoodsBanner), new(IndexTypeGoodsBanner), new(IndexPromotionBanner), new(TpshopCategory))
	db.AutoMigrate(new(OrderInfo), new(OrderGoods))
}
