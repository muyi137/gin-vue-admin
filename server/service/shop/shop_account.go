package shop

import (
    "github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/shop"
    shopReq "github.com/flipped-aurora/gin-vue-admin/server/model/shop/request"
)

type AccountService struct {
}

// CreateAccount 创建Account记录
// Author [piexlmax](https://github.com/piexlmax)
func (accountService *AccountService) CreateAccount(account shop.Account) (err error) {
    err = global.GVA_DB.Create(&account).Error
    return err
}

// DeleteAccount 删除Account记录
// Author [piexlmax](https://github.com/piexlmax)
func (accountService *AccountService) DeleteAccount(account shop.Account) (err error) {
    err = global.GVA_DB.Delete(&account).Error
    return err
}

// DeleteAccountByIds 批量删除Account记录
// Author [piexlmax](https://github.com/piexlmax)
func (accountService *AccountService) DeleteAccountByIds(ids request.IdsReq) (err error) {
    err = global.GVA_DB.Delete(&[]shop.Account{}, "id in ?", ids.Ids).Error
    return err
}

// UpdateAccount 更新Account记录
// Author [piexlmax](https://github.com/piexlmax)
func (accountService *AccountService) UpdateAccount(account shop.Account) (err error) {
    err = global.GVA_DB.Save(&account).Error
    return err
}

// GetAccount 根据id获取Account记录
// Author [piexlmax](https://github.com/piexlmax)
func (accountService *AccountService) GetAccount(id uint) (account shop.Account, err error) {
    err = global.GVA_DB.Where("id = ?", id).First(&account).Error
    return
}

// GetAccountInfoList 分页获取Account记录
// Author [piexlmax](https://github.com/piexlmax)
func (accountService *AccountService) GetAccountInfoList(info shopReq.AccountSearch) (list []shop.Account, total int64, err error) {
    limit := info.PageSize
    offset := info.PageSize * (info.Page - 1)
    // 创建db
    db := global.GVA_DB.Model(&shop.Account{})
    var accounts []shop.Account
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.Username != "" {
        db = db.Where("username = ?", info.Username)
    }

    if info.User_id != nil {
        db = db.Where("user_id = ?", info.User_id)
    }

    if info.Mall_id != nil {
        db = db.Where("mall_id = ?", info.Mall_id)
    }
    err = db.Count(&total).Error
    if err != nil {
        return
    }
    err = db.Limit(limit).Offset(offset).Find(&accounts).Error
    return accounts, total, err
}
