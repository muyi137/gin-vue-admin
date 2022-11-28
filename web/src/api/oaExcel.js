import service from '@/utils/request'
import { ElMessage } from 'element-plus'

const handleFileError = (res, fileName) => {
    if (typeof(res.data) !== 'undefined') {
        if (res.data.type === 'application/json') {
            const reader = new FileReader()
            reader.onload = function() {
                const message = JSON.parse(reader.result).msg
                ElMessage({
                    showClose: true,
                    message: message,
                    type: 'error'
                })
            }
            reader.readAsText(new Blob([res.data]))
        }
    } else {
        var downloadUrl = window.URL.createObjectURL(new Blob([res]))
        var a = document.createElement('a')
        a.style.display = 'none'
        a.href = downloadUrl
        a.download = fileName
        var event = new MouseEvent('click')
        a.dispatchEvent(event)
    }
}

// @Tags excel
// @Summary 导出Excel
// @Security ApiKeyAuth
// @accept application/json
// @Produce  application/octet-stream
// @Param data body model.ExcelInfo true "导出Excel文件信息"
// @Success 200
// @Router /oaExcel/OaExportExcel [post]
export const OaExportExcel = (tableData, fileName) => {
    service({
        url: '/excel/exportExcel',
        method: 'post',
        data: {
            fileName: fileName,
            infoList: tableData
        },
        responseType: 'blob'
    }).then((res) => {
        handleFileError(res, fileName)
    })
}

// @Tags excel
// @Summary 导入Excel文件
// @Security ApiKeyAuth
// @accept multipart/form-data
// @Produce  application/json
// @Param file formData file true "导入Excel文件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"导入成功"}"
// @Router /oaExcel/oaImportExcel [post]
export const oaImportExcel = () => {
        return service({
            url: '/oaExcel/oaImportExcel',
            method: 'post'
        })
    }
    // @Tags excel
    // @Summary 下载模板
    // @Security ApiKeyAuth
    // @accept multipart/form-data
    // @Produce  application/json
    // @Param fileName query fileName true "模板名称"
    // @Success 200
    // @Router /OaExcel/OaDownloadTemplatec [get]
export const OaDownloadTemplatec = (fileName) => {
    return service({
        url: '/excel/downloadTemplate',
        method: 'get',
        params: {
            fileName: fileName
        },
        responseType: 'blob'
    }).then((res) => {
        handleFileError(res, fileName)
    })
}

// @Tags excel
// @Summary 下载模板
// @Security ApiKeyAuth
// @accept multipart/form-data
// @Produce  application/json
// @Param fileName query fileName true "模板名称"
// @Success 200
// @Router /excel/downloadTemplate [get]
export const downloadTemplate = (fileName) => {
    return service({
        url: '/excel/downloadTemplate',
        method: 'get',
        params: {
            fileName: fileName
        },
        responseType: 'blob'
    }).then((res) => {
        handleFileError(res, fileName)
    })
}