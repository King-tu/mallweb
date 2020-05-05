package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/king-tu/mallweb/app/models"
	"github.com/king-tu/mallweb/common/database"
	//"regexp"
)

//发送短信验证码
func GetSmscd(ctx *gin.Context) {
	////获取
	//mobile := ctx.Param("mobile")
	//smsCode := ctx.Query("text")
	//uuid := ctx.Query("id")
	//
	////校验　(ctx.Param获取的数据可以不用校验)
	//if mobile == "" || smsCode == "" || uuid == "" {
	//	resp.Errno = utils.RECODE_PARAMERR
	//	resp.Errmsg = utils.RecodeText(utils.RECODE_PARAMERR)
	//	//401 Unauthorized：用户未提供身份验证凭据，或者没有通过身份验证。
	//	ctx.JSON(401, resp)
	//	//return
	//	//打印错误信息(到终端+日志), 并 调用 os.Exit(1)结束程序.
	//	errLog.Error.Println("获取不到前端数据")
	//	return
	//}
	//
	////正则校验手机号
	//regExp, _ := regexp.Compile(`^1[3-9][0-9]{9}$`)
	//isMatch := regExp.MatchString(mobile)
	//
	//if !isMatch {
	//	resp.Errno = utils.RECODE_MOBILEERR
	//	resp.Errmsg = utils.RecodeText(utils.RECODE_MOBILEERR)
	//	//401 Unauthorized：用户未提供身份验证凭据，或者没有通过身份验证。
	//	ctx.JSON(401, resp)
	//	//return
	//	//打印错误信息(到终端+日志)
	//	errLog.Error.Println("手机号码格式不对")
	//	return
	//}
	//
	////	调用ｒｐｃ服务处理数据
	//grpcService := utils.GetGRPCService()
	//client := smsCodeSrv.NewSmsCodeSrvService(utils.SmsCodeSrv_Name, grpcService.Client())
	//resp, err := client.Call(context.TODO(), &smsCodeSrv.Request{Mobile: mobile, SmsCode: smsCode, Uuid: uuid})
	//if err != nil {
	//	//打印错误信息(到终端+日志)
	//	errLog.Error.Println(utils.SmsCodeSrv_Name + ": 服务调用失败, err: ", err)
	//	resp.Errno = utils.RECODE_DATAERR
	//	resp.Errmsg = utils.RecodeText(utils.RECODE_DATAERR)
	//	ctx.JSON(500, resp)
	//	return
	//}
	//
	//resp.Errno = utils.RECODE_OK
	//resp.Errmsg = utils.RecodeText(utils.RECODE_OK)
	//ctx.JSON(200, resp)
}

func HandleRester(c *gin.Context) {
	user := models.User{
		Name: "张三",
	}

	db := database.InitDB()
	err := db.DB().Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("数据库连接成功: ", user)
	c.JSON(200, "hello world")
}
