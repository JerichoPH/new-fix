package apiRout

import (
	"new-fix/controllers"
	"new-fix/middlewares"

	"github.com/gin-gonic/gin"
)

// FactoryRout 路由
type FactoryRout struct{}

// NewFactoryRout 构造函数
func NewFactoryRout() FactoryRout {
	return FactoryRout{}
}

// Load 加载路由
func (FactoryRout) Load(engine *gin.Engine) {
	r := engine.Group(
		"api/factory",
		middlewares.CheckAuth(),
	// middlewares.CheckPermission(),
	)
	{
		// 新建
		r.POST("", controllers.NewFactoryCtrl().Store)
		// 删除
		r.POST("/destroyMany", controllers.NewFactoryCtrl().DestroyMany)
		// 删除
		r.DELETE("/:uuid", controllers.NewFactoryCtrl().Destroy)
		// 编辑
		r.PUT("/:uuid", controllers.NewFactoryCtrl().Update)
		// 详情
		r.GET("/:uuid", controllers.NewFactoryCtrl().Detail)
		// 列表
		r.GET("", controllers.NewFactoryCtrl().List)
		// jquery-dataTable数据列表
		r.GET(".jdt", controllers.NewFactoryCtrl().ListJdt)
	}
}
