import service from '@/utils/request'

// @Tags OaClassHour
// @Summary 创建OaClassHour
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.OaClassHour true "创建OaClassHour"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /oaClassHour/createOaClassHour [post]
export const createOaClassHour = (data) => {
  return service({
    url: '/oaClassHour/createOaClassHour',
    method: 'post',
    data
  })
}

// @Tags OaClassHour
// @Summary 删除OaClassHour
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.OaClassHour true "删除OaClassHour"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /oaClassHour/deleteOaClassHour [delete]
export const deleteOaClassHour = (data) => {
  return service({
    url: '/oaClassHour/deleteOaClassHour',
    method: 'delete',
    data
  })
}

// @Tags OaClassHour
// @Summary 删除OaClassHour
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除OaClassHour"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /oaClassHour/deleteOaClassHour [delete]
export const deleteOaClassHourByIds = (data) => {
  return service({
    url: '/oaClassHour/deleteOaClassHourByIds',
    method: 'delete',
    data
  })
}

// @Tags OaClassHour
// @Summary 更新OaClassHour
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.OaClassHour true "更新OaClassHour"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /oaClassHour/updateOaClassHour [put]
export const updateOaClassHour = (data) => {
  return service({
    url: '/oaClassHour/updateOaClassHour',
    method: 'put',
    data
  })
}

// @Tags OaClassHour
// @Summary 用id查询OaClassHour
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.OaClassHour true "用id查询OaClassHour"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /oaClassHour/findOaClassHour [get]
export const findOaClassHour = (params) => {
  return service({
    url: '/oaClassHour/findOaClassHour',
    method: 'get',
    params
  })
}

// @Tags OaClassHour
// @Summary 分页获取OaClassHour列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取OaClassHour列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /oaClassHour/getOaClassHourList [get]
export const getOaClassHourList = (params) => {
  return service({
    url: '/oaClassHour/getOaClassHourList',
    method: 'get',
    params
  })
}
