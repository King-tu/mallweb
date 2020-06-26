package handler

import (
	"context"
	"github.com/king-tu/mallweb/app/global"
	"github.com/king-tu/mallweb/app/services/goods/proto/goods"
	"go.uber.org/zap"
)

type GoodsService struct{}

func NewGoodsService() *GoodsService {
	return &GoodsService{}
}

func (g *GoodsService) FreshGoodsIndex(ctx context.Context, req *goods.FreshGoodsIndexRequest, resp *goods.FreshGoodsIndexResponse) error {
	var goodsTypes []*goods.GoodsType
	var indexGoodsBanners []*goods.IndexGoodsBanner
	var indexPromotionBanners []*goods.IndexPromotionBanner
	goodsMaps := make(map[string]*goods.IndexTypeGoodsBanners)

	// 商品分类
	if err := global.DB.Find(&goodsTypes).Error; err != nil {
		global.Logger.For(ctx).Error(true, "查询全部商品类型失败", zap.Error(err))
		return err
	}

	// 轮播图
	if err := global.DB.Find(&indexGoodsBanners).Error; err != nil {
		global.Logger.For(ctx).Error(true, "查询首页轮播图失败", zap.Error(err))
		return err
	}

	// 促销商品展示
	if err := global.DB.Find(&indexPromotionBanners).Error; err != nil {
		global.Logger.For(ctx).Error(true, "查询促销商品信息失败", zap.Error(err))
		return err
	}

	//首页分类商品展示
	for _, goodsType := range goodsTypes {
		var typeGoodsBanners []*goods.IndexTypeGoodsBanner
		if err := global.DB.Where("goods_type_id = ?", goodsType.Id).Find(&typeGoodsBanners).Error; err != nil {
			global.Logger.For(ctx).Error(true, "查询首页分类商品", zap.Error(err))
			continue
		}
		goodsMaps[goodsType.Name] = &goods.IndexTypeGoodsBanners{Items: typeGoodsBanners}
	}

	resp.GoodsTypes = goodsTypes
	resp.IndexGoodsBanners = indexGoodsBanners
	resp.IndexPromotionBanners = indexPromotionBanners
	resp.GoodsMaps = goodsMaps

	return nil
}

func (g *GoodsService) GetGoodsDetail(ctx context.Context, req *goods.GoodsDetailRequest, resp *goods.GoodsDetailResponse) error {
	var goodsSKU goods.GoodsSku
	var goodsSPU goods.Goods
	if err := global.DB.Where("id = ?", req.GoodsSkuId).First(&goodsSKU).Error; err != nil {
		global.Logger.For(ctx).Error(true, "查询商品详情失败", zap.Error(err))
		return err
	}

	if err := global.DB.Model(&goodsSKU).Related(&goodsSPU).Error; err != nil {
		global.Logger.For(ctx).Error(true, "查询商品SPU失败", zap.Error(err))
		return err
	}
	goodsSKU.Goods = &goodsSPU

	global.Logger.Bg().Debug("查询商品详情成功", zap.Any("goodsSKU", goodsSKU))

	//获取新品推荐
	var goodsSKUs []*goods.GoodsSku
	if err := global.DB.Where("goods_type_id = ?", goodsSKU.GoodsTypeId).Limit(2).Order("time").Find(&goodsSKUs).Error; err != nil {
		global.Logger.For(ctx).Error(true, "查询新品推荐列表失败", zap.Error(err))
		return err
	}

	resp.GoodsSku = &goodsSKU
	resp.GoodsSkus = goodsSKUs

	return nil
}
