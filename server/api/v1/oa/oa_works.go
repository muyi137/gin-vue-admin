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

type OaWorksApi struct {
}

var oaWorksService = service.ServiceGroupApp.OaServiceGroup.OaWorksService


// CreateOaWorks 创建OaWorks
// @Tags OaWorks
// @Summary 创建OaWorks
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body oa.OaWorks true "创建OaWorks"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /oaWorks/createOaWorks [post]
func (oaWorksApi *OaWorksApi) CreateOaWorks(c *gin.Context) {
	var oaWorks oa.OaWorks
	err := c.ShouldBindJSON(&oaWorks)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    oaWorks.CreatedBy = utils.GetUserID(c)
    verify := utils.Rules{
        "CardNumbaer":{utils.NotEmpty()},
        "Month":{utils.NotEmpty()},
        "Score":{utils.NotEmpty()},
    }
	if err := utils.Verify(oaWorks, verify); err != nil {
    		response.FailWithMessage(err.Error(), c)
    		return
    	}
	if err := oaWorksService.CreateOaWorks(oaWorks); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteOaWorks 删除OaWorks
// @Tags OaWorks
// @Summary 删除OaWorks
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body oa.OaWorks true "删除OaWorks"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /oaWorks/deleteOaWorks [delete]
func (oaWorksApi *OaWorksApi) DeleteOaWorks(c *gin.Context) {
	var oaWorks oa.OaWorks
	err := c.ShouldBindJSON(&oaWorks)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    oaWorks.DeletedBy = utils.GetUserID(c)
	if err := oaWorksService.DeleteOaWorks(oaWorks); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteOaWorksByIds 批量删除OaWorks
// @Tags OaWorks
// @Summary 批量删除OaWorks
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除OaWorks"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /oaWorks/deleteOaWorksByIds [delete]
func (oaWorksApi *OaWorksApi) DeleteOaWorksByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    deletedBy := utils.GetUserID(c)
	if err := oaWorksService.DeleteOaWorksByIds(IDS,deletedBy); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateOaWorks 更新OaWorks
// @Tags OaWorks
// @Summary 更新OaWorks
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body oa.OaWorks true "更新OaWorks"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /oaWorks/updateOaWorks [put]
func (oaWorksApi *OaWorksApi) UpdateOaWorks(c *gin.Context) {
	var oaWorks oa.OaWorks
	err := c.ShouldBindJSON(&oaWorks)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    oaWorks.UpdatedBy = utils.GetUserID(c)
      verify := utils.Rules{
          "CardNumbaer":{utils.NotEmpty()},
          "Month":{utils.NotEmpty()},
          "Score":{utils.NotEmpty()},
      }
    if err := utils.Verify(oaWorks, verify); err != nil {
      	response.FailWithMessage(err.Error(), c)
      	return
     }
	if err := oaWorksService.UpdateOaWorks(oaWorks); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindOaWorks 用id查询OaWorks
// @Tags OaWorks
// @Summary 用id查询OaWorks
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query oa.OaWorks true "用id查询OaWorks"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /oaWorks/findOaWorks [get]
func (oaWorksApi *OaWorksApi) FindOaWorks(c *gin.Context) {
	var oaWorks oa.OaWorks
	err := c.ShouldBindQuery(&oaWorks)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reoaWorks, err := oaWorksService.GetOaWorks(oaWorks.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reoaWorks": reoaWorks}, c)
	}
}

// GetOaWorksList 分页获取OaWorks列表
// @Tags OaWorks
// @Summary 分页获取OaWorks列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query oaReq.OaWorksSearch true "分页获取OaWorks列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /oaWorks/getOaWorksList [get]
func (oaWorksApi *OaWorksApi) GetOaWorksList(c *gin.Context) {
	var pageInfo oaReq.OaWorksSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := oaWorksService.GetOaWorksInfoList(pageInfo); err != nil {
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
