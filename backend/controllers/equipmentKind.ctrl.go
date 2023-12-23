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
	// 器材种类批量删除表单
	EquipmentKindCategoryDestroyManyForm struct {
		Uuids                   []string `json:"uuids"`
		equipmentKindCategories []*models.EquipmentKindCategoryMdl
	}

	// 器材类型控制器
	EquipmentKindTypeCtrl struct{}
	// 器材类型表单
	EquipmentKindTypeStoreForm struct {
		Name                      string `json:"name"`
		EquipmentKindCategoryUuid string `json:"equipment_kind_category_uuid"`
		equipmentKindCategory     *models.EquipmentKindCategoryMdl
	}
	// 器材类型批量删除表单
	EquipmentKindTypeDestroyManyForm struct {
		Uuids              []string `json:"uuids"`
		equipmentKindTypes []*models.EquipmentKindTypeMdl
	}

	// 器材型号控制器
	EquipmentKindModelCtrl struct{}
	// 器材型号表单
	EquipmentKindModelStoreForm struct {
		Name                  string `json:"name"`
		EquipmentKindTypeUuid string `json:"equipment_kind_type_uuid"`
		equipmentKindType     *models.EquipmentKindTypeMdl
	}
	// 器材型号批量删除表单
	EquipmentKindModelDestroyManyForm struct {
		Uuids               []string `json:"uuids"`
		equipmentKindModels []*models.EquipmentKindModelMdl
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

// ShouldBind 表单绑定（批量删除器材种类）
func (receiver EquipmentKindCategoryDestroyManyForm) ShouldBind(ctx *gin.Context) EquipmentKindCategoryDestroyManyForm {
	var err error
	if err = ctx.ShouldBind(&receiver); err != nil {
		wrongs.ThrowValidate("表单绑定失败：%s", err.Error())
	}

	if len(receiver.Uuids) == 0 {
		wrongs.ThrowValidate("器材种类编号不能为空")
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

// ShouldBind 表单绑定（批量删除器材类型）
func (receiver EquipmentKindTypeDestroyManyForm) ShouldBind(ctx *gin.Context) EquipmentKindTypeDestroyManyForm {
	var err error
	if err = ctx.ShouldBind(&receiver); err != nil {
		wrongs.ThrowValidate("表单绑定失败：%s", err.Error())
	}

	if len(receiver.Uuids) == 0 {
		wrongs.ThrowValidate("器材类型编号不能为空")
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

// ShouldBind 表单绑定（批量删除器材型号）
func (receiver EquipmentKindModelDestroyManyForm) ShouldBind(ctx *gin.Context) EquipmentKindModelDestroyManyForm {
	var err error
	if err = ctx.ShouldBind(&receiver); err != nil {
		wrongs.ThrowValidate("表单绑定失败：%s", err.Error())
	}

	if len(receiver.Uuids) == 0 {
		wrongs.ThrowValidate("器材型号编号不能为空")
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
	EquipmentKindCategory := &models.EquipmentKindCategoryMdl{
		UniqueCode: models.EquipmentKindCategoryMdl{}.GetNewUniqueCode(),
		Name:       form.Name,
	}
	if ret = models.NewEquipmentKindCategoryMdl().
		GetDb("").
		Create(&EquipmentKindCategory); ret.Error != nil {
		wrongs.ThrowForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Created(map[string]any{"equipment_kind_category": EquipmentKindCategory}).ToGinResponse())
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

// DestroyMany 批量删除
func (EquipmentKindCategoryCtrl) DestroyMany(ctx *gin.Context) {
	var ret *gorm.DB

	// 表单
	form := EquipmentKindCategoryDestroyManyForm{}.ShouldBind(ctx)

	// 删除
	if ret = models.NewEquipmentKindCategoryMdl().
		GetDb("").
		Where("uuid in ?", form.Uuids).
		Delete(nil); ret.Error != nil {
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

	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Updated(map[string]any{"equipment_kind_category": EquipmentKindCategory}).ToGinResponse())
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

	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Datum(map[string]any{"equipment_kind_category": EquipmentKindCategory}).ToGinResponse())
}

// List 列表
func (receiver EquipmentKindCategoryCtrl) List(ctx *gin.Context) {
	var equipmentKindCategories []*models.EquipmentKindCategoryMdl

	ctx.JSON(
		tools.NewCorrectWithGinContext("", ctx).
			DataForPager(
				models.EquipmentKindCategoryMdl{}.GetListByQuery(ctx),
				func(db *gorm.DB) map[string]any {
					db.Find(&equipmentKindCategories)
					return map[string]any{"equipment_kind_categories": equipmentKindCategories}
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
				func(db *gorm.DB) map[string]any {
					db.Find(&equipmentKindCategories)
					return map[string]any{"equipment_kind_categories": equipmentKindCategories}
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
	equipmentKindType := &models.EquipmentKindTypeMdl{
		UniqueCode:                models.EquipmentKindTypeMdl{}.GetNewUniqueCode(form.equipmentKindCategory),
		Name:                      form.Name,
		EquipmentKindCategoryUuid: form.equipmentKindCategory.Uuid,
	}
	if ret = models.NewEquipmentKindTypeMdl().
		GetDb("").
		Create(&equipmentKindType); ret.Error != nil {
		wrongs.ThrowForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Created(map[string]any{"equipment_kind_type": equipmentKindType}).ToGinResponse())
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

// DestroyMany 批量删除
func (EquipmentKindTypeCtrl) DestroyMany(ctx *gin.Context) {
	var ret *gorm.DB

	// 表单
	form := EquipmentKindTypeDestroyManyForm{}.ShouldBind(ctx)

	// 删除
	if ret = models.NewEquipmentKindTypeMdl().
		GetDb("").
		Where("uuid in ?", form.Uuids).
		Delete(nil); ret.Error != nil {
		wrongs.ThrowForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Deleted().ToGinResponse())
}

// Update 编辑
func (EquipmentKindTypeCtrl) Update(ctx *gin.Context) {
	var (
		ret                       *gorm.DB
		equipmentKindType, repeat *models.EquipmentKindTypeMdl
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
		First(&equipmentKindType)
	wrongs.ThrowWhenEmpty(ret, "器材类型")

	// 编辑
	equipmentKindType.Name = form.Name
	if ret = models.NewEquipmentKindTypeMdl().
		GetDb("").
		Where("uuid = ?", ctx.Param("uuid")).
		Save(&equipmentKindType); ret.Error != nil {
		wrongs.ThrowForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Updated(map[string]any{"equipment_kind_type": equipmentKindType}).ToGinResponse())
}

// Detail 详情
func (EquipmentKindTypeCtrl) Detail(ctx *gin.Context) {
	var (
		ret               *gorm.DB
		equipmentKindType *models.EquipmentKindTypeMdl
	)
	ret = models.NewEquipmentKindTypeMdl().
		SetCtx(ctx).
		GetDbUseQuery("").
		Where("uuid = ?", ctx.Param("uuid")).
		First(&equipmentKindType)
	wrongs.ThrowWhenEmpty(ret, "器材类型")

	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Datum(map[string]any{"equipment_kind_type": equipmentKindType}).ToGinResponse())
}

// List 列表
func (receiver EquipmentKindTypeCtrl) List(ctx *gin.Context) {
	var equipmentKindTypes []*models.EquipmentKindTypeMdl

	ctx.JSON(
		tools.NewCorrectWithGinContext("", ctx).
			DataForPager(
				models.EquipmentKindTypeMdl{}.GetListByQuery(ctx),
				func(db *gorm.DB) map[string]any {
					db.Find(&equipmentKindTypes)
					return map[string]any{"equipment_kind_types": equipmentKindTypes}
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
				func(db *gorm.DB) map[string]any {
					db.Find(&equipmentKindTypes)
					return map[string]any{"equipment_kind_types": equipmentKindTypes}
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
	equipmentKindModel := &models.EquipmentKindModelMdl{
		UniqueCode:            models.EquipmentKindModelMdl{}.GetNewUniqueCode(form.equipmentKindType),
		Name:                  form.Name,
		EquipmentKindTypeUuid: form.equipmentKindType.Uuid,
	}
	if ret = models.NewEquipmentKindModelMdl().
		GetDb("").
		Create(&equipmentKindModel); ret.Error != nil {
		wrongs.ThrowForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Created(map[string]any{"equipment_kind_model": equipmentKindModel}).ToGinResponse())
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

// DestroyMany 批量删除
func (EquipmentKindModelCtrl) DestroyMany(ctx *gin.Context) {
	var ret *gorm.DB

	// 表单
	form := EquipmentKindModelDestroyManyForm{}.ShouldBind(ctx)

	// 删除
	if ret = models.NewEquipmentKindModelMdl().
		GetDb("").
		Where("uuid in ?", form.Uuids).
		Delete(nil); ret.Error != nil {
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

	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Updated(map[string]any{"equipment_kind_model": equipmentKindModel}).ToGinResponse())
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

	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Datum(map[string]any{"equipment_kind_model": equipmentKindModel}).ToGinResponse())
}

// List 列表
func (receiver EquipmentKindModelCtrl) List(ctx *gin.Context) {
	var equipmentKindModels []*models.EquipmentKindModelMdl

	ctx.JSON(
		tools.NewCorrectWithGinContext("", ctx).
			DataForPager(
				models.EquipmentKindModelMdl{}.GetListByQuery(ctx),
				func(db *gorm.DB) map[string]any {
					db.Find(&equipmentKindModels)
					return map[string]any{"equipment_kind_models": equipmentKindModels}
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
				func(db *gorm.DB) map[string]any {
					db.Find(&equipmentKindModels)
					return map[string]any{"equipment_kind_models": equipmentKindModels}
				},
			).
			ToGinResponse(),
	)
}
