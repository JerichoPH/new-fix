package apiRout

import (
	"new-fix/controllers"
	"new-fix/middlewares"

	"github.com/gin-gonic/gin"
)

// SourceRout 路由
type SourceRout struct{}

// NewSourceRout 构造函数
func NewSourceRout() SourceRout {
	return SourceRout{}
}

// Load 加载路由
func (SourceRout) Load(engine *gin.Engine) {
	sourceRout := engine.Group("api/source")
	{
		// 来源类型
		typeRout := sourceRout.Group(
			"type",
			middlewares.CheckAuth(),
			middlewares.CheckPermission(),
		)
		// 新建
		typeRout.POST("", controllers.NewSourceTypeCtrl().Store)
		// 批量删除
		typeRout.POST("/destroyMany", controllers.NewSourceTypeCtrl().DestroyMany)
		// 删除
		typeRout.DELETE("/:uuid", controllers.NewSourceTypeCtrl().Destroy)
		// 编辑
		typeRout.PUT("/:uuid", controllers.NewSourceTypeCtrl().Update)
		// 详情
		typeRout.GET("/:uuid", controllers.NewSourceTypeCtrl().Detail)
		// 列表
		typeRout.GET("", controllers.NewSourceTypeCtrl().List)
		// jquery-dataTable数据列表
		typeRout.GET(".jdt", controllers.NewSourceTypeCtrl().ListJdt)
	}

	// 来源项目
	projectRout := sourceRout.Group(
		"project",
		middlewares.CheckAuth(),
		middlewares.CheckPermission(),
	)
	// 新建
	projectRout.POST("", controllers.NewSourceProjectCtrl().Store)
	// 批量删除
	projectRout.POST("/destroyMany", controllers.NewSourceProjectCtrl().DestroyMany)
	// 删除
	projectRout.DELETE("/:uuid", controllers.NewSourceProjectCtrl().Destroy)
	// 编辑
	projectRout.PUT("/:uuid", controllers.NewSourceProjectCtrl().Update)
	// 详情
	projectRout.GET("/:uuid", controllers.NewSourceProjectCtrl().Detail)
	// 列表
	projectRout.GET("", controllers.NewSourceProjectCtrl().List)
	// jquery-dataTable数据列表
	projectRout.GET(".jdt", controllers.NewSourceProjectCtrl().ListJdt)
}
