// 自动生成模板Account
package shop

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
)

// Account 结构体
type Account struct {
      global.GVA_MODEL
      Username  string `json:"username" form:"username" gorm:"column:username;comment:用户名;"`
      Password  string `json:"password" form:"password" gorm:"column:password;comment:密码;"`
      Mall_id  *int `json:"mall_id" form:"mall_id" gorm:"column:mall_id;comment:商城;"`
      User_id  *int `json:"user_id" form:"user_id" gorm:"column:user_id;comment:用户id;"`
}


// TableName Account 表名
func (Account) TableName() string {
  return "shop_accounts"
}

