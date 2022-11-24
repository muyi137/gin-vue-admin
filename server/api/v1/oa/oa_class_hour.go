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

type OaClassHourApi struct {
}

var oaClassHourService = service.ServiceGroupApp.OaServiceGroup.OaClassHourService


// CreateOaClassHour 创建OaClassHour
// @Tags OaClassHour
// @Summary 创建OaClassHour
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body oa.OaClassHour true "创建OaClassHour"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /oaClassHour/createOaClassHour [post]
func (oaClassHourApi *OaClassHourApi) CreateOaClassHour(c *gin.Context) {
	var oaClassHour oa.OaClassHour
	err := c.ShouldBindJSON(&oaClassHour)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    oaClassHour.CreatedBy = utils.GetUserID(c)
    verify := utils.Rules{
        "Card_number":{utils.NotEmpty()},
        "ClassTime":{utils.NotEmpty()},
        "ReportTime":{utils.NotEmpty()},
    }
	if err := utils.Verify(oaClassHour, verify); err != nil {
    		response.FailWithMessage(err.Error(), c)
    		return
    	}
	if err := oaClassHourService.CreateOaClassHour(oaClassHour); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteOaClassHour 删除OaClassHour
// @Tags OaClassHour
// @Summary 删除OaClassHour
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body oa.OaClassHour true "删除OaClassHour"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /oaClassHour/deleteOaClassHour [delete]
func (oaClassHourApi *OaClassHourApi) DeleteOaClassHour(c *gin.Context) {
	var oaClassHour oa.OaClassHour
	err := c.ShouldBindJSON(&oaClassHour)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    oaClassHour.DeletedBy = utils.GetUserID(c)
	if err := oaClassHourService.DeleteOaClassHour(oaClassHour); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteOaClassHourByIds 批量删除OaClassHour
// @Tags OaClassHour
// @Summary 批量删除OaClassHour
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除OaClassHour"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /oaClassHour/deleteOaClassHourByIds [delete]
func (oaClassHourApi *OaClassHourApi) DeleteOaClassHourByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    deletedBy := utils.GetUserID(c)
	if err := oaClassHourService.DeleteOaClassHourByIds(IDS,deletedBy); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateOaClassHour 更新OaClassHour
// @Tags OaClassHour
// @Summary 更新OaClassHour
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body oa.OaClassHour true "更新OaClassHour"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /oaClassHour/updateOaClassHour [put]
func (oaClassHourApi *OaClassHourApi) UpdateOaClassHour(c *gin.Context) {
	var oaClassHour oa.OaClassHour
	err := c.ShouldBindJSON(&oaClassHour)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    oaClassHour.UpdatedBy = utils.GetUserID(c)
      verify := utils.Rules{
          "Card_number":{utils.NotEmpty()},
          "ClassTime":{utils.NotEmpty()},
          "ReportTime":{utils.NotEmpty()},
      }
    if err := utils.Verify(oaClassHour, verify); err != nil {
      	response.FailWithMessage(err.Error(), c)
      	return
     }
	if err := oaClassHourService.UpdateOaClassHour(oaClassHour); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindOaClassHour 用id查询OaClassHour
// @Tags OaClassHour
// @Summary 用id查询OaClassHour
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query oa.OaClassHour true "用id查询OaClassHour"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /oaClassHour/findOaClassHour [get]
func (oaClassHourApi *OaClassHourApi) FindOaClassHour(c *gin.Context) {
	var oaClassHour oa.OaClassHour
	err := c.ShouldBindQuery(&oaClassHour)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reoaClassHour, err := oaClassHourService.GetOaClassHour(oaClassHour.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reoaClassHour": reoaClassHour}, c)
	}
}

// GetOaClassHourList 分页获取OaClassHour列表
// @Tags OaClassHour
// @Summary 分页获取OaClassHour列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query oaReq.OaClassHourSearch true "分页获取OaClassHour列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /oaClassHour/getOaClassHourList [get]
func (oaClassHourApi *OaClassHourApi) GetOaClassHourList(c *gin.Context) {
	var pageInfo oaReq.OaClassHourSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := oaClassHourService.GetOaClassHourInfoList(pageInfo); err != nil {
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
