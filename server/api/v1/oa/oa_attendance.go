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

type OaAttendanceApi struct {
}

var oaAttendanceService = service.ServiceGroupApp.OaServiceGroup.OaAttendanceService


// CreateOaAttendance 创建OaAttendance
// @Tags OaAttendance
// @Summary 创建OaAttendance
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body oa.OaAttendance true "创建OaAttendance"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /oaAttendance/createOaAttendance [post]
func (oaAttendanceApi *OaAttendanceApi) CreateOaAttendance(c *gin.Context) {
	var oaAttendance oa.OaAttendance
	err := c.ShouldBindJSON(&oaAttendance)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    oaAttendance.CreatedBy = utils.GetUserID(c)
    verify := utils.Rules{
        "CardNumber":{utils.NotEmpty()},
        "Month":{utils.NotEmpty()},
        "Need":{utils.NotEmpty()},
        "Status":{utils.NotEmpty()},
        "Signed":{utils.NotEmpty()},
    }
	if err := utils.Verify(oaAttendance, verify); err != nil {
    		response.FailWithMessage(err.Error(), c)
    		return
    	}
	if err := oaAttendanceService.CreateOaAttendance(oaAttendance); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteOaAttendance 删除OaAttendance
// @Tags OaAttendance
// @Summary 删除OaAttendance
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body oa.OaAttendance true "删除OaAttendance"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /oaAttendance/deleteOaAttendance [delete]
func (oaAttendanceApi *OaAttendanceApi) DeleteOaAttendance(c *gin.Context) {
	var oaAttendance oa.OaAttendance
	err := c.ShouldBindJSON(&oaAttendance)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    oaAttendance.DeletedBy = utils.GetUserID(c)
	if err := oaAttendanceService.DeleteOaAttendance(oaAttendance); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteOaAttendanceByIds 批量删除OaAttendance
// @Tags OaAttendance
// @Summary 批量删除OaAttendance
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除OaAttendance"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /oaAttendance/deleteOaAttendanceByIds [delete]
func (oaAttendanceApi *OaAttendanceApi) DeleteOaAttendanceByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    deletedBy := utils.GetUserID(c)
	if err := oaAttendanceService.DeleteOaAttendanceByIds(IDS,deletedBy); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateOaAttendance 更新OaAttendance
// @Tags OaAttendance
// @Summary 更新OaAttendance
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body oa.OaAttendance true "更新OaAttendance"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /oaAttendance/updateOaAttendance [put]
func (oaAttendanceApi *OaAttendanceApi) UpdateOaAttendance(c *gin.Context) {
	var oaAttendance oa.OaAttendance
	err := c.ShouldBindJSON(&oaAttendance)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    oaAttendance.UpdatedBy = utils.GetUserID(c)
      verify := utils.Rules{
          "CardNumber":{utils.NotEmpty()},
          "Month":{utils.NotEmpty()},
          "Need":{utils.NotEmpty()},
          "Status":{utils.NotEmpty()},
          "Signed":{utils.NotEmpty()},
      }
    if err := utils.Verify(oaAttendance, verify); err != nil {
      	response.FailWithMessage(err.Error(), c)
      	return
     }
	if err := oaAttendanceService.UpdateOaAttendance(oaAttendance); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindOaAttendance 用id查询OaAttendance
// @Tags OaAttendance
// @Summary 用id查询OaAttendance
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query oa.OaAttendance true "用id查询OaAttendance"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /oaAttendance/findOaAttendance [get]
func (oaAttendanceApi *OaAttendanceApi) FindOaAttendance(c *gin.Context) {
	var oaAttendance oa.OaAttendance
	err := c.ShouldBindQuery(&oaAttendance)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reoaAttendance, err := oaAttendanceService.GetOaAttendance(oaAttendance.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reoaAttendance": reoaAttendance}, c)
	}
}

// GetOaAttendanceList 分页获取OaAttendance列表
// @Tags OaAttendance
// @Summary 分页获取OaAttendance列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query oaReq.OaAttendanceSearch true "分页获取OaAttendance列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /oaAttendance/getOaAttendanceList [get]
func (oaAttendanceApi *OaAttendanceApi) GetOaAttendanceList(c *gin.Context) {
	var pageInfo oaReq.OaAttendanceSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := oaAttendanceService.GetOaAttendanceInfoList(pageInfo); err != nil {
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
