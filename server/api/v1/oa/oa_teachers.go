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

type OaTeachersApi struct {
}

var oaTeachersService = service.ServiceGroupApp.OaServiceGroup.OaTeachersService


// CreateOaTeachers 创建OaTeachers
// @Tags OaTeachers
// @Summary 创建OaTeachers
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body oa.OaTeachers true "创建OaTeachers"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /oaTeachers/createOaTeachers [post]
func (oaTeachersApi *OaTeachersApi) CreateOaTeachers(c *gin.Context) {
	var oaTeachers oa.OaTeachers
	err := c.ShouldBindJSON(&oaTeachers)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    oaTeachers.CreatedBy = utils.GetUserID(c)
    verify := utils.Rules{
        "College":{utils.NotEmpty()},
        "CardNumber":{utils.NotEmpty()},
        "Position":{utils.NotEmpty()},
        "Qualifications":{utils.NotEmpty()},
        "Professional":{utils.NotEmpty()},
        "Hiredate":{utils.NotEmpty()},
        "Jobs":{utils.NotEmpty()},
    }
	if err := utils.Verify(oaTeachers, verify); err != nil {
    		response.FailWithMessage(err.Error(), c)
    		return
    	}
	if err := oaTeachersService.CreateOaTeachers(oaTeachers); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteOaTeachers 删除OaTeachers
// @Tags OaTeachers
// @Summary 删除OaTeachers
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body oa.OaTeachers true "删除OaTeachers"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /oaTeachers/deleteOaTeachers [delete]
func (oaTeachersApi *OaTeachersApi) DeleteOaTeachers(c *gin.Context) {
	var oaTeachers oa.OaTeachers
	err := c.ShouldBindJSON(&oaTeachers)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    oaTeachers.DeletedBy = utils.GetUserID(c)
	if err := oaTeachersService.DeleteOaTeachers(oaTeachers); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteOaTeachersByIds 批量删除OaTeachers
// @Tags OaTeachers
// @Summary 批量删除OaTeachers
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除OaTeachers"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /oaTeachers/deleteOaTeachersByIds [delete]
func (oaTeachersApi *OaTeachersApi) DeleteOaTeachersByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    deletedBy := utils.GetUserID(c)
	if err := oaTeachersService.DeleteOaTeachersByIds(IDS,deletedBy); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateOaTeachers 更新OaTeachers
// @Tags OaTeachers
// @Summary 更新OaTeachers
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body oa.OaTeachers true "更新OaTeachers"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /oaTeachers/updateOaTeachers [put]
func (oaTeachersApi *OaTeachersApi) UpdateOaTeachers(c *gin.Context) {
	var oaTeachers oa.OaTeachers
	err := c.ShouldBindJSON(&oaTeachers)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    oaTeachers.UpdatedBy = utils.GetUserID(c)
      verify := utils.Rules{
          "College":{utils.NotEmpty()},
          "CardNumber":{utils.NotEmpty()},
          "Position":{utils.NotEmpty()},
          "Qualifications":{utils.NotEmpty()},
          "Professional":{utils.NotEmpty()},
          "Hiredate":{utils.NotEmpty()},
          "Jobs":{utils.NotEmpty()},
      }
    if err := utils.Verify(oaTeachers, verify); err != nil {
      	response.FailWithMessage(err.Error(), c)
      	return
     }
	if err := oaTeachersService.UpdateOaTeachers(oaTeachers); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindOaTeachers 用id查询OaTeachers
// @Tags OaTeachers
// @Summary 用id查询OaTeachers
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query oa.OaTeachers true "用id查询OaTeachers"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /oaTeachers/findOaTeachers [get]
func (oaTeachersApi *OaTeachersApi) FindOaTeachers(c *gin.Context) {
	var oaTeachers oa.OaTeachers
	err := c.ShouldBindQuery(&oaTeachers)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reoaTeachers, err := oaTeachersService.GetOaTeachers(oaTeachers.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reoaTeachers": reoaTeachers}, c)
	}
}

// GetOaTeachersList 分页获取OaTeachers列表
// @Tags OaTeachers
// @Summary 分页获取OaTeachers列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query oaReq.OaTeachersSearch true "分页获取OaTeachers列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /oaTeachers/getOaTeachersList [get]
func (oaTeachersApi *OaTeachersApi) GetOaTeachersList(c *gin.Context) {
	var pageInfo oaReq.OaTeachersSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := oaTeachersService.GetOaTeachersInfoList(pageInfo); err != nil {
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
