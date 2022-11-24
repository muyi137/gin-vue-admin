package oa

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/oa"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    oaReq "github.com/flipped-aurora/gin-vue-admin/server/model/oa/request"
    "gorm.io/gorm"
)

type OaClassHourService struct {
}

// CreateOaClassHour 创建OaClassHour记录
// Author [piexlmax](https://github.com/piexlmax)
func (oaClassHourService *OaClassHourService) CreateOaClassHour(oaClassHour oa.OaClassHour) (err error) {
	err = global.GVA_DB.Create(&oaClassHour).Error
	return err
}

// DeleteOaClassHour 删除OaClassHour记录
// Author [piexlmax](https://github.com/piexlmax)
func (oaClassHourService *OaClassHourService)DeleteOaClassHour(oaClassHour oa.OaClassHour) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&oa.OaClassHour{}).Where("id = ?", oaClassHour.ID).Update("deleted_by", oaClassHour.DeletedBy).Error; err != nil {
              return err
        }
        if err = tx.Delete(&oaClassHour).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteOaClassHourByIds 批量删除OaClassHour记录
// Author [piexlmax](https://github.com/piexlmax)
func (oaClassHourService *OaClassHourService)DeleteOaClassHourByIds(ids request.IdsReq,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&oa.OaClassHour{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", ids.Ids).Delete(&oa.OaClassHour{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateOaClassHour 更新OaClassHour记录
// Author [piexlmax](https://github.com/piexlmax)
func (oaClassHourService *OaClassHourService)UpdateOaClassHour(oaClassHour oa.OaClassHour) (err error) {
	err = global.GVA_DB.Save(&oaClassHour).Error
	return err
}

// GetOaClassHour 根据id获取OaClassHour记录
// Author [piexlmax](https://github.com/piexlmax)
func (oaClassHourService *OaClassHourService)GetOaClassHour(id uint) (oaClassHour oa.OaClassHour, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&oaClassHour).Error
	return
}

// GetOaClassHourInfoList 分页获取OaClassHour记录
// Author [piexlmax](https://github.com/piexlmax)
func (oaClassHourService *OaClassHourService)GetOaClassHourInfoList(info oaReq.OaClassHourSearch) (list []oa.OaClassHour, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&oa.OaClassHour{})
    var oaClassHours []oa.OaClassHour
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.Card_number != "" {
        db = db.Where("card_number = ?",info.Card_number)
    }
        if info.StartReportTime != nil && info.EndReportTime != nil {
            db = db.Where("report_time BETWEEN ? AND ? ",info.StartReportTime,info.EndReportTime)
        }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	err = db.Limit(limit).Offset(offset).Find(&oaClassHours).Error
	return  oaClassHours, total, err
}
