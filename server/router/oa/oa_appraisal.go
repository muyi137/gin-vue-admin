package oa

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type OaAppraisalRouter struct {
}

// InitOaAppraisalRouter 初始化 OaAppraisal 路由信息
func (s *OaAppraisalRouter) InitOaAppraisalRouter(Router *gin.RouterGroup) {
	oaAppraisalRouter := Router.Group("oaAppraisal").Use(middleware.OperationRecord())
	oaAppraisalRouterWithoutRecord := Router.Group("oaAppraisal")
	var oaAppraisalApi = v1.ApiGroupApp.OaApiGroup.OaAppraisalApi
	{
		oaAppraisalRouter.POST("createOaAppraisal", oaAppraisalApi.CreateOaAppraisal)             // 新建OaAppraisal
		oaAppraisalRouter.DELETE("deleteOaAppraisal", oaAppraisalApi.DeleteOaAppraisal)           // 删除OaAppraisal
		oaAppraisalRouter.DELETE("deleteOaAppraisalByIds", oaAppraisalApi.DeleteOaAppraisalByIds) // 批量删除OaAppraisal
		oaAppraisalRouter.PUT("updateOaAppraisal", oaAppraisalApi.UpdateOaAppraisal)              // 更新OaAppraisal

	}
	{
		oaAppraisalRouterWithoutRecord.POST("importOaAppraisal", oaAppraisalApi.ImportOaAppraisal)  // 导入
		oaAppraisalRouterWithoutRecord.GET("testOaAppraisal", oaAppraisalApi.TestOaAppraisal)       // 导入
		oaAppraisalRouterWithoutRecord.GET("findOaAppraisal", oaAppraisalApi.FindOaAppraisal)       // 根据ID获取OaAppraisal
		oaAppraisalRouterWithoutRecord.GET("getOaAppraisalList", oaAppraisalApi.GetOaAppraisalList) // 获取OaAppraisal列表
	}
}
