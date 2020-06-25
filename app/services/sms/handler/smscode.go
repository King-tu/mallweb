package handler

import (
	"context"
	"github.com/king-tu/mallweb/app/global"
	"github.com/king-tu/mallweb/app/services/customer/proto/customer"
	microErrors "github.com/micro/go-micro/v2/errors"

	"github.com/king-tu/mallweb/app/common"
	"github.com/king-tu/mallweb/app/common/redis"
	"github.com/king-tu/mallweb/app/common/utils"
	"github.com/king-tu/mallweb/app/services/sms/proto/smscode"
	"go.uber.org/zap"
)

type SmsCodeService struct{}

func NewSmsCodeService() *SmsCodeService {
	return &SmsCodeService{}
}

func (e *SmsCodeService) SendSmsCode(ctx context.Context, req *smscode.SendSmsRequest, resp *smscode.SendSmsResponse) error {
	// generate a code
	code := utils.RandomNumeric(common.SMSCODE_LEN)
	global.Logger.Bg().Debug("发送短信验证码服务", zap.String("请求的手机号码", req.Mobile), zap.String("生成的验证码", code))

	//发送短信验证码（阿里云接口）
	//if err := utils.SendSms(req.Mobile, code); err != nil {
	//	global.Logger.For(ctx).Error(true, "短信验证码发送失败!", zap.Error(err))
	//	return err
	//}

	//短信验证码发送成功，保存短信验证码到redis
	if err := redis.Setex(common.PROJECT_NAME+"_"+common.SMSCODE_GRANT+"_"+req.Mobile, code, common.SMSCODE_EXPIRATION); err != nil {
		global.Logger.For(ctx).Error(true, "短信验证码缓存失败!", zap.Error(err))
		return err
	}
	global.Logger.Bg().Debug("短信验证码缓存成功!")

	//todo **** 多重服务调用 测试 ****
	customerSrv := utils.GetGRPCService(common.SRV_NAME_CUSTOMER)
	cli := customer.NewCustomerService(common.SRV_NAME_CUSTOMER, customerSrv.Client())

	var rReq customer.RegisterRequest
	rReq.Mobile = req.Mobile
	rReq.Smscode = code
	rReq.Passwd = "662090"

	_, err := cli.Register(ctx, &rReq)
	if err != nil {
		e := microErrors.FromError(err)
		global.Logger.For(ctx).Error(false, "Register服务调用出错", zap.String("error", e.Detail))
		//utils.AbortWithError(c, errors.ConvertFrom(e.Detail))
		return err
	}

	return nil
}
