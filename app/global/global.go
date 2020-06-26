package global

import (
	"github.com/jinzhu/gorm"
	"github.com/king-tu/mallweb/app/common"
	"github.com/king-tu/mallweb/app/config"
	logger "github.com/king-tu/mallweb/app/log"
)

var (
	Config *config.App
	Logger logger.Factory
	DB     *gorm.DB
)

func init() {
	// 加载配置文件
	LoadConfig()

	// 初始化日志库
	SetLogs(Config.Log.Level, Config.Log.Format, Config.Log.FileName, common.PROJECT_NAME)

	// 初始化数据库连接池
	InitMysql()
}
