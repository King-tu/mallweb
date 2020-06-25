package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/king-tu/mallweb/app/common"
	"github.com/king-tu/mallweb/app/common/errors"
	"github.com/king-tu/mallweb/app/common/middlewares"
	"github.com/king-tu/mallweb/app/common/redis"
	"github.com/king-tu/mallweb/app/common/utils"
	"github.com/king-tu/mallweb/app/global"
	"github.com/king-tu/mallweb/app/services/sms/proto/smscode"
	"go.uber.org/zap"
	"net/http"
)

// 发送短信验证码
func SendSmsCode(c *gin.Context) {
	ctx, ok := middlewares.ContextWithSpan(c)
	if ok == false {
		global.Logger.Bg().Debug("get context err")
	}

	var req smscode.SendSmsRequest
	if err := c.BindJSON(&req); err != nil {
		global.Logger.Bg().Error(false, "c.BindJSON 出错", zap.Error(err))
		utils.AbortWithError(c, errors.ErrJsonParseError)
		return
	}

	// 校验手机号码
	if err := utils.ValidateMobileNumber(req.Mobile); err != nil {
		global.Logger.Bg().Error(false, "手机号格式错误", zap.Error(err))
		utils.AbortWithError(c, errors.ErrMobileNumberIncorrect)
		return
	}

	// 校验验证码是否还在有效期内
	if redis.Exists(common.PROJECT_NAME + "_" + common.SMSCODE_GRANT + "_" + req.Mobile) {
		global.Logger.Bg().Debug("验证码还在有效期内")
		utils.AbortWithError(c, errors.ErrRefreshNotExpires)
		return
	}

	_, err := smscodeServiceClient.SendSmsCode(ctx, &req)
	if err != nil {
		e := errors.ConvertFrom(err)
		global.Logger.For(ctx).Error(false, "SendSmsCode服务调用出错", zap.String("error", e.Detail))
		utils.AbortWithError(c, e)
		return
	}

	// Return the normal response
	c.String(http.StatusNoContent, "")
}
