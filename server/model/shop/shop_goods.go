// 自动生成模板ShopGoods
package shop

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// ShopGoods 结构体
type ShopGoods struct {
	global.GVA_MODEL
	GoodsId           *int     `json:"goodsId" form:"goodsId" gorm:"column:goodsId;comment:商品id;size:10;"`
	GoodsName         string   `json:"goodsName" form:"goodsName" gorm:"column:goodsName;comment:商品名称;size:40;"`
	GoodsNameAcronym  string   `json:"goodsNameAcronym" form:"goodsNameAcronym" gorm:"column:goodsNameAcronym;comment:商品名称首字母;size:20;"`
	CategoryId        *int     `json:"categoryId" form:"categoryId" gorm:"column:categoryId;comment:分类id;size:10;"`
	ScopeId           string   `json:"scopeId" form:"scopeId" gorm:"column:scopeId;comment:类型;size:20;"`
	AttrName          string   `json:"attrName" form:"attrName" gorm:"column:attrName;comment:规格;size:20;"`
	AttrFilter        string   `json:"attrFilter" form:"attrFilter" gorm:"column:attrFilter;comment:规格过滤;size:20;"`
	PackageUnit       string   `json:"packageUnit" form:"packageUnit" gorm:"column:packageUnit;comment:单位;size:10;"`
	Dosage            string   `json:"dosage" form:"dosage" gorm:"column:dosage;comment:剂型;size:15;"`
	ApprovalNumber    string   `json:"approvalNumber" form:"approvalNumber" gorm:"column:approvalNumber;comment:批准字号;size:30;"`
	ExpireTime        string   `json:"expireTime" form:"expireTime" gorm:"column:expireTime;comment:近效期;"`
	PackageNum        *int     `json:"packageNum" form:"packageNum" gorm:"column:packageNum;comment:包装数量;size:10;"`
	SalePrice         *float64 `json:"salePrice" form:"salePrice" gorm:"column:salePrice;comment:销售价格;size:6;"`
	PromotionPrice    *float64 `json:"promotionPrice" form:"promotionPrice" gorm:"column:promotionPrice;comment:促销价格;size:6;"`
	MaxPromotionPrice *float64 `json:"maxPromotionPrice" form:"maxPromotionPrice" gorm:"column:maxPromotionPrice;comment:最高促销价格;size:6;"`
	MaxDiscountPrice  *float64 `json:"maxDiscountPrice" form:"maxDiscountPrice" gorm:"column:maxDiscountPrice;comment:立省价格 =原价-促销价;size:6;"`
	MiddlePackage     *int     `json:"middlePackage" form:"middlePackage" gorm:"column:middlePackage;comment:中包装;size:10;"`
	StockNum          *int     `json:"stockNum" form:"stockNum" gorm:"column:stockNum;comment:库存量;size:10;"`
	ProductionName    string   `json:"productionName" form:"productionName" gorm:"column:productionName;comment:厂家;size:150;"`
	//PromotionStatus   string   `json:"promotionStatus" form:"promotionStatus" gorm:"column:promotionStatus;comment:促销状态;size:6;"`
	GoodsImage string `json:"goodsImage" form:"goodsImage" gorm:"column:goodsImage;comment:图片地址;size:150;"`
	//GoodsCode         string   `json:"goodsCode" form:"goodsCode" gorm:"column:goodsCode;comment:商品编号;size:8;"`
	//ErpCode           string   `json:"erpCode" form:"erpCode" gorm:"column:erpCode;comment:erp code;size:12;"`
	DataSources string `json:"dataSources" form:"dataSources" gorm:"column:dataSources;comment:数据来源;size:10;"`
	//Addtime           string   `json:"addtime" form:"addtime" gorm:"column:addtime;comment:添加时间;size:20;"`
	GoodsUrl string `json:"goodsUrl" form:"goodsUrl" gorm:"column:goodsUrl;comment:商品链接;size:100;"`
	Tag      string `json:"tag" form:"tag" gorm:"column:tag;comment:标签;"`
}

// TableName ShopGoods 表名
func (ShopGoods) TableName() string {
	return "shop_goods"
}
