// 自动生成模板OaAppraisal
package oa

import (
    "github.com/flipped-aurora/gin-vue-admin/server/global"
)

// OaAppraisal 结构体
type OaAppraisal struct {
    global.GVA_MODEL
    CardNumber string   `json:"cardNumber" form:"cardNumber" gorm:"column:card_number;comment:身份证号;size:22;"`
    Month      string   `json:"month" form:"month" gorm:"column:month;comment:月份;size:6;"`
    Score      *float64 `json:"score" form:"score" gorm:"column:score;comment:绩效值;size:4;"`
    Status     string   `json:"status" form:"status" gorm:"column:status;type:enum('已审核','待审核');comment:状态;"`
    CreatedBy  uint     `gorm:"column:created_by;comment:创建者"`
    UpdatedBy  uint     `gorm:"column:updated_by;comment:更新者"`
    DeletedBy  uint     `gorm:"column:deleted_by;comment:删除者"`
}

// TableName OaAppraisal 表名
func (OaAppraisal) TableName() string {
    return "oa_appraisal"
}

