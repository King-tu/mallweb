package main

import (
	"github.com/king-tu/mallweb/app/api/v1/handler"
	"github.com/king-tu/mallweb/app/api/v1/router"
	"github.com/king-tu/mallweb/app/common"
	"github.com/king-tu/mallweb/app/common/utils"
	"github.com/king-tu/mallweb/app/global"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
)

func init() {
	// 初始化日志库
	global.SetLogs(global.Config.Log.Level, global.Config.Log.Format, global.Config.Log.FileName, common.SRV_NAME_APIGATEWAY)
}

func main() {
	// new tracer
	tracer, io, err := utils.NewTracer(common.SRV_NAME_APIGATEWAY, global.Config.Jaeger.Addr)
	if err != nil {
		global.Logger.Bg().Fatal("unable to create tracer", zap.Error(err))
	}
	defer io.Close()
	opentracing.SetGlobalTracer(tracer)

	// new service
	smsSrv := utils.GetGRPCService(common.SRV_NAME_SMS)
	customerSrv := utils.GetGRPCService(common.SRV_NAME_CUSTOMER)
	// Initialise service
	smsSrv.Init()
	customerSrv.Init()
	handler.RegisterClients(smsSrv.Client())
	handler.RegisterClients(customerSrv.Client())

	// Init router
	r := router.InitRouter()

	// run
	if err := r.Run(global.Config.System.ApiAddr); err != nil {
		global.Logger.Bg().Fatal("HTTP Server启动失败", zap.Error(err))
	}
}
