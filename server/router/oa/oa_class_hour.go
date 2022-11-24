package oa

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type OaClassHourRouter struct {
}

// InitOaClassHourRouter 初始化 OaClassHour 路由信息
func (s *OaClassHourRouter) InitOaClassHourRouter(Router *gin.RouterGroup) {
	oaClassHourRouter := Router.Group("oaClassHour").Use(middleware.OperationRecord())
	oaClassHourRouterWithoutRecord := Router.Group("oaClassHour")
	var oaClassHourApi = v1.ApiGroupApp.OaApiGroup.OaClassHourApi
	{
		oaClassHourRouter.POST("createOaClassHour", oaClassHourApi.CreateOaClassHour)   // 新建OaClassHour
		oaClassHourRouter.DELETE("deleteOaClassHour", oaClassHourApi.DeleteOaClassHour) // 删除OaClassHour
		oaClassHourRouter.DELETE("deleteOaClassHourByIds", oaClassHourApi.DeleteOaClassHourByIds) // 批量删除OaClassHour
		oaClassHourRouter.PUT("updateOaClassHour", oaClassHourApi.UpdateOaClassHour)    // 更新OaClassHour
	}
	{
		oaClassHourRouterWithoutRecord.GET("findOaClassHour", oaClassHourApi.FindOaClassHour)        // 根据ID获取OaClassHour
		oaClassHourRouterWithoutRecord.GET("getOaClassHourList", oaClassHourApi.GetOaClassHourList)  // 获取OaClassHour列表
	}
}
