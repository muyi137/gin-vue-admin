// 自动生成模板TestStrutc
package shop

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
)

// TestStrutc 结构体
type TestStrutc struct {
      global.GVA_MODEL
      Testf  string `json:"testf" form:"testf" gorm:"column:testf;comment:;"`
}


// TableName TestStrutc 表名
func (TestStrutc) TableName() string {
  return "test_strutc"
}

