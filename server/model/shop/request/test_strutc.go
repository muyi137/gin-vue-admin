package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/shop"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type TestStrutcSearch struct{
    shop.TestStrutc
    request.PageInfo
}
