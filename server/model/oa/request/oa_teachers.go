package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/oa"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type OaTeachersSearch struct{
    oa.OaTeachers
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    request.PageInfo
}
