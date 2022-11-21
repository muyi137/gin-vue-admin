package oa

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/oa"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    oaReq "github.com/flipped-aurora/gin-vue-admin/server/model/oa/request"
    "gorm.io/gorm"
)

type OaSalaryService struct {
}

// CreateOaSalary 创建OaSalary记录
// Author [piexlmax](https://github.com/piexlmax)
func (oaSalaryService *OaSalaryService) CreateOaSalary(oaSalary oa.OaSalary) (err error) {
	err = global.GVA_DB.Create(&oaSalary).Error
	return err
}

// DeleteOaSalary 删除OaSalary记录
// Author [piexlmax](https://github.com/piexlmax)
func (oaSalaryService *OaSalaryService)DeleteOaSalary(oaSalary oa.OaSalary) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&oa.OaSalary{}).Where("id = ?", oaSalary.ID).Update("deleted_by", oaSalary.DeletedBy).Error; err != nil {
              return err
        }
        if err = tx.Delete(&oaSalary).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteOaSalaryByIds 批量删除OaSalary记录
// Author [piexlmax](https://github.com/piexlmax)
func (oaSalaryService *OaSalaryService)DeleteOaSalaryByIds(ids request.IdsReq,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&oa.OaSalary{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", ids.Ids).Delete(&oa.OaSalary{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateOaSalary 更新OaSalary记录
// Author [piexlmax](https://github.com/piexlmax)
func (oaSalaryService *OaSalaryService)UpdateOaSalary(oaSalary oa.OaSalary) (err error) {
	err = global.GVA_DB.Save(&oaSalary).Error
	return err
}

// GetOaSalary 根据id获取OaSalary记录
// Author [piexlmax](https://github.com/piexlmax)
func (oaSalaryService *OaSalaryService)GetOaSalary(id uint) (oaSalary oa.OaSalary, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&oaSalary).Error
	return
}

// GetOaSalaryInfoList 分页获取OaSalary记录
// Author [piexlmax](https://github.com/piexlmax)
func (oaSalaryService *OaSalaryService)GetOaSalaryInfoList(info oaReq.OaSalarySearch) (list []oa.OaSalary, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&oa.OaSalary{})
    var oaSalarys []oa.OaSalary
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.Card_number != "" {
        db = db.Where("card_number = ?",info.Card_number)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	err = db.Limit(limit).Offset(offset).Find(&oaSalarys).Error
	return  oaSalarys, total, err
}
