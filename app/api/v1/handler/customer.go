package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/king-tu/mallweb/app/common"
	"github.com/king-tu/mallweb/app/common/errors"
	"github.com/king-tu/mallweb/app/common/middlewares"
	"github.com/king-tu/mallweb/app/common/redis"
	"github.com/king-tu/mallweb/app/common/utils"
	"github.com/king-tu/mallweb/app/global"
	"github.com/king-tu/mallweb/app/models"
	"github.com/king-tu/mallweb/app/services/customer/proto/customer"
	"go.uber.org/zap"
	"net/http"
)

func CheckMobile(c *gin.Context) {
	var req customer.CheckMobileRequest
	if err := c.BindJSON(&req); err != nil {
		global.Logger.Bg().Error(false, "解析输入的JSON数据出错", zap.Error(err))
		utils.AbortWithError(c, errors.ErrJsonParseError)
		return
	}

	user := models.User{Name: req.Mobile}

	exists, err := user.Exists()
	if err != nil {
		global.Logger.Bg().Error(false, "查询用户失败", zap.Error(err))
		utils.AbortWithError(c, errors.ErrGeneralInternalFault)
		return
	}
	if exists {
		utils.AbortWithError(c, errors.ErrMobileNumberAlreadyRegistered)
		return
	}

	c.JSON(http.StatusNoContent, "")
}

func Register(c *gin.Context) {
	ctx, ok := middlewares.ContextWithSpan(c)
	if ok == false {
		global.Logger.Bg().Debug("get context err")
	}

	var req customer.RegisterRequest
	if err := c.BindJSON(&req); err != nil {
		global.Logger.Bg().Error(false, "解析输入的JSON数据出错", zap.Error(err))
		utils.AbortWithError(c, errors.ErrJsonParseError)
		return
	}

	// 校验手机号码
	if err := utils.ValidateMobileNumber(req.Mobile); err != nil {
		global.Logger.Bg().Error(false, "手机号格式错误", zap.Error(err))
		utils.AbortWithError(c, errors.ErrMobileNumberIncorrect)
		return
	}

	// 校验验证码
	smscode, err := redis.Get(common.PROJECT_NAME + "_" + common.SMSCODE_GRANT + "_" + req.Mobile)
	if err != nil {
		global.Logger.Bg().Error(false, "获取缓存中的验证码失败", zap.String("mobile", req.Mobile), zap.Error(err))
		utils.AbortWithError(c, errors.ErrSmsCodeExpires)
		return
	}
	if req.Smscode == "" || req.Smscode != smscode {
		global.Logger.Bg().Error(false, "验证码不正确", zap.String("req.Smscode", req.Smscode), zap.String("smscode", smscode), zap.Error(err))
		utils.AbortWithError(c, errors.ErrSmsCodeIncorrect)
		return
	}

	// 校验密码强度
	if err := utils.ValidatePassword(req.Mobile); err != nil {
		global.Logger.Bg().Error(false, "密码格式不正确", zap.Error(err))
		utils.AbortWithError(c, errors.ErrPasswdIncorrect)
		return
	}

	_, err = customerServiceClient.Register(ctx, &req)
	if err != nil {
		e := errors.ConvertFrom(err)
		global.Logger.Bg().Error(false, "Register服务调用出错", zap.String("error", e.Detail))
		utils.AbortWithError(c, e)
		return
	}

	c.JSON(http.StatusNoContent, "")
}

func Login(c *gin.Context) {
	ctx, ok := middlewares.ContextWithSpan(c)
	if ok == false {
		global.Logger.Bg().Debug("get context err")
	}

	var req customer.LoginRequest
	if err := c.BindJSON(&req); err != nil {
		global.Logger.For(ctx).Error(true, "解析输入的JSON数据出错", zap.Error(err))
		utils.AbortWithError(c, errors.ErrJsonParseError)
		return
	}

	// 校验手机号码
	if err := utils.ValidateMobileNumber(req.Mobile); err != nil {
		global.Logger.For(ctx).Error(true, "手机号格式错误", zap.Error(err))
		utils.AbortWithError(c, errors.ErrMobileNumberIncorrect)
		return
	}

	resp, err := customerServiceClient.Login(ctx, &req)
	if err != nil {
		e := errors.ConvertFrom(err)
		global.Logger.Bg().Error(false, "Login服务调用出错", zap.String("error", e.Detail))
		utils.AbortWithError(c, e)
		return
	}

	// Return the normal response
	c.JSON(http.StatusOK, resp)
}
