package shop

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/shop"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    shopReq "github.com/flipped-aurora/gin-vue-admin/server/model/shop/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type TestStrutcApi struct {
}

var testStrutcService = service.ServiceGroupApp.ShopServiceGroup.TestStrutcService


// CreateTestStrutc 创建TestStrutc
// @Tags TestStrutc
// @Summary 创建TestStrutc
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body shop.TestStrutc true "创建TestStrutc"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /testStrutc/createTestStrutc [post]
func (testStrutcApi *TestStrutcApi) CreateTestStrutc(c *gin.Context) {
	var testStrutc shop.TestStrutc
	_ = c.ShouldBindJSON(&testStrutc)
	if err := testStrutcService.CreateTestStrutc(testStrutc); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteTestStrutc 删除TestStrutc
// @Tags TestStrutc
// @Summary 删除TestStrutc
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body shop.TestStrutc true "删除TestStrutc"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /testStrutc/deleteTestStrutc [delete]
func (testStrutcApi *TestStrutcApi) DeleteTestStrutc(c *gin.Context) {
	var testStrutc shop.TestStrutc
	_ = c.ShouldBindJSON(&testStrutc)
	if err := testStrutcService.DeleteTestStrutc(testStrutc); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteTestStrutcByIds 批量删除TestStrutc
// @Tags TestStrutc
// @Summary 批量删除TestStrutc
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除TestStrutc"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /testStrutc/deleteTestStrutcByIds [delete]
func (testStrutcApi *TestStrutcApi) DeleteTestStrutcByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := testStrutcService.DeleteTestStrutcByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateTestStrutc 更新TestStrutc
// @Tags TestStrutc
// @Summary 更新TestStrutc
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body shop.TestStrutc true "更新TestStrutc"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /testStrutc/updateTestStrutc [put]
func (testStrutcApi *TestStrutcApi) UpdateTestStrutc(c *gin.Context) {
	var testStrutc shop.TestStrutc
	_ = c.ShouldBindJSON(&testStrutc)
	if err := testStrutcService.UpdateTestStrutc(testStrutc); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindTestStrutc 用id查询TestStrutc
// @Tags TestStrutc
// @Summary 用id查询TestStrutc
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query shop.TestStrutc true "用id查询TestStrutc"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /testStrutc/findTestStrutc [get]
func (testStrutcApi *TestStrutcApi) FindTestStrutc(c *gin.Context) {
	var testStrutc shop.TestStrutc
	_ = c.ShouldBindQuery(&testStrutc)
	if retestStrutc, err := testStrutcService.GetTestStrutc(testStrutc.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"retestStrutc": retestStrutc}, c)
	}
}

// GetTestStrutcList 分页获取TestStrutc列表
// @Tags TestStrutc
// @Summary 分页获取TestStrutc列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query shopReq.TestStrutcSearch true "分页获取TestStrutc列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /testStrutc/getTestStrutcList [get]
func (testStrutcApi *TestStrutcApi) GetTestStrutcList(c *gin.Context) {
	var pageInfo shopReq.TestStrutcSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := testStrutcService.GetTestStrutcInfoList(pageInfo); err != nil {
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
