package controllers

import (
	"new-fix/models"
	"new-fix/tools"
	"new-fix/wrongs"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type (
	// OrganizationRailwayCtrl 路局控制器
	OrganizationRailwayCtrl struct{}
	// OrganizationRailwayStoreForm 路局表单
	OrganizationRailwayStoreForm struct {
		UniqueCode string `json:"unique_code"`
		Name       string `json:"name"`
		ShortName  string `json:"short_name"`
	}
	// OrganizationRailwayDestroyManyForm 批量删除路局表单
	OrganizaitonRailwayDestroyManyForm struct {
		Uuids []string `json:"uuids"`
	}

	// OrganizationParagraphCtrl 站段控制器
	OrganizationParagraphCtrl struct{}
	// OrganizationParagraphStoreForm 站段表单
	OrganizationParagraphStoreForm struct {
		UniqueCode              string `json:"unique_code"`
		Name                    string `json:"name"`
		OrganizationRailwayUuid string `json:"organization_railway_uuid"`
		organizationRailway     *models.OrganizationRailwayMdl
	}
	// OrganizationParagraphDestroyManyForm 批量删除站段表单
	OrganizaitonParagraphDestroyManyForm struct {
		Uuids []string `json:"uuids"`
	}

	// OrganizationWorkshopCtrl 车间控制器
	OrganizationWorkshopCtrl struct{}
	// OrganizationWorkshopStoreForm 车间表单
	OrganizationWorkshopStoreForm struct {
		UniqueCode                string `json:"unique_code"`
		Name                      string `json:"name"`
		OrganizationParagraphUuid string `json:"organization_paragraph_uuid"`
		organizationParagraph     *models.OrganizationParagraphMdl
		TypeCode                  string `json:"type_code"`
	}
	// OrganizationWorkshopDestroyManyForm 批量删除车间表单
	OrganizationWorkshopDestroyManyForm struct {
		Uuids []string `json:"uuids"`
	}

	// OrganizationStationCtrl 站场控制器
	OrganizationStationCtrl struct{}
	// OrganizationStationStoreForm 站场表单
	OrganizationStationStoreForm struct {
		UniqueCode               string `json:"unique_code"`
		Name                     string `json:"name"`
		OrganizationWorkshopUuid string `json:"organization_workshop_uuid"`
		organizationWorkshop     *models.OrganizationWorkshopMdl
		OrganizationLineUuid     string `json:"organization_line_uuid"`
		organizationLine         *models.OrganizationLineMdl
	}
	// OrganizationStationDestroyManyForm 批量删除站场表单
	OrganizationStationDestroyManyForm struct {
		Uuids []string `json:"uuids"`
	}

	// OrganizationCrossroadCtrl
	OrganizationCrossroadCtrl struct{}
	// OrganizationCrossroadStoreForm 道口表单
	OrganizationCrossroadStoreForm struct {
		UniqueCode               string `json:"unique_code"`
		Name                     string `json:"name"`
		OrganizationWorkshopUuid string `json:"organization_workshop_uuid"`
		organizationWorkshop     *models.OrganizationWorkshopMdl
		OrganizationLineUuid     string `json:"organization_line_uuid"`
		organizationLine         *models.OrganizationLineMdl
	}
	// OrganizationCrossroadDestroyManyForm 批量删除道口表单
	OrganizationCrossroadDestroyManyForm struct {
		Uuids []string `json:"uuids"`
	}

	// OrganizationCenterCtrl 中心控制器
	OrganizationCenterCtrl struct{}
	// OrganizationCenterStoreForm 中心表单
	OrganizationCenterStoreForm struct {
		UniqueCode               string `json:"unique_code"`
		Name                     string `json:"name"`
		OrganizationWorkshopUuid string `json:"organization_workshop_uuid"`
		organizationWorkshop     *models.OrganizationWorkshopMdl
		OrganizationLineUuid     string `json:"organization_line_uuid"`
		organizationLine         *models.OrganizationLineMdl
	}
	// OrganizationCenterDestroyManyForm 批量删除中心表单
	OrganizationCenterDestroyManyForm struct {
		Uuids []string `json:"uuids"`
	}

	// OrganizationWorkAreaCtrl 工区控制器
	OrganizationWorkAreaCtrl struct{}
	// OrganizationWorkAreaStoreForm 工区表单
	OrganizationWorkAreaStoreForm struct {
		UniqueCode               string `json:"unique_code"`
		Name                     string `json:"name"`
		OrganizationWorkshopUuid string `json:"organization_workshop_uuid"`
		organizationWorkshop     *models.OrganizationWorkshopMdl
		TypeCode                 string `json:"type_code"`
	}
	// OrganizationWorkAreaDestroyManyForm 批量删除工区表单
	OrganizationWorkAreaDestroyManyForm struct {
		Uuids []string `json:"uuids"`
	}

	// OrganizationLineCtrl 线别控制器
	OrganizationLineCtrl struct{}
	// OrganizationLineStoreForm 线别表单
	OrganizationLineStoreForm struct {
		UniqueCode                 string `json:"unique_code"`
		Name                       string `json:"name"`
		OrganizationParagraphUuid  string `json:"organization_paragraph_uuid"`
		organizationParagraph      *models.OrganizationParagraphMdl
		OrganizationStationUuids   []string `json:"organization_station_uuids"`
		organizationStations       []*models.OrganizationStationMdl
		OrganizationCrossroadUuids []string `json:"organization_crossroad_uuids"`
		organizationCrossroads     []*models.OrganizationCrossroadMdl
		OrganizationCenterUuids    []string `json:"organization_center_uuids"`
		organizationCenters        []*models.OrganizationCenterMdl
	}
	// OrganizationLineDestroyManyForm 批量删除线别表单
	OrganizationLineDestroyManyForm struct {
		Uuids []string `json:"uuids"`
	}
)

// ShouldBind 绑定表单（路局）
func (receiver OrganizationRailwayStoreForm) ShouldBind(ctx *gin.Context) OrganizationRailwayStoreForm {
	if err := ctx.ShouldBind(&receiver); err != nil {
		wrongs.ThrowValidate("表单验证失败：%s", err.Error())
	}
	if receiver.UniqueCode == "" {
		wrongs.ThrowValidate("路局代码必填")
	} else {
		if len(receiver.UniqueCode) != 3 {
			wrongs.ThrowValidate("路局代码必须是3位")
		}
	}
	if receiver.Name == "" {
		wrongs.ThrowValidate("路局名称必填")
	}
	if receiver.ShortName == "" {
		wrongs.ThrowValidate("路局简称必填")
	}
	return receiver
}

// ShouldBind 绑定表单（批量删除路局）
func (receiver OrganizaitonRailwayDestroyManyForm) ShouldBind(ctx *gin.Context) OrganizaitonRailwayDestroyManyForm {
	if err := ctx.ShouldBind(&receiver); err != nil {
		wrongs.ThrowValidate("表单验证失败：%s", err.Error())
	}
	if len(receiver.Uuids) == 0 {
		wrongs.ThrowValidate("路局编号必填")
	}
	return receiver
}

// ShouldBind 绑定表单（站段）
func (receiver OrganizationParagraphStoreForm) ShouldBind(ctx *gin.Context) OrganizationParagraphStoreForm {
	if err := ctx.ShouldBind(&receiver); err != nil {
		wrongs.ThrowValidate("表单绑定失败：%s", err.Error())
	}
	if receiver.UniqueCode == "" {
		wrongs.ThrowValidate("站段代码必填")
	} else {
		if len(receiver.UniqueCode) != 4 {
			wrongs.ThrowValidate("站段代码必须是4位")
		}
	}
	if receiver.Name == "" {
		wrongs.ThrowValidate("站段名称必填")
	}
	if receiver.OrganizationRailwayUuid == "" {
		wrongs.ThrowValidate("路局编号必填")
	} else {
		ret := models.NewOrganizationRailwayMdl().GetDb("").Where("uuid = ?", receiver.OrganizationRailwayUuid).First(&receiver.organizationRailway)
		wrongs.ThrowWhenEmpty(ret, "所属路局")
	}
	return receiver
}

// ShouldBind 表单绑定（批量删除站段）
func (receiver OrganizaitonParagraphDestroyManyForm) ShouldBind(ctx *gin.Context) OrganizaitonParagraphDestroyManyForm {
	if err := ctx.ShouldBind(&receiver); err != nil {
		wrongs.ThrowValidate("表单绑定失败：%s", err.Error())
	}
	if len(receiver.Uuids) == 0 {
		wrongs.ThrowValidate("站段编号必填")
	}
	return receiver
}

// ShouldBind 绑定表单（车间）
func (receiver OrganizationWorkshopStoreForm) ShouldBind(ctx *gin.Context) OrganizationWorkshopStoreForm {
	if err := ctx.ShouldBind(&receiver); err != nil {
		wrongs.ThrowValidate("表单绑定失败：%s", err.Error())
	}
	if receiver.UniqueCode == "" {
		wrongs.ThrowValidate("车间代码必填")
	}
	if receiver.Name == "" {
		wrongs.ThrowValidate("车间名称必填")
	}
	if receiver.OrganizationParagraphUuid == "" {
		wrongs.ThrowValidate("站段编号必填")
	} else {
		ret := models.NewOrganizationParagraphMdl().GetDb("").Where("uuid = ?", receiver.OrganizationParagraphUuid).First(&receiver.organizationParagraph)
		wrongs.ThrowWhenEmpty(ret, "站段")
	}
	if receiver.TypeCode != "" {
		if !tools.InString(receiver.TypeCode, models.OrganizationWorkshopMdl{}.GetTypeCodes()) {
			wrongs.ThrowValidate("车间类型代码错误")
		}
	}
	return receiver
}

// ShouldBind 表单绑定（批量删除车间）
func (receiver OrganizationWorkshopDestroyManyForm) ShouldBind(ctx *gin.Context) OrganizationWorkshopDestroyManyForm {
	if err := ctx.ShouldBind(&receiver); err != nil {
		wrongs.ThrowValidate("表单绑定失败：%s", err.Error())
	}

	if len(receiver.Uuids) == 0 {
		wrongs.ThrowValidate("车间编号必填")
	}

	return receiver
}

// ShouldBind 绑定表单（站场）
func (receiver OrganizationStationStoreForm) ShouldBind(ctx *gin.Context) OrganizationStationStoreForm {
	if err := ctx.ShouldBind(&receiver); err != nil {
		wrongs.ThrowValidate("表单绑定失败：%s", err.Error())
	}
	if receiver.UniqueCode == "" {
		wrongs.ThrowValidate("站场代码必填")
	} else {
		if len(receiver.UniqueCode) != 6 {
			wrongs.ThrowValidate("站场代码必须是6位")
		}
	}
	if receiver.Name == "" {
		wrongs.ThrowValidate("站场名称必填")
	}
	if receiver.OrganizationWorkshopUuid == "" {
		wrongs.ThrowValidate("所属车间编号不能为空")
	} else {
		ret := models.NewOrganizationWorkshopMdl().SetPreloads("OrganizationParagraph", "OrganizationParagraph.OrganizationRailway").GetDb("").Where("uuid = ?", receiver.OrganizationWorkshopUuid).First(&receiver.organizationWorkshop)
		wrongs.ThrowWhenEmpty(ret, "所属车间")
	}
	if receiver.OrganizationLineUuid != "" {
		models.NewOrganizationLineMdl().GetDb("").Where("uuid = ?", receiver.OrganizationLineUuid).First(&receiver.organizationLine)
	}
	return receiver
}

// ShouldBind 表单绑定（批量删除站场）
func (receiver OrganizationStationDestroyManyForm) ShouldBind(ctx *gin.Context) OrganizationStationDestroyManyForm {
	if err := ctx.ShouldBind(&receiver); err != nil {
		wrongs.ThrowValidate("表单绑定失败：%s", err.Error())
	}
	if len(receiver.Uuids) == 0 {
		wrongs.ThrowValidate("站场编号必填")
	}
	return receiver
}

// ShouldBind 绑定表单（道口）
func (receiver OrganizationCrossroadStoreForm) ShouldBind(ctx *gin.Context) OrganizationCrossroadStoreForm {
	if err := ctx.ShouldBind(&receiver); err != nil {
		wrongs.ThrowValidate("表单绑定失败：%s", err.Error())
	}
	if receiver.UniqueCode == "" {
		wrongs.ThrowValidate("道口代码必填")
	} else {
		if len(receiver.UniqueCode) != 5 {
			wrongs.ThrowValidate("道口代码必须是6位")
		}
	}
	if receiver.Name == "" {
		wrongs.ThrowValidate("道口名称必填")
	}
	if receiver.OrganizationWorkshopUuid == "" {
		wrongs.ThrowValidate("所属车间编号不能为空")
	} else {
		ret := models.NewOrganizationWorkshopMdl().SetPreloads("OrganizationParagraph", "OrganizationParagraph.OrganizationRailway").GetDb("").Where("uuid = ?", receiver.OrganizationWorkshopUuid).First(&receiver.organizationWorkshop)
		wrongs.ThrowWhenEmpty(ret, "所属车间")
	}
	if receiver.OrganizationLineUuid != "" {
		models.NewOrganizationLineMdl().GetDb("").Where("uuid = ?", receiver.OrganizationLineUuid).First(&receiver.organizationLine)
	}
	return receiver
}

// ShouldBind 表单绑定（批量删除道口）
func (receiver OrganizationCrossroadDestroyManyForm) ShouldBind(ctx *gin.Context) OrganizationCrossroadDestroyManyForm {
	if err := ctx.ShouldBind(&receiver); err != nil {
		wrongs.ThrowValidate("表单绑定失败：%s", err.Error())
	}
	if len(receiver.Uuids) == 0 {
		wrongs.ThrowValidate("道口编号必填")
	}
	return receiver
}

// ShouldBind 绑定表单（中心）
func (receiver OrganizationCenterStoreForm) ShouldBind(ctx *gin.Context) OrganizationCenterStoreForm {
	if err := ctx.ShouldBind(&receiver); err != nil {
		wrongs.ThrowValidate("表单绑定失败：%s", err.Error())
	}
	if receiver.UniqueCode == "" {
		wrongs.ThrowValidate("中心代码必填")
	} else {
		if len(receiver.UniqueCode) != 6 {
			wrongs.ThrowValidate("中心代码必须是6位")
		}
	}
	if receiver.Name == "" {
		wrongs.ThrowValidate("中心名称必填")
	}
	if receiver.OrganizationWorkshopUuid == "" {
		wrongs.ThrowValidate("所属车间编号不能为空")
	} else {
		ret := models.NewOrganizationWorkshopMdl().SetPreloads("OrganizationParagraph", "OrganizationParagraph.OrganizationRailway").GetDb("").Where("uuid = ?", receiver.OrganizationWorkshopUuid).First(&receiver.organizationWorkshop)
		wrongs.ThrowWhenEmpty(ret, "所属车间")
	}
	if receiver.OrganizationLineUuid != "" {
		models.NewOrganizationLineMdl().GetDb("").Where("uuid = ?", receiver.OrganizationLineUuid).First(&receiver.organizationLine)
	}
	return receiver
}

// ShouldBind 表单绑定（批量删除中心）
func (receiver OrganizationCenterDestroyManyForm) ShouldBind(ctx *gin.Context) OrganizationCenterDestroyManyForm {
	if err := ctx.ShouldBind(&receiver); err != nil {
		wrongs.ThrowValidate("表单绑定失败：%s", err.Error())
	}
	if len(receiver.Uuids) == 0 {
		wrongs.ThrowValidate("中心编号必填")
	}
	return receiver
}

// ShouldBind 绑定表单（工区）
func (receiver OrganizationWorkAreaStoreForm) ShouldBind(ctx *gin.Context) OrganizationWorkAreaStoreForm {
	if err := ctx.ShouldBind(&receiver); err != nil {
		wrongs.ThrowValidate("表单绑定失败：%s", err.Error())
	}
	if receiver.UniqueCode == "" {
		wrongs.ThrowValidate("工区代码必填")
	} else {
		if len(receiver.UniqueCode) != 8 {
			wrongs.ThrowValidate("工区代码必须是8位")
		}
	}
	if receiver.Name == "" {
		wrongs.ThrowValidate("工区名称必填")
	}
	if receiver.OrganizationWorkshopUuid == "" {
		wrongs.ThrowValidate("所属车间编号不能为空")
	} else {
		ret := models.NewOrganizationWorkshopMdl().GetDb("").Where("uuid = ?", receiver.OrganizationWorkshopUuid).First(&receiver.organizationWorkshop)
		wrongs.ThrowWhenEmpty(ret, "所属车间")
	}
	if receiver.TypeCode != "" {
		if !tools.InString(receiver.TypeCode, models.OrganizationWorkAreaMdl{}.GetTypeCodes()) {
			wrongs.ThrowValidate("工区类型代码错误")
		}
	}
	return receiver
}

// ShouldBind 表单绑定（批量删除工区）
func (receiver OrganizationWorkAreaDestroyManyForm) ShouldBind(ctx *gin.Context) OrganizationWorkAreaDestroyManyForm {
	if err := ctx.ShouldBind(&receiver); err != nil {
		wrongs.ThrowValidate("表单绑定失败：%s", err.Error())
	}
	if len(receiver.Uuids) == 0 {
		wrongs.ThrowValidate("工区编号必填")
	}
	return receiver
}

// ShouldBind 表单绑定（线别）
func (receiver OrganizationLineStoreForm) ShouldBind(ctx *gin.Context) OrganizationLineStoreForm {

	if err := ctx.ShouldBind(&receiver); err != nil {
		wrongs.ThrowValidate("表单绑定失败：%s", err.Error())
	}
	if receiver.UniqueCode == "" {
		wrongs.ThrowValidate("线别代码必填")
	} else {
		if len(receiver.UniqueCode) != 5 {
			wrongs.ThrowValidate("线别代码必须是5位")
		}
	}
	if receiver.Name == "" {
		wrongs.ThrowValidate("线别名称必填")
	}
	if receiver.OrganizationParagraphUuid == "" {
		wrongs.ThrowValidate("所属站段编号不能为空")
	} else {
		ret := models.NewOrganizationParagraphMdl().GetDb("").Where("uuid = ?", receiver.OrganizationParagraphUuid).First(&receiver.organizationParagraph)
		wrongs.ThrowWhenEmpty(ret, "所属站段")
	}
	if len(receiver.OrganizationStationUuids) > 0 {
		models.NewOrganizationStationMdl().GetDb("").Where("uuid in ?", receiver.OrganizationStationUuids).Find(&receiver.organizationStations)
	}
	if len(receiver.OrganizationCrossroadUuids) > 0 {
		models.NewOrganizationCrossroadMdl().GetDb("").Where("uuid in ?", receiver.OrganizationCrossroadUuids).Find(&receiver.organizationCrossroads)
	}
	if len(receiver.OrganizationCenterUuids) > 0 {
		models.NewOrganizationCenterMdl().GetDb("").Where("uuid in ?", receiver.OrganizationCenterUuids).Find(&receiver.organizationCenters)
	}
	return receiver
}

// ShouldBind 表单绑定（批量删除线别）
func (receiver OrganizationLineDestroyManyForm) ShouldBind(ctx *gin.Context) OrganizationLineDestroyManyForm {
	if err := ctx.ShouldBind(&receiver); err != nil {
		wrongs.ThrowValidate("表单验证失败：%s", err.Error())
	}
	if len(receiver.Uuids) == 0 {
		wrongs.ThrowValidate("线别编号必填")
	}
	return receiver
}

// NewOrganizationRailwayCtrl 构造函数
func NewOrganizationRailwayCtrl() *OrganizationRailwayCtrl {
	return &OrganizationRailwayCtrl{}
}

// Store 新建
func (OrganizationRailwayCtrl) Store(ctx *gin.Context) {
	var (
		ret    *gorm.DB
		repeat *models.OrganizationRailwayMdl
	)

	// 表单
	form := OrganizationRailwayStoreForm{}.ShouldBind(ctx)

	// 查重
	ret = models.NewOrganizationRailwayMdl().
		GetDb("").
		Where("unique_code = ?", form.UniqueCode).
		First(&repeat)
	wrongs.ThrowWhenNotEmpty(ret, "路局代码")
	ret = models.NewOrganizationRailwayMdl().
		GetDb("").
		Where("name = ?", form.Name).
		First(&repeat)
	wrongs.ThrowWhenNotEmpty(ret, "路局名称")
	ret = models.NewOrganizationRailwayMdl().
		GetDb("").
		Where("short_name = ?", form.ShortName).
		First(&repeat)
	wrongs.ThrowWhenNotEmpty(ret, "路局简称")

	// 新建
	organizationRailway := &models.OrganizationRailwayMdl{
		UniqueCode: form.UniqueCode,
		Name:       form.Name,
		ShortName:  form.ShortName,
	}
	if ret = models.NewOrganizationRailwayMdl().
		GetDb("").
		Create(&organizationRailway); ret.Error != nil {
		wrongs.ThrowForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Created(map[string]any{"organization_railway": organizationRailway}).ToGinResponse())
}

// Destroy 删除
func (OrganizationRailwayCtrl) Destroy(ctx *gin.Context) {
	var (
		ret                 *gorm.DB
		organizationRailway *models.OrganizationRailwayMdl
	)

	// 查询
	ret = models.NewOrganizationRailwayMdl().
		GetDb("").
		Where("uuid = ?", ctx.Param("uuid")).
		First(&organizationRailway)
	wrongs.ThrowWhenEmpty(ret, "路局")

	// 删除
	if ret := models.NewOrganizationRailwayMdl().
		GetDb("").
		Where("uuid = ?", ctx.Param("uuid")).
		Delete(&organizationRailway); ret.Error != nil {
		wrongs.ThrowForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Deleted().ToGinResponse())
}

// DestroyMany 批量删除
func (OrganizationRailwayCtrl) DestroyMany(ctx *gin.Context) {
	form := OrganizaitonRailwayDestroyManyForm{}.ShouldBind(ctx)

	if ret := models.NewOrganizationRailwayMdl().GetDb("").Where("uuid in ?", form.Uuids).Delete(nil); ret.Error != nil {
		wrongs.ThrowForbidden("删除失败：%s", ret.Error.Error())
	}

	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Deleted().ToGinResponse())
}

// Update 编辑
func (OrganizationRailwayCtrl) Update(ctx *gin.Context) {
	var (
		ret                         *gorm.DB
		organizationRailway, repeat *models.OrganizationRailwayMdl
	)

	// 表单
	form := OrganizationRailwayStoreForm{}.ShouldBind(ctx)

	// 查重
	ret = models.NewOrganizationRailwayMdl().
		GetDb("").
		Where("unique_code = ? and uuid <> ?", form.UniqueCode, ctx.Param("uuid")).
		First(&repeat)
	wrongs.ThrowWhenNotEmpty(ret, "路局代码")
	ret = models.NewOrganizationRailwayMdl().
		GetDb("").
		Where("name = ? and uuid <> ?", form.Name, ctx.Param("uuid")).
		First(&repeat)
	wrongs.ThrowWhenNotEmpty(ret, "路局名称")
	ret = models.NewOrganizationRailwayMdl().
		GetDb("").
		Where("short_name = ? and uuid <> ?", form.ShortName, ctx.Param("uuid")).
		First(&repeat)
	wrongs.ThrowWhenNotEmpty(ret, "路局简称")

	// 查询
	ret = models.NewOrganizationRailwayMdl().
		GetDb("").
		Where("uuid = ?", ctx.Param("uuid")).
		First(&organizationRailway)
	wrongs.ThrowWhenEmpty(ret, "路局")

	// 编辑
	organizationRailway.UniqueCode = form.UniqueCode
	organizationRailway.Name = form.Name
	organizationRailway.ShortName = form.ShortName
	if ret = models.NewOrganizationRailwayMdl().
		GetDb("").
		Where("uuid = ?", ctx.Param("uuid")).
		Save(&organizationRailway); ret.Error != nil {
		wrongs.ThrowForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Updated(map[string]any{"organization_railway": organizationRailway}).ToGinResponse())
}

// Detail 详情
func (OrganizationRailwayCtrl) Detail(ctx *gin.Context) {
	var (
		ret                 *gorm.DB
		organizationRailway *models.OrganizationRailwayMdl
	)
	ret = models.NewOrganizationRailwayMdl().
		SetCtx(ctx).
		GetDbUseQuery("").
		Where("uuid = ?", ctx.Param("uuid")).
		First(&organizationRailway)
	wrongs.ThrowWhenEmpty(ret, "路局")

	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Datum(map[string]any{"organization_railway": organizationRailway}).ToGinResponse())
}

// List 列表
func (receiver OrganizationRailwayCtrl) List(ctx *gin.Context) {
	var organizationRailways []*models.OrganizationRailwayMdl

	ctx.JSON(
		tools.NewCorrectWithGinContext("", ctx).
			DataForPager(
				models.OrganizationRailwayMdl{}.GetListByQuery(ctx),
				func(db *gorm.DB) map[string]any {
					db.Find(&organizationRailways)
					return map[string]any{"organization_railways": organizationRailways}
				},
			).
			ToGinResponse(),
	)
}

// ListJdt jquery-dataTable后端分页数据
func (receiver OrganizationRailwayCtrl) ListJdt(ctx *gin.Context) {
	var organizationRailways []*models.OrganizationRailwayMdl

	ctx.JSON(
		tools.NewCorrectWithGinContext("", ctx).
			DataForJqueryDataTable(
				models.OrganizationRailwayMdl{}.GetListByQuery(ctx),
				func(db *gorm.DB) map[string]any {
					db.Find(&organizationRailways)
					return map[string]any{"organization_railways": organizationRailways}
				},
			).
			ToGinResponse(),
	)
}

// NewOrganizationParagraphCtrl 构造函数
func NewOrganizationParagraphCtrl() *OrganizationParagraphCtrl {
	return &OrganizationParagraphCtrl{}
}

// Store 新建
func (OrganizationParagraphCtrl) Store(ctx *gin.Context) {
	var (
		ret    *gorm.DB
		repeat *models.OrganizationParagraphMdl
	)

	// 表单
	form := OrganizationParagraphStoreForm{}.ShouldBind(ctx)

	// 查重
	ret = models.NewOrganizationParagraphMdl().
		GetDb("").
		Where("unique_code = ?", form.UniqueCode).
		First(&repeat)
	wrongs.ThrowWhenNotEmpty(ret, "站段代码")
	ret = models.NewOrganizationParagraphMdl().
		GetDb("").
		Where("name = ?", form.Name).
		First(&repeat)
	wrongs.ThrowWhenNotEmpty(ret, "站段名称")

	// 新建
	organizationParagraph := &models.OrganizationParagraphMdl{
		UniqueCode:              form.UniqueCode,
		Name:                    form.Name,
		OrganizationRailwayUuid: form.organizationRailway.Uuid,
	}
	if ret = models.NewOrganizationParagraphMdl().
		GetDb("").
		Create(&organizationParagraph); ret.Error != nil {
		wrongs.ThrowForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Created(map[string]any{"organization_paragraph": organizationParagraph}).ToGinResponse())
}

// Destroy 删除
func (OrganizationParagraphCtrl) Destroy(ctx *gin.Context) {
	var (
		ret                   *gorm.DB
		organizationParagraph *models.OrganizationParagraphMdl
	)

	// 查询
	ret = models.NewOrganizationParagraphMdl().
		GetDb("").
		Where("uuid = ?", ctx.Param("uuid")).
		First(&organizationParagraph)
	wrongs.ThrowWhenEmpty(ret, "站段")

	// 删除
	if ret := models.NewOrganizationParagraphMdl().
		GetDb("").
		Where("uuid = ?", ctx.Param("uuid")).
		Delete(&organizationParagraph); ret.Error != nil {
		wrongs.ThrowForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Deleted().ToGinResponse())
}

// DestroyMany 批量删除
func (OrganizationParagraphCtrl) DestroyMany(ctx *gin.Context) {
	form := OrganizaitonParagraphDestroyManyForm{}.ShouldBind(ctx)
	models.NewOrganizationRailwayMdl().GetDb("").Where("uuid in ?", form.Uuids).Delete(nil)

	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Deleted().ToGinResponse())
}

// Update 编辑
func (OrganizationParagraphCtrl) Update(ctx *gin.Context) {
	var (
		ret                           *gorm.DB
		organizationParagraph, repeat *models.OrganizationParagraphMdl
	)

	// 表单
	form := OrganizationParagraphStoreForm{}.ShouldBind(ctx)

	// 查重
	ret = models.NewOrganizationParagraphMdl().
		GetDb("").
		Where("unique_code = ? and uuid <> ?", form.UniqueCode, ctx.Param("uuid")).
		First(&repeat)
	wrongs.ThrowWhenNotEmpty(ret, "站段代码")
	ret = models.NewOrganizationParagraphMdl().
		GetDb("").
		Where("name = ? and uuid <> ?", form.Name, ctx.Param("uuid")).
		First(&repeat)
	wrongs.ThrowWhenNotEmpty(ret, "站段名称")

	// 查询
	ret = models.NewOrganizationParagraphMdl().
		GetDb("").
		Where("uuid = ?", ctx.Param("uuid")).
		First(&organizationParagraph)
	wrongs.ThrowWhenEmpty(ret, "站段")

	// 编辑
	organizationParagraph.UniqueCode = form.UniqueCode
	organizationParagraph.Name = form.Name
	organizationParagraph.OrganizationRailwayUuid = form.organizationRailway.Uuid
	if ret = models.NewOrganizationParagraphMdl().
		GetDb("").
		Where("uuid = ?", ctx.Param("uuid")).
		Save(&organizationParagraph); ret.Error != nil {
		wrongs.ThrowForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Updated(map[string]any{"organization_paragraph": organizationParagraph}).ToGinResponse())
}

// Detail 详情
func (OrganizationParagraphCtrl) Detail(ctx *gin.Context) {
	var (
		ret                   *gorm.DB
		organizationParagraph *models.OrganizationParagraphMdl
	)
	ret = models.NewOrganizationParagraphMdl().
		SetCtx(ctx).
		GetDbUseQuery("").
		Where("uuid = ?", ctx.Param("uuid")).
		First(&organizationParagraph)
	wrongs.ThrowWhenEmpty(ret, "站段")

	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Datum(map[string]any{"organization_paragraph": organizationParagraph}).ToGinResponse())
}

// List 列表
func (receiver OrganizationParagraphCtrl) List(ctx *gin.Context) {
	var organizationParagraphs []*models.OrganizationParagraphMdl

	ctx.JSON(
		tools.NewCorrectWithGinContext("", ctx).
			DataForPager(
				models.OrganizationParagraphMdl{}.GetListByQuery(ctx),
				func(db *gorm.DB) map[string]any {
					db.Find(&organizationParagraphs)
					return map[string]any{"organization_paragraphs": organizationParagraphs}
				},
			).
			ToGinResponse(),
	)
}

// ListJdt jquery-dataTable后端分页数据
func (receiver OrganizationParagraphCtrl) ListJdt(ctx *gin.Context) {
	var organizationParagraphs []*models.OrganizationParagraphMdl

	ctx.JSON(
		tools.NewCorrectWithGinContext("", ctx).
			DataForJqueryDataTable(
				models.OrganizationParagraphMdl{}.GetListByQuery(ctx),
				func(db *gorm.DB) map[string]any {
					db.Find(&organizationParagraphs)
					return map[string]any{"organization_paragraphs": organizationParagraphs}
				},
			).
			ToGinResponse(),
	)
}

// NewOrganizationWorkshopCtrl 构造函数
func NewOrganizationWorkshopCtrl() *OrganizationWorkshopCtrl {
	return &OrganizationWorkshopCtrl{}
}

// Store 新建
func (OrganizationWorkshopCtrl) Store(ctx *gin.Context) {
	var (
		ret    *gorm.DB
		repeat *models.OrganizationWorkshopMdl
	)

	// 表单
	form := OrganizationWorkshopStoreForm{}.ShouldBind(ctx)

	// 查重
	ret = models.NewOrganizationWorkshopMdl().
		GetDb("").
		Where("unique_code = ?", form.UniqueCode).
		First(&repeat)
	wrongs.ThrowWhenNotEmpty(ret, "车间代码")
	ret = models.NewOrganizationWorkshopMdl().
		GetDb("").
		Where("name = ?", form.Name).
		First(&repeat)
	wrongs.ThrowWhenNotEmpty(ret, "车间名称")

	// 新建
	organizationWorkshop := &models.OrganizationWorkshopMdl{
		UniqueCode:                form.UniqueCode,
		Name:                      form.Name,
		OrganizationParagraphUuid: form.organizationParagraph.Uuid,
		TypeCode:                  form.TypeCode,
	}
	if ret = models.NewOrganizationWorkshopMdl().
		GetDb("").
		Create(&organizationWorkshop); ret.Error != nil {
		wrongs.ThrowForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Created(map[string]any{"organization_workshop": organizationWorkshop}).ToGinResponse())
}

// Destroy 删除
func (OrganizationWorkshopCtrl) Destroy(ctx *gin.Context) {
	var (
		ret                  *gorm.DB
		organizationWorkshop *models.OrganizationWorkshopMdl
	)

	// 查询
	ret = models.NewOrganizationWorkshopMdl().
		GetDb("").
		Where("uuid = ?", ctx.Param("uuid")).
		First(&organizationWorkshop)
	wrongs.ThrowWhenEmpty(ret, "车间")

	// 删除
	if ret := models.NewOrganizationWorkshopMdl().
		GetDb("").
		Where("uuid = ?", ctx.Param("uuid")).
		Delete(&organizationWorkshop); ret.Error != nil {
		wrongs.ThrowForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Deleted().ToGinResponse())
}

// DestroyMany 批量删除
func (OrganizationWorkshopCtrl) DestroyMany(ctx *gin.Context) {
	form := OrganizationWorkshopDestroyManyForm{}.ShouldBind(ctx)
	if ret := models.NewOrganizationWorkshopMdl().GetDb("").Where("uuid in ?", form.Uuids).Delete(nil); ret.Error != nil {
		wrongs.ThrowForbidden("删除失败：%s", ret.Error.Error())
	}
	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Deleted().ToGinResponse())
}

// Update 编辑
func (OrganizationWorkshopCtrl) Update(ctx *gin.Context) {
	var (
		ret                          *gorm.DB
		organizationWorkshop, repeat *models.OrganizationWorkshopMdl
	)

	// 表单
	form := OrganizationWorkshopStoreForm{}.ShouldBind(ctx)

	// 查重
	ret = models.NewOrganizationWorkshopMdl().
		GetDb("").
		Where("unique_code = ? and uuid <> ?", form.UniqueCode, ctx.Param("uuid")).
		First(&repeat)
	wrongs.ThrowWhenNotEmpty(ret, "车间代码")
	ret = models.NewOrganizationWorkshopMdl().
		GetDb("").
		Where("name = ? and uuid <> ?", form.Name, ctx.Param("uuid")).
		First(&repeat)
	wrongs.ThrowWhenNotEmpty(ret, "车间名称")

	// 查询
	ret = models.NewOrganizationWorkshopMdl().
		GetDb("").
		Where("uuid = ?", ctx.Param("uuid")).
		First(&organizationWorkshop)
	wrongs.ThrowWhenEmpty(ret, "车间")

	// 编辑
	organizationWorkshop.UniqueCode = form.UniqueCode
	organizationWorkshop.Name = form.Name
	organizationWorkshop.OrganizationParagraphUuid = form.organizationParagraph.Uuid
	organizationWorkshop.TypeCode = form.TypeCode
	if ret = models.NewOrganizationWorkshopMdl().
		GetDb("").
		Where("uuid = ?", ctx.Param("uuid")).
		Save(&organizationWorkshop); ret.Error != nil {
		wrongs.ThrowForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Updated(map[string]any{"organization_workshop": organizationWorkshop}).ToGinResponse())
}

// Detail 详情
func (OrganizationWorkshopCtrl) Detail(ctx *gin.Context) {
	var (
		ret                  *gorm.DB
		organizationWorkshop *models.OrganizationWorkshopMdl
	)
	ret = models.NewOrganizationWorkshopMdl().
		SetCtx(ctx).
		GetDbUseQuery("").
		Where("uuid = ?", ctx.Param("uuid")).
		First(&organizationWorkshop)
	wrongs.ThrowWhenEmpty(ret, "车间")

	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Datum(map[string]any{"organization_workshop": organizationWorkshop}).ToGinResponse())
}

// List 列表
func (receiver OrganizationWorkshopCtrl) List(ctx *gin.Context) {
	var organizationWorkshops []*models.OrganizationWorkshopMdl

	ctx.JSON(
		tools.NewCorrectWithGinContext("", ctx).
			DataForPager(
				models.OrganizationWorkshopMdl{}.GetListByQuery(ctx),
				func(db *gorm.DB) map[string]any {
					db.Find(&organizationWorkshops)
					return map[string]any{"organization_workshops": organizationWorkshops}
				},
			).
			ToGinResponse(),
	)
}

// ListJdt jquery-dataTable后端分页数据
func (receiver OrganizationWorkshopCtrl) ListJdt(ctx *gin.Context) {
	var organizationWorkshops []*models.OrganizationWorkshopMdl

	ctx.JSON(
		tools.NewCorrectWithGinContext("", ctx).
			DataForJqueryDataTable(
				models.OrganizationWorkshopMdl{}.GetListByQuery(ctx),
				func(db *gorm.DB) map[string]any {
					db.Find(&organizationWorkshops)
					return map[string]any{"organization_workshops": organizationWorkshops}
				},
			).
			ToGinResponse(),
	)
}

// GetTypeCodes 获取车间类型代码
func (OrganizationWorkshopCtrl) GetTypeCodesMap(ctx *gin.Context) {
	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Datum(map[string]any{"type_codes_map": models.OrganizationWorkshopMdl{}.GetTypeCodesMap()}).ToGinResponse())
}

// NewOrganizationStationCtrl 构造函数
func NewOrganizationStationCtrl() *OrganizationStationCtrl {
	return &OrganizationStationCtrl{}
}

// Store 新建
func (OrganizationStationCtrl) Store(ctx *gin.Context) {
	var (
		ret    *gorm.DB
		repeat *models.OrganizationStationMdl
	)

	// 表单
	form := OrganizationStationStoreForm{}.ShouldBind(ctx)

	// 查重
	ret = models.NewOrganizationStationMdl().
		GetDb("").
		Where("unique_code = ?", form.UniqueCode).
		First(&repeat)
	wrongs.ThrowWhenNotEmpty(ret, "站场代码")
	ret = models.NewOrganizationStationMdl().
		GetDb("").
		Where("name = ?", form.Name).
		First(&repeat)
	wrongs.ThrowWhenNotEmpty(ret, "站场名称")

	// 新建
	organizationStation := &models.OrganizationStationMdl{
		UniqueCode:               form.UniqueCode,
		Name:                     form.Name,
		OrganizationWorkshopUuid: form.organizationWorkshop.Uuid,
		OrganizationLineUuid:     form.OrganizationLineUuid,
	}
	if ret = models.NewOrganizationStationMdl().
		GetDb("").
		Create(&organizationStation); ret.Error != nil {
		wrongs.ThrowForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Created(map[string]any{"organization_station": organizationStation}).ToGinResponse())
}

// Destroy 删除
func (OrganizationStationCtrl) Destroy(ctx *gin.Context) {
	var (
		ret                 *gorm.DB
		organizationStation *models.OrganizationStationMdl
	)

	// 查询
	ret = models.NewOrganizationStationMdl().
		GetDb("").
		Where("uuid = ?", ctx.Param("uuid")).
		First(&organizationStation)
	wrongs.ThrowWhenEmpty(ret, "站场")

	// 删除
	if ret := models.NewOrganizationStationMdl().
		GetDb("").
		Where("uuid = ?", ctx.Param("uuid")).
		Delete(&organizationStation); ret.Error != nil {
		wrongs.ThrowForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Deleted().ToGinResponse())
}

// DestroyMany 批量删除
func (OrganizationStationCtrl) DestroyMany(ctx *gin.Context) {
	form := OrganizationStationDestroyManyForm{}.ShouldBind(ctx)
	if ret := models.NewOrganizationStationMdl().GetDb("").Where("uuid in ?", form.Uuids).Delete(nil); ret.Error != nil {
		wrongs.ThrowForbidden("删除失败：%s", ret.Error.Error())
	}
	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Deleted().ToGinResponse())
}

// Update 编辑
func (OrganizationStationCtrl) Update(ctx *gin.Context) {
	var (
		ret                         *gorm.DB
		organizationStation, repeat *models.OrganizationStationMdl
	)

	// 表单
	form := OrganizationStationStoreForm{}.ShouldBind(ctx)

	// 查重
	ret = models.NewOrganizationStationMdl().
		GetDb("").
		Where("unique_code = ? and uuid <> ?", form.UniqueCode, ctx.Param("uuid")).
		First(&repeat)
	wrongs.ThrowWhenNotEmpty(ret, "站场代码")
	ret = models.NewOrganizationStationMdl().
		GetDb("").
		Where("name = ? and uuid <> ?", form.Name, ctx.Param("uuid")).
		First(&repeat)
	wrongs.ThrowWhenNotEmpty(ret, "站场名称")

	// 查询
	ret = models.NewOrganizationStationMdl().
		GetDb("").
		Where("uuid = ?", ctx.Param("uuid")).
		First(&organizationStation)
	wrongs.ThrowWhenEmpty(ret, "站场")

	// 编辑
	organizationStation.UniqueCode = form.UniqueCode
	organizationStation.Name = form.Name
	organizationStation.OrganizationWorkshopUuid = form.organizationWorkshop.Uuid
	organizationStation.OrganizationLineUuid = form.OrganizationLineUuid
	if ret = models.NewOrganizationStationMdl().
		GetDb("").
		Where("uuid = ?", ctx.Param("uuid")).
		Save(&organizationStation); ret.Error != nil {
		wrongs.ThrowForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Updated(map[string]any{"organization_station": organizationStation}).ToGinResponse())
}

// Detail 详情
func (OrganizationStationCtrl) Detail(ctx *gin.Context) {
	var (
		ret                 *gorm.DB
		organizationStation *models.OrganizationStationMdl
	)
	ret = models.NewOrganizationStationMdl().
		SetCtx(ctx).
		GetDbUseQuery("").
		Where("uuid = ?", ctx.Param("uuid")).
		First(&organizationStation)
	wrongs.ThrowWhenEmpty(ret, "站场")

	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Datum(map[string]any{"organization_station": organizationStation}).ToGinResponse())
}

// List 列表
func (receiver OrganizationStationCtrl) List(ctx *gin.Context) {
	var organizationStations []*models.OrganizationStationMdl

	ctx.JSON(
		tools.NewCorrectWithGinContext("", ctx).
			DataForPager(
				models.OrganizationStationMdl{}.GetListByQuery(ctx),
				func(db *gorm.DB) map[string]any {
					db.Find(&organizationStations)
					return map[string]any{"organization_stations": organizationStations}
				},
			).
			ToGinResponse(),
	)
}

// ListJdt jquery-dataTable后端分页数据
func (receiver OrganizationStationCtrl) ListJdt(ctx *gin.Context) {
	var organizationStations []*models.OrganizationStationMdl

	ctx.JSON(
		tools.NewCorrectWithGinContext("", ctx).
			DataForJqueryDataTable(
				models.OrganizationStationMdl{}.GetListByQuery(ctx),
				func(db *gorm.DB) map[string]any {
					db.Find(&organizationStations)
					return map[string]any{"organization_stations": organizationStations}
				},
			).
			ToGinResponse(),
	)
}

// NewOrganizationCrossroadCtrl 构造函数
func NewOrganizationCrossroadCtrl() *OrganizationCrossroadCtrl {
	return &OrganizationCrossroadCtrl{}
}

// Store 新建
func (OrganizationCrossroadCtrl) Store(ctx *gin.Context) {
	var (
		ret    *gorm.DB
		repeat *models.OrganizationCrossroadMdl
	)

	// 表单
	form := OrganizationCrossroadStoreForm{}.ShouldBind(ctx)

	// 查重
	ret = models.NewOrganizationCrossroadMdl().
		GetDb("").
		Where("unique_code = ?", form.UniqueCode).
		First(&repeat)
	wrongs.ThrowWhenNotEmpty(ret, "道口代码")
	ret = models.NewOrganizationCrossroadMdl().
		GetDb("").
		Where("name = ?", form.Name).
		First(&repeat)
	wrongs.ThrowWhenNotEmpty(ret, "道口名称")

	// 新建
	organizationCrossroad := &models.OrganizationCrossroadMdl{
		UniqueCode:               form.UniqueCode,
		Name:                     form.Name,
		OrganizationWorkshopUuid: form.organizationWorkshop.Uuid,
		OrganizationLineUuid:     form.OrganizationLineUuid,
	}
	if ret = models.NewOrganizationCrossroadMdl().
		GetDb("").
		Create(&organizationCrossroad); ret.Error != nil {
		wrongs.ThrowForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Created(map[string]any{"organization_crossroad": organizationCrossroad}).ToGinResponse())
}

// Destroy 删除
func (OrganizationCrossroadCtrl) Destroy(ctx *gin.Context) {
	var (
		ret                   *gorm.DB
		organizationCrossroad *models.OrganizationCrossroadMdl
	)

	// 查询
	ret = models.NewOrganizationCrossroadMdl().
		GetDb("").
		Where("uuid = ?", ctx.Param("uuid")).
		First(&organizationCrossroad)
	wrongs.ThrowWhenEmpty(ret, "道口")

	// 删除
	if ret := models.NewOrganizationCrossroadMdl().
		GetDb("").
		Where("uuid = ?", ctx.Param("uuid")).
		Delete(&organizationCrossroad); ret.Error != nil {
		wrongs.ThrowForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Deleted().ToGinResponse())
}

// 批量删除
func (OrganizationCrossroadCtrl) DestroyMany(ctx *gin.Context) {
	form := OrganizationCrossroadDestroyManyForm{}.ShouldBind(ctx)
	if ret := models.NewOrganizationCrossroadMdl().GetDb("").Where("uuid in ?", form.Uuids).Delete(nil); ret.Error != nil {
		wrongs.ThrowForbidden("删除失败：%s", ret.Error.Error())
	}
	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Deleted().ToGinResponse())
}

// Update 编辑
func (OrganizationCrossroadCtrl) Update(ctx *gin.Context) {
	var (
		ret                           *gorm.DB
		organizationCrossroad, repeat *models.OrganizationCrossroadMdl
	)

	// 表单
	form := OrganizationCrossroadStoreForm{}.ShouldBind(ctx)

	// 查重
	ret = models.NewOrganizationCrossroadMdl().
		GetDb("").
		Where("unique_code = ? and uuid <> ?", form.UniqueCode, ctx.Param("uuid")).
		First(&repeat)
	wrongs.ThrowWhenNotEmpty(ret, "道口代码")
	ret = models.NewOrganizationCrossroadMdl().
		GetDb("").
		Where("name = ? and uuid <> ?", form.Name, ctx.Param("uuid")).
		First(&repeat)
	wrongs.ThrowWhenNotEmpty(ret, "道口名称")

	// 查询
	ret = models.NewOrganizationCrossroadMdl().
		GetDb("").
		Where("uuid = ?", ctx.Param("uuid")).
		First(&organizationCrossroad)
	wrongs.ThrowWhenEmpty(ret, "道口")

	// 编辑
	organizationCrossroad.UniqueCode = form.UniqueCode
	organizationCrossroad.Name = form.Name
	organizationCrossroad.OrganizationWorkshopUuid = form.organizationWorkshop.Uuid
	organizationCrossroad.OrganizationLineUuid = form.OrganizationLineUuid
	if ret = models.NewOrganizationCrossroadMdl().
		GetDb("").
		Where("uuid = ?", ctx.Param("uuid")).
		Save(&organizationCrossroad); ret.Error != nil {
		wrongs.ThrowForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Updated(map[string]any{"organization_crossroad": organizationCrossroad}).ToGinResponse())
}

// Detail 详情
func (OrganizationCrossroadCtrl) Detail(ctx *gin.Context) {
	var (
		ret                   *gorm.DB
		organizationCrossroad *models.OrganizationCrossroadMdl
	)
	ret = models.NewOrganizationCrossroadMdl().
		SetCtx(ctx).
		GetDbUseQuery("").
		Where("uuid = ?", ctx.Param("uuid")).
		First(&organizationCrossroad)
	wrongs.ThrowWhenEmpty(ret, "道口")

	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Datum(map[string]any{"organization_crossroad": organizationCrossroad}).ToGinResponse())
}

// List 列表
func (receiver OrganizationCrossroadCtrl) List(ctx *gin.Context) {
	var organizationCrossroads []*models.OrganizationCrossroadMdl

	ctx.JSON(
		tools.NewCorrectWithGinContext("", ctx).
			DataForPager(
				models.OrganizationCrossroadMdl{}.GetListByQuery(ctx),
				func(db *gorm.DB) map[string]any {
					db.Find(&organizationCrossroads)
					return map[string]any{"organization_crossroads": organizationCrossroads}
				},
			).
			ToGinResponse(),
	)
}

// ListJdt jquery-dataTable后端分页数据
func (receiver OrganizationCrossroadCtrl) ListJdt(ctx *gin.Context) {
	var organizationCrossroads []*models.OrganizationCrossroadMdl

	ctx.JSON(
		tools.NewCorrectWithGinContext("", ctx).
			DataForJqueryDataTable(
				models.OrganizationCrossroadMdl{}.GetListByQuery(ctx),
				func(db *gorm.DB) map[string]any {
					db.Find(&organizationCrossroads)
					return map[string]any{"organization_crossroads": organizationCrossroads}
				},
			).
			ToGinResponse(),
	)
}

// NewOrganizationCenterCtrl 构造函数
func NewOrganizationCenterCtrl() *OrganizationCenterCtrl {
	return &OrganizationCenterCtrl{}
}

// Store 新建
func (OrganizationCenterCtrl) Store(ctx *gin.Context) {
	var (
		ret    *gorm.DB
		repeat *models.OrganizationCenterMdl
	)

	// 表单
	form := OrganizationCenterStoreForm{}.ShouldBind(ctx)

	// 查重
	ret = models.NewOrganizationCenterMdl().
		GetDb("").
		Where("unique_code = ?", form.UniqueCode).
		First(&repeat)
	wrongs.ThrowWhenNotEmpty(ret, "中心代码")
	ret = models.NewOrganizationCenterMdl().
		GetDb("").
		Where("name = ?", form.Name).
		First(&repeat)
	wrongs.ThrowWhenNotEmpty(ret, "中心名称")

	// 新建
	organizationCenter := &models.OrganizationCenterMdl{
		UniqueCode:               form.UniqueCode,
		Name:                     form.Name,
		OrganizationWorkshopUuid: form.organizationWorkshop.Uuid,
		OrganizationLineUuid:     form.OrganizationLineUuid,
	}
	if ret = models.NewOrganizationCenterMdl().
		GetDb("").
		Create(&organizationCenter); ret.Error != nil {
		wrongs.ThrowForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Created(map[string]any{"organization_center": organizationCenter}).ToGinResponse())
}

// Destroy 删除
func (OrganizationCenterCtrl) Destroy(ctx *gin.Context) {
	var (
		ret                *gorm.DB
		organizationCenter *models.OrganizationCenterMdl
	)

	// 查询
	ret = models.NewOrganizationCenterMdl().
		GetDb("").
		Where("uuid = ?", ctx.Param("uuid")).
		First(&organizationCenter)
	wrongs.ThrowWhenEmpty(ret, "中心")

	// 删除
	if ret := models.NewOrganizationCenterMdl().
		GetDb("").
		Where("uuid = ?", ctx.Param("uuid")).
		Delete(&organizationCenter); ret.Error != nil {
		wrongs.ThrowForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Deleted().ToGinResponse())
}

// DestroyMany 批量删除
func (OrganizationCenterCtrl) DestroyMany(ctx *gin.Context) {
	form := OrganizationCenterDestroyManyForm{}.ShouldBind(ctx)
	if ret := models.NewOrganizationCenterMdl().GetDb("").Where("uuid in ?", form.Uuids).Delete(nil); ret.Error != nil {
		wrongs.ThrowForbidden("删除失败：%s", ret.Error.Error())
	}
	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Deleted().ToGinResponse())
}

// Update 编辑
func (OrganizationCenterCtrl) Update(ctx *gin.Context) {
	var (
		ret                        *gorm.DB
		organizationCenter, repeat *models.OrganizationCenterMdl
	)

	// 表单
	form := OrganizationCenterStoreForm{}.ShouldBind(ctx)

	// 查重
	ret = models.NewOrganizationCenterMdl().
		GetDb("").
		Where("unique_code = ? and uuid <> ?", form.UniqueCode, ctx.Param("uuid")).
		First(&repeat)
	wrongs.ThrowWhenNotEmpty(ret, "中心代码")
	ret = models.NewOrganizationCenterMdl().
		GetDb("").
		Where("name = ? and uuid <> ?", form.Name, ctx.Param("uuid")).
		First(&repeat)
	wrongs.ThrowWhenNotEmpty(ret, "中心名称")

	// 查询
	ret = models.NewOrganizationCenterMdl().
		GetDb("").
		Where("uuid = ?", ctx.Param("uuid")).
		First(&organizationCenter)
	wrongs.ThrowWhenEmpty(ret, "中心")

	// 编辑
	organizationCenter.UniqueCode = form.UniqueCode
	organizationCenter.Name = form.Name
	organizationCenter.OrganizationWorkshopUuid = form.organizationWorkshop.Uuid
	organizationCenter.OrganizationLineUuid = form.OrganizationLineUuid
	if ret = models.NewOrganizationCenterMdl().
		GetDb("").
		Where("uuid = ?", ctx.Param("uuid")).
		Save(&organizationCenter); ret.Error != nil {
		wrongs.ThrowForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Updated(map[string]any{"organization_center": organizationCenter}).ToGinResponse())
}

// Detail 详情
func (OrganizationCenterCtrl) Detail(ctx *gin.Context) {
	var (
		ret                *gorm.DB
		organizationCenter *models.OrganizationCenterMdl
	)
	ret = models.NewOrganizationCenterMdl().
		SetCtx(ctx).
		GetDbUseQuery("").
		Where("uuid = ?", ctx.Param("uuid")).
		First(&organizationCenter)
	wrongs.ThrowWhenEmpty(ret, "中心")

	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Datum(map[string]any{"organization_center": organizationCenter}).ToGinResponse())
}

// List 列表
func (receiver OrganizationCenterCtrl) List(ctx *gin.Context) {
	var organizationCenters []*models.OrganizationCenterMdl

	ctx.JSON(
		tools.NewCorrectWithGinContext("", ctx).
			DataForPager(
				models.OrganizationCenterMdl{}.GetListByQuery(ctx),
				func(db *gorm.DB) map[string]any {
					db.Find(&organizationCenters)
					return map[string]any{"organization_centers": organizationCenters}
				},
			).
			ToGinResponse(),
	)
}

// ListJdt jquery-dataTable后端分页数据
func (receiver OrganizationCenterCtrl) ListJdt(ctx *gin.Context) {
	var organizationCenters []*models.OrganizationCenterMdl

	ctx.JSON(
		tools.NewCorrectWithGinContext("", ctx).
			DataForJqueryDataTable(
				models.OrganizationCenterMdl{}.GetListByQuery(ctx),
				func(db *gorm.DB) map[string]any {
					db.Find(&organizationCenters)
					return map[string]any{"organization_centers": organizationCenters}
				},
			).
			ToGinResponse(),
	)
}

// NewOrganizationWorkAreaCtrl 构造函数
func NewOrganizationWorkAreaCtrl() *OrganizationWorkAreaCtrl {
	return &OrganizationWorkAreaCtrl{}
}

// Store 新建
func (OrganizationWorkAreaCtrl) Store(ctx *gin.Context) {
	var (
		ret    *gorm.DB
		repeat *models.OrganizationWorkAreaMdl
	)

	// 表单
	form := OrganizationWorkAreaStoreForm{}.ShouldBind(ctx)

	// 查重
	ret = models.NewOrganizationWorkAreaMdl().
		GetDb("").
		Where("unique_code = ?", form.UniqueCode).
		First(&repeat)
	wrongs.ThrowWhenNotEmpty(ret, "工区代码")
	ret = models.NewOrganizationWorkAreaMdl().
		GetDb("").
		Where("name = ?", form.Name).
		First(&repeat)
	wrongs.ThrowWhenNotEmpty(ret, "工区名称")

	// 新建
	organizationWorkArea := &models.OrganizationWorkAreaMdl{
		UniqueCode:               form.UniqueCode,
		Name:                     form.Name,
		OrganizationWorkshopUuid: form.organizationWorkshop.Uuid,
		TypeCode:                 form.TypeCode,
	}
	if ret = models.NewOrganizationWorkAreaMdl().
		GetDb("").
		Create(&organizationWorkArea); ret.Error != nil {
		wrongs.ThrowForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Created(map[string]any{"organization_work_area": organizationWorkArea}).ToGinResponse())
}

// Destroy 删除
func (OrganizationWorkAreaCtrl) Destroy(ctx *gin.Context) {
	var (
		ret                  *gorm.DB
		organizationWorkArea *models.OrganizationWorkAreaMdl
	)

	// 查询
	ret = models.NewOrganizationWorkAreaMdl().
		GetDb("").
		Where("uuid = ?", ctx.Param("uuid")).
		First(&organizationWorkArea)
	wrongs.ThrowWhenEmpty(ret, "工区")

	// 删除
	if ret := models.NewOrganizationWorkAreaMdl().
		GetDb("").
		Where("uuid = ?", ctx.Param("uuid")).
		Delete(&organizationWorkArea); ret.Error != nil {
		wrongs.ThrowForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Deleted().ToGinResponse())
}

// DestroyMany 批量删除
func (OrganizationWorkAreaCtrl) DestroyMany(ctx *gin.Context) {
	form := OrganizationWorkAreaDestroyManyForm{}.ShouldBind(ctx)
	if ret := models.NewOrganizationWorkAreaMdl().GetDb("").Where("uuid in ?", form.Uuids).Delete(nil); ret.Error != nil {
		wrongs.ThrowForbidden("删除失败：%s", ret.Error.Error())
	}
	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Deleted().ToGinResponse())
}

// Update 编辑
func (OrganizationWorkAreaCtrl) Update(ctx *gin.Context) {
	var (
		ret                          *gorm.DB
		organizationWorkArea, repeat *models.OrganizationWorkAreaMdl
	)

	// 表单
	form := OrganizationWorkAreaStoreForm{}.ShouldBind(ctx)

	// 查重
	ret = models.NewOrganizationWorkAreaMdl().
		GetDb("").
		Where("unique_code = ? and uuid <> ?", form.UniqueCode, ctx.Param("uuid")).
		First(&repeat)
	wrongs.ThrowWhenNotEmpty(ret, "工区代码")
	ret = models.NewOrganizationWorkAreaMdl().
		GetDb("").
		Where("name = ? and uuid <> ?", form.Name, ctx.Param("uuid")).
		First(&repeat)
	wrongs.ThrowWhenNotEmpty(ret, "工区名称")

	// 查询
	ret = models.NewOrganizationWorkAreaMdl().
		GetDb("").
		Where("uuid = ?", ctx.Param("uuid")).
		First(&organizationWorkArea)
	wrongs.ThrowWhenEmpty(ret, "工区")

	// 编辑
	organizationWorkArea.UniqueCode = form.UniqueCode
	organizationWorkArea.Name = form.Name
	organizationWorkArea.OrganizationWorkshopUuid = form.organizationWorkshop.Uuid
	organizationWorkArea.TypeCode = form.TypeCode
	if ret = models.NewOrganizationWorkAreaMdl().
		GetDb("").
		Where("uuid = ?", ctx.Param("uuid")).
		Save(&organizationWorkArea); ret.Error != nil {
		wrongs.ThrowForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Updated(map[string]any{"organization_work_area": organizationWorkArea}).ToGinResponse())
}

// Detail 详情
func (OrganizationWorkAreaCtrl) Detail(ctx *gin.Context) {
	var (
		ret                  *gorm.DB
		organizationWorkArea *models.OrganizationWorkAreaMdl
	)
	ret = models.NewOrganizationWorkAreaMdl().
		SetCtx(ctx).
		GetDbUseQuery("").
		Where("uuid = ?", ctx.Param("uuid")).
		First(&organizationWorkArea)
	wrongs.ThrowWhenEmpty(ret, "工区")

	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Datum(map[string]any{"organization_work_area": organizationWorkArea}).ToGinResponse())
}

// List 列表
func (receiver OrganizationWorkAreaCtrl) List(ctx *gin.Context) {
	var organizationWorkAreas []*models.OrganizationWorkAreaMdl

	ctx.JSON(
		tools.NewCorrectWithGinContext("", ctx).
			DataForPager(
				models.OrganizationWorkAreaMdl{}.GetListByQuery(ctx),
				func(db *gorm.DB) map[string]any {
					db.Find(&organizationWorkAreas)
					return map[string]any{"organization_work_areas": organizationWorkAreas}
				},
			).
			ToGinResponse(),
	)
}

// ListJdt jquery-dataTable后端分页数据
func (receiver OrganizationWorkAreaCtrl) ListJdt(ctx *gin.Context) {
	var organizationWorkAreas []*models.OrganizationWorkAreaMdl

	ctx.JSON(
		tools.NewCorrectWithGinContext("", ctx).
			DataForJqueryDataTable(
				models.OrganizationWorkAreaMdl{}.GetListByQuery(ctx),
				func(db *gorm.DB) map[string]any {
					db.Find(&organizationWorkAreas)
					return map[string]any{"organization_work_areas": organizationWorkAreas}
				},
			).
			ToGinResponse(),
	)
}

// GetTypeCodes 获取工区类型代码
func (OrganizationWorkAreaCtrl) GetTypeCodesMap(ctx *gin.Context) {
	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Datum(map[string]any{"type_codes_map": models.OrganizationWorkAreaMdl{}.GetTypeCodesMap()}).ToGinResponse())
}

// NewOrganizationLineCtrl 构造函数
func NewOrganizationLineCtrl() *OrganizationLineCtrl {
	return &OrganizationLineCtrl{}
}

// Store 新建
func (OrganizationLineCtrl) Store(ctx *gin.Context) {
	var (
		ret    *gorm.DB
		repeat *models.OrganizationLineMdl
	)

	// 表单
	form := OrganizationLineStoreForm{}.ShouldBind(ctx)

	// 查重
	ret = models.NewOrganizationLineMdl().
		GetDb("").
		Where("unique_code = ?", form.UniqueCode).
		First(&repeat)
	wrongs.ThrowWhenNotEmpty(ret, "线别代码")
	ret = models.NewOrganizationLineMdl().
		GetDb("").
		Where("name = ?", form.Name).
		First(&repeat)
	wrongs.ThrowWhenNotEmpty(ret, "线别名称")

	// 新建
	organizationLine := &models.OrganizationLineMdl{
		UniqueCode:                form.UniqueCode,
		Name:                      form.Name,
		OrganizationParagraphUuid: form.organizationParagraph.Uuid,
	}
	if ret = models.NewOrganizationLineMdl().
		GetDb("").
		Create(&organizationLine); ret.Error != nil {
		wrongs.ThrowForbidden(ret.Error.Error())
	}

	// 线别绑定站场
	organizationLine.BindOrganizationStations(form.organizationStations)
	// 线别绑定道口
	organizationLine.BindOrganizationCrossroads(form.organizationCrossroads)
	// 线别绑定中心
	organizationLine.BindOrganizationCenters(form.organizationCenters)

	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Created(map[string]any{"organization_line": organizationLine}).ToGinResponse())
}

// Destroy 删除
func (OrganizationLineCtrl) Destroy(ctx *gin.Context) {
	var (
		ret              *gorm.DB
		organizationLine *models.OrganizationLineMdl
	)

	// 查询
	ret = models.NewOrganizationLineMdl().
		GetDb("").
		Where("uuid = ?", ctx.Param("uuid")).
		First(&organizationLine)
	wrongs.ThrowWhenEmpty(ret, "线别")

	// 删除
	if ret := models.NewOrganizationLineMdl().
		GetDb("").
		Where("uuid = ?", ctx.Param("uuid")).
		Delete(&organizationLine); ret.Error != nil {
		wrongs.ThrowForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Deleted().ToGinResponse())
}

// DestroyMany 批量删除
func (OrganizationLineCtrl) DestroyMany(ctx *gin.Context) {
	form := OrganizationLineDestroyManyForm{}.ShouldBind(ctx)
	if ret := models.NewOrganizationLineMdl().GetDb("").Where("uuid in ?", form.Uuids).Delete(nil); ret.Error != nil {
		wrongs.ThrowForbidden("删除失败：%s", ret.Error.Error())
	}
	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Deleted().ToGinResponse())
}

// Update 编辑
func (OrganizationLineCtrl) Update(ctx *gin.Context) {
	var (
		ret                      *gorm.DB
		organizationLine, repeat *models.OrganizationLineMdl
	)

	// 表单
	form := OrganizationLineStoreForm{}.ShouldBind(ctx)

	// 查重
	ret = models.NewOrganizationLineMdl().
		GetDb("").
		Where("unique_code = ? and uuid <> ?", form.UniqueCode, ctx.Param("uuid")).
		First(&repeat)
	wrongs.ThrowWhenNotEmpty(ret, "线别代码")
	ret = models.NewOrganizationLineMdl().
		GetDb("").
		Where("name = ? and uuid <> ?", form.Name, ctx.Param("uuid")).
		First(&repeat)
	wrongs.ThrowWhenNotEmpty(ret, "线别名称")

	// 查询
	ret = models.NewOrganizationLineMdl().
		GetDb("").
		Where("uuid = ?", ctx.Param("uuid")).
		First(&organizationLine)
	wrongs.ThrowWhenEmpty(ret, "线别")

	// 编辑
	organizationLine.UniqueCode = form.UniqueCode
	organizationLine.Name = form.Name
	organizationLine.OrganizationParagraphUuid = form.organizationParagraph.Uuid
	if ret = models.NewOrganizationLineMdl().
		GetDb("").
		Where("uuid = ?", ctx.Param("uuid")).
		Save(&organizationLine); ret.Error != nil {
		wrongs.ThrowForbidden(ret.Error.Error())
	}

	// 线别绑定站场
	organizationLine.BindOrganizationStations(form.organizationStations)
	// 线别绑定道口
	organizationLine.BindOrganizationCrossroads(form.organizationCrossroads)
	// 线别绑定中心
	organizationLine.BindOrganizationCenters(form.organizationCenters)

	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Updated(map[string]any{"organization_line": organizationLine}).ToGinResponse())
}

// Detail 详情
func (OrganizationLineCtrl) Detail(ctx *gin.Context) {
	var (
		ret              *gorm.DB
		organizationLine *models.OrganizationLineMdl
	)
	ret = models.NewOrganizationLineMdl().
		SetCtx(ctx).
		GetDbUseQuery("").
		Where("uuid = ?", ctx.Param("uuid")).
		First(&organizationLine)
	wrongs.ThrowWhenEmpty(ret, "线别")

	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Datum(map[string]any{"organization_line": organizationLine}).ToGinResponse())
}

// List 列表
func (receiver OrganizationLineCtrl) List(ctx *gin.Context) {
	var organizationLines []*models.OrganizationLineMdl

	ctx.JSON(
		tools.NewCorrectWithGinContext("", ctx).
			DataForPager(
				models.OrganizationLineMdl{}.GetListByQuery(ctx),
				func(db *gorm.DB) map[string]any {
					db.Find(&organizationLines)
					return map[string]any{"organization_lines": organizationLines}
				},
			).
			ToGinResponse(),
	)
}

// ListJdt jquery-dataTable后端分页数据
func (receiver OrganizationLineCtrl) ListJdt(ctx *gin.Context) {
	var organizationLines []*models.OrganizationLineMdl

	ctx.JSON(
		tools.NewCorrectWithGinContext("", ctx).
			DataForJqueryDataTable(
				models.OrganizationLineMdl{}.GetListByQuery(ctx),
				func(db *gorm.DB) map[string]any {
					db.Find(&organizationLines)
					return map[string]any{"organization_lines": organizationLines}
				},
			).
			ToGinResponse(),
	)
}