package shop

import (
    "fmt"
    "github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/example"
    "github.com/flipped-aurora/gin-vue-admin/server/model/shop"
    shopReq "github.com/flipped-aurora/gin-vue-admin/server/model/shop/request"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/flipped-aurora/gin-vue-admin/server/utils"
    "github.com/gin-gonic/gin"
    "github.com/xuri/excelize/v2"
    "go.uber.org/zap"
    "math/rand"
    "os"
    "path"
    "strconv"
    "strings"
    "time"
)

type ShopGoodsApi struct {
}

var shopGoodsService = service.ServiceGroupApp.ShopServiceGroup.ShopGoodsService
var excelService = service.ServiceGroupApp.ExampleServiceGroup.ExcelService

// CreateShopGoods 创建ShopGoods
// @Tags ShopGoods
// @Summary 创建ShopGoods
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body shop.ShopGoods true "创建ShopGoods"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /shopGoods/createShopGoods [post]
func (shopGoodsApi *ShopGoodsApi) CreateShopGoods(c *gin.Context) {
    var shopGoods shop.ShopGoods
    _ = c.ShouldBindJSON(&shopGoods)
    verify := utils.Rules{
        "GoodsName": {utils.NotEmpty()},
    }
    if err := utils.Verify(shopGoods, verify); err != nil {
        response.FailWithMessage(err.Error(), c)
        return
    }
    if err := shopGoodsService.CreateShopGoods(shopGoods); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
        response.FailWithMessage("创建失败", c)
    } else {
        response.OkWithMessage("创建成功", c)
    }
}

// DeleteShopGoods 删除ShopGoods
// @Tags ShopGoods
// @Summary 删除ShopGoods
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body shop.ShopGoods true "删除ShopGoods"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /shopGoods/deleteShopGoods [delete]
func (shopGoodsApi *ShopGoodsApi) DeleteShopGoods(c *gin.Context) {
    var shopGoods shop.ShopGoods
    _ = c.ShouldBindJSON(&shopGoods)
    if err := shopGoodsService.DeleteShopGoods(shopGoods); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
        response.FailWithMessage("删除失败", c)
    } else {
        response.OkWithMessage("删除成功", c)
    }
}

// DeleteShopGoodsByIds 批量删除ShopGoods
// @Tags ShopGoods
// @Summary 批量删除ShopGoods
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ShopGoods"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /shopGoods/deleteShopGoodsByIds [delete]
func (shopGoodsApi *ShopGoodsApi) DeleteShopGoodsByIds(c *gin.Context) {
    var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
    if err := shopGoodsService.DeleteShopGoodsByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
        response.FailWithMessage("批量删除失败", c)
    } else {
        response.OkWithMessage("批量删除成功", c)
    }
}

// UpdateShopGoods 更新ShopGoods
// @Tags ShopGoods
// @Summary 更新ShopGoods
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body shop.ShopGoods true "更新ShopGoods"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /shopGoods/updateShopGoods [put]
func (shopGoodsApi *ShopGoodsApi) UpdateShopGoods(c *gin.Context) {
    var shopGoods shop.ShopGoods
    _ = c.ShouldBindJSON(&shopGoods)
    verify := utils.Rules{
        "GoodsName": {utils.NotEmpty()},
    }
    if err := utils.Verify(shopGoods, verify); err != nil {
        response.FailWithMessage(err.Error(), c)
        return
    }
    if err := shopGoodsService.UpdateShopGoods(shopGoods); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
        response.FailWithMessage("更新失败", c)
    } else {
        response.OkWithMessage("更新成功", c)
    }
}

// FindShopGoods 用id查询ShopGoods
// @Tags ShopGoods
// @Summary 用id查询ShopGoods
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query shop.ShopGoods true "用id查询ShopGoods"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /shopGoods/findShopGoods [get]
func (shopGoodsApi *ShopGoodsApi) FindShopGoods(c *gin.Context) {
    var shopGoods shop.ShopGoods
    _ = c.ShouldBindQuery(&shopGoods)
    if reshopGoods, err := shopGoodsService.GetShopGoods(shopGoods.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
        response.FailWithMessage("查询失败", c)
    } else {
        response.OkWithData(gin.H{"reshopGoods": reshopGoods}, c)
    }
}

// GetShopGoodsList 分页获取ShopGoods列表
// @Tags ShopGoods
// @Summary 分页获取ShopGoods列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query shopReq.ShopGoodsSearch true "分页获取ShopGoods列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /shopGoods/getShopGoodsList [get]
func (shopGoodsApi *ShopGoodsApi) GetShopGoodsList(c *gin.Context) {
    var pageInfo shopReq.ShopGoodsSearch
    _ = c.ShouldBindQuery(&pageInfo)
    if list, total, err := shopGoodsService.GetShopGoodsInfoList(pageInfo); err != nil {
        global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "获取成功", c)
    }
}

// @Tags excel
// @Summary 导入Excel文件
// @Security ApiKeyAuth
// @accept multipart/form-data
// @Produce  application/json
// @Param file formData file true "导入Excel文件"
// @Success 200 {object} response.Response{msg=string} "导入Excel文件"
// @Router /excel/importExcel [post]
func (shopGoodsApi *ShopGoodsApi) ImportGoodsExcel(c *gin.Context) {
    _, header, err := c.Request.FormFile("file")
    if err != nil {
        global.GVA_LOG.Error("接收文件失败!", zap.Error(err))
        response.FailWithMessage("接收文件失败", c)
        return
    }

    t := time.Now()
    date := t.Format("200601")
    pathTmp := global.GVA_CONFIG.Excel.Dir + date + "/"
    if isDirExists(pathTmp) {
        fmt.Println("目录存在")
    } else {
        fmt.Println("目录不存在")
        err := os.Mkdir(pathTmp, 0777)
        if err != nil {
            global.GVA_LOG.Error("创建目录失败!")
        }
    }

    file_name := strconv.FormatInt(time.Now().Unix(), 10) + strconv.Itoa(rand.Intn(999999-100000)+100000) + path.Ext(header.Filename)
    _ = c.SaveUploadedFile(header, pathTmp+file_name)
    response.OkWithData(gin.H{"fileName": file_name}, c)
    //response.OkWithData(gin.H{"reshopGoods": reshopGoods}, c)
    //response.OkWithMessage("导入成功", c)
}

func isDirExists(filename string) bool {
    _, err := os.Stat(filename)
    if os.IsNotExist(err) {
        return false
    }
    return true
}

// @Tags excel
// @Summary 加载Excel数据
// @Security ApiKeyAuth
// @Produce  application/json
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "加载Excel数据,返回包括列表,总数,页码,每页数量"
// @Router /excel/loadExcel [get]
func (shopGoodsApi *ShopGoodsApi) LoadGoodsExcel(c *gin.Context) {

    var fileName = c.Query("fileName")

    menus, err := shopGoodsApi.ParseExcel2InfoList(fileName)
    if err != nil {
        global.GVA_LOG.Error("加载数据失败!", zap.Error(err))
        response.FailWithMessage("加载数据失败", c)
        return
    }
    response.OkWithDetailed(response.PageResult{
        List:     menus,
        Total:    int64(len(menus)),
        Page:     1,
        PageSize: 999,
    }, "加载数据成功", c)
}

// @Tags excel
// @Summary 导出Excel
// @Security ApiKeyAuth
// @accept application/json
// @Produce  application/octet-stream
// @Param data body example.ExcelInfo true "导出Excel文件信息"
// @Success 200
// @Router /excel/exportExcel [post]
func (shopGoodsApi *ShopGoodsApi) ExportGoodsExcel(c *gin.Context) {
    var excelInfo example.ExcelInfo
    _ = c.ShouldBindJSON(&excelInfo)
    if strings.Index(excelInfo.FileName, "..") > -1 {
        response.FailWithMessage("包含非法字符", c)
        return
    }
    filePath := global.GVA_CONFIG.Excel.Dir + excelInfo.FileName
    err := excelService.ParseInfoList2Excel(excelInfo.InfoList, filePath)
    if err != nil {
        global.GVA_LOG.Error("转换Excel失败!", zap.Error(err))
        response.FailWithMessage("转换Excel失败", c)
        return
    }
    c.Writer.Header().Add("success", "true")
    c.File(filePath)
}

func (shopGoodsApi *ShopGoodsApi) ParseExcel2InfoList(fileName string) ([]shop.GoodsExcel, error) {
    //skipHeader := true
    //fixedHeader := []string{"ID", "路由Name", "路由Path", "是否隐藏", "父节点", "排序", "文件名称"}
    t := time.Now()
    date := t.Format("200601")
    file, err := excelize.OpenFile(global.GVA_CONFIG.Excel.Dir + date + "/" + fileName)
    if err != nil {
        return nil, err
    }
    menus := make([]shop.GoodsExcel, 0)
    rows, err := file.Rows("Sheet1")
    if err != nil {
        return nil, err
    }
    for rows.Next() {
        row, err := rows.Columns()
        if err != nil {
            return nil, err
        }
        //if skipHeader {   //严格检查表格 标题
        //	if shopGoodsApi.compareStrSlice(row, fixedHeader) {
        //		skipHeader = false
        //		continue
        //	} else {
        //		return nil, errors.New("Excel格式错误")
        //	}
        //}
        //if len(row) != len(fixedHeader) {
        //	continue
        //}
        id, _ := strconv.Atoi(row[0])
        menu := shop.GoodsExcel{
            GVA_MODEL: global.GVA_MODEL{
                ID: uint(id),
            },
            Name:       row[1],
            Attr:       row[2],
            Production: row[3],
        }
        menus = append(menus, menu)
    }
    return menus, nil
}

func (shopGoodsApi *ShopGoodsApi) compareStrSlice(a, b []string) bool {
    if len(a) != len(b) {
        return false
    }
    if (b == nil) != (a == nil) {
        return false
    }
    for key, value := range a {
        if value != b[key] {
            return false
        }
    }
    return true
}
