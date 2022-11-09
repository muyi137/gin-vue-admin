import service from '@/utils/request'

// @Tags OaAppraisal
// @Summary 创建OaAppraisal
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.OaAppraisal true "创建OaAppraisal"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /oaAppraisal/createOaAppraisal [post]
export const createOaAppraisal = (data) => {
  return service({
    url: '/oaAppraisal/createOaAppraisal',
    method: 'post',
    data
  })
}

// @Tags OaAppraisal
// @Summary 删除OaAppraisal
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.OaAppraisal true "删除OaAppraisal"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /oaAppraisal/deleteOaAppraisal [delete]
export const deleteOaAppraisal = (data) => {
  return service({
    url: '/oaAppraisal/deleteOaAppraisal',
    method: 'delete',
    data
  })
}

// @Tags OaAppraisal
// @Summary 删除OaAppraisal
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除OaAppraisal"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /oaAppraisal/deleteOaAppraisal [delete]
export const deleteOaAppraisalByIds = (data) => {
  return service({
    url: '/oaAppraisal/deleteOaAppraisalByIds',
    method: 'delete',
    data
  })
}

// @Tags OaAppraisal
// @Summary 更新OaAppraisal
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.OaAppraisal true "更新OaAppraisal"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /oaAppraisal/updateOaAppraisal [put]
export const updateOaAppraisal = (data) => {
  return service({
    url: '/oaAppraisal/updateOaAppraisal',
    method: 'put',
    data
  })
}

// @Tags OaAppraisal
// @Summary 用id查询OaAppraisal
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.OaAppraisal true "用id查询OaAppraisal"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /oaAppraisal/findOaAppraisal [get]
export const findOaAppraisal = (params) => {
  return service({
    url: '/oaAppraisal/findOaAppraisal',
    method: 'get',
    params
  })
}

// @Tags OaAppraisal
// @Summary 分页获取OaAppraisal列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取OaAppraisal列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /oaAppraisal/getOaAppraisalList [get]
export const getOaAppraisalList = (params) => {
  return service({
    url: '/oaAppraisal/getOaAppraisalList',
    method: 'get',
    params
  })
}
