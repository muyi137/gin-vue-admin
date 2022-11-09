package oa

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type OaAttendanceRouter struct {
}

// InitOaAttendanceRouter 初始化 OaAttendance 路由信息
func (s *OaAttendanceRouter) InitOaAttendanceRouter(Router *gin.RouterGroup) {
	oaAttendanceRouter := Router.Group("oaAttendance").Use(middleware.OperationRecord())
	oaAttendanceRouterWithoutRecord := Router.Group("oaAttendance")
	var oaAttendanceApi = v1.ApiGroupApp.OaApiGroup.OaAttendanceApi
	{
		oaAttendanceRouter.POST("createOaAttendance", oaAttendanceApi.CreateOaAttendance)   // 新建OaAttendance
		oaAttendanceRouter.DELETE("deleteOaAttendance", oaAttendanceApi.DeleteOaAttendance) // 删除OaAttendance
		oaAttendanceRouter.DELETE("deleteOaAttendanceByIds", oaAttendanceApi.DeleteOaAttendanceByIds) // 批量删除OaAttendance
		oaAttendanceRouter.PUT("updateOaAttendance", oaAttendanceApi.UpdateOaAttendance)    // 更新OaAttendance
	}
	{
		oaAttendanceRouterWithoutRecord.GET("findOaAttendance", oaAttendanceApi.FindOaAttendance)        // 根据ID获取OaAttendance
		oaAttendanceRouterWithoutRecord.GET("getOaAttendanceList", oaAttendanceApi.GetOaAttendanceList)  // 获取OaAttendance列表
	}
}
