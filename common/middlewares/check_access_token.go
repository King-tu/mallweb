package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/king-tu/mallweb/common/conf"
	"github.com/king-tu/mallweb/common/errors"
	"github.com/king-tu/mallweb/common/utils"
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
			zap.L().Debug("从请求中获取的token为空")
			context.JSON(http.StatusUnauthorized, errors.ErrAuthenticationFail)
			return
		}

		auth = strings.Fields(auth)[1]
		zap.L().Sugar().Debugf("auth %s", auth)

		// 校验token
		token, err := utils.ParseToken(auth, conf.TOKEN_SECRECT)
		if err != nil || time.Now().Unix() > token.ExpiresAt {
			context.Abort()
			zap.L().Error("token解析失败", zap.Error(err))
			context.JSON(http.StatusUnauthorized, errors.ErrAuthenticationFail)
			return
		}

		zap.L().Debug("token认证成功", zap.String("userName", token.Audience))

		context.Next()
	}
}
