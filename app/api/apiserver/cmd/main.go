package main

import (
	"github.com/king-tu/mallweb/app/api/apiserver/handler"
	"github.com/king-tu/mallweb/app/api/apiserver/router"
	"github.com/king-tu/mallweb/app/log"
	"github.com/king-tu/mallweb/common"
	"github.com/king-tu/mallweb/common/conf"
	"github.com/king-tu/mallweb/common/utils"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
)

func init() {
	// 初始化日志库
	log.SetLogs(zap.DebugLevel, log.LOGFORMAT_CONSOLE, conf.LogFileName)
}

func main() {
	// new tracer
	tracer, io, err := utils.NewTracer(common.SRV_NAME_SMS, common.JAEGER_ADDR)
	if err != nil {
		zap.L().Sugar().Fatalf("unable to create tracer: %+v\n", err)
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
	if err := r.Run(":8888"); err != nil {
		zap.L().Fatal("HTTP Server启动失败", zap.Error(err))
	}
}
