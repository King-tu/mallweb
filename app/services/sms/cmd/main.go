package main

import (
	"github.com/king-tu/mallweb/app/common"
	"github.com/king-tu/mallweb/app/common/utils"
	"github.com/king-tu/mallweb/app/global"
	"github.com/king-tu/mallweb/app/services/sms/handler"
	. "github.com/king-tu/mallweb/app/services/sms/proto/smscode"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
)

func init() {
	// 初始化日志库
	global.SetLogs(global.Config.Log.Level, global.Config.Log.Format, global.Config.Log.FileName, common.SRV_NAME_SMS)
}

func main() {
	// new tracer
	tracer, closer, err := utils.NewTracer(common.SRV_NAME_SMS, global.Config.Jaeger.Addr)
	if err != nil {
		global.Logger.Bg().Fatal("unable to create tracer", zap.Error(err))
	}
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)

	// new service
	smsSrv := utils.GetGRPCService(common.SRV_NAME_SMS)
	// Initialise service
	smsSrv.Init()

	// Register Handler
	if err := RegisterSmsCodeServiceHandler(smsSrv.Server(), handler.NewSmsCodeService()); err != nil {
		global.Logger.Bg().Error(false, "Fail to run  register SmsCodeService", zap.Error(err))
	}

	// Run service
	if err := smsSrv.Run(); err != nil {
		global.Logger.Bg().Fatal("Fail to run service ", zap.Error(err))
	}
}
