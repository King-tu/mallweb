package router

import (
	"github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/king-tu/mallweb/app/api/v1/handler"
	"github.com/king-tu/mallweb/app/common/middlewares"
	"go.uber.org/zap"
	"time"
)

func InitRouter() *gin.Engine {
	engine := gin.New()
	// 使用zap日志库
	engine.Use(ginzap.Ginzap(zap.L(), time.RFC3339, true))
	engine.Use(ginzap.RecoveryWithZap(zap.L(), true))

	g := engine.Group("/api/v1")
	g.Use(middlewares.TracerWrapper)
	{
		/******* 短信服务 *******/
		// 短信验证码
		g.POST("/smscode", handler.SendSmsCode)

		/******* 用户相关的服务 *******/
		// 检查手机号(用户名)是否已注册
		g.POST("/user/checkMobile", handler.CheckMobile)
		// 注册
		g.POST("/user/register", handler.Register)
		// 登录
		g.POST("/user/login", handler.Login)

		g.Use(middlewares.Auth())

		g.GET("/hello", func(c *gin.Context) {
			c.JSON(200, "hello world !!!")
		})

		/******* 商品服务 *******/

	}

	return engine
}
