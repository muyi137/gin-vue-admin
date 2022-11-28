package oa

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type OaExcelRouter struct{}

func (e *OaExcelRouter) InitOaExcelRouter(Router *gin.RouterGroup) {
	OaExcelRouter := Router.Group("oaExcel")
	OaExcelApi := v1.ApiGroupApp.OaApiGroup.OaExcelApi

	{
		OaExcelRouter.POST("oaImportExcel", OaExcelApi.ImportExcel) // 导入Excel
		//excelRouter.GET("oaLoadExcel", exaExcelApi.LoadExcel)               // 加载Excel数据
		OaExcelRouter.POST("oaExportExcel", OaExcelApi.ExportExcel)          // 导出Excel
		OaExcelRouter.GET("oaDownloadTemplate", OaExcelApi.DownloadTemplate) // 下载模板文件
	}
}
