package model

import (
	"time"
)

type OrderInfo struct { //订单表
	Id           int
	OrderId      string `gorm:"unique"` //订单号
	UserID       int
	AddressID    int
	PayMethod    int       //付款方式
	TotalCount   int       `gorm:"default(1)"` //商品数量
	TotalPrice   int       //商品总价
	TransitPrice int       //运费
	Orderstatus  int       `gorm:"default(0)"`    //订单状态
	TradeNo      string    `gorm:"default('')"`   //支付编号
	Time         time.Time `gorm:"type:datetime"` //订单时间
	OrderGoods   []*OrderGoods
}

type OrderGoods struct { //订单商品表
	Id          int
	OrderInfoID int
	GoodsSKUID  int
	Count       int    `gorm:"default(1)"` //商品数量
	Price       int    //商品价格
	Comment     string `gorm:"default('')"` //评论
}
