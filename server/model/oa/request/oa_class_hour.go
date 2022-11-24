package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/oa"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type OaClassHourSearch struct{
    oa.OaClassHour
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    StartReportTime  *time.Time  `json:"startReportTime" form:"startReportTime"`
    EndReportTime  *time.Time  `json:"endReportTime" form:"endReportTime"`
    request.PageInfo
}
