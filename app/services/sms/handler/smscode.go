package handler

import (
	"context"
	"github.com/king-tu/mallweb/app/log"
	"github.com/king-tu/mallweb/app/services/customer/proto/customer"
	microErrors "github.com/micro/go-micro/v2/errors"

	//"fmt"
	//"github.com/king-tu/mallweb/app/services/sms/errors"
	"github.com/king-tu/mallweb/app/services/sms/proto/smscode"
	"github.com/king-tu/mallweb/common"
	"github.com/king-tu/mallweb/common/redis"
	"github.com/king-tu/mallweb/common/utils"
	"go.uber.org/zap"
)

type SmsCodeService struct{}

func NewSmsCodeService() *SmsCodeService {
	return &SmsCodeService{}
}

func (e *SmsCodeService) SendSmsCode(ctx context.Context, req *smscode.SendSmsRequest, resp *smscode.SendSmsResponse) error {
	// generate a code
	code := utils.RandomNumeric(common.SMSCODE_LEN)

	zap.L().Sugar().Debugf("请求的手机号码：%s， 生成的验证码：%s", req.Mobile, code)

	//发送短信验证码（阿里云接口）
	//if err := utils.SendSms(req.Mobile, code); err != nil {
	//	zap.L().Error("短信验证码发送失败!", zap.Error(err))
	//	return err
	//}

	//短信验证码发送成功，保存短信验证码到redis
	if err := redis.Setex(common.PROJECT_NAME+"_"+common.SMSCODE_GRANT+"_"+req.Mobile, code, common.SMSCODE_EXPIRATION); err != nil {
		log.L().For(ctx).Error(true, "短信验证码缓存失败!", zap.Error(err))
		return err
	}
	zap.L().Debug("短信验证码缓存成功!")

	//c, ok := middlewares.ContextWithSpan(ctx)
	//if ok == false {
	//	zap.L().Debug("get context err")
	//}

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
		log.L().For(ctx).Error(false, "Register服务调用出错", zap.String("error", e.Detail))
		//utils.AbortWithError(c, errors.ConvertFrom(e.Detail))
		return err
	}

	return nil
}
