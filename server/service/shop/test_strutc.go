package shop

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/shop"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    shopReq "github.com/flipped-aurora/gin-vue-admin/server/model/shop/request"
)

type TestStrutcService struct {
}

// CreateTestStrutc 创建TestStrutc记录
// Author [piexlmax](https://github.com/piexlmax)
func (testStrutcService *TestStrutcService) CreateTestStrutc(testStrutc shop.TestStrutc) (err error) {
	err = global.GVA_DB.Create(&testStrutc).Error
	return err
}

// DeleteTestStrutc 删除TestStrutc记录
// Author [piexlmax](https://github.com/piexlmax)
func (testStrutcService *TestStrutcService)DeleteTestStrutc(testStrutc shop.TestStrutc) (err error) {
	err = global.GVA_DB.Delete(&testStrutc).Error
	return err
}

// DeleteTestStrutcByIds 批量删除TestStrutc记录
// Author [piexlmax](https://github.com/piexlmax)
func (testStrutcService *TestStrutcService)DeleteTestStrutcByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]shop.TestStrutc{},"id in ?",ids.Ids).Error
	return err
}

// UpdateTestStrutc 更新TestStrutc记录
// Author [piexlmax](https://github.com/piexlmax)
func (testStrutcService *TestStrutcService)UpdateTestStrutc(testStrutc shop.TestStrutc) (err error) {
	err = global.GVA_DB.Save(&testStrutc).Error
	return err
}

// GetTestStrutc 根据id获取TestStrutc记录
// Author [piexlmax](https://github.com/piexlmax)
func (testStrutcService *TestStrutcService)GetTestStrutc(id uint) (testStrutc shop.TestStrutc, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&testStrutc).Error
	return
}

// GetTestStrutcInfoList 分页获取TestStrutc记录
// Author [piexlmax](https://github.com/piexlmax)
func (testStrutcService *TestStrutcService)GetTestStrutcInfoList(info shopReq.TestStrutcSearch) (list []shop.TestStrutc, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&shop.TestStrutc{})
    var testStrutcs []shop.TestStrutc
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&testStrutcs).Error
	return  testStrutcs, total, err
}
