package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/king-tu/mallweb/common/database"
	"time"
)

func init() {
	db := database.InitDB()
	// 自动迁移(创建)表
	db.AutoMigrate(new(User), new(Address), new(Address), new(TpshopCategory), new(Goods), new(GoodsType),
		new(Goods), new(GoodsType), new(GoodsSKU), new(GoodsImage),
		new(IndexGoodsBanner), new(IndexTypeGoodsBanner), new(IndexPromotionBanner), new(OrderInfo), new(OrderGoods))
}

type User struct {
	Id        int
	Name      string       `gorm:"size(20);unique"` //用户名
	PassWord  string       `gorm:"size(20)"`        //登陆密码
	Email     string       `gorm:"size(50);null"`   //邮箱
	Active    bool         `gorm:"default(false)"`  //是否激活
	Power     int          `gorm:"default(0)"`      //权限设置  0 表示未激活  1表示激活
	Address   []*Address   // `orm:"reverse(many)"` 一对多
	OrderInfo []*OrderInfo // `orm:"reverse(many)"` 一对多
}

type Address struct {
	Id        int
	Receiver  string `gorm:"size(20)"`      //收件人
	Addr      string `gorm:"size(50)"`      //收件地址
	Zipcode   string `gorm:"size(20)"`      //邮编
	Phone     string `gorm:"size(20)"`      //联系方式
	Isdefault bool   `gorm:"defalt(false)"` //是否默认 0 为非默认  1为默认
	//User *User `orm:"rel(fk)"` 	//用户ID
	UserID    int
	OrderInfo []*OrderInfo // `orm:"reverse(many)"`  一对多
}

type TpshopCategory struct {
	Id         int
	CateName   string `gorm:"default('')"`
	Pid        int    `gorm:"default(0)"`
	IsShow     int    `gorm:"default(0)"`
	CreateTime int    `gorm:"null"`
	UpdateTime int    `gorm:"null"`
	DeleteTime int    `gorm:"null"`
}

type Goods struct { //商品SPU表
	Id       int
	Name     string      `gorm:"size(20)"`  //商品名称
	Detail   string      `gorm:"size(200)"` //详细描述
	GoodsSKU []*GoodsSKU // `orm:"reverse(many)"`   一对多
}

type GoodsType struct { //商品类型表
	Id                   int
	Name                 string                  //种类名称
	Logo                 string                  //logo
	Image                string                  //图片
	GoodsSKU             []*GoodsSKU             // `orm:"reverse(many)"`  一对多
	IndexTypeGoodsBanner []*IndexTypeGoodsBanner // `orm:"reverse(many)"`	一对多
}

type GoodsSKU struct { //商品SKU表
	Id int
	//Goods              *Goods                  `gorm:"rel(fk)"` //商品SPU
	GoodsID int
	//GoodsType          *GoodsType              `orm:"rel(fk)"` //商品所属种类
	GoodsTypeID          int
	Name                 string                  //商品名称
	Desc                 string                  //商品简介
	Price                int                     //商品价格
	Unite                string                  //商品单位
	Image                string                  //商品图片
	Stock                int                     `gorm:"default(1)"`    //商品库存
	Sales                int                     `gorm:"default(0)"`    //商品销量
	Status               int                     `gorm:"default(1)"`    //商品状态
	Time                 time.Time               `gorm:"type:datetime"` //添加时间
	GoodsImage           []*GoodsImage           // `orm:"reverse(many)"`
	IndexGoodsBanner     []*IndexGoodsBanner     // `orm:"reverse(many)"`
	IndexTypeGoodsBanner []*IndexTypeGoodsBanner // `orm:"reverse(many)"`

	OrderGoods []*OrderGoods // `orm:"reverse(many)"`
}

type GoodsImage struct { //商品图片表
	Id    int
	Image string //商品图片
	//GoodsSKU *GoodsSKU `orm:"rel(fk)"` //商品SKU
	GoodsSKUID int
}
type IndexGoodsBanner struct { //首页轮播商品展示表
	Id int
	//GoodsSKU *GoodsSKU `orm:"rel(fk)"` //商品sku
	GoodsSKUID int
	Image      string //商品图片
	Index      int    `gorm:"default(0)"` //展示顺序
}

type IndexTypeGoodsBanner struct { //首页分类商品展示表
	Id int
	//GoodsType   *GoodsType `orm:"rel(fk)"`    //商品类型
	GoodsTypeID int
	//GoodsSKU    *GoodsSKU `orm:"rel(fk)"`    //商品sku
	GoodsSKUID  int
	DisplayType int `gorm:"default(1)"` //展示类型 0代表文字，1代表图片
	Index       int `gorm:"default(0)"` //展示顺序
}

type IndexPromotionBanner struct { //首页促销商品展示表
	Id    int
	Name  string `gorm:"size(20)"` //活动名称
	Url   string `gorm:"size(50)"` //活动链接
	Image string //活动图片
	Index int    `gorm:"default(0)"` //展示顺序
}

type OrderInfo struct { //订单表
	Id      int
	OrderId string `gorm:"unique"` //订单号
	//User 			*User	`orm:"rel(fk)"`		//用户
	UserID int
	//Address      *Address  `orm:"rel(fk)"` //地址
	AddressID    int
	PayMethod    int           //付款方式
	TotalCount   int           `gorm:"default(1)"` //商品数量
	TotalPrice   int           //商品总价
	TransitPrice int           //运费
	Orderstatus  int           `gorm:"default(0)"`    //订单状态
	TradeNo      string        `gorm:"default('')"`   //支付编号
	Time         time.Time     `gorm:"type:datetime"` //订单时间
	OrderGoods   []*OrderGoods // `orm:"reverse(many)"`
}

type OrderGoods struct { //订单商品表
	Id int
	//OrderInfo *OrderInfo `gorm:"rel(fk)"`    //订单
	OrderInfoID int
	//GoodsSKU  *GoodsSKU  `orm:"rel(fk)"`    //商品
	GoodsSKUID int
	Count      int    `gorm:"default(1)"` //商品数量
	Price      int    //商品价格
	Comment    string `gorm:"default('')"` //评论
}
