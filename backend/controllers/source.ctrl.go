package controllers

import (
	"new-fix/models"
	"new-fix/tools"
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

	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Created(map[string]any{"source_type": sourceType}).ToGinResponse())
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

	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Deleted().ToGinResponse())
}

// DestroyMany 批量删除
func (SourceTypeCtrl) DestroyMany(ctx *gin.Context) {
	form := SourceTypeDestroyManyForm{}.ShouldBind(ctx)
	if ret := models.NewSourceTypeMdl().GetDb("").Where("uuid in ?", form.Uuids).Delete(nil); ret.Error != nil {
		wrongs.ThrowForbidden("批量删除失败：%s", ret.Error.Error())
	}
	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Deleted().ToGinResponse())
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

	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Updated(map[string]any{"source_type": sourceType}).ToGinResponse())
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

	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Datum(map[string]any{"source_type": sourceType}).ToGinResponse())
}

// List 列表
func (receiver SourceTypeCtrl) List(ctx *gin.Context) {
	var sourceTypes []*models.SourceTypeMdl

	ctx.JSON(
		tools.NewCorrectWithGinContext("", ctx).
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
		tools.NewCorrectWithGinContext("", ctx).
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
