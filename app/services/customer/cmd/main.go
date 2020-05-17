package main

import (
	"github.com/king-tu/mallweb/app/log"
	"github.com/king-tu/mallweb/app/services/customer/handler"
	. "github.com/king-tu/mallweb/app/services/customer/proto/customer"
	"github.com/king-tu/mallweb/common"
	"github.com/king-tu/mallweb/common/conf"
	"github.com/king-tu/mallweb/common/utils"
	"go.uber.org/zap"
)

func init() {
	// 初始化日志库
	log.SetLogs(zap.DebugLevel, log.LOGFORMAT_CONSOLE, conf.LogFileName)
}

func main() {
	customerSrv := utils.GetGRPCService(common.SRV_NAME_CUSTOMER)

	// Initialise service
	customerSrv.Init()

	// Register Handler
	if err := RegisterCustomerServiceHandler(customerSrv.Server(), handler.NewRegisterService()); err != nil {
		zap.L().Error("Fail to run  register SmsCodeService", zap.Error(err))
	}

	// Run service
	if err := customerSrv.Run(); err != nil {
		zap.L().Fatal("Fail to run service ", zap.Error(err))
	}
}
