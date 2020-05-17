package handler

import (
	"context"
	"errors"
	"github.com/king-tu/mallweb/app/models"
	"github.com/king-tu/mallweb/app/services/customer/proto/customer"
	"github.com/king-tu/mallweb/common/conf"
	"github.com/king-tu/mallweb/common/utils"
	"go.uber.org/zap"
)

type CustomerService struct{}

func NewRegisterService() *CustomerService {
	return &CustomerService{}
}

func (e *CustomerService) Register(ctx context.Context, req *customer.RegisterRequest, resp *customer.RegisterResponse) error {
	hashPwd, err := utils.HashPassword(req.Passwd)
	if err != nil {
		return err
	}

	user := new(models.User)
	// 手机号作为用户名
	user.Name = req.Mobile
	user.PassWord = hashPwd

	if err := user.Add(); err != nil {
		zap.L().Error("添加用户失败", zap.Error(err))
		return err
	}

	zap.L().Info("用户注册成功", zap.String("userName", user.Name))

	return nil
}

func (e *CustomerService) Login(ctx context.Context, req *customer.LoginRequest, resp *customer.LoginResponse) error {
	user := new(models.User)
	user.Name = req.Mobile

	if err := user.Get(); err != nil {
		zap.L().Error("获取用户失败", zap.Error(err))
		return err
	}

	if !utils.VerifyPassword(req.Passwd, user.PassWord) {
		return errors.New("用户名与密码不匹配")
	}

	// jwt
	token, err := utils.GenerateToken(user, conf.TOKEN_EXPIRES_TIME, conf.TOKEN_SECRECT)
	if err != nil {
		return err
	}

	resp.Token = token
	zap.L().Info("用户登录成功", zap.String("userName", user.Name))

	return nil
}
