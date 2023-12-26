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
	}
}
