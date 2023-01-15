import service from '@/utils/request'

// @Tags TestStrutc
// @Summary 创建TestStrutc
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TestStrutc true "创建TestStrutc"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /testStrutc/createTestStrutc [post]
export const createTestStrutc = (data) => {
  return service({
    url: '/testStrutc/createTestStrutc',
    method: 'post',
    data
  })
}

// @Tags TestStrutc
// @Summary 删除TestStrutc
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TestStrutc true "删除TestStrutc"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /testStrutc/deleteTestStrutc [delete]
export const deleteTestStrutc = (data) => {
  return service({
    url: '/testStrutc/deleteTestStrutc',
    method: 'delete',
    data
  })
}

// @Tags TestStrutc
// @Summary 删除TestStrutc
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除TestStrutc"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /testStrutc/deleteTestStrutc [delete]
export const deleteTestStrutcByIds = (data) => {
  return service({
    url: '/testStrutc/deleteTestStrutcByIds',
    method: 'delete',
    data
  })
}

// @Tags TestStrutc
// @Summary 更新TestStrutc
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TestStrutc true "更新TestStrutc"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /testStrutc/updateTestStrutc [put]
export const updateTestStrutc = (data) => {
  return service({
    url: '/testStrutc/updateTestStrutc',
    method: 'put',
    data
  })
}

// @Tags TestStrutc
// @Summary 用id查询TestStrutc
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.TestStrutc true "用id查询TestStrutc"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /testStrutc/findTestStrutc [get]
export const findTestStrutc = (params) => {
  return service({
    url: '/testStrutc/findTestStrutc',
    method: 'get',
    params
  })
}

// @Tags TestStrutc
// @Summary 分页获取TestStrutc列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取TestStrutc列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /testStrutc/getTestStrutcList [get]
export const getTestStrutcList = (params) => {
  return service({
    url: '/testStrutc/getTestStrutcList',
    method: 'get',
    params
  })
}
