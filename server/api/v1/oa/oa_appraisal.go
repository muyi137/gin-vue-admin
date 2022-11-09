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

type OaAppraisalApi struct {
}

var oaAppraisalService = service.ServiceGroupApp.OaServiceGroup.OaAppraisalService


// CreateOaAppraisal 创建OaAppraisal
// @Tags OaAppraisal
// @Summary 创建OaAppraisal
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body oa.OaAppraisal true "创建OaAppraisal"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /oaAppraisal/createOaAppraisal [post]
func (oaAppraisalApi *OaAppraisalApi) CreateOaAppraisal(c *gin.Context) {
	var oaAppraisal oa.OaAppraisal
	err := c.ShouldBindJSON(&oaAppraisal)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    oaAppraisal.CreatedBy = utils.GetUserID(c)
    verify := utils.Rules{
        "CardNumber":{utils.NotEmpty()},
        "Month":{utils.NotEmpty()},
        "Score":{utils.NotEmpty()},
    }
	if err := utils.Verify(oaAppraisal, verify); err != nil {
    		response.FailWithMessage(err.Error(), c)
    		return
    	}
	if err := oaAppraisalService.CreateOaAppraisal(oaAppraisal); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteOaAppraisal 删除OaAppraisal
// @Tags OaAppraisal
// @Summary 删除OaAppraisal
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body oa.OaAppraisal true "删除OaAppraisal"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /oaAppraisal/deleteOaAppraisal [delete]
func (oaAppraisalApi *OaAppraisalApi) DeleteOaAppraisal(c *gin.Context) {
	var oaAppraisal oa.OaAppraisal
	err := c.ShouldBindJSON(&oaAppraisal)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    oaAppraisal.DeletedBy = utils.GetUserID(c)
	if err := oaAppraisalService.DeleteOaAppraisal(oaAppraisal); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteOaAppraisalByIds 批量删除OaAppraisal
// @Tags OaAppraisal
// @Summary 批量删除OaAppraisal
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除OaAppraisal"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /oaAppraisal/deleteOaAppraisalByIds [delete]
func (oaAppraisalApi *OaAppraisalApi) DeleteOaAppraisalByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    deletedBy := utils.GetUserID(c)
	if err := oaAppraisalService.DeleteOaAppraisalByIds(IDS,deletedBy); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateOaAppraisal 更新OaAppraisal
// @Tags OaAppraisal
// @Summary 更新OaAppraisal
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body oa.OaAppraisal true "更新OaAppraisal"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /oaAppraisal/updateOaAppraisal [put]
func (oaAppraisalApi *OaAppraisalApi) UpdateOaAppraisal(c *gin.Context) {
	var oaAppraisal oa.OaAppraisal
	err := c.ShouldBindJSON(&oaAppraisal)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    oaAppraisal.UpdatedBy = utils.GetUserID(c)
      verify := utils.Rules{
          "CardNumber":{utils.NotEmpty()},
          "Month":{utils.NotEmpty()},
          "Score":{utils.NotEmpty()},
      }
    if err := utils.Verify(oaAppraisal, verify); err != nil {
      	response.FailWithMessage(err.Error(), c)
      	return
     }
	if err := oaAppraisalService.UpdateOaAppraisal(oaAppraisal); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindOaAppraisal 用id查询OaAppraisal
// @Tags OaAppraisal
// @Summary 用id查询OaAppraisal
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query oa.OaAppraisal true "用id查询OaAppraisal"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /oaAppraisal/findOaAppraisal [get]
func (oaAppraisalApi *OaAppraisalApi) FindOaAppraisal(c *gin.Context) {
	var oaAppraisal oa.OaAppraisal
	err := c.ShouldBindQuery(&oaAppraisal)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reoaAppraisal, err := oaAppraisalService.GetOaAppraisal(oaAppraisal.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reoaAppraisal": reoaAppraisal}, c)
	}
}

// GetOaAppraisalList 分页获取OaAppraisal列表
// @Tags OaAppraisal
// @Summary 分页获取OaAppraisal列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query oaReq.OaAppraisalSearch true "分页获取OaAppraisal列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /oaAppraisal/getOaAppraisalList [get]
func (oaAppraisalApi *OaAppraisalApi) GetOaAppraisalList(c *gin.Context) {
	var pageInfo oaReq.OaAppraisalSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := oaAppraisalService.GetOaAppraisalInfoList(pageInfo); err != nil {
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
