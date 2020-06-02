package handler

import (
	"context"
	"errors"
	"github.com/king-tu/mallweb/app/log"
	"github.com/king-tu/mallweb/app/models"
	"github.com/king-tu/mallweb/app/services/customer/proto/customer"
	"github.com/king-tu/mallweb/common/conf"
	"github.com/king-tu/mallweb/common/utils"
	"go.uber.org/zap"
)

type CustomerService struct{}

func NewCustomerService() *CustomerService {
	return &CustomerService{}
}

func (e *CustomerService) Register(ctx context.Context, req *customer.RegisterRequest, resp *customer.RegisterResponse) error {
	hashPwd, err := utils.HashPassword(req.Passwd)
	if err != nil {
		log.L().For(ctx).Error(true, "计算密码的hash值出错", zap.Error(err))
		return err
	}

	user := new(models.User)
	// 手机号作为用户名
	user.Name = req.Mobile
	user.PassWord = hashPwd

	if err := user.Add(); err != nil {
		log.L().For(ctx).Error(true, "添加用户失败", zap.Error(err))
		return err
	}

	zap.L().Info("用户注册成功", zap.String("userName", user.Name))

	return nil
}

func (e *CustomerService) Login(ctx context.Context, req *customer.LoginRequest, resp *customer.LoginResponse) error {
	user := new(models.User)
	user.Name = req.Mobile

	if err := user.Get(); err != nil {
		log.L().For(ctx).Error(true, "获取用户失败", zap.Error(err))
		return err
	}

	if !utils.VerifyPassword(req.Passwd, user.PassWord) {
		log.L().For(ctx).Error(true, "用户名与密码不匹配")
		return errors.New("用户名与密码不匹配")
	}

	// jwt
	token, err := utils.GenerateToken(user, conf.TOKEN_EXPIRES_TIME, conf.TOKEN_SECRECT)
	if err != nil {
		log.L().For(ctx).Error(true, "生成token失败", zap.Error(err))
		return err
	}

	resp.Token = token
	zap.L().Info("用户登录成功", zap.String("userName", user.Name))

	return nil
}
