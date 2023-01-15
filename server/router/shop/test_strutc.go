package shop

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type TestStrutcRouter struct {
}

// InitTestStrutcRouter 初始化 TestStrutc 路由信息
func (s *TestStrutcRouter) InitTestStrutcRouter(Router *gin.RouterGroup) {
	testStrutcRouter := Router.Group("testStrutc").Use(middleware.OperationRecord())
	testStrutcRouterWithoutRecord := Router.Group("testStrutc")
	var testStrutcApi = v1.ApiGroupApp.ShopApiGroup.TestStrutcApi
	{
		testStrutcRouter.POST("createTestStrutc", testStrutcApi.CreateTestStrutc)   // 新建TestStrutc
		testStrutcRouter.DELETE("deleteTestStrutc", testStrutcApi.DeleteTestStrutc) // 删除TestStrutc
		testStrutcRouter.DELETE("deleteTestStrutcByIds", testStrutcApi.DeleteTestStrutcByIds) // 批量删除TestStrutc
		testStrutcRouter.PUT("updateTestStrutc", testStrutcApi.UpdateTestStrutc)    // 更新TestStrutc
	}
	{
		testStrutcRouterWithoutRecord.GET("findTestStrutc", testStrutcApi.FindTestStrutc)        // 根据ID获取TestStrutc
		testStrutcRouterWithoutRecord.GET("getTestStrutcList", testStrutcApi.GetTestStrutcList)  // 获取TestStrutc列表
	}
}
