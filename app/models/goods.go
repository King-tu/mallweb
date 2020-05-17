package models

import (
	"time"
)

type Goods struct { //商品SPU表
	Id       int
	Name     string      `gorm:"size(20)"`  //商品名称
	Detail   string      `gorm:"size(200)"` //详细描述
	GoodsSKU []*GoodsSKU // 一对多
}

type GoodsType struct { //商品类型表
	Id                   int
	Name                 string                  //种类名称
	Logo                 string                  //logo
	Image                string                  //图片
	GoodsSKU             []*GoodsSKU             // 一对多
	IndexTypeGoodsBanner []*IndexTypeGoodsBanner // 一对多
}

type GoodsSKU struct { //商品SKU表
	Id                   int
	GoodsID              int
	GoodsTypeID          int
	Name                 string    //商品名称
	Desc                 string    //商品简介
	Price                int       //商品价格
	Unite                string    //商品单位
	Image                string    //商品图片
	Stock                int       `gorm:"default(1)"`    //商品库存
	Sales                int       `gorm:"default(0)"`    //商品销量
	Status               int       `gorm:"default(1)"`    //商品状态
	Time                 time.Time `gorm:"type:datetime"` //添加时间
	GoodsImage           []*GoodsImage
	IndexGoodsBanner     []*IndexGoodsBanner
	IndexTypeGoodsBanner []*IndexTypeGoodsBanner

	OrderGoods []*OrderGoods
}

type GoodsImage struct { //商品图片表
	Id         int
	Image      string //商品图片
	GoodsSKUID int
}

type IndexGoodsBanner struct { //首页轮播商品展示表
	Id         int
	GoodsSKUID int
	Image      string //商品图片链接
	Index      int    `gorm:"default(0)"` //展示顺序
}

type IndexTypeGoodsBanner struct { //首页分类商品展示表
	Id          int
	GoodsTypeID int
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

type TpshopCategory struct {
	Id         int
	CateName   string `gorm:"default('')"`
	Pid        int    `gorm:"default(0)"`
	IsShow     int    `gorm:"default(0)"`
	CreateTime int    `gorm:"null"`
	UpdateTime int    `gorm:"null"`
	DeleteTime int    `gorm:"null"`
}

func init() {
	// 自动迁移(创建)表
	DB().AutoMigrate(new(Goods), new(GoodsType), new(GoodsSKU), new(GoodsImage),
		new(IndexGoodsBanner), new(IndexTypeGoodsBanner), new(IndexPromotionBanner), new(TpshopCategory))
}
