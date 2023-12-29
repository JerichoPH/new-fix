package apiRout

import (
	"new-fix/controllers"
	"new-fix/middlewares"

	"github.com/gin-gonic/gin"
)

// BreakdownRout 路由
type BreakdownRout struct{}

// NewBreakdownRout 构造函数
func NewBreakdownRout() BreakdownRout {
	return BreakdownRout{}
}

// Load 加载路由
func (BreakdownRout) Load(engine *gin.Engine) {
	breakdownRout := engine.Group("api/breakdown")
	{
		typeRout := breakdownRout.Group(
			"type",
			middlewares.CheckAuth(),
			middlewares.CheckPermission(),
		)
		{
			// 新建
			typeRout.POST("", controllers.NewBreakdownTypeCtrl().Store)
			// 批量删除
			typeRout.POST("/destroyMany", controllers.NewBreakdownTypeCtrl().DestroyMany)
			// 删除
			typeRout.DELETE("/:uuid", controllers.NewBreakdownTypeCtrl().Destroy)
			// 编辑
			typeRout.PUT("/:uuid", controllers.NewBreakdownTypeCtrl().Update)
			// 详情
			typeRout.GET("/:uuid", controllers.NewBreakdownTypeCtrl().Detail)
			// 列表
			typeRout.GET("", controllers.NewBreakdownTypeCtrl().List)
			// jquery-dataTable数据列表
			typeRout.GET(".jdt", controllers.NewBreakdownTypeCtrl().ListJdt)
		}

		logRout := breakdownRout.Group(
			"log",
			middlewares.CheckAuth(),
			middlewares.CheckPermission(),
		)
		{
			// 新建
			logRout.POST("", controllers.NewBreakdownLogCtrl().Store)
			// 详情
			logRout.GET("/:uuid", controllers.NewBreakdownLogCtrl().Detail)
			// 列表
			logRout.GET("", controllers.NewBreakdownLogCtrl().List)
			// jquery-dataTable数据列表
			logRout.GET(".jdt", controllers.NewBreakdownLogCtrl().ListJdt)
			// 获取故障类型代码映射
			logRout.GET("/typeCodesMap", controllers.NewBreakdownLogCtrl().GetTypeCodesMap)
		}
	}

}
