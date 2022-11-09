package oa

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type OaWorksRouter struct {
}

// InitOaWorksRouter 初始化 OaWorks 路由信息
func (s *OaWorksRouter) InitOaWorksRouter(Router *gin.RouterGroup) {
	oaWorksRouter := Router.Group("oaWorks").Use(middleware.OperationRecord())
	oaWorksRouterWithoutRecord := Router.Group("oaWorks")
	var oaWorksApi = v1.ApiGroupApp.OaApiGroup.OaWorksApi
	{
		oaWorksRouter.POST("createOaWorks", oaWorksApi.CreateOaWorks)   // 新建OaWorks
		oaWorksRouter.DELETE("deleteOaWorks", oaWorksApi.DeleteOaWorks) // 删除OaWorks
		oaWorksRouter.DELETE("deleteOaWorksByIds", oaWorksApi.DeleteOaWorksByIds) // 批量删除OaWorks
		oaWorksRouter.PUT("updateOaWorks", oaWorksApi.UpdateOaWorks)    // 更新OaWorks
	}
	{
		oaWorksRouterWithoutRecord.GET("findOaWorks", oaWorksApi.FindOaWorks)        // 根据ID获取OaWorks
		oaWorksRouterWithoutRecord.GET("getOaWorksList", oaWorksApi.GetOaWorksList)  // 获取OaWorks列表
	}
}
