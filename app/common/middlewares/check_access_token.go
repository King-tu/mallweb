package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/king-tu/mallweb/app/common/errors"
	"github.com/king-tu/mallweb/app/common/utils"
	"github.com/king-tu/mallweb/app/global"
	"go.uber.org/zap"
	"net/http"
	"strings"
	"time"
)

func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		// 从请求中获取token
		auth := context.Request.Header.Get("Authorization")
		if len(auth) == 0 {
			context.Abort()
			global.Logger.Bg().Debug("从请求中获取的token为空")
			context.JSON(http.StatusUnauthorized, errors.ErrAuthenticationFail)
			return
		}

		auth = strings.Fields(auth)[1]

		// 校验token
		token, err := utils.ParseToken(auth, global.Config.JWT.Secrect)
		if err != nil || time.Now().Unix() > token.ExpiresAt {
			context.Abort()
			global.Logger.Bg().Error(false, "token解析失败", zap.Error(err))
			context.JSON(http.StatusUnauthorized, errors.ErrAuthenticationFail)
			return
		}

		global.Logger.Bg().Debug("token认证成功", zap.String("userName", token.Audience))

		context.Next()
	}
}
