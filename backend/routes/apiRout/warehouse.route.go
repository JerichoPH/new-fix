package apiRout

import (
	"new-fix/controllers"
	"new-fix/middlewares"

	"github.com/gin-gonic/gin"
)

// WarehouseRout 路由
type WarehouseRout struct{}

// NewWarehouseRout 构造函数
func NewWarehouseRout() WarehouseRout {
	return WarehouseRout{}
}

// Load 加载路由
func (WarehouseRout) Load(engine *gin.Engine) {
	warehouseRout := engine.Group(
		"api/warehouse",
	)
	{
		// 仓库-库房
		storehouseRout := warehouseRout.Group(
			"storehouse",
			middlewares.CheckAuth(),
			middlewares.CheckPermission(),
		)
		{
			// 新建
			storehouseRout.POST("", controllers.NewWarehouseStorehouseCtrl().Store)
			// 批量删除
			storehouseRout.POST("/destroyMany", controllers.NewWarehouseStorehouseCtrl().DestroyMany)
			// 删除
			storehouseRout.DELETE("/:uuid", controllers.NewWarehouseStorehouseCtrl().Destroy)
			// 编辑
			storehouseRout.PUT("/:uuid", controllers.NewWarehouseStorehouseCtrl().Update)
			// 详情
			storehouseRout.GET("/:uuid", controllers.NewWarehouseStorehouseCtrl().Detail)
			// 列表
			storehouseRout.GET("", controllers.NewWarehouseStorehouseCtrl().List)
			// jquery-dataTable数据列表
			storehouseRout.GET(".jdt", controllers.NewWarehouseStorehouseCtrl().ListJdt)
		}

		// 仓库-库区
		areaRout := warehouseRout.Group(
			"area",
			middlewares.CheckAuth(),
			middlewares.CheckPermission(),
		)
		{
			// 新建
			areaRout.POST("", controllers.NewWarehouseAreaCtrl().Store)
			// 批量删除
			areaRout.POST("/destroyMany", controllers.NewWarehouseAreaCtrl().DestroyMany)
			// 删除
			areaRout.DELETE("/:uuid", controllers.NewWarehouseAreaCtrl().Destroy)
			// 编辑
			areaRout.PUT("/:uuid", controllers.NewWarehouseAreaCtrl().Update)
			// 详情
			areaRout.GET("/:uuid", controllers.NewWarehouseAreaCtrl().Detail)
			// 列表
			areaRout.GET("", controllers.NewWarehouseAreaCtrl().List)
			// jquery-dataTable数据列表
			areaRout.GET(".jdt", controllers.NewWarehouseAreaCtrl().ListJdt)
		}

		// 仓库-排
		platoonRout := warehouseRout.Group(
			"platoon",
			middlewares.CheckAuth(),
			middlewares.CheckPermission(),
		)
		{
			// 新建
			platoonRout.POST("", controllers.NewWarehousePlatoonCtrl().Store)
			// 批量删除
			platoonRout.POST("/destroyMany", controllers.NewWarehousePlatoonCtrl().DestroyMany)
			// 删除
			platoonRout.DELETE("/:uuid", controllers.NewWarehousePlatoonCtrl().Destroy)
			// 编辑
			platoonRout.PUT("/:uuid", controllers.NewWarehousePlatoonCtrl().Update)
			// 详情
			platoonRout.GET("/:uuid", controllers.NewWarehousePlatoonCtrl().Detail)
			// 列表
			platoonRout.GET("", controllers.NewWarehousePlatoonCtrl().List)
			// jquery-dataTable数据列表
			platoonRout.GET(".jdt", controllers.NewWarehousePlatoonCtrl().ListJdt)
		}

		// 仓库-架
		shelfRout := warehouseRout.Group(
			"shelf",
			middlewares.CheckAuth(),
			middlewares.CheckPermission(),
		)
		{
			// 新建
			shelfRout.POST("", controllers.NewWarehouseShelfCtrl().Store)
			// 批量删除
			shelfRout.POST("/destroyMany", controllers.NewWarehouseShelfCtrl().DestroyMany)
			// 删除
			shelfRout.DELETE("/:uuid", controllers.NewWarehouseShelfCtrl().Destroy)
			// 编辑
			shelfRout.PUT("/:uuid", controllers.NewWarehouseShelfCtrl().Update)
			// 详情
			shelfRout.GET("/:uuid", controllers.NewWarehouseShelfCtrl().Detail)
			// 列表
			shelfRout.GET("", controllers.NewWarehouseShelfCtrl().List)
			// jquery-dataTable数据列表
			shelfRout.GET(".jdt", controllers.NewWarehouseShelfCtrl().ListJdt)
		}

		// 仓库-层
		tierRout := warehouseRout.Group(
			"tier",
			middlewares.CheckAuth(),
			middlewares.CheckPermission(),
		)
		{
			tierRout.POST("/storeMany", controllers.NewWarehouseTierCtrl().StoreMany)
			// 新建
			tierRout.POST("", controllers.NewWarehouseTierCtrl().Store)
			// 批量删除
			tierRout.POST("/destroyMany", controllers.NewWarehouseTierCtrl().DestroyMany)
			// 删除
			tierRout.DELETE("/:uuid", controllers.NewWarehouseTierCtrl().Destroy)
			// 编辑
			tierRout.PUT("/:uuid", controllers.NewWarehouseTierCtrl().Update)
			// 详情
			tierRout.GET("/:uuid", controllers.NewWarehouseTierCtrl().Detail)
			// 列表
			tierRout.GET("", controllers.NewWarehouseTierCtrl().List)
			// jquery-dataTable数据列表
			tierRout.GET(".jdt", controllers.NewWarehouseTierCtrl().ListJdt)
		}

		// 仓库-位
		cellRout := warehouseRout.Group(
			"cell",
			middlewares.CheckAuth(),
			middlewares.CheckPermission(),
		)
		{
			// 批量新建
			cellRout.POST("/storeMany", controllers.NewWarehouseCellCtrl().StoreMany)
			// 新建
			cellRout.POST("", controllers.NewWarehouseCellCtrl().Store)
			// 批量删除
			cellRout.POST("/destroyMany", controllers.NewWarehouseCellCtrl().DestroyMany)
			// 删除
			cellRout.DELETE("/:uuid", controllers.NewWarehouseCellCtrl().Destroy)
			// 编辑
			cellRout.PUT("/:uuid", controllers.NewWarehouseCellCtrl().Update)
			// 详情
			cellRout.GET("/:uuid", controllers.NewWarehouseCellCtrl().Detail)
			// 列表
			cellRout.GET("", controllers.NewWarehouseCellCtrl().List)
			// jquery-dataTable数据列表
			cellRout.GET(".jdt", controllers.NewWarehouseCellCtrl().ListJdt)
		}
	}
}
