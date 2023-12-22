package controllers

import (
	"new-fix/models"
	"new-fix/tools"
	"new-fix/wrongs"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type (
	// 器材种类控制器
	EquipmentKindCategoryCtrl struct{}
	// 器材种类表单
	EquipmentKindCategoryStoreForm struct {
		Name string `json:"name"`
	}

	// 器材类型控制器
	EquipmentKindTypeCtrl struct{}
	// 器材类型表单
	EquipmentKindTypeStoreForm struct {
		Name                      string `json:"name"`
		EquipmentKindCategoryUuid string `json:"equipment_kind_category_uuid"`
		equipmentKindCategory     *models.EquipmentKindCategoryMdl
	}

	// 器材型号控制器
	EquipmentKindModelCtrl struct{}
	// 器材型号表单
	EquipmentKindModelStoreForm struct {
		Name                  string `json:"name"`
		EquipmentKindTypeUuid string `json:"equipment_kind_type_uuid"`
		equipmentKindType     *models.EquipmentKindTypeMdl
	}
)

// ShouldBind 表单绑定（器材种类）
func (receiver EquipmentKindCategoryStoreForm) ShouldBind(ctx *gin.Context) EquipmentKindCategoryStoreForm {
	var err error
	if err = ctx.ShouldBind(&receiver); err != nil {
		wrongs.ThrowValidate("表单绑定失败：%s", err.Error())
	}

	if receiver.Name == "" {
		wrongs.ThrowValidate("器材种类名称不能为空")
	}

	return receiver
}

// ShouldBind 表单绑定（器材类型）
func (receiver EquipmentKindTypeStoreForm) ShouldBind(ctx *gin.Context) EquipmentKindTypeStoreForm {
	var (
		err error
		ret *gorm.DB
	)
	if err = ctx.ShouldBind(&receiver); err != nil {
		wrongs.ThrowValidate("表单绑定失败：%s", err.Error())
	}

	if receiver.Name == "" {
		wrongs.ThrowValidate("器材类型名称不能为空")
	}

	if receiver.EquipmentKindCategoryUuid == "" {
		wrongs.ThrowValidate("所属器材种类不能为空")
	} else {
		ret = models.NewEquipmentKindCategoryMdl().GetDb("").Where("uuid = ?", receiver.EquipmentKindCategoryUuid).First(&receiver.equipmentKindCategory)
		wrongs.ThrowWhenEmpty(ret, "所属器材种类")
	}

	return receiver
}

// 表单绑定（器材类型）
func (receiver EquipmentKindModelStoreForm) ShouldBind(ctx *gin.Context) EquipmentKindModelStoreForm {
	var ret *gorm.DB
	if err := ctx.ShouldBind(&receiver); err != nil {
		wrongs.ThrowValidate("表单绑定失败：%s", err.Error())
	}
	if receiver.Name == "" {
		wrongs.ThrowValidate("器材型号名称必填")
	}
	if receiver.EquipmentKindTypeUuid == "" {
		wrongs.ThrowValidate("所属器材类型不能为空")
	} else {
		ret = models.NewEquipmentKindTypeMdl().GetDb("").Where("uuid = ?", receiver.EquipmentKindTypeUuid).First(&receiver.equipmentKindType)
		wrongs.ThrowWhenEmpty(ret, "所属器材类型")
	}
	return receiver
}

// NewEquipmentKindCategoryCtrl 构造函数
func NewEquipmentKindCategoryCtrl() *EquipmentKindCategoryCtrl {
	return &EquipmentKindCategoryCtrl{}
}

// Store 新建
func (EquipmentKindCategoryCtrl) Store(ctx *gin.Context) {
	var (
		ret    *gorm.DB
		repeat *models.EquipmentKindCategoryMdl
	)

	// 表单
	form := EquipmentKindCategoryStoreForm{}.ShouldBind(ctx)

	// 查重
	ret = models.NewEquipmentKindCategoryMdl().
		GetDb("").
		Where("name = ?", form.Name).
		First(&repeat)
	wrongs.ThrowWhenNotEmpty(ret, "种类名称")

	// 新建
	EquipmentKindCategory := &models.EquipmentKindCategoryMdl{}
	if ret = models.NewEquipmentKindCategoryMdl().
		GetDb("").
		Create(&EquipmentKindCategory); ret.Error != nil {
		wrongs.ThrowForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Created(map[string]any{"EquipmentKindCategory": EquipmentKindCategory}).ToGinResponse())
}

// Destroy 删除
func (EquipmentKindCategoryCtrl) Destroy(ctx *gin.Context) {
	var (
		ret                   *gorm.DB
		EquipmentKindCategory *models.EquipmentKindCategoryMdl
	)

	// 查询
	ret = models.NewEquipmentKindCategoryMdl().
		GetDb("").
		Where("uuid = ?", ctx.Param("uuid")).
		First(&EquipmentKindCategory)
	wrongs.ThrowWhenEmpty(ret, "种类")

	// 删除
	if ret := models.NewEquipmentKindCategoryMdl().
		GetDb("").
		Where("uuid = ?", ctx.Param("uuid")).
		Delete(&EquipmentKindCategory); ret.Error != nil {
		wrongs.ThrowForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Deleted().ToGinResponse())
}

// Update 编辑
func (EquipmentKindCategoryCtrl) Update(ctx *gin.Context) {
	var (
		ret                           *gorm.DB
		EquipmentKindCategory, repeat *models.EquipmentKindCategoryMdl
	)

	// 表单
	form := EquipmentKindCategoryStoreForm{}.ShouldBind(ctx)

	// 查重
	ret = models.NewEquipmentKindCategoryMdl().
		GetDb("").
		Where("name = ? and uuid <> ?", form.Name, ctx.Param("uuid")).
		First(&repeat)
	wrongs.ThrowWhenNotEmpty(ret, "种类名称")

	// 查询
	ret = models.NewEquipmentKindCategoryMdl().
		GetDb("").
		Where("uuid = ?", ctx.Param("uuid")).
		First(&EquipmentKindCategory)
	wrongs.ThrowWhenEmpty(ret, "种类")

	// 编辑
	EquipmentKindCategory.Name = form.Name
	if ret = models.NewEquipmentKindCategoryMdl().
		GetDb("").
		Where("uuid = ?", ctx.Param("uuid")).
		Save(&EquipmentKindCategory); ret.Error != nil {
		wrongs.ThrowForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Updated(map[string]any{"EquipmentKindCategory": EquipmentKindCategory}).ToGinResponse())
}

// Detail 详情
func (EquipmentKindCategoryCtrl) Detail(ctx *gin.Context) {
	var (
		ret                   *gorm.DB
		EquipmentKindCategory *models.EquipmentKindCategoryMdl
	)
	ret = models.NewEquipmentKindCategoryMdl().
		SetCtx(ctx).
		GetDbUseQuery("").
		Where("uuid = ?", ctx.Param("uuid")).
		First(&EquipmentKindCategory)
	wrongs.ThrowWhenEmpty(ret, "种类")

	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Datum(map[string]any{"EquipmentKindCategory": EquipmentKindCategory}).ToGinResponse())
}

// List 列表
func (receiver EquipmentKindCategoryCtrl) List(ctx *gin.Context) {
	var equipmentKindCategories []*models.EquipmentKindCategoryMdl

	ctx.JSON(
		tools.NewCorrectWithGinContext("", ctx).
			DataForPager(
				models.EquipmentKindCategoryMdl{}.GetListByQuery(ctx),
				func(db *gorm.DB) map[string]interface{} {
					db.Find(&equipmentKindCategories)
					return map[string]any{"equipmentKindCategories": equipmentKindCategories}
				},
			).
			ToGinResponse(),
	)
}

// ListJdt jquery-dataTable后端分页数据
func (receiver EquipmentKindCategoryCtrl) ListJdt(ctx *gin.Context) {
	var equipmentKindCategories []*models.EquipmentKindCategoryMdl

	ctx.JSON(
		tools.NewCorrectWithGinContext("", ctx).
			DataForJqueryDataTable(
				models.EquipmentKindCategoryMdl{}.GetListByQuery(ctx),
				func(db *gorm.DB) map[string]interface{} {
					db.Find(&equipmentKindCategories)
					return map[string]any{"equipmentKindCategories": equipmentKindCategories}
				},
			).
			ToGinResponse(),
	)
}

// NewEquipmentKindTypeCtrl 构造函数
func NewEquipmentKindTypeCtrl() *EquipmentKindTypeCtrl {
	return &EquipmentKindTypeCtrl{}
}

// Store 新建
func (EquipmentKindTypeCtrl) Store(ctx *gin.Context) {
	var (
		ret    *gorm.DB
		repeat *models.EquipmentKindTypeMdl
	)

	// 表单
	form := EquipmentKindTypeStoreForm{}.ShouldBind(ctx)

	// 查重
	ret = models.NewEquipmentKindTypeMdl().
		GetDb("").
		Where("name = ?", form.Name).
		Where("equipment_kind_category_uuid = ?", form.equipmentKindCategory.Uuid).
		First(&repeat)
	wrongs.ThrowWhenNotEmpty(ret, "器材类型名称")

	// 新建
	equipmentKind := &models.EquipmentKindTypeMdl{}
	if ret = models.NewEquipmentKindTypeMdl().
		GetDb("").
		Create(&equipmentKind); ret.Error != nil {
		wrongs.ThrowForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Created(map[string]any{"equipmentKind": equipmentKind}).ToGinResponse())
}

// Destroy 删除
func (EquipmentKindTypeCtrl) Destroy(ctx *gin.Context) {
	var (
		ret           *gorm.DB
		equipmentKind *models.EquipmentKindTypeMdl
	)

	// 查询
	ret = models.NewEquipmentKindTypeMdl().
		GetDb("").
		Where("uuid = ?", ctx.Param("uuid")).
		First(&equipmentKind)
	wrongs.ThrowWhenEmpty(ret, "器材类型")

	// 删除
	if ret := models.NewEquipmentKindTypeMdl().
		GetDb("").
		Where("uuid = ?", ctx.Param("uuid")).
		Delete(&equipmentKind); ret.Error != nil {
		wrongs.ThrowForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Deleted().ToGinResponse())
}

// Update 编辑
func (EquipmentKindTypeCtrl) Update(ctx *gin.Context) {
	var (
		ret                   *gorm.DB
		equipmentKind, repeat *models.EquipmentKindTypeMdl
	)

	// 表单
	form := EquipmentKindTypeStoreForm{}.ShouldBind(ctx)

	// 查重
	ret = models.NewEquipmentKindTypeMdl().
		GetDb("").
		Where("name = ? and uuid <> ?", form.Name, ctx.Param("uuid")).
		Where("equipment_kind_category_uuid = ?", form.equipmentKindCategory.Uuid).
		First(&repeat)
	wrongs.ThrowWhenNotEmpty(ret, "器材类型名称")

	// 查询
	ret = models.NewEquipmentKindTypeMdl().
		GetDb("").
		Where("uuid = ?", ctx.Param("uuid")).
		First(&equipmentKind)
	wrongs.ThrowWhenEmpty(ret, "器材类型")

	// 编辑
	equipmentKind.Name = form.Name
	equipmentKind.EquipmentKindCategoryUuid = form.equipmentKindCategory.Uuid
	if ret = models.NewEquipmentKindTypeMdl().
		GetDb("").
		Where("uuid = ?", ctx.Param("uuid")).
		Save(&equipmentKind); ret.Error != nil {
		wrongs.ThrowForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Updated(map[string]any{"equipmentKind": equipmentKind}).ToGinResponse())
}

// Detail 详情
func (EquipmentKindTypeCtrl) Detail(ctx *gin.Context) {
	var (
		ret           *gorm.DB
		equipmentKind *models.EquipmentKindTypeMdl
	)
	ret = models.NewEquipmentKindTypeMdl().
		SetCtx(ctx).
		GetDbUseQuery("").
		Where("uuid = ?", ctx.Param("uuid")).
		First(&equipmentKind)
	wrongs.ThrowWhenEmpty(ret, "器材类型")

	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Datum(map[string]any{"equipmentKind": equipmentKind}).ToGinResponse())
}

// List 列表
func (receiver EquipmentKindTypeCtrl) List(ctx *gin.Context) {
	var equipmentKindTypes []*models.EquipmentKindTypeMdl

	ctx.JSON(
		tools.NewCorrectWithGinContext("", ctx).
			DataForPager(
				models.EquipmentKindTypeMdl{}.GetListByQuery(ctx),
				func(db *gorm.DB) map[string]interface{} {
					db.Find(&equipmentKindTypes)
					return map[string]any{"equipmentKindTypes": equipmentKindTypes}
				},
			).
			ToGinResponse(),
	)
}

// ListJdt jquery-dataTable后端分页数据
func (receiver EquipmentKindTypeCtrl) ListJdt(ctx *gin.Context) {
	var equipmentKindTypes []*models.EquipmentKindTypeMdl

	ctx.JSON(
		tools.NewCorrectWithGinContext("", ctx).
			DataForJqueryDataTable(
				models.EquipmentKindTypeMdl{}.GetListByQuery(ctx),
				func(db *gorm.DB) map[string]interface{} {
					db.Find(&equipmentKindTypes)
					return map[string]any{"equipmentKindTypes": equipmentKindTypes}
				},
			).
			ToGinResponse(),
	)
}

// NewEquipmentKindModelCtrl 构造函数
func NewEquipmentKindModelCtrl() *EquipmentKindModelCtrl {
	return &EquipmentKindModelCtrl{}
}

// Store 新建
func (EquipmentKindModelCtrl) Store(ctx *gin.Context) {
	var (
		ret    *gorm.DB
		repeat *models.EquipmentKindModelMdl
	)

	// 表单
	form := EquipmentKindModelStoreForm{}.ShouldBind(ctx)

	// 查重
	ret = models.NewEquipmentKindModelMdl().
		GetDb("").
		Where("name = ?", form.Name).
		Where("equipment_kind_type_uuid = ?", form.equipmentKindType.Uuid).
		First(&repeat)
	wrongs.ThrowWhenNotEmpty(ret, "器材型号名称")

	// 新建
	equipmentKindModel := &models.EquipmentKindModelMdl{}
	if ret = models.NewEquipmentKindModelMdl().
		GetDb("").
		Create(&equipmentKindModel); ret.Error != nil {
		wrongs.ThrowForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Created(map[string]any{"equipmentKindModel": equipmentKindModel}).ToGinResponse())
}

// Destroy 删除
func (EquipmentKindModelCtrl) Destroy(ctx *gin.Context) {
	var (
		ret                *gorm.DB
		equipmentKindModel *models.EquipmentKindModelMdl
	)

	// 查询
	ret = models.NewEquipmentKindModelMdl().
		GetDb("").
		Where("uuid = ?", ctx.Param("uuid")).
		First(&equipmentKindModel)
	wrongs.ThrowWhenEmpty(ret, "器材型号")

	// 删除
	if ret := models.NewEquipmentKindModelMdl().
		GetDb("").
		Where("uuid = ?", ctx.Param("uuid")).
		Delete(&equipmentKindModel); ret.Error != nil {
		wrongs.ThrowForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Deleted().ToGinResponse())
}

// Update 编辑
func (EquipmentKindModelCtrl) Update(ctx *gin.Context) {
	var (
		ret                        *gorm.DB
		equipmentKindModel, repeat *models.EquipmentKindModelMdl
	)

	// 表单
	form := EquipmentKindModelStoreForm{}.ShouldBind(ctx)

	// 查重
	ret = models.NewEquipmentKindModelMdl().
		GetDb("").
		Where("name = ? and uuid <> ?", form.Name, ctx.Param("uuid")).
		Where("equipment_kind_type_uuid = ?", form.equipmentKindType.Uuid).
		First(&repeat)
	wrongs.ThrowWhenNotEmpty(ret, "器材型号名称")

	// 查询
	ret = models.NewEquipmentKindModelMdl().
		GetDb("").
		Where("uuid = ?", ctx.Param("uuid")).
		First(&equipmentKindModel)
	wrongs.ThrowWhenEmpty(ret, "器材型号")

	// 编辑
	equipmentKindModel.Name = form.Name
	if ret = models.NewEquipmentKindModelMdl().
		GetDb("").
		Where("uuid = ?", ctx.Param("uuid")).
		Save(&equipmentKindModel); ret.Error != nil {
		wrongs.ThrowForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Updated(map[string]any{"equipmentKindModel": equipmentKindModel}).ToGinResponse())
}

// Detail 详情
func (EquipmentKindModelCtrl) Detail(ctx *gin.Context) {
	var (
		ret                *gorm.DB
		equipmentKindModel *models.EquipmentKindModelMdl
	)
	ret = models.NewEquipmentKindModelMdl().
		SetCtx(ctx).
		GetDbUseQuery("").
		Where("uuid = ?", ctx.Param("uuid")).
		First(&equipmentKindModel)
	wrongs.ThrowWhenEmpty(ret, "器材型号")

	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Datum(map[string]any{"equipmentKindModel": equipmentKindModel}).ToGinResponse())
}

// List 列表
func (receiver EquipmentKindModelCtrl) List(ctx *gin.Context) {
	var equipmentKindModels []*models.EquipmentKindModelMdl

	ctx.JSON(
		tools.NewCorrectWithGinContext("", ctx).
			DataForPager(
				models.EquipmentKindModelMdl{}.GetListByQuery(ctx),
				func(db *gorm.DB) map[string]interface{} {
					db.Find(&equipmentKindModels)
					return map[string]any{"equipmentKindModels": equipmentKindModels}
				},
			).
			ToGinResponse(),
	)
}

// ListJdt jquery-dataTable后端分页数据
func (receiver EquipmentKindModelCtrl) ListJdt(ctx *gin.Context) {
	var equipmentKindModels []*models.EquipmentKindModelMdl

	ctx.JSON(
		tools.NewCorrectWithGinContext("", ctx).
			DataForJqueryDataTable(
				models.EquipmentKindModelMdl{}.GetListByQuery(ctx),
				func(db *gorm.DB) map[string]interface{} {
					db.Find(&equipmentKindModels)
					return map[string]any{"equipmentKindModels": equipmentKindModels}
				},
			).
			ToGinResponse(),
	)
}
