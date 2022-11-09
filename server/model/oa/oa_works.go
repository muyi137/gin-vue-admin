// 自动生成模板OaWorks
package oa

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
)

// OaWorks 结构体
type OaWorks struct {
      global.GVA_MODEL
      CardNumbaer  string `json:"cardNumbaer" form:"cardNumbaer" gorm:"column:card_numbaer;comment:身份证号;size:22;"`
      Month  *int `json:"month" form:"month" gorm:"column:month;comment:月份;size:4;"`
      Score  *float64 `json:"score" form:"score" gorm:"column:score;comment:绩效值;size:6;"`
      Status  string `json:"status" form:"status" gorm:"column:status;type:enum('已审核','待审核');comment:审核状态;"`
      CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
      UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
      DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName OaWorks 表名
func (OaWorks) TableName() string {
  return "oa_works"
}

