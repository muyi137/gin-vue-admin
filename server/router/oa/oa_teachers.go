package oa

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type OaTeachersRouter struct {
}

// InitOaTeachersRouter 初始化 OaTeachers 路由信息
func (s *OaTeachersRouter) InitOaTeachersRouter(Router *gin.RouterGroup) {
	oaTeachersRouter := Router.Group("oaTeachers").Use(middleware.OperationRecord())
	oaTeachersRouterWithoutRecord := Router.Group("oaTeachers")
	var oaTeachersApi = v1.ApiGroupApp.OaApiGroup.OaTeachersApi
	{
		oaTeachersRouter.POST("createOaTeachers", oaTeachersApi.CreateOaTeachers)   // 新建OaTeachers
		oaTeachersRouter.DELETE("deleteOaTeachers", oaTeachersApi.DeleteOaTeachers) // 删除OaTeachers
		oaTeachersRouter.DELETE("deleteOaTeachersByIds", oaTeachersApi.DeleteOaTeachersByIds) // 批量删除OaTeachers
		oaTeachersRouter.PUT("updateOaTeachers", oaTeachersApi.UpdateOaTeachers)    // 更新OaTeachers
	}
	{
		oaTeachersRouterWithoutRecord.GET("findOaTeachers", oaTeachersApi.FindOaTeachers)        // 根据ID获取OaTeachers
		oaTeachersRouterWithoutRecord.GET("getOaTeachersList", oaTeachersApi.GetOaTeachersList)  // 获取OaTeachers列表
	}
}
