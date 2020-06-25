package handler

import (
	"context"
	"errors"
	"github.com/king-tu/mallweb/app/common/utils"
	"github.com/king-tu/mallweb/app/global"
	"github.com/king-tu/mallweb/app/models"
	"github.com/king-tu/mallweb/app/services/customer/proto/customer"
	"go.uber.org/zap"
)

type CustomerService struct{}

func NewCustomerService() *CustomerService {
	return &CustomerService{}
}

func (e *CustomerService) Register(ctx context.Context, req *customer.RegisterRequest, resp *customer.RegisterResponse) error {
	hashPwd, err := utils.HashPassword(req.Passwd)
	if err != nil {
		global.Logger.For(ctx).Error(true, "计算密码的hash值出错", zap.Error(err))
		return err
	}

	user := new(models.User)
	// 手机号作为用户名
	user.Name = req.Mobile
	user.PassWord = hashPwd

	if err := user.Add(); err != nil {
		global.Logger.For(ctx).Error(true, "添加用户失败", zap.Error(err))
		return err
	}

	global.Logger.Bg().Info("用户注册成功", zap.String("userName", user.Name))

	return nil
}

func (e *CustomerService) Login(ctx context.Context, req *customer.LoginRequest, resp *customer.LoginResponse) error {
	user := new(models.User)
	user.Name = req.Mobile

	if err := user.Get(); err != nil {
		global.Logger.For(ctx).Error(true, "获取用户失败", zap.Error(err))
		return err
	}

	if !utils.VerifyPassword(req.Passwd, user.PassWord) {
		global.Logger.For(ctx).Error(true, "用户名与密码不匹配")
		return errors.New("用户名或密码不正确")

	}

	// jwt
	token, err := utils.GenerateToken(user, global.Config.JWT.ExpiresTime, global.Config.JWT.Secrect)
	if err != nil {
		global.Logger.For(ctx).Error(true, "生成token失败", zap.Error(err))
		return err
	}

	resp.Token = token
	global.Logger.Bg().Info("用户登录成功", zap.String("userName", user.Name))

	return nil
}
