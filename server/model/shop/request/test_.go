package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/shop"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type TestSearch struct{
    shop.Test
    request.PageInfo
}
