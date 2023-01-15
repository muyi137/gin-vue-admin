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
// @Router /excel/exportExcel [post]
export const exportExcel = (tableData, fileName) => {
    service({
        url: '/shopGoods/exportGoodsExcel',
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
// @Router /excel/importExcel [post]
export const loadExcelData = (fileName) => {
    console.log(fileName)
    return service({
        url: '/shopGoods/loadGoodsExcel?fileName=' + fileName,
        method: 'get'
    })
}