package shop

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ShopGoodsRouter struct {
}

// InitShopGoodsRouter 初始化 ShopGoods 路由信息
func (s *ShopGoodsRouter) InitShopGoodsRouter(Router *gin.RouterGroup) {
	shopGoodsRouter := Router.Group("shopGoods").Use(middleware.OperationRecord())
	shopGoodsRouterWithoutRecord := Router.Group("shopGoods")
	var shopGoodsApi = v1.ApiGroupApp.ShopApiGroup.ShopGoodsApi
	{
		shopGoodsRouter.POST("createShopGoods", shopGoodsApi.CreateShopGoods)             // 新建ShopGoods
		shopGoodsRouter.DELETE("deleteShopGoods", shopGoodsApi.DeleteShopGoods)           // 删除ShopGoods
		shopGoodsRouter.DELETE("deleteShopGoodsByIds", shopGoodsApi.DeleteShopGoodsByIds) // 批量删除ShopGoods
		shopGoodsRouter.PUT("updateShopGoods", shopGoodsApi.UpdateShopGoods)              // 更新ShopGoods

	}
	{
		shopGoodsRouterWithoutRecord.POST("importGoodsExcel", shopGoodsApi.ImportGoodsExcel) // 导入商品excel
		shopGoodsRouterWithoutRecord.POST("exportGoodsExcel", shopGoodsApi.ExportGoodsExcel) // 导出商品excel
		shopGoodsRouterWithoutRecord.GET("loadGoodsExcel", shopGoodsApi.LoadGoodsExcel)      //加载商品excel

		shopGoodsRouterWithoutRecord.GET("findShopGoods", shopGoodsApi.FindShopGoods)       // 根据ID获取ShopGoods
		shopGoodsRouterWithoutRecord.GET("getShopGoodsList", shopGoodsApi.GetShopGoodsList) // 获取ShopGoods列表
	}

}
