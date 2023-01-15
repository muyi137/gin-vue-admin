package shop

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/shop"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    shopReq "github.com/flipped-aurora/gin-vue-admin/server/model/shop/request"
)

type TestService struct {
}

// CreateTest 创建Test记录
// Author [piexlmax](https://github.com/piexlmax)
func (testService *TestService) CreateTest(test shop.Test) (err error) {
	err = global.GVA_DB.Create(&test).Error
	return err
}

// DeleteTest 删除Test记录
// Author [piexlmax](https://github.com/piexlmax)
func (testService *TestService)DeleteTest(test shop.Test) (err error) {
	err = global.GVA_DB.Delete(&test).Error
	return err
}

// DeleteTestByIds 批量删除Test记录
// Author [piexlmax](https://github.com/piexlmax)
func (testService *TestService)DeleteTestByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]shop.Test{},"id in ?",ids.Ids).Error
	return err
}

// UpdateTest 更新Test记录
// Author [piexlmax](https://github.com/piexlmax)
func (testService *TestService)UpdateTest(test shop.Test) (err error) {
	err = global.GVA_DB.Save(&test).Error
	return err
}

// GetTest 根据id获取Test记录
// Author [piexlmax](https://github.com/piexlmax)
func (testService *TestService)GetTest(id uint) (test shop.Test, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&test).Error
	return
}

// GetTestInfoList 分页获取Test记录
// Author [piexlmax](https://github.com/piexlmax)
func (testService *TestService)GetTestInfoList(info shopReq.TestSearch) (list []shop.Test, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&shop.Test{})
    var tests []shop.Test
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&tests).Error
	return  tests, total, err
}
