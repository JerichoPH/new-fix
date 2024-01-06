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

	// InstallIndoorRoomCtrl 室内上道位置-机房控制器
	InstallIndoorRoomCtrl struct{}
	// InstallIndoorRoomStoreForm 室内上道位置-机房表单
	InstallIndoorRoomStoreForm struct {
		Name                      string `json:"name"`
		InstallIndoorRoomTypeUuid string `json:"install_indoor_room_type_uuid"`
		installIndoorRoomType     *models.InstallIndoorRoomTypeMdl
		OrganizationStationUuid   string `json:"organization_station_uuid"`
		OrganizationCenterUuid    string `json:"organization_center_uuid"`
		OrganizationCrossroadUuid string `json:"organization_crossroad_uuid"`
	}
	InstallIndoorRoomDestroyManyForm struct {
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

// ShouldBind 室内上道位置-机房表单验证
func (receiver InstallIndoorRoomStoreForm) ShouldBind(ctx *gin.Context) InstallIndoorRoomStoreForm {
	if err := ctx.ShouldBind(&receiver); err != nil {
		wrongs.ThrowValidate("表单绑定失败：%s", err.Error())
	}
	if receiver.Name == "" {
		wrongs.ThrowValidate("室内上道位置-机房名称必填")
	}
	if receiver.InstallIndoorRoomTypeUuid == "" {
		wrongs.ThrowValidate("室内上道位置-机房类型必填")
	} else {
		if ret := models.NewInstallIndoorRoomMdl().GetDb("").Where("uuid = ?", receiver.InstallIndoorRoomTypeUuid).First(&receiver.installIndoorRoomType); ret.Error != nil {
			wrongs.ThrowValidate("室内上道位置-机房类型不存在")
		}
	}

	parents := tools.RemoveEmptyStrings([]string{
		receiver.OrganizationStationUuid,
		receiver.OrganizationCenterUuid,
		receiver.OrganizationCrossroadUuid,
	})
	if len(parents) == 0 {
		wrongs.ThrowValidate("室内上道位置-机房所属单位必填")
	} else if len(parents) > 1 {
		wrongs.ThrowValidate("室内上道位置-机房所属单位只能选择一个")
	}

	if receiver.OrganizationStationUuid != "" {
		if ret := models.NewOrganizationStationMdl().GetDb("").Where("uuid = ?", receiver.OrganizationStationUuid).First(&models.OrganizationStationMdl{}); ret.Error != nil {
			wrongs.ThrowValidate("室内上道位置-机房所属站场不存在")
		}
	}
	if receiver.OrganizationCenterUuid != "" {
		if ret := models.NewOrganizationCenterMdl().GetDb("").Where("uuid = ?", receiver.OrganizationCenterUuid).First(&models.OrganizationCenterMdl{}); ret.Error != nil {
			wrongs.ThrowValidate("室内上道位置-机房所属中心不存在")
		}
	}
	if receiver.OrganizationCrossroadUuid != "" {
		if ret := models.NewOrganizationCrossroadMdl().GetDb("").Where("uuid = ?", receiver.OrganizationCrossroadUuid).First(&models.OrganizationCrossroadMdl{}); ret.Error != nil {
			wrongs.ThrowValidate("室内上道位置-机房所属路口不存在")
		}
	}
	return receiver
}

// ShouldBind 批量删除室内上道位置-机房表单验证
func (receiver InstallIndoorRoomDestroyManyForm) ShouldBind(ctx *gin.Context) InstallIndoorRoomDestroyManyForm {
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
	if ret := models.NewInstallIndoorRoomTypeMdl().GetDb("").Where("uuid in ?", form.Uuids).Delete(nil); ret.Error != nil {
		wrongs.ThrowForbidden("删除失败：%s", ret.Error.Error())
	}
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

// NewInstallIndoorRoomCtrl 构造函数
func NewInstallIndoorRoomCtrl() *InstallIndoorRoomCtrl {
	return &InstallIndoorRoomCtrl{}
}

// Store 新建
func (InstallIndoorRoomCtrl) Store(ctx *gin.Context) {
	var (
		ret    *gorm.DB
		repeat *models.InstallIndoorRoomMdl
	)

	// 表单
	form := InstallIndoorRoomStoreForm{}.ShouldBind(ctx)

	// 查重
	ret = models.NewInstallIndoorRoomMdl().
		GetDb("").
		Where("name = ?", form.Name).
		First(&repeat)
	wrongs.ThrowWhenNotEmpty(ret, "室内上道位置-机房名称")

	// 新建
	installIndoorRoom := &models.InstallIndoorRoomMdl{
		Name:                      form.Name,
		InstallIndoorRoomTypeUuid: form.installIndoorRoomType.Uuid,
		OrganizationStationUuid:   form.OrganizationStationUuid,
		OrganizationCenterUuid:    form.OrganizationCenterUuid,
		OrganizationCrossroadUuid: form.OrganizationCrossroadUuid,
	}
	if ret = models.NewInstallIndoorRoomMdl().
		GetDb("").
		Create(&installIndoorRoom); ret.Error != nil {
		wrongs.ThrowForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Created(map[string]any{"install_indoor_room": installIndoorRoom}).ToGinResponse())
}

// Destroy 删除
func (InstallIndoorRoomCtrl) Destroy(ctx *gin.Context) {
	var (
		ret               *gorm.DB
		installIndoorRoom *models.InstallIndoorRoomMdl
	)

	// 查询
	ret = models.NewInstallIndoorRoomMdl().
		GetDb("").
		Where("uuid = ?", ctx.Param("uuid")).
		First(&installIndoorRoom)
	wrongs.ThrowWhenEmpty(ret, "室内上道位置-机房")

	// 删除
	if ret := models.NewInstallIndoorRoomMdl().
		GetDb("").
		Where("uuid = ?", ctx.Param("uuid")).
		Delete(&installIndoorRoom); ret.Error != nil {
		wrongs.ThrowForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Deleted().ToGinResponse())
}

// DestroyMany 批量删除
func (InstallIndoorRoomCtrl) DestroyMany(ctx *gin.Context) {
	form := InstallIndoorRoomDestroyManyForm{}.ShouldBind(ctx)
	if ret := models.NewInstallIndoorRoomMdl().GetDb("").Where("uuid in ?", form.Uuids).Delete(nil); ret.Error != nil {
		wrongs.ThrowForbidden("删除失败：%s", ret.Error.Error())
	}
	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Deleted().ToGinResponse())
}

// Update 编辑
func (InstallIndoorRoomCtrl) Update(ctx *gin.Context) {
	var (
		ret                       *gorm.DB
		installIndoorRoom, repeat *models.InstallIndoorRoomMdl
	)

	// 表单
	form := InstallIndoorRoomStoreForm{}.ShouldBind(ctx)

	// 查重
	ret = models.NewInstallIndoorRoomMdl().
		GetDb("").
		Where("name = ? and uuid <> ?", form.Name, ctx.Param("uuid")).
		First(&repeat)
	wrongs.ThrowWhenNotEmpty(ret, "室内上道位置-机房名称")

	// 查询
	ret = models.NewInstallIndoorRoomMdl().
		GetDb("").
		Where("uuid = ?", ctx.Param("uuid")).
		First(&installIndoorRoom)
	wrongs.ThrowWhenEmpty(ret, "室内上道位置-机房")

	// 编辑
	installIndoorRoom.Name = form.Name
	installIndoorRoom.InstallIndoorRoomTypeUuid = form.installIndoorRoomType.Uuid
	installIndoorRoom.OrganizationStationUuid = form.OrganizationStationUuid
	installIndoorRoom.OrganizationCenterUuid = form.OrganizationCenterUuid
	installIndoorRoom.OrganizationCrossroadUuid = form.OrganizationCrossroadUuid
	if ret = models.NewInstallIndoorRoomMdl().
		GetDb("").
		Where("uuid = ?", ctx.Param("uuid")).
		Save(&installIndoorRoom); ret.Error != nil {
		wrongs.ThrowForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Updated(map[string]any{"install_indoor_room": installIndoorRoom}).ToGinResponse())
}

// Detail 详情
func (InstallIndoorRoomCtrl) Detail(ctx *gin.Context) {
	var (
		ret               *gorm.DB
		installIndoorRoom *models.InstallIndoorRoomMdl
	)
	ret = models.NewInstallIndoorRoomMdl().
		SetCtx(ctx).
		GetDbUseQuery("").
		Where("uuid = ?", ctx.Param("uuid")).
		First(&installIndoorRoom)
	wrongs.ThrowWhenEmpty(ret, "室内上道位置-机房")

	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Datum(map[string]any{"install_indoor_room": installIndoorRoom}).ToGinResponse())
}

// List 列表
func (receiver InstallIndoorRoomCtrl) List(ctx *gin.Context) {
	var installIndoorRooms []*models.InstallIndoorRoomMdl

	ctx.JSON(
		tools.NewCorrectWithGinContext("", ctx).
			DataForPager(
				models.InstallIndoorRoomMdl{}.GetListByQuery(ctx),
				func(db *gorm.DB) map[string]any {
					db.Find(&installIndoorRooms)
					return map[string]any{"install_indoor_rooms": installIndoorRooms}
				},
			).
			ToGinResponse(),
	)
}

// ListJdt jquery-dataTable后端分页数据
func (receiver InstallIndoorRoomCtrl) ListJdt(ctx *gin.Context) {
	var installIndoorRooms []*models.InstallIndoorRoomMdl

	ctx.JSON(
		tools.NewCorrectWithGinContext("", ctx).
			DataForJqueryDataTable(
				models.InstallIndoorRoomMdl{}.GetListByQuery(ctx),
				func(db *gorm.DB) map[string]any {
					db.Find(&installIndoorRooms)
					return map[string]any{"install_indoor_rooms": installIndoorRooms}
				},
			).
			ToGinResponse(),
	)
}
