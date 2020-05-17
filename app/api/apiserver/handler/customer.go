package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/king-tu/mallweb/app/models"
	"github.com/king-tu/mallweb/app/services/customer/proto/customer"
	"github.com/king-tu/mallweb/common"
	"github.com/king-tu/mallweb/common/errors"
	"github.com/king-tu/mallweb/common/redis"
	"github.com/king-tu/mallweb/common/utils"
	microErrors "github.com/micro/go-micro/v2/errors"
	"go.uber.org/zap"
	"net/http"
)

func CheckMobile(c *gin.Context) {
	var req customer.CheckMobileRequest
	if err := c.BindJSON(&req); err != nil {
		zap.L().Error("解析输入的JSON数据出错", zap.Error(err))
		utils.AbortWithError(c, errors.ErrJsonParseError)
		return
	}

	user := models.User{Name: req.Mobile}

	exists, err := user.Exists()
	if err != nil {
		zap.L().Error("查询用户失败", zap.Error(err))
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
	var req customer.RegisterRequest
	if err := c.BindJSON(&req); err != nil {
		zap.L().Error("解析输入的JSON数据出错", zap.Error(err))
		utils.AbortWithError(c, errors.ErrJsonParseError)
		return
	}

	// 校验手机号码
	if err := utils.ValidateMobileNumber(req.Mobile); err != nil {
		zap.L().Error("手机号格式错误", zap.Error(err))
		utils.AbortWithError(c, errors.ErrMobileNumberIncorrect)
		return
	}

	// 校验验证码
	smscode, err := redis.Get(common.PROJECT_NAME + "_" + common.SMSCODE_GRANT + "_" + req.Mobile)
	if err != nil {
		zap.L().Error("获取缓存中的验证码失败", zap.String("mobile", req.Mobile), zap.Error(err))
		utils.AbortWithError(c, errors.ErrSmsCodeExpires)
		return
	}
	if req.Smscode == "" || req.Smscode != smscode {
		zap.L().Error("验证码不正确", zap.String("req.Smscode", req.Smscode), zap.String("smscode", smscode), zap.Error(err))
		utils.AbortWithError(c, errors.ErrSmsCodeIncorrect)
		return
	}

	// 校验密码强度
	if err := utils.ValidatePassword(req.Mobile); err != nil {
		zap.L().Error("密码格式不正确", zap.Error(err))
		utils.AbortWithError(c, errors.ErrPasswdIncorrect)
		return
	}

	_, err = customerServiceClient.Register(c, &req)
	if err != nil {
		e := microErrors.FromError(err)
		zap.L().Error("Register服务调用出错", zap.String("error", e.Detail))
		utils.AbortWithError(c, errors.ConvertFrom(e.Detail))
		return
	}

	c.JSON(http.StatusNoContent, "")
}

func Login(c *gin.Context) {
	var req customer.LoginRequest
	if err := c.BindJSON(&req); err != nil {
		zap.L().Error("解析输入的JSON数据出错", zap.Error(err))
		utils.AbortWithError(c, errors.ErrJsonParseError)
		return
	}

	// 校验手机号码
	if err := utils.ValidateMobileNumber(req.Mobile); err != nil {
		zap.L().Error("手机号格式错误", zap.Error(err))
		utils.AbortWithError(c, errors.ErrMobileNumberIncorrect)
		return
	}

	resp, err := customerServiceClient.Login(c, &req)
	if err != nil {
		e := microErrors.FromError(err)
		zap.L().Error("Register服务调用出错", zap.String("error", e.Detail))
		utils.AbortWithError(c, errors.ConvertFrom(e.Detail))
		return
	}

	// Return the normal response
	c.JSON(http.StatusOK, resp)
}
