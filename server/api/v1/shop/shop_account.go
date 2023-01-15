package shop

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/shop"
	shopReq "github.com/flipped-aurora/gin-vue-admin/server/model/shop/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

type AccountApi struct {
}

var accountService = service.ServiceGroupApp.ShopServiceGroup.AccountService

// CreateAccount 创建Account
// @Tags Account
// @Summary 创建Account
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body shop.Account true "创建Account"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /account/createAccount [post]
func (accountApi *AccountApi) CreateAccount(c *gin.Context) {
	var account shop.Account
	_ = c.ShouldBindJSON(&account)
	verify := utils.Rules{
		"Username": {utils.NotEmpty()},
		"Password": {utils.NotEmpty()},
	}
	if err := utils.Verify(account, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := accountService.CreateAccount(account); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteAccount 删除Account
// @Tags Account
// @Summary 删除Account
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body shop.Account true "删除Account"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /account/deleteAccount [delete]
func (accountApi *AccountApi) DeleteAccount(c *gin.Context) {
	var account shop.Account
	_ = c.ShouldBindJSON(&account)
	if err := accountService.DeleteAccount(account); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteAccountByIds 批量删除Account
// @Tags Account
// @Summary 批量删除Account
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Account"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /account/deleteAccountByIds [delete]
func (accountApi *AccountApi) DeleteAccountByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := accountService.DeleteAccountByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateAccount 更新Account
// @Tags Account
// @Summary 更新Account
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body shop.Account true "更新Account"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /account/updateAccount [put]
func (accountApi *AccountApi) UpdateAccount(c *gin.Context) {
	var account shop.Account
	_ = c.ShouldBindJSON(&account)
	verify := utils.Rules{
		"Username": {utils.NotEmpty()},
		"Password": {utils.NotEmpty()},
	}
	if err := utils.Verify(account, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := accountService.UpdateAccount(account); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindAccount 用id查询Account
// @Tags Account
// @Summary 用id查询Account
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query shop.Account true "用id查询Account"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /account/findAccount [get]
func (accountApi *AccountApi) FindAccount(c *gin.Context) {
	var account shop.Account
	_ = c.ShouldBindQuery(&account)
	if reaccount, err := accountService.GetAccount(account.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reaccount": reaccount}, c)
	}
}

// GetAccountList 分页获取Account列表
// @Tags Account
// @Summary 分页获取Account列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query shopReq.AccountSearch true "分页获取Account列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /account/getAccountList [get]
func (accountApi *AccountApi) GetAccountList(c *gin.Context) {
	var pageInfo shopReq.AccountSearch
	_ = c.ShouldBindQuery(&pageInfo)

	userId, err := strconv.Atoi(c.Request.Header.Get("x-user-id"))
	if err == nil {
		if userId != 1 {
			pageInfo.User_id = &userId
		}
	}

	if list, total, err := accountService.GetAccountInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
