package oa

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/oa"
	oaReq "github.com/flipped-aurora/gin-vue-admin/server/model/oa/request"
	"gorm.io/gorm"
)

type OaWorksService struct {
}

// CreateOaWorks 创建OaWorks记录
// Author [piexlmax](https://github.com/piexlmax)
func (oaWorksService *OaWorksService) CreateOaWorks(oaWorks oa.OaWorks) (err error) {
	err = global.GVA_DB.Create(&oaWorks).Error
	return err
}

// DeleteOaWorks 删除OaWorks记录
// Author [piexlmax](https://github.com/piexlmax)
func (oaWorksService *OaWorksService) DeleteOaWorks(oaWorks oa.OaWorks) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&oa.OaWorks{}).Where("id = ?", oaWorks.ID).Update("deleted_by", oaWorks.DeletedBy).Error; err != nil {
			return err
		}
		if err = tx.Delete(&oaWorks).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteOaWorksByIds 批量删除OaWorks记录
// Author [piexlmax](https://github.com/piexlmax)
func (oaWorksService *OaWorksService) DeleteOaWorksByIds(ids request.IdsReq, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&oa.OaWorks{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", ids.Ids).Delete(&oa.OaWorks{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateOaWorks 更新OaWorks记录
// Author [piexlmax](https://github.com/piexlmax)
func (oaWorksService *OaWorksService) UpdateOaWorks(oaWorks oa.OaWorks) (err error) {
	err = global.GVA_DB.Save(&oaWorks).Error
	return err
}

// GetOaWorks 根据id获取OaWorks记录
// Author [piexlmax](https://github.com/piexlmax)
func (oaWorksService *OaWorksService) GetOaWorks(id uint) (oaWorks oa.OaWorks, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&oaWorks).Error
	return
}

// GetOaWorksInfoList 分页获取OaWorks记录
// Author [piexlmax](https://github.com/piexlmax)
func (oaWorksService *OaWorksService) GetOaWorksInfoList(info oaReq.OaWorksSearch) (list []oa.OaWorks, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&oa.OaWorks{})
	var oaWorkss []oa.OaWorks
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.CardNumbaer != "" {
		db = db.Where("card_numbaer = ?", info.CardNumbaer)
	}
	if info.Month != nil {
		db = db.Where("month = ?", info.Month)
	}
	if info.Status != "" {
		db = db.Where("status = ?", info.Status)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&oaWorkss).Error
	return oaWorkss, total, err
}
