import service from '@/utils/request'

// @Tags OaSalary
// @Summary 创建OaSalary
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.OaSalary true "创建OaSalary"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /oaSalary/createOaSalary [post]
export const createOaSalary = (data) => {
  return service({
    url: '/oaSalary/createOaSalary',
    method: 'post',
    data
  })
}

// @Tags OaSalary
// @Summary 删除OaSalary
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.OaSalary true "删除OaSalary"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /oaSalary/deleteOaSalary [delete]
export const deleteOaSalary = (data) => {
  return service({
    url: '/oaSalary/deleteOaSalary',
    method: 'delete',
    data
  })
}

// @Tags OaSalary
// @Summary 删除OaSalary
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除OaSalary"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /oaSalary/deleteOaSalary [delete]
export const deleteOaSalaryByIds = (data) => {
  return service({
    url: '/oaSalary/deleteOaSalaryByIds',
    method: 'delete',
    data
  })
}

// @Tags OaSalary
// @Summary 更新OaSalary
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.OaSalary true "更新OaSalary"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /oaSalary/updateOaSalary [put]
export const updateOaSalary = (data) => {
  return service({
    url: '/oaSalary/updateOaSalary',
    method: 'put',
    data
  })
}

// @Tags OaSalary
// @Summary 用id查询OaSalary
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.OaSalary true "用id查询OaSalary"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /oaSalary/findOaSalary [get]
export const findOaSalary = (params) => {
  return service({
    url: '/oaSalary/findOaSalary',
    method: 'get',
    params
  })
}

// @Tags OaSalary
// @Summary 分页获取OaSalary列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取OaSalary列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /oaSalary/getOaSalaryList [get]
export const getOaSalaryList = (params) => {
  return service({
    url: '/oaSalary/getOaSalaryList',
    method: 'get',
    params
  })
}
