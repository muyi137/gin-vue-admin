package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/shop"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type ShopGoodsSearch struct{
    shop.ShopGoods
    request.PageInfo
}
