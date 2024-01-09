package apiRout

import (
	"new-fix/controllers"
	"new-fix/middlewares"

	"github.com/gin-gonic/gin"
)

// InstallRout 路由
type InstallRout struct{}

// NewInstallRout 构造函数
func NewInstallRout() InstallRout {
	return InstallRout{}
}

// Load 加载路由
func (InstallRout) Load(engine *gin.Engine) {
	installRout := engine.Group("api/install")
	{
		// 室内上道位置-机房类型
		indoorRoomTypeRout := installRout.Group(
			"indoorRoomType",
			middlewares.CheckAuth(),
			middlewares.CheckPermission(),
		)
		// 新建
		indoorRoomTypeRout.POST("", controllers.NewInstallIndoorRoomTypeCtrl().Store)
		// 批量删除
		indoorRoomTypeRout.POST("/destroyMany", controllers.NewInstallIndoorRoomTypeCtrl().DestroyMany)
		// 删除
		indoorRoomTypeRout.DELETE("/:uuid", controllers.NewInstallIndoorRoomTypeCtrl().Destroy)
		// 编辑
		indoorRoomTypeRout.PUT("/:uuid", controllers.NewInstallIndoorRoomTypeCtrl().Update)
		// 详情
		indoorRoomTypeRout.GET("/:uuid", controllers.NewInstallIndoorRoomTypeCtrl().Detail)
		// 列表
		indoorRoomTypeRout.GET("", controllers.NewInstallIndoorRoomTypeCtrl().List)
		// jquery-dataTable数据列表
		indoorRoomTypeRout.GET(".jdt", controllers.NewInstallIndoorRoomTypeCtrl().ListJdt)

		// 室内上道位置-机房
		indoorRoomRout := installRout.Group(
			"indoorRoom",
			middlewares.CheckAuth(),
			middlewares.CheckPermission(),
		)
		// 新建
		indoorRoomRout.POST("", controllers.NewInstallIndoorRoomCtrl().Store)
		// 批量删除
		indoorRoomRout.POST("/destroyMany", controllers.NewInstallIndoorRoomCtrl().DestroyMany)
		// 删除
		indoorRoomRout.DELETE("/:uuid", controllers.NewInstallIndoorRoomCtrl().Destroy)
		// 编辑
		indoorRoomRout.PUT("/:uuid", controllers.NewInstallIndoorRoomCtrl().Update)
		// 详情
		indoorRoomRout.GET("/:uuid", controllers.NewInstallIndoorRoomCtrl().Detail)
		// 列表
		indoorRoomRout.GET("", controllers.NewInstallIndoorRoomCtrl().List)
		// jquery-dataTable数据列表
		indoorRoomRout.GET(".jdt", controllers.NewInstallIndoorRoomCtrl().ListJdt)

		// 室内上道位置-排
		indoorPlatoonRout := installRout.Group(
			"indoorPlatoon",
			middlewares.CheckAuth(),
			middlewares.CheckPermission(),
		)
		// 新建
		indoorPlatoonRout.POST("", controllers.NewInstallIndoorPlatoonCtrl().Store)
		// 批量删除
		indoorPlatoonRout.POST("/destroyMany", controllers.NewInstallIndoorPlatoonCtrl().DestroyMany)
		// 删除
		indoorPlatoonRout.DELETE("/:uuid", controllers.NewInstallIndoorPlatoonCtrl().Destroy)
		// 编辑
		indoorPlatoonRout.PUT("/:uuid", controllers.NewInstallIndoorPlatoonCtrl().Update)
		// 详情
		indoorPlatoonRout.GET("/:uuid", controllers.NewInstallIndoorPlatoonCtrl().Detail)
		// 列表
		indoorPlatoonRout.GET("", controllers.NewInstallIndoorPlatoonCtrl().List)
		// jquery-dataTable数据列表
		indoorPlatoonRout.GET(".jdt", controllers.NewInstallIndoorPlatoonCtrl().ListJdt)

		// 室内上道位置-架
		indoorShelfRout := installRout.Group(
			"indoorShelf",
			middlewares.CheckAuth(),
			middlewares.CheckPermission(),
		)
		// 新建
		indoorShelfRout.POST("", controllers.NewInstallIndoorShelfCtrl().Store)
		// 批量删除
		indoorShelfRout.POST("/destroyMany", controllers.NewInstallIndoorShelfCtrl().DestroyMany)
		// 删除
		indoorShelfRout.DELETE("/:uuid", controllers.NewInstallIndoorShelfCtrl().Destroy)
		// 编辑
		indoorShelfRout.PUT("/:uuid", controllers.NewInstallIndoorShelfCtrl().Update)
		// 详情
		indoorShelfRout.GET("/:uuid", controllers.NewInstallIndoorShelfCtrl().Detail)
		// 列表
		indoorShelfRout.GET("", controllers.NewInstallIndoorShelfCtrl().List)
		// jquery-dataTable数据列表
		indoorShelfRout.GET(".jdt", controllers.NewInstallIndoorShelfCtrl().ListJdt)

		// 室内上道位置-层
		indoorTierRout := installRout.Group(
			"indoorTier",
			middlewares.CheckAuth(),
			middlewares.CheckPermission(),
		)
		// 新建
		indoorTierRout.POST("", controllers.NewInstallIndoorTierCtrl().Store)
		// 批量新建
		indoorTierRout.POST("/storeMany", controllers.NewInstallIndoorTierCtrl().StoreMany)
		// 批量删除
		indoorTierRout.POST("/destroyMany", controllers.NewInstallIndoorTierCtrl().DestroyMany)
		// 删除
		indoorTierRout.DELETE("/:uuid", controllers.NewInstallIndoorTierCtrl().Destroy)
		// 编辑
		indoorTierRout.PUT("/:uuid", controllers.NewInstallIndoorTierCtrl().Update)
		// 详情
		indoorTierRout.GET("/:uuid", controllers.NewInstallIndoorTierCtrl().Detail)
		// 列表
		indoorTierRout.GET("", controllers.NewInstallIndoorTierCtrl().List)
		// jquery-dataTable数据列表
		indoorTierRout.GET(".jdt", controllers.NewInstallIndoorTierCtrl().ListJdt)

		// 室内上道位置-位
		indoorCellRout := installRout.Group(
			"indoorCell",
			middlewares.CheckAuth(),
			middlewares.CheckPermission(),
		)
		// 新建
		indoorCellRout.POST("", controllers.NewInstallIndoorCellCtrl().Store)
		// 批量新建
		indoorCellRout.POST("/storeMany", controllers.NewInstallIndoorCellCtrl().StoreMany)
		// 批量删除
		indoorCellRout.POST("/destroyMany", controllers.NewInstallIndoorCellCtrl().DestroyMany)
		// 删除
		indoorCellRout.DELETE("/:uuid", controllers.NewInstallIndoorCellCtrl().Destroy)
		// 编辑
		indoorCellRout.PUT("/:uuid", controllers.NewInstallIndoorCellCtrl().Update)
		// 详情
		indoorCellRout.GET("/:uuid", controllers.NewInstallIndoorCellCtrl().Detail)
		// 列表
		indoorCellRout.GET("", controllers.NewInstallIndoorCellCtrl().List)
		// jquery-dataTable数据列表
		indoorCellRout.GET(".jdt", controllers.NewInstallIndoorCellCtrl().ListJdt)
	}
}
