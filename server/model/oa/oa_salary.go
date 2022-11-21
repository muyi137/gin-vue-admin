// 自动生成模板OaSalary
package oa

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
)

// OaSalary 结构体
type OaSalary struct {
      global.GVA_MODEL
      Basis  *float64 `json:"basis" form:"basis" gorm:"column:basis;comment:基本工资;"`
      Card_number  string `json:"card_number" form:"card_number" gorm:"column:card_number;comment:身份证号;size:20;"`
      Jobs  *float64 `json:"jobs" form:"jobs" gorm:"column:jobs;comment:岗位津贴;"`
      EduLevel  *float64 `json:"eduLevel" form:"eduLevel" gorm:"column:edu_level;comment:学历;"`
      ProTitles  *float64 `json:"proTitles" form:"proTitles" gorm:"column:pro_titles;comment:专业职称;"`
      WorkYears  *float64 `json:"workYears" form:"workYears" gorm:"column:work_years;comment:工作年限;"`
      PlanWhole  *float64 `json:"planWhole" form:"planWhole" gorm:"column:plan_whole;comment:统筹;"`
      Custom1  *float64 `json:"custom1" form:"custom1" gorm:"column:custom1;comment:;"`
      Custom2  *float64 `json:"custom2" form:"custom2" gorm:"column:custom2;comment:自定义2;"`
      Custom3  *float64 `json:"custom3" form:"custom3" gorm:"column:custom3;comment:自定义3;"`
      Custom4  *float64 `json:"custom4" form:"custom4" gorm:"column:custom4;comment:自定义4;"`
      CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
      UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
      DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName OaSalary 表名
func (OaSalary) TableName() string {
  return "oa_salary"
}

