package apiRout

import (
	"new-fix/controllers"
	"new-fix/middlewares"

	"github.com/gin-gonic/gin"
)

// OrganizationRout 路由
type OrganizationRout struct{}

// NewOrganizationRout 构造函数
func NewOrganizationRout() OrganizationRout {
	return OrganizationRout{}
}

// Load 加载路由
func (OrganizationRout) Load(engine *gin.Engine) {
	organizationRout := engine.Group("api/organization")
	{
		railwayRout := organizationRout.Group(
			"railway",
			middlewares.CheckAuth(),
			middlewares.CheckPermission(),
		)
		{
			// 新建
			railwayRout.POST("", controllers.NewOrganizationRailwayCtrl().Store)
			// 批量删除
			railwayRout.POST("destroyMany", controllers.NewOrganizationRailwayCtrl().DestroyMany)
			// 删除
			railwayRout.DELETE(":uuid", controllers.NewOrganizationRailwayCtrl().Destroy)
			// 编辑
			railwayRout.PUT(":uuid", controllers.NewOrganizationRailwayCtrl().Update)
			// 详情
			railwayRout.GET(":uuid", controllers.NewOrganizationRailwayCtrl().Detail)
			// 列表
			railwayRout.GET("", controllers.NewOrganizationRailwayCtrl().List)
			// jquery-dataTable数据列表
			railwayRout.GET(".jdt", controllers.NewOrganizationRailwayCtrl().ListJdt)
		}

		paragraphRout := organizationRout.Group(
			"paragraph",
			middlewares.CheckAuth(),
			middlewares.CheckPermission(),
		)
		{
			// 新建
			paragraphRout.POST("", controllers.NewOrganizationParagraphCtrl().Store)
			// 批量删除
			paragraphRout.POST("destroyMany", controllers.NewOrganizationParagraphCtrl().DestroyMany)
			// 删除
			paragraphRout.DELETE(":uuid", controllers.NewOrganizationParagraphCtrl().Destroy)
			// 编辑
			paragraphRout.PUT(":uuid", controllers.NewOrganizationParagraphCtrl().Update)
			// 详情
			paragraphRout.GET(":uuid", controllers.NewOrganizationParagraphCtrl().Detail)
			// 列表
			paragraphRout.GET("", controllers.NewOrganizationParagraphCtrl().List)
			// jquery-dataTable数据列表
			paragraphRout.GET(".jdt", controllers.NewOrganizationParagraphCtrl().ListJdt)
		}

		workshopRout := organizationRout.Group(
			"workshop",
			middlewares.CheckAuth(),
			middlewares.CheckPermission(),
		)
		{
			// 新建
			workshopRout.POST("", controllers.NewOrganizationWorkshopCtrl().Store)
			// 批量删除
			workshopRout.POST("destroyMany", controllers.NewOrganizationWorkshopCtrl().DestroyMany)
			// 删除
			workshopRout.DELETE(":uuid", controllers.NewOrganizationWorkshopCtrl().Destroy)
			// 编辑
			workshopRout.PUT(":uuid", controllers.NewOrganizationWorkshopCtrl().Update)
			// 详情
			workshopRout.GET(":uuid", controllers.NewOrganizationWorkshopCtrl().Detail)
			// 列表
			workshopRout.GET("", controllers.NewOrganizationWorkshopCtrl().List)
			// jquery-dataTable数据列表
			workshopRout.GET(".jdt", controllers.NewOrganizationWorkshopCtrl().ListJdt)
			// 获取车间类型代码映射
			workshopRout.GET("typeCodesMap", controllers.NewOrganizationWorkshopCtrl().GetTypeCodesMap)
		}

		stationRout := organizationRout.Group(
			"station",
			middlewares.CheckAuth(),
			middlewares.CheckPermission(),
		)
		{
			// 新建
			stationRout.POST("", controllers.NewOrganizationStationCtrl().Store)
			// 批量删除
			stationRout.POST("destroyMany", controllers.NewOrganizationStationCtrl().DestroyMany)
			// 删除
			stationRout.DELETE(":uuid", controllers.NewOrganizationStationCtrl().Destroy)
			// 编辑
			stationRout.PUT(":uuid", controllers.NewOrganizationStationCtrl().Update)
			// 详情
			stationRout.GET(":uuid", controllers.NewOrganizationStationCtrl().Detail)
			// 列表
			stationRout.GET("", controllers.NewOrganizationStationCtrl().List)
			// jquery-dataTable数据列表
			stationRout.GET(".jdt", controllers.NewOrganizationStationCtrl().ListJdt)
		}

		crossroadRout := organizationRout.Group(
			"crossroad",
			middlewares.CheckAuth(),
			middlewares.CheckPermission(),
		)
		{
			// 新建
			crossroadRout.POST("", controllers.NewOrganizationCrossroadCtrl().Store)
			// 批量删除
			crossroadRout.POST("destroyMany", controllers.NewOrganizationCrossroadCtrl().DestroyMany)
			// 删除
			crossroadRout.DELETE(":uuid", controllers.NewOrganizationCrossroadCtrl().Destroy)
			// 编辑
			crossroadRout.PUT(":uuid", controllers.NewOrganizationCrossroadCtrl().Update)
			// 详情
			crossroadRout.GET(":uuid", controllers.NewOrganizationCrossroadCtrl().Detail)
			// 列表
			crossroadRout.GET("", controllers.NewOrganizationCrossroadCtrl().List)
			// jquery-dataTable数据列表
			crossroadRout.GET(".jdt", controllers.NewOrganizationCrossroadCtrl().ListJdt)
		}

		centerRout := organizationRout.Group(
			"center",
			middlewares.CheckAuth(),
			middlewares.CheckPermission(),
		)
		{
			// 新建
			centerRout.POST("", controllers.NewOrganizationCenterCtrl().Store)
			// 批量删除
			centerRout.POST("destroyMany", controllers.NewOrganizationCenterCtrl().DestroyMany)
			// 删除
			centerRout.DELETE(":uuid", controllers.NewOrganizationCenterCtrl().Destroy)
			// 编辑
			centerRout.PUT(":uuid", controllers.NewOrganizationCenterCtrl().Update)
			// 详情
			centerRout.GET(":uuid", controllers.NewOrganizationCenterCtrl().Detail)
			// 列表
			centerRout.GET("", controllers.NewOrganizationCenterCtrl().List)
			// jquery-dataTable数据列表
			centerRout.GET(".jdt", controllers.NewOrganizationCenterCtrl().ListJdt)
		}

		workAreaRout := organizationRout.Group(
			"workArea",
			middlewares.CheckAuth(),
			middlewares.CheckPermission(),
		)
		{
			// 新建
			workAreaRout.POST("", controllers.NewOrganizationWorkAreaCtrl().Store)
			// 批量删除
			workAreaRout.POST("destroyMany", controllers.NewOrganizationWorkAreaCtrl().DestroyMany)
			// 删除
			workAreaRout.DELETE(":uuid", controllers.NewOrganizationWorkAreaCtrl().Destroy)
			// 编辑
			workAreaRout.PUT(":uuid", controllers.NewOrganizationWorkAreaCtrl().Update)
			// 详情
			workAreaRout.GET(":uuid", controllers.NewOrganizationWorkAreaCtrl().Detail)
			// 列表
			workAreaRout.GET("", controllers.NewOrganizationWorkAreaCtrl().List)
			// jquery-dataTable数据列表
			workAreaRout.GET(".jdt", controllers.NewOrganizationWorkAreaCtrl().ListJdt)
			// 获取车间类型代码映射
			workAreaRout.GET("typeCodesMap", controllers.NewOrganizationWorkAreaCtrl().GetTypeCodesMap)
		}

		lineRout := organizationRout.Group(
			"line",
			middlewares.CheckAuth(),
			middlewares.CheckPermission(),
		)
		{
			// 新建
			lineRout.POST("", controllers.NewOrganizationLineCtrl().Store)
			// 批量删除
			lineRout.POST("destroyMany", controllers.NewOrganizationLineCtrl().DestroyMany)
			// 删除
			lineRout.DELETE(":uuid", controllers.NewOrganizationLineCtrl().Destroy)
			// 编辑
			lineRout.PUT(":uuid", controllers.NewOrganizationLineCtrl().Update)
			// 详情
			lineRout.GET(":uuid", controllers.NewOrganizationLineCtrl().Detail)
			// 列表
			lineRout.GET("", controllers.NewOrganizationLineCtrl().List)
			// jquery-dataTable数据列表
			lineRout.GET(".jdt", controllers.NewOrganizationLineCtrl().ListJdt)
		}
	}
}
