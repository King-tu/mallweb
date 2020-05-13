package router

import (
	"github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/king-tu/mallweb/app/api/apiserver/handler"
	"go.uber.org/zap"
	"time"
)

func InitRouter() *gin.Engine {
	engine := gin.New()
	// 使用zap日志库
	engine.Use(ginzap.Ginzap(zap.L(), time.RFC3339, true))
	engine.Use(ginzap.RecoveryWithZap(zap.L(), true))

	g := engine.Group("/api/v1")
	{
		/******* 短信服务 *******/
		// 短信验证码
		g.POST("/smscode", handler.SendSmsCode)

		/******* 用户相关的服务 *******/
		// 注册
		g.POST("/user/register", handler.HandleRester)
		// 登录

		/******* 商品服务 *******/

	}

	return engine
}
