package oa

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

//type OaExcelRouter struct{}
//
//func (s *OaExcelRouter) InitOaExcelRouter(Router *gin.RouterGroup) {
//	oaExcelRouter := Router.Group("oaExcel")
//	var OaExcelApi = v1.ApiGroupApp.OaApiGroup.OaExcelApi
//	{
//		oaExcelRouter.POST("oaImportExcel", OaExcelApi.OaImportExcel) // 导入Excel
//		//excelRouter.GET("oaLoadExcel", exaExcelApi.LoadExcel)               // 加载Excel数据
//		oaExcelRouter.POST("oaExportExcel", OaExcelApi.OaExportExcel)          // 导出Excel
//		oaExcelRouter.GET("oaDownloadTemplate", OaExcelApi.OaDownloadTemplate) // 下载模板文件
//	}
//}

type OaExcelRouter struct {
}

// InitOaClassHourRouter 初始化 OaClassHour 路由信息
func (s *OaExcelRouter) InitOaExcelRouter(Router *gin.RouterGroup) {

	oaExcelRouterWithoutRecord := Router.Group("oaExcel")
	var oaExcelApi = v1.ApiGroupApp.OaApiGroup.OaExcelApi
	{
		//oaExcelRouterWithoutRecord.GET("findOaClassHour", oaExcelApi.FindOaClassHour)        // 根据ID获取OaClassHour
		oaExcelRouterWithoutRecord.GET("oaDownloadTemplate", oaExcelApi.OaDownloadTemplate) // 获取OaClassHour列表
	}
}
