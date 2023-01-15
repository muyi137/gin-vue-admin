package shop

import (
	"database/sql"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/shop"
	shopReq "github.com/flipped-aurora/gin-vue-admin/server/model/shop/request"
	"strings"
)

type ShopGoodsService struct {
}

// CreateShopGoods 创建ShopGoods记录
// Author [piexlmax](https://github.com/piexlmax)
func (shopGoodsService *ShopGoodsService) CreateShopGoods(shopGoods shop.ShopGoods) (err error) {
	err = global.GVA_DB.Create(&shopGoods).Error
	return err
}

// DeleteShopGoods 删除ShopGoods记录
// Author [piexlmax](https://github.com/piexlmax)
func (shopGoodsService *ShopGoodsService) DeleteShopGoods(shopGoods shop.ShopGoods) (err error) {
	err = global.GVA_DB.Delete(&shopGoods).Error
	return err
}

// DeleteShopGoodsByIds 批量删除ShopGoods记录
// Author [piexlmax](https://github.com/piexlmax)
func (shopGoodsService *ShopGoodsService) DeleteShopGoodsByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]shop.ShopGoods{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateShopGoods 更新ShopGoods记录
// Author [piexlmax](https://github.com/piexlmax)
func (shopGoodsService *ShopGoodsService) UpdateShopGoods(shopGoods shop.ShopGoods) (err error) {
	err = global.GVA_DB.Save(&shopGoods).Error
	return err
}

// GetShopGoods 根据id获取ShopGoods记录
// Author [piexlmax](https://github.com/piexlmax)
func (shopGoodsService *ShopGoodsService) GetShopGoods(id uint) (shopGoods shop.ShopGoods, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&shopGoods).Error
	return
}

// GetShopGoodsInfoList 分页获取ShopGoods记录
// Author [piexlmax](https://github.com/piexlmax)
func (shopGoodsService *ShopGoodsService) GetShopGoodsInfoList(info shopReq.ShopGoodsSearch) (list []shop.ShopGoods, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&shop.ShopGoods{})

	var shopGoodss []shop.ShopGoods
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.GoodsName != "" {
		db = db.Where(" goodsName LIKE @name or goodsNameAcronym LIKE @name ", sql.Named("name", "%"+info.GoodsName+"%"))
	}

	if info.AttrName != "" {
		db = db.Where("attrName LIKE ?", "%"+info.AttrName+"%")
	}
	if info.AttrFilter != "" {
		var attrStr = info.AttrFilter
		//字符串转数组

		if attrStr != "" {
			var a = strings.Split(attrStr, ",")
			for key := range a {
				db = db.Where("attrFilter LIKE ?", "%"+a[key]+",%")
			}
			//var c = attrStr.split(" ")
			//for i, c := range collection {
			//	db.Where("attrFilter LIKE ?", "%"+i+"%")
			//}
			//db.Where("attrFilter LIKE ?", "%"+c+"%")
		}

		//db = db.Where("attrFilter LIKE ?", "%"+info.AttrFilter+"%")
		//var goodsAttrs = info.AttrFilter.split(",")

		//for attr in goodsAttrs:
		//    db = db.Where("attrFilter LIKE ?", "%"+attr+",%")
	}
	if info.ApprovalNumber != "" {
		db = db.Where("approvalNumber = ?", info.ApprovalNumber)
	}
	if info.ProductionName != "" {
		db = db.Where("productionName LIKE ?", "%"+info.ProductionName+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	db.Debug().Limit(limit).Offset(offset).Find(&shopGoodss)

	err = db.Limit(limit).Offset(offset).Order("created_at desc").Find(&shopGoodss).Error

	return shopGoodss, total, err
}
