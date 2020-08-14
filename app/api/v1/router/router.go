package router

import (
	"github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/king-tu/mallweb/app/api/v1/handler"
	"github.com/king-tu/mallweb/app/common/middlewares"
	"go.uber.org/zap"
	"net/http"
	"time"
)

func InitRouter() *gin.Engine {
	engine := gin.New()
	// 使用zap日志库
	engine.Use(ginzap.Ginzap(zap.L(), time.RFC3339, true))
	engine.Use(ginzap.RecoveryWithZap(zap.L(), true))

	v1 := engine.Group("/api/v1")
	v1.Use(middlewares.RateLimiter)
	v1.Use(middlewares.TracerWrapper)
	{
		// test
		v1.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, "OK!!!\n")
		})

		/******* 短信服务 *******/
		// 短信验证码
		v1.POST("/smscode", handler.SendSmsCode)

		/******* 用户相关的服务 *******/
		u := v1.Group("user")
		{
			// 检查手机号(用户名)是否已注册
			u.POST("checkMobile", handler.CheckMobile)
			// 注册
			u.POST("register", handler.Register)
			// 登录
			u.POST("login", handler.Login)
		}

		/******* 商品服务 *******/
		g := v1.Group("goods")
		// token 校验
		g.Use(middlewares.Auth())
		{
			g.GET("freshGoodsIndex", handler.GetFreshGoodsIndex)
			g.GET("goodsDetail/:id", handler.GetGoodsDetail)
			g.GET("searchGoods", handler.SearchGoods)
		}
	}

	return engine
}
