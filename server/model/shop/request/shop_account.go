package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/shop"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type AccountSearch struct{
    shop.Account
    request.PageInfo
}
