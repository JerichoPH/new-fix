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
	organizationParagraph := &models.OrganizationParagraphMdl{}
	if ret = models.NewOrganizationParagraphMdl().
		GetDb("").
		Create(&organizationParagraph); ret.Error != nil {
		wrongs.ThrowForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Created(map[string]any{"organizationParagraph": organizationParagraph}).ToGinResponse())
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
	if ret = models.NewOrganizationParagraphMdl().
		GetDb("").
		Where("uuid = ?", ctx.Param("uuid")).
		Save(&organizationParagraph); ret.Error != nil {
		wrongs.ThrowForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Updated(map[string]any{"organizationParagraph": organizationParagraph}).ToGinResponse())
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

	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Datum(map[string]any{"organizationParagraph": organizationParagraph}).ToGinResponse())
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
					return map[string]any{"organizationParagraphs": organizationParagraphs}
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
					return map[string]any{"organizationParagraphs": organizationParagraphs}
				},
			).
			ToGinResponse(),
	)
}
