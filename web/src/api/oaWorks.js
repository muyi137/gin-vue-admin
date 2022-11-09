import service from '@/utils/request'

// @Tags OaWorks
// @Summary 创建OaWorks
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.OaWorks true "创建OaWorks"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /oaWorks/createOaWorks [post]
export const createOaWorks = (data) => {
  return service({
    url: '/oaWorks/createOaWorks',
    method: 'post',
    data
  })
}

// @Tags OaWorks
// @Summary 删除OaWorks
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.OaWorks true "删除OaWorks"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /oaWorks/deleteOaWorks [delete]
export const deleteOaWorks = (data) => {
  return service({
    url: '/oaWorks/deleteOaWorks',
    method: 'delete',
    data
  })
}

// @Tags OaWorks
// @Summary 删除OaWorks
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除OaWorks"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /oaWorks/deleteOaWorks [delete]
export const deleteOaWorksByIds = (data) => {
  return service({
    url: '/oaWorks/deleteOaWorksByIds',
    method: 'delete',
    data
  })
}

// @Tags OaWorks
// @Summary 更新OaWorks
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.OaWorks true "更新OaWorks"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /oaWorks/updateOaWorks [put]
export const updateOaWorks = (data) => {
  return service({
    url: '/oaWorks/updateOaWorks',
    method: 'put',
    data
  })
}

// @Tags OaWorks
// @Summary 用id查询OaWorks
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.OaWorks true "用id查询OaWorks"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /oaWorks/findOaWorks [get]
export const findOaWorks = (params) => {
  return service({
    url: '/oaWorks/findOaWorks',
    method: 'get',
    params
  })
}

// @Tags OaWorks
// @Summary 分页获取OaWorks列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取OaWorks列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /oaWorks/getOaWorksList [get]
export const getOaWorksList = (params) => {
  return service({
    url: '/oaWorks/getOaWorksList',
    method: 'get',
    params
  })
}
