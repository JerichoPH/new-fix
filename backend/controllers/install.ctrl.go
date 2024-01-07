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

	// InstallIndoorPlatoonCtrl 室内上道位置-排控制器
	InstallIndoorPlatoonCtrl struct{}
	// InstallIndoorPlatoonStoreForm 室内上道位置-排表单
	InstallIndoorPlatoonStoreForm struct {
		Name                  string `json:"name"`
		InstallIndoorRoomUuid string `json:"install_indoor_room_uuid"`
		installIndoorRoom     *models.InstallIndoorRoomMdl
	}
	// InstallIndoorPlatoonDestroyManyForm 批量删除室内上道位置-排表单
	InstallIndoorPlatoonDestroyManyForm struct {
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
		ret := models.NewInstallIndoorRoomTypeMdl().GetDb("").Where("uuid = ?", receiver.InstallIndoorRoomTypeUuid).First(&receiver.installIndoorRoomType)
		wrongs.ThrowWhenEmpty(ret, "室内上道位置-机房类型")
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
		ret := models.NewOrganizationStationMdl().GetDb("").Where("uuid = ?", receiver.OrganizationStationUuid).First(&models.OrganizationStationMdl{})
		wrongs.ThrowWhenEmpty(ret, "室内上道位置-机房所属站场")

	}
	if receiver.OrganizationCenterUuid != "" {
		ret := models.NewOrganizationCenterMdl().GetDb("").Where("uuid = ?", receiver.OrganizationCenterUuid).First(&models.OrganizationCenterMdl{})
		wrongs.ThrowWhenEmpty(ret, "室内上道位置-机房所属中心")

	}
	if receiver.OrganizationCrossroadUuid != "" {
		ret := models.NewOrganizationCrossroadMdl().GetDb("").Where("uuid = ?", receiver.OrganizationCrossroadUuid).First(&models.OrganizationCrossroadMdl{})
		wrongs.ThrowWhenEmpty(ret, "室内上道位置-机房所属路口")

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

// ShouldBind 室内上道位置-排表单验证
func (receiver InstallIndoorPlatoonStoreForm) ShouldBind(ctx *gin.Context) InstallIndoorPlatoonStoreForm {
	if err := ctx.ShouldBind(&receiver); err != nil {
		wrongs.ThrowValidate("表单绑定失败：%s", err.Error())
	}
	if receiver.Name == "" {
		wrongs.ThrowValidate("室内上道位置-排名称必填")
	}
	if receiver.InstallIndoorRoomUuid == "" {
		wrongs.ThrowValidate("室内上道位置-排所属机房必填")
	} else {
		ret := models.NewInstallIndoorRoomMdl().GetDb("").Where("uuid = ?", receiver.InstallIndoorRoomUuid).First(&receiver.installIndoorRoom)
		wrongs.ThrowWhenEmpty(ret, "室内上道位置-排所属机房")
	}
	return receiver
}

// ShouldBind 批量删除室内上道位置-排表单验证
func (receiver InstallIndoorPlatoonDestroyManyForm) ShouldBind(ctx *gin.Context) InstallIndoorPlatoonDestroyManyForm {
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

// NewInstallIndoorPlatoonCtrl 构造函数
func NewInstallIndoorPlatoonCtrl() *InstallIndoorPlatoonCtrl {
	return &InstallIndoorPlatoonCtrl{}
}

// Store 新建
func (InstallIndoorPlatoonCtrl) Store(ctx *gin.Context) {
	var (
		ret    *gorm.DB
		repeat *models.InstallIndoorPlatoonMdl
	)

	// 表单
	form := InstallIndoorPlatoonStoreForm{}.ShouldBind(ctx)

	// 查重
	ret = models.NewInstallIndoorPlatoonMdl().
		GetDb("").
		Where("name = ?", form.Name).
		First(&repeat)
	wrongs.ThrowWhenNotEmpty(ret, "室内上道位置-排名称")

	// 新建
	installIndoorPlatoon := &models.InstallIndoorPlatoonMdl{
		Name:                  form.Name,
		InstallIndoorRoomUuid: form.installIndoorRoom.Uuid,
	}
	if ret = models.NewInstallIndoorPlatoonMdl().
		GetDb("").
		Create(&installIndoorPlatoon); ret.Error != nil {
		wrongs.ThrowForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Created(map[string]any{"install_indoor_platoon": installIndoorPlatoon}).ToGinResponse())
}

// Destroy 删除
func (InstallIndoorPlatoonCtrl) Destroy(ctx *gin.Context) {
	var (
		ret                  *gorm.DB
		installIndoorPlatoon *models.InstallIndoorPlatoonMdl
	)

	// 查询
	ret = models.NewInstallIndoorPlatoonMdl().
		GetDb("").
		Where("uuid = ?", ctx.Param("uuid")).
		First(&installIndoorPlatoon)
	wrongs.ThrowWhenEmpty(ret, "室内上道位置-排")

	// 删除
	if ret := models.NewInstallIndoorPlatoonMdl().
		GetDb("").
		Where("uuid = ?", ctx.Param("uuid")).
		Delete(&installIndoorPlatoon); ret.Error != nil {
		wrongs.ThrowForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Deleted().ToGinResponse())
}

// DestroyMany 批量删除
func (InstallIndoorPlatoonCtrl) DestroyMany(ctx *gin.Context) {
	form := InstallIndoorPlatoonDestroyManyForm{}.ShouldBind(ctx)
	if ret := models.NewInstallIndoorPlatoonMdl().GetDb("").Where("uuid in ?", form.Uuids).Delete(nil); ret.Error != nil {
		wrongs.ThrowForbidden("删除失败：%s", ret.Error.Error())
	}
	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Deleted().ToGinResponse())
}

// Update 编辑
func (InstallIndoorPlatoonCtrl) Update(ctx *gin.Context) {
	var (
		ret                          *gorm.DB
		installIndoorPlatoon, repeat *models.InstallIndoorPlatoonMdl
	)

	// 表单
	form := InstallIndoorPlatoonStoreForm{}.ShouldBind(ctx)

	// 查重
	ret = models.NewInstallIndoorPlatoonMdl().
		GetDb("").
		Where("name = ? and uuid <> ?", form.Name, ctx.Param("uuid")).
		First(&repeat)
	wrongs.ThrowWhenNotEmpty(ret, "室内上道位置-排名称")

	// 查询
	ret = models.NewInstallIndoorPlatoonMdl().
		GetDb("").
		Where("uuid = ?", ctx.Param("uuid")).
		First(&installIndoorPlatoon)
	wrongs.ThrowWhenEmpty(ret, "室内上道位置-排")

	// 编辑
	installIndoorPlatoon.Name = form.Name
	installIndoorPlatoon.InstallIndoorRoomUuid = form.installIndoorRoom.Uuid
	if ret = models.NewInstallIndoorPlatoonMdl().
		GetDb("").
		Where("uuid = ?", ctx.Param("uuid")).
		Save(&installIndoorPlatoon); ret.Error != nil {
		wrongs.ThrowForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Updated(map[string]any{"install_indoor_platoon": installIndoorPlatoon}).ToGinResponse())
}

// Detail 详情
func (InstallIndoorPlatoonCtrl) Detail(ctx *gin.Context) {
	var (
		ret                  *gorm.DB
		installIndoorPlatoon *models.InstallIndoorPlatoonMdl
	)
	ret = models.NewInstallIndoorPlatoonMdl().
		SetCtx(ctx).
		GetDbUseQuery("").
		Where("uuid = ?", ctx.Param("uuid")).
		First(&installIndoorPlatoon)
	wrongs.ThrowWhenEmpty(ret, "室内上道位置-排")

	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Datum(map[string]any{"install_indoor_platoon": installIndoorPlatoon}).ToGinResponse())
}

// List 列表
func (receiver InstallIndoorPlatoonCtrl) List(ctx *gin.Context) {
	var installIndoorPlatoons []*models.InstallIndoorPlatoonMdl

	ctx.JSON(
		tools.NewCorrectWithGinContext("", ctx).
			DataForPager(
				models.InstallIndoorPlatoonMdl{}.GetListByQuery(ctx),
				func(db *gorm.DB) map[string]any {
					db.Find(&installIndoorPlatoons)
					return map[string]any{"install_indoor_platoons": installIndoorPlatoons}
				},
			).
			ToGinResponse(),
	)
}

// ListJdt jquery-dataTable后端分页数据
func (receiver InstallIndoorPlatoonCtrl) ListJdt(ctx *gin.Context) {
	var installIndoorPlatoons []*models.InstallIndoorPlatoonMdl

	ctx.JSON(
		tools.NewCorrectWithGinContext("", ctx).
			DataForJqueryDataTable(
				models.InstallIndoorPlatoonMdl{}.GetListByQuery(ctx),
				func(db *gorm.DB) map[string]any {
					db.Find(&installIndoorPlatoons)
					return map[string]any{"install_indoor_platoons": installIndoorPlatoons}
				},
			).
			ToGinResponse(),
	)
}
