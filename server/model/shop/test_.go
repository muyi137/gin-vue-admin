// 自动生成模板Test
package shop

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
)

// Test 结构体
type Test struct {
      global.GVA_MODEL
      Name  string `json:"name" form:"name" gorm:"column:name;comment:;"`
}


// TableName Test 表名
func (Test) TableName() string {
  return "test"
}

