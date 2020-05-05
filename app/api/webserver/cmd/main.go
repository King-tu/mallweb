package main

import (
	"github.com/king-tu/mallweb/app/api/webserver/router"
	"github.com/king-tu/mallweb/app/log"
	"go.uber.org/zap"
)

func init() {
	// 初始化日志库
	logfile := "/tmp/log/mallweb/mallweb.log"
	log.SetLogs(zap.DebugLevel, log.LOGFORMAT_CONSOLE, logfile)
}

func main() {
	r := router.InitRoutes()

	//启动监听
	if err := r.Run(":8888"); err != nil {
		zap.L().Fatal("HTTP Server启动失败", zap.Error(err))
	}
}
