package oa

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/oa"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    oaReq "github.com/flipped-aurora/gin-vue-admin/server/model/oa/request"
    "gorm.io/gorm"
)

type OaAppraisalService struct {
}

// CreateOaAppraisal 创建OaAppraisal记录
// Author [piexlmax](https://github.com/piexlmax)
func (oaAppraisalService *OaAppraisalService) CreateOaAppraisal(oaAppraisal oa.OaAppraisal) (err error) {
	err = global.GVA_DB.Create(&oaAppraisal).Error
	return err
}

// DeleteOaAppraisal 删除OaAppraisal记录
// Author [piexlmax](https://github.com/piexlmax)
func (oaAppraisalService *OaAppraisalService)DeleteOaAppraisal(oaAppraisal oa.OaAppraisal) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&oa.OaAppraisal{}).Where("id = ?", oaAppraisal.ID).Update("deleted_by", oaAppraisal.DeletedBy).Error; err != nil {
              return err
        }
        if err = tx.Delete(&oaAppraisal).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteOaAppraisalByIds 批量删除OaAppraisal记录
// Author [piexlmax](https://github.com/piexlmax)
func (oaAppraisalService *OaAppraisalService)DeleteOaAppraisalByIds(ids request.IdsReq,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&oa.OaAppraisal{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", ids.Ids).Delete(&oa.OaAppraisal{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateOaAppraisal 更新OaAppraisal记录
// Author [piexlmax](https://github.com/piexlmax)
func (oaAppraisalService *OaAppraisalService)UpdateOaAppraisal(oaAppraisal oa.OaAppraisal) (err error) {
	err = global.GVA_DB.Save(&oaAppraisal).Error
	return err
}

// GetOaAppraisal 根据id获取OaAppraisal记录
// Author [piexlmax](https://github.com/piexlmax)
func (oaAppraisalService *OaAppraisalService)GetOaAppraisal(id uint) (oaAppraisal oa.OaAppraisal, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&oaAppraisal).Error
	return
}

// GetOaAppraisalInfoList 分页获取OaAppraisal记录
// Author [piexlmax](https://github.com/piexlmax)
func (oaAppraisalService *OaAppraisalService)GetOaAppraisalInfoList(info oaReq.OaAppraisalSearch) (list []oa.OaAppraisal, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&oa.OaAppraisal{})
    var oaAppraisals []oa.OaAppraisal
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.CardNumber != "" {
        db = db.Where("card_number = ?",info.CardNumber)
    }
    if info.Month != "" {
        db = db.Where("month = ?",info.Month)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	err = db.Limit(limit).Offset(offset).Find(&oaAppraisals).Error
	return  oaAppraisals, total, err
}
