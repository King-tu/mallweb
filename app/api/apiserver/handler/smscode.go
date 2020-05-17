package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/king-tu/mallweb/app/services/sms/proto/smscode"
	"github.com/king-tu/mallweb/common"
	"github.com/king-tu/mallweb/common/errors"
	"github.com/king-tu/mallweb/common/redis"
	"github.com/king-tu/mallweb/common/utils"
	microErrors "github.com/micro/go-micro/v2/errors"
	"go.uber.org/zap"
	"net/http"
)

// 发送短信验证码
func SendSmsCode(c *gin.Context) {
	var req smscode.SendSmsRequest
	if err := c.BindJSON(&req); err != nil {
		zap.L().Error("c.BindJSON 出错", zap.Error(err))
		utils.AbortWithError(c, errors.ErrJsonParseError)
		return
	}

	// 校验手机号码
	if err := utils.ValidateMobileNumber(req.Mobile); err != nil {
		zap.L().Error("手机号格式错误", zap.Error(err))
		utils.AbortWithError(c, errors.ErrMobileNumberIncorrect)
		return
	}

	// 校验验证码是否还在有效期内
	if redis.Exists(common.PROJECT_NAME + "_" + common.SMSCODE_GRANT + "_" + req.Mobile) {
		zap.L().Debug("验证码还在有效期内")
		utils.AbortWithError(c, errors.ErrRefreshNotExpires)
		return
	}

	_, err := smscodeServiceClient.SendSmsCode(c, &req)
	if err != nil {
		e := microErrors.FromError(err)
		zap.L().Error("SendSmsCode服务调用出错", zap.String("error", e.Detail))
		utils.AbortWithError(c, errors.ConvertFrom(e.Detail))
		return
	}

	// Return the normal response
	c.String(http.StatusNoContent, "")
}
