package controllers

import (
	"new-fix/models"
	"new-fix/tools"
	"new-fix/wrongs"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type (
	// InstallIndoorRoomTypeCtrl 室内上道位置-机房类型控制器
	InstallIndoorRoomTypeCtrl struct{}
	// InstallIndoorRoomTypeStoreForm 室内上道位置-机房类型表单
	InstallIndoorRoomTypeStoreForm struct {
		UniqueCode string `json:"unique_code"`
		Name       string `json:"name"`
	}
	// InstallIndoorRoomTypeDestroyManyForm 批量删除室内上道位置-机房类型表单
	InstallIndoorRoomTypeDestroyManyForm struct {
		Uuids []string `json:"uuids"`
	}
)

// ShouldBind 室内上道位置-机房类型表单验证
func (receiver InstallIndoorRoomTypeStoreForm) ShouldBind(ctx *gin.Context) InstallIndoorRoomTypeStoreForm {
	if err := ctx.ShouldBind(&receiver); err != nil {
		wrongs.ThrowValidate("表单绑定失败：%s", err.Error())
	}
	if receiver.UniqueCode == "" {
		wrongs.ThrowValidate("室内上道位置-机房类型代码必填")
	}
	if receiver.Name == "" {
		wrongs.ThrowValidate("室内上道位置-机房类型名称必填")
	}
	return receiver
}

// ShouldBind 批量删除室内上道位置-机房类型表单验证
func (receiver InstallIndoorRoomTypeDestroyManyForm) ShouldBind(ctx *gin.Context) InstallIndoorRoomTypeDestroyManyForm {
	if err := ctx.ShouldBind(&receiver); err != nil {
		wrongs.ThrowValidate("表单绑定失败：%s", err.Error())
	}
	if len(receiver.Uuids) == 0 {
		wrongs.ThrowValidate("请选择需要删除的内容")
	}
	return receiver
}

// NewInstallIndoorRoomTypeCtrl 构造函数
func NewInstallIndoorRoomTypeCtrl() *InstallIndoorRoomTypeCtrl {
	return &InstallIndoorRoomTypeCtrl{}
}

// Store 新建
func (InstallIndoorRoomTypeCtrl) Store(ctx *gin.Context) {
	var (
		ret    *gorm.DB
		repeat *models.InstallIndoorRoomTypeMdl
	)

	// 表单
	form := InstallIndoorRoomTypeStoreForm{}.ShouldBind(ctx)

	// 查重
	ret = models.NewInstallIndoorRoomTypeMdl().
		GetDb("").
		Where("unique_code = ?", form.UniqueCode).
		First(&repeat)
	wrongs.ThrowWhenNotEmpty(ret, "室内上道位置-机房类型代码")
	ret = models.NewInstallIndoorRoomTypeMdl().
		GetDb("").
		Where("name = ?", form.Name).
		First(&repeat)
	wrongs.ThrowWhenNotEmpty(ret, "室内上道位置-机房类型名称")

	// 新建
	installIndoorRoomType := &models.InstallIndoorRoomTypeMdl{
		UniqueCode: form.UniqueCode,
		Name:       form.Name,
	}
	if ret = models.NewInstallIndoorRoomTypeMdl().
		GetDb("").
		Create(&installIndoorRoomType); ret.Error != nil {
		wrongs.ThrowForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Created(map[string]any{"install_indoor_room_type": installIndoorRoomType}).ToGinResponse())
}

// Destroy 删除
func (InstallIndoorRoomTypeCtrl) Destroy(ctx *gin.Context) {
	var (
		ret                   *gorm.DB
		installIndoorRoomType *models.InstallIndoorRoomTypeMdl
	)

	// 查询
	ret = models.NewInstallIndoorRoomTypeMdl().
		GetDb("").
		Where("uuid = ?", ctx.Param("uuid")).
		First(&installIndoorRoomType)
	wrongs.ThrowWhenEmpty(ret, "室内上道位置-机房类型")

	// 删除
	if ret := models.NewInstallIndoorRoomTypeMdl().
		GetDb("").
		Where("uuid = ?", ctx.Param("uuid")).
		Delete(&installIndoorRoomType); ret.Error != nil {
		wrongs.ThrowForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Deleted().ToGinResponse())
}

// DestroyMany 批量删除
func (InstallIndoorRoomTypeCtrl) DestroyMany(ctx *gin.Context) {
	form := InstallIndoorRoomTypeDestroyManyForm{}.ShouldBind(ctx)
	models.NewInstallIndoorRoomTypeMdl().GetDb("").Where("uuid in ?", form.Uuids).Delete(nil)
	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Deleted().ToGinResponse())
}

// Update 编辑
func (InstallIndoorRoomTypeCtrl) Update(ctx *gin.Context) {
	var (
		ret                           *gorm.DB
		installIndoorRoomType, repeat *models.InstallIndoorRoomTypeMdl
	)

	// 表单
	form := InstallIndoorRoomTypeStoreForm{}.ShouldBind(ctx)

	// 查重
	ret = models.NewInstallIndoorRoomTypeMdl().
		GetDb("").
		Where("unique_code = ? and uuid <> ?", form.UniqueCode, ctx.Param("uuid")).
		First(&repeat)
	wrongs.ThrowWhenNotEmpty(ret, "室内上道位置-机房类型代码")
	ret = models.NewInstallIndoorRoomTypeMdl().
		GetDb("").
		Where("name = ? and uuid <> ?", form.Name, ctx.Param("uuid")).
		First(&repeat)
	wrongs.ThrowWhenNotEmpty(ret, "室内上道位置-机房类型名称")

	// 查询
	ret = models.NewInstallIndoorRoomTypeMdl().
		GetDb("").
		Where("uuid = ?", ctx.Param("uuid")).
		First(&installIndoorRoomType)
	wrongs.ThrowWhenEmpty(ret, "室内上道位置-机房类型")

	// 编辑
	installIndoorRoomType.UniqueCode = form.UniqueCode
	installIndoorRoomType.Name = form.Name
	if ret = models.NewInstallIndoorRoomTypeMdl().
		GetDb("").
		Where("uuid = ?", ctx.Param("uuid")).
		Save(&installIndoorRoomType); ret.Error != nil {
		wrongs.ThrowForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Updated(map[string]any{"install_indoor_room_type": installIndoorRoomType}).ToGinResponse())
}

// Detail 详情
func (InstallIndoorRoomTypeCtrl) Detail(ctx *gin.Context) {
	var (
		ret                   *gorm.DB
		installIndoorRoomType *models.InstallIndoorRoomTypeMdl
	)
	ret = models.NewInstallIndoorRoomTypeMdl().
		SetCtx(ctx).
		GetDbUseQuery("").
		Where("uuid = ?", ctx.Param("uuid")).
		First(&installIndoorRoomType)
	wrongs.ThrowWhenEmpty(ret, "室内上道位置-机房类型")

	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Datum(map[string]any{"install_indoor_room_type": installIndoorRoomType}).ToGinResponse())
}

// List 列表
func (receiver InstallIndoorRoomTypeCtrl) List(ctx *gin.Context) {
	var installIndoorRoomTypes []*models.InstallIndoorRoomTypeMdl

	ctx.JSON(
		tools.NewCorrectWithGinContext("", ctx).
			DataForPager(
				models.InstallIndoorRoomTypeMdl{}.GetListByQuery(ctx),
				func(db *gorm.DB) map[string]any {
					db.Find(&installIndoorRoomTypes)
					return map[string]any{"install_indoor_room_types": installIndoorRoomTypes}
				},
			).
			ToGinResponse(),
	)
}

// ListJdt jquery-dataTable后端分页数据
func (receiver InstallIndoorRoomTypeCtrl) ListJdt(ctx *gin.Context) {
	var installIndoorRoomTypes []*models.InstallIndoorRoomTypeMdl

	ctx.JSON(
		tools.NewCorrectWithGinContext("", ctx).
			DataForJqueryDataTable(
				models.InstallIndoorRoomTypeMdl{}.GetListByQuery(ctx),
				func(db *gorm.DB) map[string]any {
					db.Find(&installIndoorRoomTypes)
					return map[string]any{"install_indoor_room_types": installIndoorRoomTypes}
				},
			).
			ToGinResponse(),
	)
}
