package handler

import (
	"github.com/king-tu/mallweb/app/common"
	"github.com/king-tu/mallweb/app/services/customer/proto/customer"
	"github.com/king-tu/mallweb/app/services/goods/proto/goods"
	"github.com/king-tu/mallweb/app/services/sms/proto/smscode"
	"github.com/micro/go-micro/v2/client"
)

var (
	smscodeServiceClient  smscode.SmsCodeService
	customerServiceClient customer.CustomerService
	goodsServiceClient    goods.GoodsService
)

func RegisterClients(client client.Client) {
	smscodeServiceClient = smscode.NewSmsCodeService(common.SRV_NAME_SMS, client)
	customerServiceClient = customer.NewCustomerService(common.SRV_NAME_CUSTOMER, client)
	goodsServiceClient = goods.NewGoodsService(common.SRV_NAME_GOODS, client)
}
