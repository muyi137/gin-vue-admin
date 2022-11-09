package oa

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/oa"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    oaReq "github.com/flipped-aurora/gin-vue-admin/server/model/oa/request"
    "gorm.io/gorm"
)

type OaTeachersService struct {
}

// CreateOaTeachers 创建OaTeachers记录
// Author [piexlmax](https://github.com/piexlmax)
func (oaTeachersService *OaTeachersService) CreateOaTeachers(oaTeachers oa.OaTeachers) (err error) {
	err = global.GVA_DB.Create(&oaTeachers).Error
	return err
}

// DeleteOaTeachers 删除OaTeachers记录
// Author [piexlmax](https://github.com/piexlmax)
func (oaTeachersService *OaTeachersService)DeleteOaTeachers(oaTeachers oa.OaTeachers) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&oa.OaTeachers{}).Where("id = ?", oaTeachers.ID).Update("deleted_by", oaTeachers.DeletedBy).Error; err != nil {
              return err
        }
        if err = tx.Delete(&oaTeachers).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteOaTeachersByIds 批量删除OaTeachers记录
// Author [piexlmax](https://github.com/piexlmax)
func (oaTeachersService *OaTeachersService)DeleteOaTeachersByIds(ids request.IdsReq,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&oa.OaTeachers{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", ids.Ids).Delete(&oa.OaTeachers{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateOaTeachers 更新OaTeachers记录
// Author [piexlmax](https://github.com/piexlmax)
func (oaTeachersService *OaTeachersService)UpdateOaTeachers(oaTeachers oa.OaTeachers) (err error) {
	err = global.GVA_DB.Save(&oaTeachers).Error
	return err
}

// GetOaTeachers 根据id获取OaTeachers记录
// Author [piexlmax](https://github.com/piexlmax)
func (oaTeachersService *OaTeachersService)GetOaTeachers(id uint) (oaTeachers oa.OaTeachers, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&oaTeachers).Error
	return
}

// GetOaTeachersInfoList 分页获取OaTeachers记录
// Author [piexlmax](https://github.com/piexlmax)
func (oaTeachersService *OaTeachersService)GetOaTeachersInfoList(info oaReq.OaTeachersSearch) (list []oa.OaTeachers, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&oa.OaTeachers{})
    var oaTeacherss []oa.OaTeachers
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.College != nil {
        db = db.Where("college = ?",info.College)
    }
    if info.Name != "" {
        db = db.Where("name LIKE ?","%"+ info.Name+"%")
    }
    if info.Qualifications != nil {
        db = db.Where("qualifications = ?",info.Qualifications)
    }
    if info.Jobs != nil {
        db = db.Where("jobs = ?",info.Jobs)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	err = db.Limit(limit).Offset(offset).Find(&oaTeacherss).Error
	return  oaTeacherss, total, err
}
