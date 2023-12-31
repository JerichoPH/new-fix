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
}
