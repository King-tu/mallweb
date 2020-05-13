package main

import (
	"github.com/king-tu/mallweb/app/log"
	"github.com/king-tu/mallweb/app/services/sms/handler"
	"github.com/king-tu/mallweb/app/services/sms/proto/smscode"
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

	// Register Handler
	if err := smscode.RegisterSmsCodeServiceHandler(grpcService.Server(), handler.NewSmsCodeService()); err != nil {
		zap.L().Error("Fail to run  register SmsCodeService", zap.Error(err))
	}

	// Run service
	if err := grpcService.Run(); err != nil {
		zap.L().Fatal("Fail to run service ", zap.Error(err))
	}
}
