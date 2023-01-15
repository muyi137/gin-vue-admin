// 自动生成模板Test
package shop

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type GoodsExcel struct {
	global.GVA_MODEL
	Name       string      `json:"name"`       // 名称
	Attr       string      `json:"attr"`       // 规格
	Production string      `json:"production"` // 厂家
	GoodsInfo  []ShopGoods `json:"goodsInfo"`
}
