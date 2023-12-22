package apiRout

import (
	"new-fix/controllers"
	"new-fix/middlewares"

	"github.com/gin-gonic/gin"
)

// EquipmentKindRout 路由
type EquipmentKindRout struct{}

// NewEquipmentKindRout 构造函数
func NewEquipmentKindRout() EquipmentKindRout {
	return EquipmentKindRout{}
}

// Load 加载路由
func (EquipmentKindRout) Load(engine *gin.Engine) {
	equipmentKindRout := engine.Group(
		"api/equipmentKind",
	)
	{
		// 器材种类
		categoryRout := equipmentKindRout.Group(
			"category",
			middlewares.CheckAuth(),
			middlewares.CheckPermission(),
		)
		{
			// 新建
			categoryRout.POST("", controllers.NewEquipmentKindCategoryCtrl().Store)
			// 批量删除
			categoryRout.POST("/destroyMany", controllers.NewEquipmentKindCategoryCtrl().DestroyMany)
			// 删除
			categoryRout.DELETE("/:uuid", controllers.NewEquipmentKindCategoryCtrl().Destroy)
			// 编辑
			categoryRout.PUT("/:uuid", controllers.NewEquipmentKindCategoryCtrl().Update)
			// 详情
			categoryRout.GET("/:uuid", controllers.NewEquipmentKindCategoryCtrl().Detail)
			// 列表
			categoryRout.GET("", controllers.NewEquipmentKindCategoryCtrl().List)
			// jquery-dataTable数据列表
			categoryRout.GET(".jdt", controllers.NewEquipmentKindCategoryCtrl().ListJdt)
		}

		// 器材类型
		typeRout := equipmentKindRout.Group(
			"type",
			middlewares.CheckAuth(),
			middlewares.CheckPermission(),
		)
		{
			// 新建
			typeRout.POST("", controllers.NewEquipmentKindTypeCtrl().Store)
			// 删除
			typeRout.DELETE("/:uuid", controllers.NewEquipmentKindTypeCtrl().Destroy)
			// 编辑
			typeRout.PUT("/:uuid", controllers.NewEquipmentKindTypeCtrl().Update)
			// 详情
			typeRout.GET("/:uuid", controllers.NewEquipmentKindTypeCtrl().Detail)
			// 列表
			typeRout.GET("", controllers.NewEquipmentKindTypeCtrl().List)
			// jquery-dataTable数据列表
			typeRout.GET(".jdt", controllers.NewEquipmentKindTypeCtrl().ListJdt)
		}

		// 器材型号
		modelRout := equipmentKindRout.Group(
			"model",
			middlewares.CheckAuth(),
			middlewares.CheckPermission(),
		)
		{
			// 新建
			modelRout.POST("", controllers.NewEquipmentKindModelCtrl().Store)
			// 删除
			modelRout.DELETE("/:uuid", controllers.NewEquipmentKindModelCtrl().Destroy)
			// 编辑
			modelRout.PUT("/:uuid", controllers.NewEquipmentKindModelCtrl().Update)
			// 详情
			modelRout.GET("/:uuid", controllers.NewEquipmentKindModelCtrl().Detail)
			// 列表
			modelRout.GET("", controllers.NewEquipmentKindModelCtrl().List)
			// jquery-dataTable数据列表
			modelRout.GET(".jdt", controllers.NewEquipmentKindModelCtrl().ListJdt)
		}
	}
}
