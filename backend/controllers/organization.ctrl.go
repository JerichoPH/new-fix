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
	}
	// OrganizationStationDestroyManyForm 批量删除站场表单
	OrganizationStationDestroyManyForm struct {
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
		models.NewOrganizationRailwayMdl().GetDb("").Where("uuid = ?", receiver.OrganizationRailwayUuid).First(&receiver.organizationRailway)
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
		models.NewOrganizationParagraphMdl().GetDb("").Where("uuid = ?", receiver.OrganizationParagraphUuid).First(&receiver.organizationParagraph)
	}
	if !tools.InString(receiver.TypeCode, models.OrganizationWorkshopMdl{}.GetTypeCodes()) {
		wrongs.ThrowValidate("车间类型代码错误")
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
		if len(receiver.UniqueCode) != 5 {
			wrongs.ThrowValidate("站场代码必须是8位")
		}
	}
	if receiver.Name == "" {
		wrongs.ThrowValidate("站场名称必填")
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
