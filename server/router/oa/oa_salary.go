package oa

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type OaSalaryRouter struct {
}

// InitOaSalaryRouter 初始化 OaSalary 路由信息
func (s *OaSalaryRouter) InitOaSalaryRouter(Router *gin.RouterGroup) {
	oaSalaryRouter := Router.Group("oaSalary").Use(middleware.OperationRecord())
	oaSalaryRouterWithoutRecord := Router.Group("oaSalary")
	var oaSalaryApi = v1.ApiGroupApp.OaApiGroup.OaSalaryApi
	{
		oaSalaryRouter.POST("createOaSalary", oaSalaryApi.CreateOaSalary)   // 新建OaSalary
		oaSalaryRouter.DELETE("deleteOaSalary", oaSalaryApi.DeleteOaSalary) // 删除OaSalary
		oaSalaryRouter.DELETE("deleteOaSalaryByIds", oaSalaryApi.DeleteOaSalaryByIds) // 批量删除OaSalary
		oaSalaryRouter.PUT("updateOaSalary", oaSalaryApi.UpdateOaSalary)    // 更新OaSalary
	}
	{
		oaSalaryRouterWithoutRecord.GET("findOaSalary", oaSalaryApi.FindOaSalary)        // 根据ID获取OaSalary
		oaSalaryRouterWithoutRecord.GET("getOaSalaryList", oaSalaryApi.GetOaSalaryList)  // 获取OaSalary列表
	}
}
