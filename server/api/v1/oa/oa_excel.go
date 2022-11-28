package oa

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/example"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"go.uber.org/zap"
	"os"
	"strings"
	"time"
)

type OaExcelApi struct{}

// ImportExcel
// @Tags      excel
// @Summary   导入Excel文件
// @Security  ApiKeyAuth
// @accept    multipart/form-data
// @Produce   application/json
// @Param     file  formData  file                           true  "导入Excel文件"
// @Success   200   {object}  response.Response{msg=string}  "导入Excel文件"
// @Router    /Oaexcel/importExcel [post]
func (e *OaExcelApi) ImportExcel(c *gin.Context) {

	_, header, err := c.Request.FormFile("file")

	if err != nil {
		global.GVA_LOG.Error("接收文件失败!", zap.Error(err))
		response.FailWithMessage("接收文件失败", c)
		return
	}

	excelType := c.PostForm("type")

	switch {
	case excelType == "appraisal":
		fmt.Printf("优秀!\n")
	case excelType == "B":
		fmt.Printf("良好\n")
	case excelType == "D":
		fmt.Printf("及格\n")
	default:
		fmt.Printf("差\n")
	}

	xlsxPath := global.GVA_CONFIG.Local.Path + "/" + excelType + "/" + cast.ToString(time.Now().Unix()) + ".xlsx"

	_ = c.SaveUploadedFile(header, xlsxPath)

	//打开文件路径
	//xlsxFile, err := xlsx.OpenFile(xlsxPath)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}

	response.OkWithMessage("导入成功", c)
}

// ExportExcel
// @Tags      excel
// @Summary   导出Excel
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/octet-stream
// @Param     data  body  example.ExcelInfo  true  "导出Excel文件信息"
// @Success   200
// @Router    /excel/exportExcel [post]
func (e *OaExcelApi) ExportExcel(c *gin.Context) {
	var excelInfo example.ExcelInfo
	err := c.ShouldBindJSON(&excelInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if strings.Index(excelInfo.FileName, "..") > -1 {
		response.FailWithMessage("包含非法字符", c)
		return
	}
	filePath := global.GVA_CONFIG.Excel.Dir + excelInfo.FileName
	//err = excelService.ParseInfoList2Excel(excelInfo.InfoList, filePath)
	//if err != nil {
	//    global.GVA_LOG.Error("转换Excel失败!", zap.Error(err))
	//    response.FailWithMessage("转换Excel失败", c)
	//    return
	//}
	c.Writer.Header().Add("success", "true")
	c.File(filePath)
}

// DownloadTemplate
// @Tags      excel
// @Summary   下载模板
// @Security  ApiKeyAuth
// @accept    multipart/form-data
// @Produce   application/json
// @Param     fileName  query  string  true  "模板名称"
// @Success   200
// @Router    /excel/downloadTemplate [get]
func (e *OaExcelApi) DownloadTemplate(c *gin.Context) {
	fileName := c.Query("fileName")
	filePath := global.GVA_CONFIG.Excel.Dir + fileName

	fi, err := os.Stat(filePath)
	if err != nil {
		global.GVA_LOG.Error("文件不存在!", zap.Error(err))
		response.FailWithMessage("文件不存在", c)
		return
	}
	if fi.IsDir() {
		global.GVA_LOG.Error("不支持下载文件夹!", zap.Error(err))
		response.FailWithMessage("不支持下载文件夹", c)
		return
	}
	c.Writer.Header().Add("success", "true")
	c.File(filePath)
}
