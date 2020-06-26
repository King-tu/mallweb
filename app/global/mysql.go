package global

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/king-tu/mallweb/app/model"
	"go.uber.org/zap"
	"time"
)

func InitMysql() {
	var err error
	//	连接mysql
	connStr := Config.Mysql.Username + ":" + Config.Mysql.Password
	connStr = connStr + "@tcp(" + Config.Mysql.Addr + ")/" + Config.Mysql.Dbname
	connStr = connStr + "?charset=utf8&parseTime=True&loc=Local"

	Logger.Bg().Debug("Connect to the database", zap.String("connStr", connStr))
	DB, err = gorm.Open("mysql", connStr)
	if err != nil {
		Logger.Bg().Fatal("models.Setup err", zap.Error(err))
	}

	DB.DB().SetMaxIdleConns(Config.Mysql.MaxIdleConns)
	DB.DB().SetMaxOpenConns(Config.Mysql.MaxOpenConns)
	DB.DB().SetConnMaxLifetime(time.Second * time.Duration(Config.Mysql.MaxConnLifetime))
	// 表名使用单数（默认复数）
	DB.SingularTable(true)
	//迁移表
	MigrateTables(DB)
}

func MigrateTables(db *gorm.DB) {
	// 自动迁移(创建)表
	db.AutoMigrate(new(model.User), new(model.Address))
	db.AutoMigrate(new(model.Goods), new(model.GoodsType), new(model.GoodsSKU), new(model.GoodsImage), new(model.IndexGoodsBanner), new(model.IndexTypeGoodsBanner), new(model.IndexPromotionBanner), new(model.TpshopCategory))
	db.AutoMigrate(new(model.OrderInfo), new(model.OrderGoods))
}
