import service from '@/utils/request'

// @Tags ShopGoods
// @Summary 创建ShopGoods
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ShopGoods true "创建ShopGoods"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /shopGoods/createShopGoods [post]
export const createShopGoods = (data) => {
    return service({
        url: '/shopGoods/createShopGoods',
        method: 'post',
        data
    })
}

// @Tags ShopGoods
// @Summary 删除ShopGoods
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ShopGoods true "删除ShopGoods"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /shopGoods/deleteShopGoods [delete]
export const deleteShopGoods = (data) => {
    return service({
        url: '/shopGoods/deleteShopGoods',
        method: 'delete',
        data
    })
}

// @Tags ShopGoods
// @Summary 删除ShopGoods
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ShopGoods"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /shopGoods/deleteShopGoods [delete]
export const deleteShopGoodsByIds = (data) => {
    return service({
        url: '/shopGoods/deleteShopGoodsByIds',
        method: 'delete',
        data
    })
}

// @Tags ShopGoods
// @Summary 更新ShopGoods
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ShopGoods true "更新ShopGoods"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /shopGoods/updateShopGoods [put]
export const updateShopGoods = (data) => {
    return service({
        url: '/shopGoods/updateShopGoods',
        method: 'put',
        data
    })
}

// @Tags ShopGoods
// @Summary 用id查询ShopGoods
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.ShopGoods true "用id查询ShopGoods"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /shopGoods/findShopGoods [get]
export const findShopGoods = (params) => {
    return service({
        url: '/shopGoods/findShopGoods',
        method: 'get',
        params
    })
}

// @Tags ShopGoods
// @Summary 分页获取ShopGoods列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取ShopGoods列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /shopGoods/getShopGoodsList [get]
export const getShopGoodsList = (params) => {
    return service({
        url: '/shopGoods/getShopGoodsList',
        method: 'get',
        params
    })
}