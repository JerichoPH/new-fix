package controllers

import (
	"new-fix/models"
	"new-fix/utils"
	"new-fix/wrongs"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type (
	// SourceTypeCtrl 来源类型控制器
	SourceTypeCtrl struct{}
	// SourceTypeStoreForm 来源类型表单
	SourceTypeStoreForm struct {
		UniqueCode string `json:"unique_code"`
		Name       string `json:"name"`
	}
	// SourceTypeDestroyManyForm 批量删除来源类型表单
	SourceTypeDestroyManyForm struct {
		Uuids []string `json:"uuids"`
	}

	// SourceProjectCtrl 来源项目控制器
	SourceProjectCtrl struct{}
	// SourceProjectStoreForm 来源项目表单
	SourceProjectStoreForm struct {
		Name           string `json:"name"`
		SourceTypeUuid string `json:"source_type_uuid"`
		sourceType     *models.SourceTypeMdl
	}
	// SourceProjectDestroyManyForm 批量删除来源项目表单
	SourceProjectDestroyManyForm struct {
		Uuids []string `json:"uuids"`
	}
)

// ShouldBind 表单绑定（来源类型）
func (receiver SourceTypeStoreForm) ShouldBind(ctx *gin.Context) SourceTypeStoreForm {
	if err := ctx.ShouldBind(&receiver); err != nil {
		wrongs.ThrowValidate("表单绑定失败：%s", err.Error())
	}
	if receiver.UniqueCode == "" {
		wrongs.ThrowValidate("来源类型代码必填")
	} else {
		if len(receiver.UniqueCode) != 2 {
			wrongs.ThrowValidate("来源类型代码长度必须为2位")
		}
	}
	if receiver.Name == "" {
		wrongs.ThrowValidate("来源类型名称必填")
	}
	return receiver
}

// ShouldBind 表单绑定（批量删除来源类型）
func (receiver SourceTypeDestroyManyForm) ShouldBind(ctx *gin.Context) SourceTypeDestroyManyForm {
	if err := ctx.ShouldBind(&receiver); err != nil {
		wrongs.ThrowValidate("表单绑定失败：%s", err.Error())
	}
	if len(receiver.Uuids) == 0 {
		wrongs.ThrowValidate("来源类型uuids必填")
	}
	return receiver
}

// ShouldBind 表单绑定（来源项目）
func (receiver SourceProjectStoreForm) ShouldBind(ctx *gin.Context) SourceProjectStoreForm {
	if err := ctx.ShouldBind(&receiver); err != nil {
		wrongs.ThrowValidate("表单绑定失败：%s", err.Error())
	}
	if receiver.Name == "" {
		wrongs.ThrowValidate("来源项目名称必填")
	}
	if receiver.SourceTypeUuid == "" {
		wrongs.ThrowValidate("来源类型必选")
	} else {
		ret := models.NewSourceTypeMdl().GetDb("").Where("uuid = ?", receiver.SourceTypeUuid).First(&receiver.sourceType)
		wrongs.ThrowWhenEmpty(ret, "来源类型")
	}
	return receiver
}

// ShouldBind 表单绑定（批量删除来源项目）
func (receiver SourceProjectDestroyManyForm) ShouldBind(ctx *gin.Context) SourceProjectDestroyManyForm {
	if err := ctx.ShouldBind(&receiver); err != nil {
		wrongs.ThrowValidate("表单绑定失败：%s", err.Error())
	}
	if len(receiver.Uuids) == 0 {
		wrongs.ThrowValidate("来源项目uuids必填")
	}
	return receiver
}

// NewSourceTypeCtrl 构造函数
func NewSourceTypeCtrl() *SourceTypeCtrl {
	return &SourceTypeCtrl{}
}

// Store 新建
func (SourceTypeCtrl) Store(ctx *gin.Context) {
	var (
		ret    *gorm.DB
		repeat *models.SourceTypeMdl
	)

	// 表单
	form := SourceTypeStoreForm{}.ShouldBind(ctx)

	// 查重
	ret = models.NewSourceTypeMdl().
		GetDb("").
		Where("unique_code = ?", form.UniqueCode).
		First(&repeat)
	wrongs.ThrowWhenNotEmpty(ret, "来源类型代码")
	ret = models.NewSourceTypeMdl().
		GetDb("").
		Where("name = ?", form.Name).
		First(&repeat)
	wrongs.ThrowWhenNotEmpty(ret, "来源类型名称")

	// 新建
	sourceType := &models.SourceTypeMdl{
		UniqueCode: form.UniqueCode,
		Name:       form.Name,
	}
	if ret = models.NewSourceTypeMdl().
		GetDb("").
		Create(&sourceType); ret.Error != nil {
		wrongs.ThrowForbidden(ret.Error.Error())
	}

	ctx.JSON(utils.NewCorrectWithGinContext("", ctx).Created(map[string]any{"source_type": sourceType}).ToGinResponse())
}

// Destroy 删除
func (SourceTypeCtrl) Destroy(ctx *gin.Context) {
	var (
		ret        *gorm.DB
		sourceType *models.SourceTypeMdl
	)

	// 查询
	ret = models.NewSourceTypeMdl().
		GetDb("").
		Where("uuid = ?", ctx.Param("uuid")).
		First(&sourceType)
	wrongs.ThrowWhenEmpty(ret, "来源类型")

	// 删除
	if ret := models.NewSourceTypeMdl().
		GetDb("").
		Where("uuid = ?", ctx.Param("uuid")).
		Delete(&sourceType); ret.Error != nil {
		wrongs.ThrowForbidden(ret.Error.Error())
	}

	ctx.JSON(utils.NewCorrectWithGinContext("", ctx).Deleted().ToGinResponse())
}

// DestroyMany 批量删除
func (SourceTypeCtrl) DestroyMany(ctx *gin.Context) {
	form := SourceTypeDestroyManyForm{}.ShouldBind(ctx)
	if ret := models.NewSourceTypeMdl().GetDb("").Where("uuid in ?", form.Uuids).Delete(nil); ret.Error != nil {
		wrongs.ThrowForbidden("批量删除失败：%s", ret.Error.Error())
	}
	ctx.JSON(utils.NewCorrectWithGinContext("", ctx).Deleted().ToGinResponse())
}

// Update 编辑
func (SourceTypeCtrl) Update(ctx *gin.Context) {
	var (
		ret                *gorm.DB
		sourceType, repeat *models.SourceTypeMdl
	)

	// 表单
	form := SourceTypeStoreForm{}.ShouldBind(ctx)

	// 查重
	ret = models.NewSourceTypeMdl().
		GetDb("").
		Where("unique_code = ? and uuid <> ?", form.UniqueCode, ctx.Param("uuid")).
		First(&repeat)
	wrongs.ThrowWhenNotEmpty(ret, "来源类型代码")
	ret = models.NewSourceTypeMdl().
		GetDb("").
		Where("name = ? and uuid <> ?", form.Name, ctx.Param("uuid")).
		First(&repeat)
	wrongs.ThrowWhenNotEmpty(ret, "来源类型名称")

	// 查询
	ret = models.NewSourceTypeMdl().
		GetDb("").
		Where("uuid = ?", ctx.Param("uuid")).
		First(&sourceType)
	wrongs.ThrowWhenEmpty(ret, "来源类型")

	// 编辑
	sourceType.UniqueCode = form.UniqueCode
	sourceType.Name = form.Name
	if ret = models.NewSourceTypeMdl().
		GetDb("").
		Where("uuid = ?", ctx.Param("uuid")).
		Save(&sourceType); ret.Error != nil {
		wrongs.ThrowForbidden(ret.Error.Error())
	}

	ctx.JSON(utils.NewCorrectWithGinContext("", ctx).Updated(map[string]any{"source_type": sourceType}).ToGinResponse())
}

// Detail 详情
func (SourceTypeCtrl) Detail(ctx *gin.Context) {
	var (
		ret        *gorm.DB
		sourceType *models.SourceTypeMdl
	)
	ret = models.NewSourceTypeMdl().
		SetCtx(ctx).
		GetDbUseQuery("").
		Where("uuid = ?", ctx.Param("uuid")).
		First(&sourceType)
	wrongs.ThrowWhenEmpty(ret, "来源类型")

	ctx.JSON(utils.NewCorrectWithGinContext("", ctx).Datum(map[string]any{"source_type": sourceType}).ToGinResponse())
}

// List 列表
func (receiver SourceTypeCtrl) List(ctx *gin.Context) {
	var sourceTypes []*models.SourceTypeMdl

	ctx.JSON(
		utils.NewCorrectWithGinContext("", ctx).
			DataForPager(
				models.SourceTypeMdl{}.GetListByQuery(ctx),
				func(db *gorm.DB) map[string]any {
					db.Find(&sourceTypes)
					return map[string]any{"source_types": sourceTypes}
				},
			).
			ToGinResponse(),
	)
}

// ListJdt jquery-dataTable后端分页数据
func (receiver SourceTypeCtrl) ListJdt(ctx *gin.Context) {
	var sourceTypes []*models.SourceTypeMdl

	ctx.JSON(
		utils.NewCorrectWithGinContext("", ctx).
			DataForJqueryDataTable(
				models.SourceTypeMdl{}.GetListByQuery(ctx),
				func(db *gorm.DB) map[string]any {
					db.Find(&sourceTypes)
					return map[string]any{"source_types": sourceTypes}
				},
			).
			ToGinResponse(),
	)
}

// NewSourceProjectCtrl 构造函数
func NewSourceProjectCtrl() *SourceProjectCtrl {
	return &SourceProjectCtrl{}
}

// Store 新建
func (SourceProjectCtrl) Store(ctx *gin.Context) {
	var (
		ret    *gorm.DB
		repeat *models.SourceProjectMdl
	)

	// 表单
	form := SourceProjectStoreForm{}.ShouldBind(ctx)

	// 查重
	ret = models.NewSourceProjectMdl().
		GetDb("").
		Where("name = ?", form.Name).
		First(&repeat)
	wrongs.ThrowWhenNotEmpty(ret, "来源项目名称")

	// 新建
	sourceProject := &models.SourceProjectMdl{
		Name:           form.Name,
		SourceTypeUuid: form.sourceType.Uuid,
	}
	if ret = models.NewSourceProjectMdl().
		GetDb("").
		Create(&sourceProject); ret.Error != nil {
		wrongs.ThrowForbidden(ret.Error.Error())
	}

	ctx.JSON(utils.NewCorrectWithGinContext("", ctx).Created(map[string]any{"source_project": sourceProject}).ToGinResponse())
}

// Destroy 删除
func (SourceProjectCtrl) Destroy(ctx *gin.Context) {
	var (
		ret           *gorm.DB
		sourceProject *models.SourceProjectMdl
	)

	// 查询
	ret = models.NewSourceProjectMdl().
		GetDb("").
		Where("uuid = ?", ctx.Param("uuid")).
		First(&sourceProject)
	wrongs.ThrowWhenEmpty(ret, "来源项目")

	// 删除
	if ret := models.NewSourceProjectMdl().
		GetDb("").
		Where("uuid = ?", ctx.Param("uuid")).
		Delete(&sourceProject); ret.Error != nil {
		wrongs.ThrowForbidden(ret.Error.Error())
	}

	ctx.JSON(utils.NewCorrectWithGinContext("", ctx).Deleted().ToGinResponse())
}

// DestroyMany 批量删除
func (SourceProjectCtrl) DestroyMany(ctx *gin.Context) {
	form := SourceProjectDestroyManyForm{}.ShouldBind(ctx)
	if ret := models.NewSourceProjectMdl().GetDb("").Where("uuid in ?", form.Uuids).Delete(nil); ret.Error != nil {
		wrongs.ThrowForbidden("批量删除失败：%s", ret.Error.Error())
	}
	ctx.JSON(utils.NewCorrectWithGinContext("", ctx).Deleted().ToGinResponse())
}

// Update 编辑
func (SourceProjectCtrl) Update(ctx *gin.Context) {
	var (
		ret                   *gorm.DB
		sourceProject, repeat *models.SourceProjectMdl
	)

	// 表单
	form := SourceProjectStoreForm{}.ShouldBind(ctx)

	// 查重
	ret = models.NewSourceProjectMdl().
		GetDb("").
		Where("name = ? and uuid <> ?", form.Name, ctx.Param("uuid")).
		First(&repeat)
	wrongs.ThrowWhenNotEmpty(ret, "来源项目名称")

	// 查询
	ret = models.NewSourceProjectMdl().
		GetDb("").
		Where("uuid = ?", ctx.Param("uuid")).
		First(&sourceProject)
	wrongs.ThrowWhenEmpty(ret, "来源项目")

	// 编辑
	sourceProject.Name = form.Name
	sourceProject.SourceTypeUuid = form.sourceType.Uuid
	if ret = models.NewSourceProjectMdl().
		GetDb("").
		Where("uuid = ?", ctx.Param("uuid")).
		Save(&sourceProject); ret.Error != nil {
		wrongs.ThrowForbidden(ret.Error.Error())
	}

	ctx.JSON(utils.NewCorrectWithGinContext("", ctx).Updated(map[string]any{"source_project": sourceProject}).ToGinResponse())
}

// Detail 详情
func (SourceProjectCtrl) Detail(ctx *gin.Context) {
	var (
		ret           *gorm.DB
		sourceProject *models.SourceProjectMdl
	)
	ret = models.NewSourceProjectMdl().
		SetCtx(ctx).
		GetDbUseQuery("").
		Where("uuid = ?", ctx.Param("uuid")).
		First(&sourceProject)
	wrongs.ThrowWhenEmpty(ret, "来源项目")

	ctx.JSON(utils.NewCorrectWithGinContext("", ctx).Datum(map[string]any{"source_project": sourceProject}).ToGinResponse())
}

// List 列表
func (receiver SourceProjectCtrl) List(ctx *gin.Context) {
	var sourceProjects []*models.SourceProjectMdl

	ctx.JSON(
		utils.NewCorrectWithGinContext("", ctx).
			DataForPager(
				models.SourceProjectMdl{}.GetListByQuery(ctx),
				func(db *gorm.DB) map[string]any {
					db.Find(&sourceProjects)
					return map[string]any{"source_projects": sourceProjects}
				},
			).
			ToGinResponse(),
	)
}

// ListJdt jquery-dataTable后端分页数据
func (receiver SourceProjectCtrl) ListJdt(ctx *gin.Context) {
	var sourceProjects []*models.SourceProjectMdl

	ctx.JSON(
		utils.NewCorrectWithGinContext("", ctx).
			DataForJqueryDataTable(
				models.SourceProjectMdl{}.GetListByQuery(ctx),
				func(db *gorm.DB) map[string]any {
					db.Find(&sourceProjects)
					return map[string]any{"source_projects": sourceProjects}
				},
			).
			ToGinResponse(),
	)
}
