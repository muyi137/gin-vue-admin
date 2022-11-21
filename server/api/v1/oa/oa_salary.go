package oa

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/oa"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    oaReq "github.com/flipped-aurora/gin-vue-admin/server/model/oa/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type OaSalaryApi struct {
}

var oaSalaryService = service.ServiceGroupApp.OaServiceGroup.OaSalaryService


// CreateOaSalary 创建OaSalary
// @Tags OaSalary
// @Summary 创建OaSalary
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body oa.OaSalary true "创建OaSalary"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /oaSalary/createOaSalary [post]
func (oaSalaryApi *OaSalaryApi) CreateOaSalary(c *gin.Context) {
	var oaSalary oa.OaSalary
	err := c.ShouldBindJSON(&oaSalary)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    oaSalary.CreatedBy = utils.GetUserID(c)
    verify := utils.Rules{
        "Basis":{utils.NotEmpty()},
        "Card_number":{utils.NotEmpty()},
    }
	if err := utils.Verify(oaSalary, verify); err != nil {
    		response.FailWithMessage(err.Error(), c)
    		return
    	}
	if err := oaSalaryService.CreateOaSalary(oaSalary); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteOaSalary 删除OaSalary
// @Tags OaSalary
// @Summary 删除OaSalary
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body oa.OaSalary true "删除OaSalary"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /oaSalary/deleteOaSalary [delete]
func (oaSalaryApi *OaSalaryApi) DeleteOaSalary(c *gin.Context) {
	var oaSalary oa.OaSalary
	err := c.ShouldBindJSON(&oaSalary)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    oaSalary.DeletedBy = utils.GetUserID(c)
	if err := oaSalaryService.DeleteOaSalary(oaSalary); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteOaSalaryByIds 批量删除OaSalary
// @Tags OaSalary
// @Summary 批量删除OaSalary
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除OaSalary"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /oaSalary/deleteOaSalaryByIds [delete]
func (oaSalaryApi *OaSalaryApi) DeleteOaSalaryByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    deletedBy := utils.GetUserID(c)
	if err := oaSalaryService.DeleteOaSalaryByIds(IDS,deletedBy); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateOaSalary 更新OaSalary
// @Tags OaSalary
// @Summary 更新OaSalary
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body oa.OaSalary true "更新OaSalary"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /oaSalary/updateOaSalary [put]
func (oaSalaryApi *OaSalaryApi) UpdateOaSalary(c *gin.Context) {
	var oaSalary oa.OaSalary
	err := c.ShouldBindJSON(&oaSalary)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    oaSalary.UpdatedBy = utils.GetUserID(c)
      verify := utils.Rules{
          "Basis":{utils.NotEmpty()},
          "Card_number":{utils.NotEmpty()},
      }
    if err := utils.Verify(oaSalary, verify); err != nil {
      	response.FailWithMessage(err.Error(), c)
      	return
     }
	if err := oaSalaryService.UpdateOaSalary(oaSalary); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindOaSalary 用id查询OaSalary
// @Tags OaSalary
// @Summary 用id查询OaSalary
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query oa.OaSalary true "用id查询OaSalary"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /oaSalary/findOaSalary [get]
func (oaSalaryApi *OaSalaryApi) FindOaSalary(c *gin.Context) {
	var oaSalary oa.OaSalary
	err := c.ShouldBindQuery(&oaSalary)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reoaSalary, err := oaSalaryService.GetOaSalary(oaSalary.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reoaSalary": reoaSalary}, c)
	}
}

// GetOaSalaryList 分页获取OaSalary列表
// @Tags OaSalary
// @Summary 分页获取OaSalary列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query oaReq.OaSalarySearch true "分页获取OaSalary列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /oaSalary/getOaSalaryList [get]
func (oaSalaryApi *OaSalaryApi) GetOaSalaryList(c *gin.Context) {
	var pageInfo oaReq.OaSalarySearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := oaSalaryService.GetOaSalaryInfoList(pageInfo); err != nil {
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
