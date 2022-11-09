import service from '@/utils/request'

// @Tags OaTeachers
// @Summary 创建OaTeachers
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.OaTeachers true "创建OaTeachers"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /oaTeachers/createOaTeachers [post]
export const createOaTeachers = (data) => {
  return service({
    url: '/oaTeachers/createOaTeachers',
    method: 'post',
    data
  })
}

// @Tags OaTeachers
// @Summary 删除OaTeachers
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.OaTeachers true "删除OaTeachers"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /oaTeachers/deleteOaTeachers [delete]
export const deleteOaTeachers = (data) => {
  return service({
    url: '/oaTeachers/deleteOaTeachers',
    method: 'delete',
    data
  })
}

// @Tags OaTeachers
// @Summary 删除OaTeachers
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除OaTeachers"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /oaTeachers/deleteOaTeachers [delete]
export const deleteOaTeachersByIds = (data) => {
  return service({
    url: '/oaTeachers/deleteOaTeachersByIds',
    method: 'delete',
    data
  })
}

// @Tags OaTeachers
// @Summary 更新OaTeachers
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.OaTeachers true "更新OaTeachers"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /oaTeachers/updateOaTeachers [put]
export const updateOaTeachers = (data) => {
  return service({
    url: '/oaTeachers/updateOaTeachers',
    method: 'put',
    data
  })
}

// @Tags OaTeachers
// @Summary 用id查询OaTeachers
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.OaTeachers true "用id查询OaTeachers"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /oaTeachers/findOaTeachers [get]
export const findOaTeachers = (params) => {
  return service({
    url: '/oaTeachers/findOaTeachers',
    method: 'get',
    params
  })
}

// @Tags OaTeachers
// @Summary 分页获取OaTeachers列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取OaTeachers列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /oaTeachers/getOaTeachersList [get]
export const getOaTeachersList = (params) => {
  return service({
    url: '/oaTeachers/getOaTeachersList',
    method: 'get',
    params
  })
}
