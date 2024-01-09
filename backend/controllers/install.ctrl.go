package controllers

import (
	"fmt"
	"new-fix/models"
	"new-fix/utils"
	"new-fix/wrongs"
	"strconv"

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

	// InstallIndoorShelfCtrl 室内上道位置-架控制器
	InstallIndoorShelfCtrl struct{}
	// InstallIndoorShelfStoreForm 室内上道位置-架表单
	InstallIndoorShelfStoreForm struct {
		Name                     string `json:"name"`
		InstallIndoorPlatoonUuid string `json:"install_indoor_platoon_uuid"`
		installIndoorPlatoon     *models.InstallIndoorPlatoonMdl
	}
	// InstlalIndoorShelfDestroyManyForm 批量删除室内上道位置-架表单
	InstlalIndoorShelfDestroyManyForm struct {
		Uuids []string `json:"uuids"`
	}

	// InstallIndoorTierCtrl 室内上道位置-层控制器
	InstallIndoorTierCtrl struct{}
	// InstallIndoorTierStoreForm 室内上道位置-层表单
	InstallIndoorTierStoreForm struct {
		Name                   string `json:"name"`
		InstallIndoorShelfUuid string `json:"install_indoor_shelf_uuid"`
		installIndoorShelf     *models.InstallIndoorShelfMdl
	}
	// InstallIndoorTierStoreManyForm 批量新建室内上道位置-层表单
	InstallIndoorTierStoreManyForm struct {
		InstallIndoorShelfUuid string `json:"install_indoor_shelf_uuid"`
		installIndoorShelf     *models.InstallIndoorShelfMdl
		Number                 uint `json:"number"`
	}
	// InstallIndoorTierDestroyManyForm 批量删除室内上道位置-层表单
	InstallIndoorTierDestroyManyForm struct {
		Uuids []string `json:"uuids"`
	}

	// InstallIndoorCellCtrl 室内上道位置-位控制器
	InstallIndoorCellCtrl struct{}
	// InstallIndoorCellStoreForm 室内上道位置-位表单
	InstallIndoorCellStoreForm struct {
		Name                  string `json:"name"`
		InstallIndoorTierUuid string `json:"install_indoor_tier_uuid"`
		installIndoorTier     *models.InstallIndoorTierMdl
	}
	// InstallIndoorCellStoreManyForm 批量新建室内上道位置-位表单
	InstallIndoorCellStoreManyForm struct {
		InstallIndoorTierUuid string `json:"install_indoor_tier_uuid"`
		installIndoorTier     *models.InstallIndoorTierMdl
		Number                uint `json:"number"`
	}
	// InstallIndoorCellDestroyManyForm 批量删除室内上道位置-位表单
	InstallIndoorCellDestroyManyForm struct {
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

	parents := utils.RemoveEmptyStrings([]string{
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

// ShouldBind 室内上道位置-架表单验证
func (receiver InstallIndoorShelfStoreForm) ShouldBind(ctx *gin.Context) InstallIndoorShelfStoreForm {
	if err := ctx.ShouldBind(&receiver); err != nil {
		wrongs.ThrowValidate("表单绑定失败：%s", err.Error())
	}
	if receiver.Name == "" {
		wrongs.ThrowValidate("室内上道位置-架名称必填")
	}
	if receiver.InstallIndoorPlatoonUuid == "" {
		wrongs.ThrowValidate("室内上道位置-架所属排必填")
	} else {
		if ret := models.NewInstallIndoorPlatoonMdl().GetDb("").Where("uuid = ?", receiver.InstallIndoorPlatoonUuid).First(&receiver.installIndoorPlatoon); ret.Error != nil {
			wrongs.ThrowValidate("室内上道位置-架所属排")
		}
	}
	return receiver
}

// ShouldBind 批量删除室内上道位置-架表单验证
func (receiver InstlalIndoorShelfDestroyManyForm) ShouldBind(ctx *gin.Context) InstlalIndoorShelfDestroyManyForm {
	if err := ctx.ShouldBind(&receiver); err != nil {
		wrongs.ThrowValidate("表单绑定失败：%s", err.Error())
	}
	if len(receiver.Uuids) == 0 {
		wrongs.ThrowValidate("请选择需要删除的内容")
	}
	return receiver
}

// ShouldBind 室内上道位置-层表单验证
func (receiver InstallIndoorTierStoreForm) ShouldBind(ctx *gin.Context) InstallIndoorTierStoreForm {
	if err := ctx.ShouldBind(&receiver); err != nil {
		wrongs.ThrowValidate("表单绑定失败：%s", err.Error())
	}
	if receiver.Name == "" {
		wrongs.ThrowValidate("室内上道位置-层名称必填")
	}
	if receiver.InstallIndoorShelfUuid == "" {
		wrongs.ThrowValidate("室内上道位置-层所属架必填")
	} else {
		if ret := models.NewInstallIndoorShelfMdl().GetDb("").Where("uuid = ?", receiver.InstallIndoorShelfUuid).First(&receiver.installIndoorShelf); ret.Error != nil {
			wrongs.ThrowValidate("室内上道位置-层所属架")
		}
	}
	return receiver
}

// InstallIndoorTierStoreManyForm 批量新建室内上道位置-层表单
func (receiver InstallIndoorTierStoreManyForm) ShouldBind(ctx *gin.Context) InstallIndoorTierStoreManyForm {
	if err := ctx.ShouldBind(&receiver); err != nil {
		wrongs.ThrowValidate("表单绑定失败：%s", err.Error())
	}
	if receiver.Number < 1 {
		wrongs.ThrowValidate("增加的层数不能小于1")
	}
	if receiver.InstallIndoorShelfUuid == "" {
		wrongs.ThrowValidate("室内上道位置-架必选")
	} else {
		if ret := models.NewInstallIndoorShelfMdl().GetDb("").Where("uuid = ?", receiver.InstallIndoorShelfUuid).First(&receiver.installIndoorShelf); ret.Error != nil {
			wrongs.ThrowValidate("仓库-架不存在")
		}
	}
	return receiver
}

// ShouldBind 批量删除室内上道位置-层表单验证
func (receiver InstallIndoorTierDestroyManyForm) ShouldBind(ctx *gin.Context) InstallIndoorTierDestroyManyForm {
	if err := ctx.ShouldBind(&receiver); err != nil {
		wrongs.ThrowValidate("表单绑定失败：%s", err.Error())
	}
	if len(receiver.Uuids) == 0 {
		wrongs.ThrowValidate("请选择需要删除的内容")
	}
	return receiver
}

// ShouldBind 室内上道位置-位表单验证
func (receiver InstallIndoorCellStoreForm) ShouldBind(ctx *gin.Context) InstallIndoorCellStoreForm {
	if err := ctx.ShouldBind(&receiver); err != nil {
		wrongs.ThrowValidate("表单绑定失败：%s", err.Error())
	}
	if receiver.Name == "" {
		wrongs.ThrowValidate("室内上道位置-位名称必填")
	}
	if receiver.InstallIndoorTierUuid == "" {
		wrongs.ThrowValidate("室内上道位置-位所属层必填")
	} else {
		if ret := models.NewInstallIndoorTierMdl().GetDb("").Where("uuid = ?", receiver.InstallIndoorTierUuid).First(&receiver.installIndoorTier); ret.Error != nil {
			wrongs.ThrowValidate("室内上道位置-位所属层")
		}
	}
	return receiver
}

// InstallIndoorCellStoreManyForm 批量新建室内上道位置-位表单
func (receiver InstallIndoorCellStoreManyForm) ShouldBind(ctx *gin.Context) InstallIndoorCellStoreManyForm {
	if err := ctx.ShouldBind(&receiver); err != nil {
		wrongs.ThrowValidate("表单绑定失败：%s", err.Error())
	}
	if receiver.Number < 1 {
		wrongs.ThrowValidate("增加的层数不能小于1")
	}
	if receiver.InstallIndoorTierUuid == "" {
		wrongs.ThrowValidate("室内上道位置-层必选")
	} else {
		if ret := models.NewInstallIndoorTierMdl().GetDb("").Where("uuid = ?", receiver.InstallIndoorTierUuid).First(&receiver.installIndoorTier); ret.Error != nil {
			wrongs.ThrowValidate("仓库-架不存在")
		}
	}
	return receiver
}

// ShouldBind 批量删除室内上道位置-位表单验证
func (receiver InstallIndoorCellDestroyManyForm) ShouldBind(ctx *gin.Context) InstallIndoorCellDestroyManyForm {
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

	ctx.JSON(utils.NewCorrectWithGinContext("", ctx).Created(map[string]any{"install_indoor_room_type": installIndoorRoomType}).ToGinResponse())
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

	ctx.JSON(utils.NewCorrectWithGinContext("", ctx).Deleted().ToGinResponse())
}

// DestroyMany 批量删除
func (InstallIndoorRoomTypeCtrl) DestroyMany(ctx *gin.Context) {
	form := InstallIndoorRoomTypeDestroyManyForm{}.ShouldBind(ctx)
	if ret := models.NewInstallIndoorRoomTypeMdl().GetDb("").Where("uuid in ?", form.Uuids).Delete(nil); ret.Error != nil {
		wrongs.ThrowForbidden("删除失败：%s", ret.Error.Error())
	}
	ctx.JSON(utils.NewCorrectWithGinContext("", ctx).Deleted().ToGinResponse())
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

	ctx.JSON(utils.NewCorrectWithGinContext("", ctx).Updated(map[string]any{"install_indoor_room_type": installIndoorRoomType}).ToGinResponse())
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

	ctx.JSON(utils.NewCorrectWithGinContext("", ctx).Datum(map[string]any{"install_indoor_room_type": installIndoorRoomType}).ToGinResponse())
}

// List 列表
func (receiver InstallIndoorRoomTypeCtrl) List(ctx *gin.Context) {
	var installIndoorRoomTypes []*models.InstallIndoorRoomTypeMdl

	ctx.JSON(
		utils.NewCorrectWithGinContext("", ctx).
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
		utils.NewCorrectWithGinContext("", ctx).
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
		ret, db *gorm.DB
		repeat  *models.InstallIndoorRoomMdl
	)

	// 表单
	form := InstallIndoorRoomStoreForm{}.ShouldBind(ctx)

	// 查重
	db = models.NewInstallIndoorRoomMdl().GetDb("")
	if form.OrganizationStationUuid != "" {
		db = db.Where("(name = ? and organization_station_uuid = ?)", form.Name, form.OrganizationStationUuid)
	}
	if form.OrganizationCenterUuid != "" {
		db = db.Where("(name = ? and organization_center_uuid = ?)", form.Name, form.OrganizationCenterUuid)
	}
	if form.OrganizationCrossroadUuid != "" {
		db = db.Where("(name = ? and organization_crossroad_uuid = ?)", form.Name, form.OrganizationCrossroadUuid)
	}
	ret = db.First(&repeat)
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

	ctx.JSON(utils.NewCorrectWithGinContext("", ctx).Created(map[string]any{"install_indoor_room": installIndoorRoom}).ToGinResponse())
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

	ctx.JSON(utils.NewCorrectWithGinContext("", ctx).Deleted().ToGinResponse())
}

// DestroyMany 批量删除
func (InstallIndoorRoomCtrl) DestroyMany(ctx *gin.Context) {
	form := InstallIndoorRoomDestroyManyForm{}.ShouldBind(ctx)
	if ret := models.NewInstallIndoorRoomMdl().GetDb("").Where("uuid in ?", form.Uuids).Delete(nil); ret.Error != nil {
		wrongs.ThrowForbidden("删除失败：%s", ret.Error.Error())
	}
	ctx.JSON(utils.NewCorrectWithGinContext("", ctx).Deleted().ToGinResponse())
}

// Update 编辑
func (InstallIndoorRoomCtrl) Update(ctx *gin.Context) {
	var (
		ret, db                   *gorm.DB
		installIndoorRoom, repeat *models.InstallIndoorRoomMdl
	)

	// 表单
	form := InstallIndoorRoomStoreForm{}.ShouldBind(ctx)

	// 查重
	db = models.NewInstallIndoorRoomMdl().GetDb("")
	if form.OrganizationStationUuid != "" {
		db = db.Where("(name = ? and uuid <> ? and organization_station_uuid = ?)", form.Name, ctx.Param("uuid"), form.OrganizationStationUuid)
	}
	if form.OrganizationCenterUuid != "" {
		db = db.Where("(name = ? and uuid <> ? and organization_center_uuid = ?)", form.Name, ctx.Param("uuid"), form.OrganizationCenterUuid)
	}
	if form.OrganizationCrossroadUuid != "" {
		db = db.Where("(name = ? and uuid <> ? and organization_crossroad_uuid = ?)", form.Name, ctx.Param("uuid"), form.OrganizationCrossroadUuid)
	}
	ret = db.First(&repeat)
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

	ctx.JSON(utils.NewCorrectWithGinContext("", ctx).Updated(map[string]any{"install_indoor_room": installIndoorRoom}).ToGinResponse())
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

	ctx.JSON(utils.NewCorrectWithGinContext("", ctx).Datum(map[string]any{"install_indoor_room": installIndoorRoom}).ToGinResponse())
}

// List 列表
func (receiver InstallIndoorRoomCtrl) List(ctx *gin.Context) {
	var installIndoorRooms []*models.InstallIndoorRoomMdl

	ctx.JSON(
		utils.NewCorrectWithGinContext("", ctx).
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
		utils.NewCorrectWithGinContext("", ctx).
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
		Where("install_indoor_room_uuid = ?", form.installIndoorRoom.Uuid).
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

	ctx.JSON(utils.NewCorrectWithGinContext("", ctx).Created(map[string]any{"install_indoor_platoon": installIndoorPlatoon}).ToGinResponse())
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

	ctx.JSON(utils.NewCorrectWithGinContext("", ctx).Deleted().ToGinResponse())
}

// DestroyMany 批量删除
func (InstallIndoorPlatoonCtrl) DestroyMany(ctx *gin.Context) {
	form := InstallIndoorPlatoonDestroyManyForm{}.ShouldBind(ctx)
	if ret := models.NewInstallIndoorPlatoonMdl().GetDb("").Where("uuid in ?", form.Uuids).Delete(nil); ret.Error != nil {
		wrongs.ThrowForbidden("删除失败：%s", ret.Error.Error())
	}
	ctx.JSON(utils.NewCorrectWithGinContext("", ctx).Deleted().ToGinResponse())
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
		Where("name = ?", form.Name).
		Where("uuid <> ?", ctx.Param("uuid")).
		Where("install_indoor_room_uuid = ?", form.installIndoorRoom.Uuid).
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

	ctx.JSON(utils.NewCorrectWithGinContext("", ctx).Updated(map[string]any{"install_indoor_platoon": installIndoorPlatoon}).ToGinResponse())
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

	ctx.JSON(utils.NewCorrectWithGinContext("", ctx).Datum(map[string]any{"install_indoor_platoon": installIndoorPlatoon}).ToGinResponse())
}

// List 列表
func (receiver InstallIndoorPlatoonCtrl) List(ctx *gin.Context) {
	var installIndoorPlatoons []*models.InstallIndoorPlatoonMdl

	ctx.JSON(
		utils.NewCorrectWithGinContext("", ctx).
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
		utils.NewCorrectWithGinContext("", ctx).
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

// NewInstallIndoorShelfCtrl 构造函数
func NewInstallIndoorShelfCtrl() *InstallIndoorShelfCtrl {
	return &InstallIndoorShelfCtrl{}
}

// Store 新建
func (InstallIndoorShelfCtrl) Store(ctx *gin.Context) {
	var (
		ret    *gorm.DB
		repeat *models.InstallIndoorShelfMdl
	)

	// 表单
	form := InstallIndoorShelfStoreForm{}.ShouldBind(ctx)

	// 查重
	ret = models.NewInstallIndoorShelfMdl().
		GetDb("").
		Where("name = ?", form.Name).
		Where("install_indoor_platoon_uuid = ?", form.installIndoorPlatoon.Uuid).
		First(&repeat)
	wrongs.ThrowWhenNotEmpty(ret, "室内上道位置-架名称")

	// 新建
	installIndoorShelf := &models.InstallIndoorShelfMdl{
		Name:                     form.Name,
		InstallIndoorPlatoonUuid: form.installIndoorPlatoon.Uuid,
	}
	if ret = models.NewInstallIndoorShelfMdl().
		GetDb("").
		Create(&installIndoorShelf); ret.Error != nil {
		wrongs.ThrowForbidden(ret.Error.Error())
	}

	ctx.JSON(utils.NewCorrectWithGinContext("", ctx).Created(map[string]any{"install_indoor_shelf": installIndoorShelf}).ToGinResponse())
}

// Destroy 删除
func (InstallIndoorShelfCtrl) Destroy(ctx *gin.Context) {
	var (
		ret                *gorm.DB
		installIndoorShelf *models.InstallIndoorShelfMdl
	)

	// 查询
	ret = models.NewInstallIndoorShelfMdl().
		GetDb("").
		Where("uuid = ?", ctx.Param("uuid")).
		First(&installIndoorShelf)
	wrongs.ThrowWhenEmpty(ret, "室内上道位置-架")

	// 删除
	if ret := models.NewInstallIndoorShelfMdl().
		GetDb("").
		Where("uuid = ?", ctx.Param("uuid")).
		Delete(&installIndoorShelf); ret.Error != nil {
		wrongs.ThrowForbidden(ret.Error.Error())
	}

	ctx.JSON(utils.NewCorrectWithGinContext("", ctx).Deleted().ToGinResponse())
}

// DestroyMany 批量删除
func (InstallIndoorShelfCtrl) DestroyMany(ctx *gin.Context) {
	form := InstlalIndoorShelfDestroyManyForm{}.ShouldBind(ctx)
	if ret := models.NewInstallIndoorShelfMdl().GetDb("").Where("uuid in ?", form.Uuids).Delete(nil); ret.Error != nil {
		wrongs.ThrowForbidden("删除失败：%s", ret.Error.Error())
	}
	ctx.JSON(utils.NewCorrectWithGinContext("", ctx).Deleted().ToGinResponse())
}

// Update 编辑
func (InstallIndoorShelfCtrl) Update(ctx *gin.Context) {
	var (
		ret                        *gorm.DB
		installIndoorShelf, repeat *models.InstallIndoorShelfMdl
	)

	// 表单
	form := InstallIndoorShelfStoreForm{}.ShouldBind(ctx)

	// 查重
	ret = models.NewInstallIndoorShelfMdl().
		GetDb("").
		Where("name = ?", form.Name).
		Where("uuid <> ?", ctx.Param("uuid")).
		Where("install_indoor_platoon_uuid = ?", form.installIndoorPlatoon.Uuid).
		First(&repeat)
	wrongs.ThrowWhenNotEmpty(ret, "室内上道位置-架名称")

	// 查询
	ret = models.NewInstallIndoorShelfMdl().
		GetDb("").
		Where("uuid = ?", ctx.Param("uuid")).
		First(&installIndoorShelf)
	wrongs.ThrowWhenEmpty(ret, "室内上道位置-架")

	// 编辑
	installIndoorShelf.Name = form.Name
	if ret = models.NewInstallIndoorShelfMdl().
		GetDb("").
		Where("uuid = ?", ctx.Param("uuid")).
		Save(&installIndoorShelf); ret.Error != nil {
		wrongs.ThrowForbidden(ret.Error.Error())
	}

	ctx.JSON(utils.NewCorrectWithGinContext("", ctx).Updated(map[string]any{"install_indoor_shelf": installIndoorShelf}).ToGinResponse())
}

// Detail 详情
func (InstallIndoorShelfCtrl) Detail(ctx *gin.Context) {
	var (
		ret                *gorm.DB
		installIndoorShelf *models.InstallIndoorShelfMdl
	)
	ret = models.NewInstallIndoorShelfMdl().
		SetCtx(ctx).
		GetDbUseQuery("").
		Where("uuid = ?", ctx.Param("uuid")).
		First(&installIndoorShelf)
	wrongs.ThrowWhenEmpty(ret, "室内上道位置-架")

	ctx.JSON(utils.NewCorrectWithGinContext("", ctx).Datum(map[string]any{"install_indoor_shelf": installIndoorShelf}).ToGinResponse())
}

// List 列表
func (receiver InstallIndoorShelfCtrl) List(ctx *gin.Context) {
	var installIndoorShelves []*models.InstallIndoorShelfMdl

	ctx.JSON(
		utils.NewCorrectWithGinContext("", ctx).
			DataForPager(
				models.InstallIndoorShelfMdl{}.GetListByQuery(ctx),
				func(db *gorm.DB) map[string]any {
					db.Find(&installIndoorShelves)
					return map[string]any{"install_indoor_shelves": installIndoorShelves}
				},
			).
			ToGinResponse(),
	)
}

// ListJdt jquery-dataTable后端分页数据
func (receiver InstallIndoorShelfCtrl) ListJdt(ctx *gin.Context) {
	var installIndoorShelves []*models.InstallIndoorShelfMdl

	ctx.JSON(
		utils.NewCorrectWithGinContext("", ctx).
			DataForJqueryDataTable(
				models.InstallIndoorShelfMdl{}.GetListByQuery(ctx),
				func(db *gorm.DB) map[string]any {
					db.Find(&installIndoorShelves)
					return map[string]any{"install_indoor_shelves": installIndoorShelves}
				},
			).
			ToGinResponse(),
	)
}

// NewInstallIndoorTierCtrl 构造函数
func NewInstallIndoorTierCtrl() *InstallIndoorTierCtrl {
	return &InstallIndoorTierCtrl{}
}

// Store 新建
func (InstallIndoorTierCtrl) Store(ctx *gin.Context) {
	var (
		ret    *gorm.DB
		repeat *models.InstallIndoorTierMdl
	)

	// 表单
	form := InstallIndoorTierStoreForm{}.ShouldBind(ctx)

	// 查重
	ret = models.NewInstallIndoorTierMdl().
		GetDb("").
		Where("name = ?", form.Name).
		Where("install_indoor_shelf_uuid", form.installIndoorShelf.Uuid).
		First(&repeat)
	wrongs.ThrowWhenNotEmpty(ret, "室内上道位置-层名称")

	// 新建
	installIndoorTier := &models.InstallIndoorTierMdl{
		Name:                   form.Name,
		InstallIndoorShelfUuid: form.installIndoorShelf.Uuid,
	}
	if ret = models.NewInstallIndoorTierMdl().
		GetDb("").
		Create(&installIndoorTier); ret.Error != nil {
		wrongs.ThrowForbidden(ret.Error.Error())
	}

	ctx.JSON(utils.NewCorrectWithGinContext("", ctx).Created(map[string]any{"install_indoor_tier": installIndoorTier}).ToGinResponse())
}

// StoreMany 批量新建
func (InstallIndoorTierCtrl) StoreMany(ctx *gin.Context) {
	var (
		lastInstallIndoorTierNameInt int
		err                          error
		newInstallIndoorTiers        []*models.InstallIndoorTierMdl
	)
	form := InstallIndoorTierStoreManyForm{}.ShouldBind(ctx)
	lastInstallIndoorTier := models.InstallIndoorTierMdl{}.GetLastByInstallIndoorShelf(form.installIndoorShelf)
	lastInstallIndoorTierNameInt, err = strconv.Atoi(lastInstallIndoorTier.Name)
	if err != nil {
		lastInstallIndoorTierNameInt = 0
	}
	for i := 0; i < int(form.Number); i++ {
		newInstallIndoorTiers = append(newInstallIndoorTiers, &models.InstallIndoorTierMdl{
			Name:                   strconv.Itoa(lastInstallIndoorTierNameInt + i + 1),
			InstallIndoorShelfUuid: form.installIndoorShelf.Uuid,
		})
	}

	models.NewInstallIndoorTierMdl().GetDb("").Create(&newInstallIndoorTiers)

	ctx.JSON(utils.NewCorrectWithGinContext(fmt.Sprintf("成功新建：%d层", len(newInstallIndoorTiers)), ctx).Created(map[string]any{"install_indoor_tiers": newInstallIndoorTiers}).ToGinResponse())
}

// Destroy 删除
func (InstallIndoorTierCtrl) Destroy(ctx *gin.Context) {
	var (
		ret               *gorm.DB
		installIndoorTier *models.InstallIndoorTierMdl
	)

	// 查询
	ret = models.NewInstallIndoorTierMdl().
		GetDb("").
		Where("uuid = ?", ctx.Param("uuid")).
		First(&installIndoorTier)
	wrongs.ThrowWhenEmpty(ret, "室内上道位置-层")

	// 删除
	if ret := models.NewInstallIndoorTierMdl().
		GetDb("").
		Where("uuid = ?", ctx.Param("uuid")).
		Delete(&installIndoorTier); ret.Error != nil {
		wrongs.ThrowForbidden(ret.Error.Error())
	}

	ctx.JSON(utils.NewCorrectWithGinContext("", ctx).Deleted().ToGinResponse())
}

// DestroyMany 批量删除
func (InstallIndoorTierCtrl) DestroyMany(ctx *gin.Context) {
	form := InstallIndoorTierDestroyManyForm{}.ShouldBind(ctx)
	if ret := models.NewInstallIndoorTierMdl().GetDb("").Where("uuid in ?", form.Uuids).Delete(nil); ret.Error != nil {
		wrongs.ThrowForbidden("删除失败：%s", ret.Error.Error())
	}
	ctx.JSON(utils.NewCorrectWithGinContext("", ctx).Deleted().ToGinResponse())
}

// Update 编辑
func (InstallIndoorTierCtrl) Update(ctx *gin.Context) {
	var (
		ret                       *gorm.DB
		installIndoorTier, repeat *models.InstallIndoorTierMdl
	)

	// 表单
	form := InstallIndoorTierStoreForm{}.ShouldBind(ctx)

	// 查重
	ret = models.NewInstallIndoorTierMdl().
		GetDb("").
		Where("name = ?", form.Name).
		Where("uuid <> ?", ctx.Param("uuid")).
		Where("install_indoor_shelf_uuid = ?", form.installIndoorShelf.Uuid).
		First(&repeat)
	wrongs.ThrowWhenNotEmpty(ret, "室内上道位置-层名称")

	// 查询
	ret = models.NewInstallIndoorTierMdl().
		GetDb("").
		Where("uuid = ?", ctx.Param("uuid")).
		First(&installIndoorTier)
	wrongs.ThrowWhenEmpty(ret, "室内上道位置-层")

	// 编辑
	installIndoorTier.Name = form.Name
	installIndoorTier.InstallIndoorShelfUuid = form.installIndoorShelf.Uuid
	if ret = models.NewInstallIndoorTierMdl().
		GetDb("").
		Where("uuid = ?", ctx.Param("uuid")).
		Save(&installIndoorTier); ret.Error != nil {
		wrongs.ThrowForbidden(ret.Error.Error())
	}

	ctx.JSON(utils.NewCorrectWithGinContext("", ctx).Updated(map[string]any{"install_indoor_tier": installIndoorTier}).ToGinResponse())
}

// Detail 详情
func (InstallIndoorTierCtrl) Detail(ctx *gin.Context) {
	var (
		ret               *gorm.DB
		installIndoorTier *models.InstallIndoorTierMdl
	)
	ret = models.NewInstallIndoorTierMdl().
		SetCtx(ctx).
		GetDbUseQuery("").
		Where("uuid = ?", ctx.Param("uuid")).
		First(&installIndoorTier)
	wrongs.ThrowWhenEmpty(ret, "室内上道位置-层")

	ctx.JSON(utils.NewCorrectWithGinContext("", ctx).Datum(map[string]any{"install_indoor_tier": installIndoorTier}).ToGinResponse())
}

// List 列表
func (receiver InstallIndoorTierCtrl) List(ctx *gin.Context) {
	var installIndoorTiers []*models.InstallIndoorTierMdl

	ctx.JSON(
		utils.NewCorrectWithGinContext("", ctx).
			DataForPager(
				models.InstallIndoorTierMdl{}.GetListByQuery(ctx),
				func(db *gorm.DB) map[string]any {
					db.Find(&installIndoorTiers)
					return map[string]any{"install_indoor_tiers": installIndoorTiers}
				},
			).
			ToGinResponse(),
	)
}

// ListJdt jquery-dataTable后端分页数据
func (receiver InstallIndoorTierCtrl) ListJdt(ctx *gin.Context) {
	var installIndoorTiers []*models.InstallIndoorTierMdl

	ctx.JSON(
		utils.NewCorrectWithGinContext("", ctx).
			DataForJqueryDataTable(
				models.InstallIndoorTierMdl{}.GetListByQuery(ctx),
				func(db *gorm.DB) map[string]any {
					db.Find(&installIndoorTiers)
					return map[string]any{"install_indoor_tiers": installIndoorTiers}
				},
			).
			ToGinResponse(),
	)
}

// NewInstallIndoorCellCtrl 构造函数
func NewInstallIndoorCellCtrl() *InstallIndoorCellCtrl {
	return &InstallIndoorCellCtrl{}
}

// Store 新建
func (InstallIndoorCellCtrl) Store(ctx *gin.Context) {
	var (
		ret    *gorm.DB
		repeat *models.InstallIndoorCellMdl
	)

	// 表单
	form := InstallIndoorCellStoreForm{}.ShouldBind(ctx)

	// 查重
	ret = models.NewInstallIndoorCellMdl().
		GetDb("").
		Where("name = ?", form.Name).
		First(&repeat)
	wrongs.ThrowWhenNotEmpty(ret, "室内上道位置-位名称")

	// 新建
	installIndoorCell := &models.InstallIndoorCellMdl{
		Name:                  form.Name,
		InstallIndoorTierUuid: form.installIndoorTier.Uuid,
	}
	if ret = models.NewInstallIndoorCellMdl().
		GetDb("").
		Create(&installIndoorCell); ret.Error != nil {
		wrongs.ThrowForbidden(ret.Error.Error())
	}

	ctx.JSON(utils.NewCorrectWithGinContext("", ctx).Created(map[string]any{"install_indoor_cell": installIndoorCell}).ToGinResponse())
}

// StoreMany 批量新建
func (InstallIndoorCellCtrl) StoreMany(ctx *gin.Context) {
	var (
		lastInstallIndoorCellNameInt int
		err                          error
		newInstallIndoorCells        []*models.InstallIndoorCellMdl
	)
	form := InstallIndoorCellStoreManyForm{}.ShouldBind(ctx)
	lastInstallIndoorTier := models.InstallIndoorCellMdl{}.GetLastByInstallIndoorTier(form.installIndoorTier)
	lastInstallIndoorCellNameInt, err = strconv.Atoi(lastInstallIndoorTier.Name)
	if err != nil {
		lastInstallIndoorCellNameInt = 0
	}
	for i := 0; i < int(form.Number); i++ {
		newInstallIndoorCells = append(newInstallIndoorCells, &models.InstallIndoorCellMdl{
			Name:                  strconv.Itoa(lastInstallIndoorCellNameInt + i + 1),
			InstallIndoorTierUuid: form.installIndoorTier.Uuid,
		})
	}

	models.NewInstallIndoorTierMdl().GetDb("").Create(&newInstallIndoorCells)

	ctx.JSON(utils.NewCorrectWithGinContext(fmt.Sprintf("成功新建：%d位", len(newInstallIndoorCells)), ctx).Created(map[string]any{"install_indoorj_cells": newInstallIndoorCells}).ToGinResponse())
}

// Destroy 删除
func (InstallIndoorCellCtrl) Destroy(ctx *gin.Context) {
	var (
		ret               *gorm.DB
		installIndoorCell *models.InstallIndoorCellMdl
	)

	// 查询
	ret = models.NewInstallIndoorCellMdl().
		GetDb("").
		Where("uuid = ?", ctx.Param("uuid")).
		First(&installIndoorCell)
	wrongs.ThrowWhenEmpty(ret, "室内上道位置-位")

	// 删除
	if ret := models.NewInstallIndoorCellMdl().
		GetDb("").
		Where("uuid = ?", ctx.Param("uuid")).
		Delete(&installIndoorCell); ret.Error != nil {
		wrongs.ThrowForbidden(ret.Error.Error())
	}

	ctx.JSON(utils.NewCorrectWithGinContext("", ctx).Deleted().ToGinResponse())
}

// DestroyMany 批量删除
func (InstallIndoorCellCtrl) DestroyMany(ctx *gin.Context) {
	form := InstallIndoorCellDestroyManyForm{}.ShouldBind(ctx)
	if ret := models.NewInstallIndoorCellMdl().GetDb("").Where("uuid in ?", form.Uuids).Delete(nil); ret.Error != nil {
		wrongs.ThrowForbidden("删除失败：%s", ret.Error.Error())
	}
	ctx.JSON(utils.NewCorrectWithGinContext("", ctx).Deleted().ToGinResponse())
}

// Update 编辑
func (InstallIndoorCellCtrl) Update(ctx *gin.Context) {
	var (
		ret                       *gorm.DB
		installIndoorCell, repeat *models.InstallIndoorCellMdl
	)

	// 表单
	form := InstallIndoorCellStoreForm{}.ShouldBind(ctx)

	// 查重
	ret = models.NewInstallIndoorCellMdl().
		GetDb("").
		Where("name = ? and uuid <> ?", form.Name).
		Where("uuid <> ?", ctx.Param("uuid")).
		Where("install_indoor_tier_uuid = ?", form.installIndoorTier.Uuid).
		First(&repeat)
	wrongs.ThrowWhenNotEmpty(ret, "室内上道位置-位名称")

	// 查询
	ret = models.NewInstallIndoorCellMdl().
		GetDb("").
		Where("uuid = ?", ctx.Param("uuid")).
		First(&installIndoorCell)
	wrongs.ThrowWhenEmpty(ret, "室内上道位置-位")

	// 编辑
	installIndoorCell.Name = form.Name
	installIndoorCell.InstallIndoorTierUuid = form.installIndoorTier.Uuid
	if ret = models.NewInstallIndoorCellMdl().
		GetDb("").
		Where("uuid = ?", ctx.Param("uuid")).
		Save(&installIndoorCell); ret.Error != nil {
		wrongs.ThrowForbidden(ret.Error.Error())
	}

	ctx.JSON(utils.NewCorrectWithGinContext("", ctx).Updated(map[string]any{"install_indoor_cell": installIndoorCell}).ToGinResponse())
}

// Detail 详情
func (InstallIndoorCellCtrl) Detail(ctx *gin.Context) {
	var (
		ret               *gorm.DB
		installIndoorCell *models.InstallIndoorCellMdl
	)
	ret = models.NewInstallIndoorCellMdl().
		SetCtx(ctx).
		GetDbUseQuery("").
		Where("uuid = ?", ctx.Param("uuid")).
		First(&installIndoorCell)
	wrongs.ThrowWhenEmpty(ret, "室内上道位置-位")

	ctx.JSON(utils.NewCorrectWithGinContext("", ctx).Datum(map[string]any{"install_indoor_cell": installIndoorCell}).ToGinResponse())
}

// List 列表
func (receiver InstallIndoorCellCtrl) List(ctx *gin.Context) {
	var installIndoorCells []*models.InstallIndoorCellMdl

	ctx.JSON(
		utils.NewCorrectWithGinContext("", ctx).
			DataForPager(
				models.InstallIndoorCellMdl{}.GetListByQuery(ctx),
				func(db *gorm.DB) map[string]any {
					db.Find(&installIndoorCells)
					return map[string]any{"install_indoor_cells": installIndoorCells}
				},
			).
			ToGinResponse(),
	)
}

// ListJdt jquery-dataTable后端分页数据
func (receiver InstallIndoorCellCtrl) ListJdt(ctx *gin.Context) {
	var installIndoorCells []*models.InstallIndoorCellMdl

	ctx.JSON(
		utils.NewCorrectWithGinContext("", ctx).
			DataForJqueryDataTable(
				models.InstallIndoorCellMdl{}.GetListByQuery(ctx),
				func(db *gorm.DB) map[string]any {
					db.Find(&installIndoorCells)
					return map[string]any{"install_indoor_cells": installIndoorCells}
				},
			).
			ToGinResponse(),
	)
}
