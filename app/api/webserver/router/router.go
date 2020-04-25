package router

import (
	"fmt"
	"github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

func InitRoutes() *gin.Engine {
	engine := gin.New()
	// 使用zap日志库
	engine.Use(ginzap.Ginzap(zap.L(), time.RFC3339, true))
	engine.Use(ginzap.RecoveryWithZap(zap.L(), true))

	// Example ping request.
	engine.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong "+fmt.Sprint(time.Now().Unix()))
	})

	// Example when panic happen.
	engine.GET("/panic", func(c *gin.Context) {
		panic("An unexpected error happen!")
	})

	g := engine.Group("/api/v1")
	{
		/***用户相关***/
		// 短信验证码
		g.GET("/smscode/:mobile")
		//// 注册
		g.POST("/user/register")
		//// 登录
		//g.GET("/user/login", authSrv.User, login)
		//g.POST("/user/logout", authSrv.Guest, bm.Mobile(), pushLogout)

	}

	return engine
}
