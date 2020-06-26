package main

import (
	"github.com/king-tu/mallweb/app/common"
	"github.com/king-tu/mallweb/app/common/utils"
	"github.com/king-tu/mallweb/app/global"
	"github.com/king-tu/mallweb/app/services/goods/handler"
	"github.com/king-tu/mallweb/app/services/goods/proto/goods"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
)

func init() {
	// 初始化日志库
	global.SetLogs(global.Config.Log.Level, global.Config.Log.Format, global.Config.Log.FileName, common.SRV_NAME_CUSTOMER)
}

func main() {
	// new tracer
	tracer, closer, err := utils.NewTracer(common.SRV_NAME_CUSTOMER, global.Config.Jaeger.Addr)
	if err != nil {
		global.Logger.Bg().Fatal("unable to create tracer", zap.Error(err))
	}
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)

	goodsSrv := utils.GetGRPCService(common.SRV_NAME_GOODS)

	// Initialise service
	goodsSrv.Init()

	// Register Handler
	if err := goods.RegisterGoodsServiceHandler(goodsSrv.Server(), handler.NewGoodsService()); err != nil {
		global.Logger.Bg().Error(false, "Fail to run  register SmsCodeService", zap.Error(err))
	}

	// Run service
	if err := goodsSrv.Run(); err != nil {
		global.Logger.Bg().Fatal("Fail to run service ", zap.Error(err))
	}
}
