import service from '@/utils/request'

// @Tags OaAttendance
// @Summary 创建OaAttendance
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.OaAttendance true "创建OaAttendance"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /oaAttendance/createOaAttendance [post]
export const createOaAttendance = (data) => {
  return service({
    url: '/oaAttendance/createOaAttendance',
    method: 'post',
    data
  })
}

// @Tags OaAttendance
// @Summary 删除OaAttendance
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.OaAttendance true "删除OaAttendance"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /oaAttendance/deleteOaAttendance [delete]
export const deleteOaAttendance = (data) => {
  return service({
    url: '/oaAttendance/deleteOaAttendance',
    method: 'delete',
    data
  })
}

// @Tags OaAttendance
// @Summary 删除OaAttendance
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除OaAttendance"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /oaAttendance/deleteOaAttendance [delete]
export const deleteOaAttendanceByIds = (data) => {
  return service({
    url: '/oaAttendance/deleteOaAttendanceByIds',
    method: 'delete',
    data
  })
}

// @Tags OaAttendance
// @Summary 更新OaAttendance
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.OaAttendance true "更新OaAttendance"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /oaAttendance/updateOaAttendance [put]
export const updateOaAttendance = (data) => {
  return service({
    url: '/oaAttendance/updateOaAttendance',
    method: 'put',
    data
  })
}

// @Tags OaAttendance
// @Summary 用id查询OaAttendance
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.OaAttendance true "用id查询OaAttendance"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /oaAttendance/findOaAttendance [get]
export const findOaAttendance = (params) => {
  return service({
    url: '/oaAttendance/findOaAttendance',
    method: 'get',
    params
  })
}

// @Tags OaAttendance
// @Summary 分页获取OaAttendance列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取OaAttendance列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /oaAttendance/getOaAttendanceList [get]
export const getOaAttendanceList = (params) => {
  return service({
    url: '/oaAttendance/getOaAttendanceList',
    method: 'get',
    params
  })
}
