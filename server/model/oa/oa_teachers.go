// 自动生成模板OaTeachers
package oa

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// OaTeachers 结构体
type OaTeachers struct {
	global.GVA_MODEL
	College        *int       `json:"college" form:"college" gorm:"column:college;comment:学院;"`
	Department     string     `json:"department" form:"department" gorm:"column:department;comment:;"`
	Name           string     `json:"name" form:"name" gorm:"column:name;comment:;"`
	Gender         *int       `json:"gender" form:"gender" gorm:"column:gender;comment:;"`
	CardNumber     string     `json:"cardNumber" form:"cardNumber" gorm:"column:card_number;comment:;"`
	Position       *int       `json:"position" form:"position" gorm:"column:position;comment:;"`
	Qualifications *int       `json:"qualifications" form:"qualifications" gorm:"column:qualifications;comment:;size:2;"`
	Professional   *int       `json:"professional" form:"professional" gorm:"column:professional;comment:;"`
	Hiredate       *time.Time `json:"hiredate" form:"hiredate" gorm:"column:hiredate;comment:来校时间;"`
	Jobs           *int       `json:"jobs" form:"jobs" gorm:"column:jobs;comment:;"`
	CreatedBy      uint       `gorm:"column:created_by;comment:创建者"`
	UpdatedBy      uint       `gorm:"column:updated_by;comment:更新者"`
	DeletedBy      uint       `gorm:"column:deleted_by;comment:删除者"`
}

// TableName OaTeachers 表名
func (OaTeachers) TableName() string {
	return "oa_teachers"
}
