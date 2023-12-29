package controllers

import (
	"new-fix/models"
	"new-fix/tools"
	"new-fix/wrongs"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type (
	// BreakdownTypeCtrl 故障类型控制器
	BreakdownTypeCtrl struct{}
	// BreakdownTypeStoreForm 故障类型表单
	BreakdownTypeStoreForm struct {
		Name                      string `json:"name"`
		EquipmentKindCategoryUuid string `json:"equipment_kind_category_uuid"`
		equipmentKindCategory     *models.EquipmentKindCategoryMdl
	}
	// BreakdownTypeDestroyManyForm 故障类型批量删除表单
	BreakdownTypeDestroyManyForm struct {
		Uuids []string `json:"uuids"`
	}

	// BreakdownLogCtrl 故障日志控制器
	BreakdownLogCtrl struct{}
	// BreakdownLogStoreForm 故障日志表单
	BreakdownLogStoreForm struct {
		EquipmentUuid     string `json:"equipment_uuid"`
		equipment         *models.EquipmentMdl
		BreakdownTypeUuid string `json:"breakdown_type_uuid"`
		breakdownType     *models.BreakdownTypeMdl
		TypeCode          string `json:"type_code"`
	}
)

// ShouldBind	表单绑定（故障类型）
func (receiver BreakdownTypeStoreForm) ShouldBind(ctx *gin.Context) BreakdownTypeStoreForm {
	if err := ctx.ShouldBind(&receiver); err != nil {
		wrongs.ThrowValidate("表单绑定失败：%s", err.Error())
	}
	if receiver.Name == "" {
		wrongs.ThrowValidate("故障类型名称必填")
	}
	if receiver.EquipmentKindCategoryUuid == "" {
		wrongs.ThrowValidate("所属器材种类编号必填")
	} else {
		ret := models.NewEquipmentKindCategoryMdl().GetDb("").Where("uuid = ?", receiver.EquipmentKindCategoryUuid).First(&receiver.equipmentKindCategory)
		wrongs.ThrowWhenEmpty(ret, "所属器材种类")
	}
	return receiver
}

// ShouldBind	表单绑定（故障类型批量删除）
func (receiver BreakdownTypeDestroyManyForm) ShouldBind(ctx *gin.Context) BreakdownTypeDestroyManyForm {
	if err := ctx.ShouldBind(&receiver); err != nil {
		wrongs.ThrowValidate("表单绑定失败：%s", err.Error())
	}
	if len(receiver.Uuids) == 0 {
		wrongs.ThrowValidate("故障类型编号不能为空")
	}
	return receiver
}

// ShouldBind	表单绑定（故障日志）
func (receiver BreakdownLogStoreForm) ShouldBind(ctx *gin.Context) BreakdownLogStoreForm {
	if err := ctx.ShouldBind(&receiver); err != nil {
		wrongs.ThrowValidate("表单绑定失败：%s", err.Error())
	}
	if receiver.EquipmentUuid == "" {
		wrongs.ThrowValidate("器材编号不能为空")
	} else {
		ret := models.NewEquipmentMdl().GetDb("").Where("uuid = ?", receiver.EquipmentUuid).First(&receiver.equipment)
		wrongs.ThrowWhenEmpty(ret, "器材")
	}
	if receiver.BreakdownTypeUuid == "" {
		wrongs.ThrowValidate("故障类型编号不能为空")
	} else {
		ret := models.NewBreakdownTypeMdl().GetDb("").Where("uuid = ?", receiver.BreakdownTypeUuid).First(&receiver.breakdownType)
		wrongs.ThrowWhenEmpty(ret, "故障类型")
	}
	if !tools.InString(receiver.TypeCode, models.BreakdownLogMdl{}.GetTypeCodes()) {
		wrongs.ThrowValidate("故障日志类型错误")
	}
	return receiver
}

// NewBreakdownTypeCtrl 构造函数
func NewBreakdownTypeCtrl() *BreakdownTypeCtrl {
	return &BreakdownTypeCtrl{}
}

// Store 新建
func (BreakdownTypeCtrl) Store(ctx *gin.Context) {
	var (
		ret    *gorm.DB
		repeat *models.BreakdownTypeMdl
	)

	// 表单
	form := BreakdownTypeStoreForm{}.ShouldBind(ctx)

	// 查重
	ret = models.NewBreakdownTypeMdl().
		GetDb("").
		Where("name = ?", form.Name).
		Where("equipment_kind_category_uuid = ?", form.equipmentKindCategory.Uuid).
		First(&repeat)
	wrongs.ThrowWhenNotEmpty(ret, "故障类型名称")

	// 新建
	breakdownType := &models.BreakdownTypeMdl{
		Name:                      form.Name,
		EquipmentKindCategoryUuid: form.equipmentKindCategory.Uuid,
	}
	if ret = models.NewBreakdownTypeMdl().
		GetDb("").
		Create(&breakdownType); ret.Error != nil {
		wrongs.ThrowForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Created(map[string]any{"breakdown_type": breakdownType}).ToGinResponse())
}

// Destroy 删除
func (BreakdownTypeCtrl) Destroy(ctx *gin.Context) {
	var (
		ret           *gorm.DB
		breakdownType *models.BreakdownTypeMdl
	)

	// 查询
	ret = models.NewBreakdownTypeMdl().
		GetDb("").
		Where("uuid = ?", ctx.Param("uuid")).
		First(&breakdownType)
	wrongs.ThrowWhenEmpty(ret, "故障类型")

	// 删除
	if ret := models.NewBreakdownTypeMdl().
		GetDb("").
		Where("uuid = ?", ctx.Param("uuid")).
		Delete(&breakdownType); ret.Error != nil {
		wrongs.ThrowForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Deleted().ToGinResponse())
}

// DestroyMany 批量删除
func (BreakdownTypeCtrl) DestroyMany(ctx *gin.Context) {
	form := BreakdownTypeDestroyManyForm{}.ShouldBind(ctx)
	if ret := models.NewBreakdownTypeMdl().GetDb("").Where("uuid in ?", form.Uuids).Delete(nil); ret.Error != nil {
		wrongs.ThrowForbidden("批量删除失败：%s", ret.Error.Error())
	}
	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Deleted().ToGinResponse())
}

// Update 编辑
func (BreakdownTypeCtrl) Update(ctx *gin.Context) {
	var (
		ret                   *gorm.DB
		breakdownType, repeat *models.BreakdownTypeMdl
	)

	// 表单
	form := BreakdownTypeStoreForm{}.ShouldBind(ctx)

	// 查重
	ret = models.NewBreakdownTypeMdl().
		GetDb("").
		Where("name = ?", form.Name).
		Where("uuid <> ?", ctx.Param("uuid")).
		Where("equipment_kind_category_uuid = ?", form.equipmentKindCategory.Uuid).
		First(&repeat)
	wrongs.ThrowWhenNotEmpty(ret, "故障类型名称")

	// 查询
	ret = models.NewBreakdownTypeMdl().
		GetDb("").
		Where("uuid = ?", ctx.Param("uuid")).
		First(&breakdownType)
	wrongs.ThrowWhenEmpty(ret, "故障类型")

	// 编辑
	breakdownType.Name = form.Name
	breakdownType.EquipmentKindCategoryUuid = form.equipmentKindCategory.Uuid
	if ret = models.NewBreakdownTypeMdl().
		GetDb("").
		Where("uuid = ?", ctx.Param("uuid")).
		Save(&breakdownType); ret.Error != nil {
		wrongs.ThrowForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Updated(map[string]any{"breakdown_type": breakdownType}).ToGinResponse())
}

// Detail 详情
func (BreakdownTypeCtrl) Detail(ctx *gin.Context) {
	var (
		ret           *gorm.DB
		breakdownType *models.BreakdownTypeMdl
	)
	ret = models.NewBreakdownTypeMdl().
		SetCtx(ctx).
		GetDbUseQuery("").
		Where("uuid = ?", ctx.Param("uuid")).
		First(&breakdownType)
	wrongs.ThrowWhenEmpty(ret, "故障类型")

	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Datum(map[string]any{"breakdown_type": breakdownType}).ToGinResponse())
}

// List 列表
func (receiver BreakdownTypeCtrl) List(ctx *gin.Context) {
	var breakdownTypes []*models.BreakdownTypeMdl

	ctx.JSON(
		tools.NewCorrectWithGinContext("", ctx).
			DataForPager(
				models.BreakdownTypeMdl{}.GetListByQuery(ctx),
				func(db *gorm.DB) map[string]any {
					db.Find(&breakdownTypes)
					return map[string]any{"breakdown_types": breakdownTypes}
				},
			).
			ToGinResponse(),
	)
}

// ListJdt jquery-dataTable后端分页数据
func (receiver BreakdownTypeCtrl) ListJdt(ctx *gin.Context) {
	var breakdownTypes []*models.BreakdownTypeMdl

	ctx.JSON(
		tools.NewCorrectWithGinContext("", ctx).
			DataForJqueryDataTable(
				models.BreakdownTypeMdl{}.GetListByQuery(ctx),
				func(db *gorm.DB) map[string]any {
					db.Find(&breakdownTypes)
					return map[string]any{"breakdown_types": breakdownTypes}
				},
			).
			ToGinResponse(),
	)
}

// NewBreakdownLogCtrl 构造函数
func NewBreakdownLogCtrl() *BreakdownLogCtrl {
	return &BreakdownLogCtrl{}
}

// Store 新建
func (BreakdownLogCtrl) Store(ctx *gin.Context) {
	var ret *gorm.DB

	// 表单
	form := BreakdownLogStoreForm{}.ShouldBind(ctx)

	// 新建
	breakdownLog := &models.BreakdownLogMdl{
		EquipmentUuid:     form.equipment.Uuid,
		BreakdownTypeUuid: form.breakdownType.Uuid,
		TypeCode:          form.TypeCode,
	}
	if ret = models.NewBreakdownLogMdl().
		GetDb("").
		Create(&breakdownLog); ret.Error != nil {
		wrongs.ThrowForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Created(map[string]any{"breakdown_log": breakdownLog}).ToGinResponse())
}

// Detail 详情
func (BreakdownLogCtrl) Detail(ctx *gin.Context) {
	var (
		ret          *gorm.DB
		breakdownLog *models.BreakdownLogMdl
	)
	ret = models.NewBreakdownLogMdl().
		SetCtx(ctx).
		GetDbUseQuery("").
		Where("uuid = ?", ctx.Param("uuid")).
		First(&breakdownLog)
	wrongs.ThrowWhenEmpty(ret, "故障日志")

	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Datum(map[string]any{"breakdown_log": breakdownLog}).ToGinResponse())
}

// List 列表
func (receiver BreakdownLogCtrl) List(ctx *gin.Context) {
	var breakdownLogs []*models.BreakdownLogMdl

	ctx.JSON(
		tools.NewCorrectWithGinContext("", ctx).
			DataForPager(
				models.BreakdownLogMdl{}.GetListByQuery(ctx),
				func(db *gorm.DB) map[string]any {
					db.Find(&breakdownLogs)
					return map[string]any{"breakdown_logs": breakdownLogs}
				},
			).
			ToGinResponse(),
	)
}

// ListJdt jquery-dataTable后端分页数据
func (receiver BreakdownLogCtrl) ListJdt(ctx *gin.Context) {
	var breakdownLogs []*models.BreakdownLogMdl

	ctx.JSON(
		tools.NewCorrectWithGinContext("", ctx).
			DataForJqueryDataTable(
				models.BreakdownLogMdl{}.GetListByQuery(ctx),
				func(db *gorm.DB) map[string]any {
					db.Find(&breakdownLogs)
					return map[string]any{"breakdown_logs": breakdownLogs}
				},
			).
			ToGinResponse(),
	)
}

// GetTypeCodesMap 获取故障日志类型映射
func (BreakdownLogCtrl) GetTypeCodesMap(ctx *gin.Context) {
	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Datum(map[string]any{"type_codes_map": models.BreakdownLogMdl{}.GetTypeCodesMap()}).ToGinResponse())
}
