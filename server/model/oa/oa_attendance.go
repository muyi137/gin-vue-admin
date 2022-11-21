// 自动生成模板OaAttendance
package oa

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
)

// OaAttendance 结构体
type OaAttendance struct {
      global.GVA_MODEL
      CardNumber  string `json:"cardNumber" form:"cardNumber" gorm:"column:card_number;comment:;size:22;"`
      Month  *int `json:"month" form:"month" gorm:"column:month;comment:月份;size:4;"`
      Need  *float64 `json:"need" form:"need" gorm:"column:need;comment:应出勤天数;size:6;"`
      Leave  *float64 `json:"leave" form:"leave" gorm:"column:leave;comment:请假天数;size:6;"`
      Absent  *float64 `json:"absent" form:"absent" gorm:"column:absent;comment:旷工天数;size:6;"`
      Cdzt  *float64 `json:"cdzt" form:"cdzt" gorm:"column:cdzt;comment:迟到早退天数;size:6;"`
      Status  string `json:"status" form:"status" gorm:"column:status;type:enum('已审核','待审核');comment:状态;"`
      Signed  *float64 `json:"signed" form:"signed" gorm:"column:signed;comment:实际签到天数;size:6;"`
      CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
      UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
      DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName OaAttendance 表名
func (OaAttendance) TableName() string {
  return "oa_attendance"
}

