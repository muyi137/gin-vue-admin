package oa

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/oa"
	oaReq "github.com/flipped-aurora/gin-vue-admin/server/model/oa/request"
	"gorm.io/gorm"
)

type OaAttendanceService struct {
}

// CreateOaAttendance 创建OaAttendance记录
// Author [piexlmax](https://github.com/piexlmax)
func (oaAttendanceService *OaAttendanceService) CreateOaAttendance(oaAttendance oa.OaAttendance) (err error) {
	err = global.GVA_DB.Create(&oaAttendance).Error
	return err
}

// DeleteOaAttendance 删除OaAttendance记录
// Author [piexlmax](https://github.com/piexlmax)
func (oaAttendanceService *OaAttendanceService) DeleteOaAttendance(oaAttendance oa.OaAttendance) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&oa.OaAttendance{}).Where("id = ?", oaAttendance.ID).Update("deleted_by", oaAttendance.DeletedBy).Error; err != nil {
			return err
		}
		if err = tx.Delete(&oaAttendance).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteOaAttendanceByIds 批量删除OaAttendance记录
// Author [piexlmax](https://github.com/piexlmax)
func (oaAttendanceService *OaAttendanceService) DeleteOaAttendanceByIds(ids request.IdsReq, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&oa.OaAttendance{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", ids.Ids).Delete(&oa.OaAttendance{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateOaAttendance 更新OaAttendance记录
// Author [piexlmax](https://github.com/piexlmax)
func (oaAttendanceService *OaAttendanceService) UpdateOaAttendance(oaAttendance oa.OaAttendance) (err error) {
	err = global.GVA_DB.Save(&oaAttendance).Error
	return err
}

// GetOaAttendance 根据id获取OaAttendance记录
// Author [piexlmax](https://github.com/piexlmax)
func (oaAttendanceService *OaAttendanceService) GetOaAttendance(id uint) (oaAttendance oa.OaAttendance, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&oaAttendance).Error
	return
}

// GetOaAttendanceInfoList 分页获取OaAttendance记录
// Author [piexlmax](https://github.com/piexlmax)
func (oaAttendanceService *OaAttendanceService) GetOaAttendanceInfoList(info oaReq.OaAttendanceSearch) (list []oa.OaAttendance, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&oa.OaAttendance{})
	var oaAttendances []oa.OaAttendance
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.CardNumber != "" {
		db = db.Where("card_number = ?", info.CardNumber)
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

	err = db.Limit(limit).Offset(offset).Find(&oaAttendances).Error
	return oaAttendances, total, err
}
