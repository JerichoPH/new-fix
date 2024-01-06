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
	}
}
