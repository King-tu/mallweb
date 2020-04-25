package main

import (
	"github.com/king-tu/mallweb/app/api/webserver/router"
	"github.com/king-tu/mallweb/app/log"
	"go.uber.org/zap"
)

func init() {
	// 初始化日志库
	log.SetLogs(zap.DebugLevel, log.LOGFORMAT_CONSOLE, "../tmp/log/mallweb.log")
}
func main() {
	r := router.InitRoutes()

	//启动监听
	if err := r.Run(); err != nil {
		zap.L().Fatal("HTTP Server启动失败", zap.Error(err))
	}
}
