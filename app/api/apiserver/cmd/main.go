package main

import (
	"github.com/king-tu/mallweb/app/api/apiserver/handler"
	"github.com/king-tu/mallweb/app/api/apiserver/router"
	"github.com/king-tu/mallweb/app/log"

	"github.com/king-tu/mallweb/common/conf"
	"github.com/king-tu/mallweb/common/utils"
	"go.uber.org/zap"
)

func init() {
	// 初始化日志库
	log.SetLogs(zap.DebugLevel, log.LOGFORMAT_CONSOLE, conf.LogFileName)
}

func main() {
	grpcService := utils.GetGRPCService()

	// Initialise service
	grpcService.Init()

	handler.RegisterClients(grpcService.Client())

	r := router.InitRouter()
	//启动监听
	if err := r.Run(":8888"); err != nil {
		zap.L().Fatal("HTTP Server启动失败", zap.Error(err))
	}
}
