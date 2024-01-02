package controllers

import (
	"fmt"
	"new-fix/models"
	"new-fix/tools"
	"new-fix/wrongs"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type (
	// WarehouseStorehouseCtrl 仓库-库房控制器
	WarehouseStorehouseCtrl struct{}
	// WarehouseStorehouseStoreForm 仓库-库房表单
	WarehouseStorehouseStoreForm struct {
		Name                     string `json:"name"`
		OrganizationWorkshopUuid string `json:"organization_workshop_uuid"`
		organizationWorkshop     *models.OrganizationWorkshopMdl
		OrganizationWorkAreaUuid string `json:"organization_work_area_uuid"`
		organizationWorkArea     *models.OrganizationWorkAreaMdl
	}
	// WarehouseStorehouseDestroyManyForm 仓库-库房批量删除表单
	WarehouseStorehouseDestroyManyForm struct {
		Uuids []string `json:"uuids"`
	}

	// WarehouseAreaCtrl 仓库-库区控制器
	WarehouseAreaCtrl struct{}
	// WarehouseAreaStoreForm 仓库-库区表单
	WarehouseAreaStoreForm struct {
		Name                    string `json:"name"`
		WarehouseStorehouseUuid string `json:"warehouse_storehouse_uuid"`
		warehouseStorehouse     *models.WarehouseStorehouseMdl
	}
	// WarehouseAreaDestroyManyForm 仓库-库区批量删除表单
	WarehouseAreaDestroyManyForm struct {
		Uuids []string `json:"uuids"`
	}

	// WarehousePlatoonCtrl 仓库-排控制器
	WarehousePlatoonCtrl struct{}
	// WarehousePlatoonStoreForm 仓库-排表单
	WarehousePlatoonStoreForm struct {
		Name              string `json:"name"`
		WarehouseAreaUuid string `json:"warehouse_area_uuid"`
		warehouseArea     *models.WarehouseAreaMdl
		TypeCode          string `json:"type_code"`
	}
	// WarehousePlatoonDestroyManyForm 仓库-排批量删除表单
	WarehousePlatoonDestroyManyForm struct {
		Uuids []string `json:"uuids"`
	}

	// WarehouseShelfCtrl 仓库-架控制器
	WarehouseShelfCtrl struct{}
	// WarehouseShelfStoreForm 仓库-架表单
	WarehouseShelfStoreForm struct {
		Name                 string `json:"name"`
		WarehousePlatoonUuid string `json:"warehouse_platoon_uuid"`
		warehousePlatoon     *models.WarehousePlatoonMdl
	}
	// WarehouseShelfDestroyManyForm 仓库-架批量删除表单
	WarehouseShelfDestroyManyForm struct {
		Uuids []string `json:"uuids"`
	}

	// WarehouseTierCtrl 仓库-层控制器
	WarehouseTierCtrl struct{}
	// WarehouseTierStoreForm 仓库-层表单
	WarehouseTierStoreForm struct {
		Name               string `json:"name"`
		WarehouseShelfUuid string `json:"warehouse_shelf_uuid"`
		warehouseShelf     *models.WarehouseShelfMdl
	}
	// WarehouseTierStoreManyForm 仓库-层批量新建表单
	WarehouseTierStoreManyForm struct {
		WarehouseShelfUuid string `json:"warehouse_shelf_uuid"`
		warehouseShelf     *models.WarehouseShelfMdl
		Number             uint `json:"number"`
	}
	// WarehouseTierDestroyManyForm 仓库-层批量删除表单
	WarehouseTierDestroyManyForm struct {
		Uuids []string `json:"uuids"`
	}

	// WarehouseCellCtrl 仓库-位控制器
	WarehouseCellCtrl struct{}
	// WarehouseCellStoreForm 仓库-位表单
	WarehouseCellStoreForm struct {
		Name              string `json:"name"`
		WarehouseTierUuid string `json:"warehouse_tier_uuid"`
		warehouseTier     *models.WarehouseTierMdl
	}
	// WarehouseCellStoreManyForm 仓库-位批量新建表单
	WarehouseCellStoreManyForm struct {
		WarehouseTierUuid string `json:"warehouse_tier_uuid"`
		warehouseTier     *models.WarehouseTierMdl
		Number            uint `json:"number"`
	}
	// WarehouseCellDestroyManyForm 仓库-位批量删除表单
	WarehouseCellDestroyManyForm struct {
		Uuids []string `json:"uuids"`
	}
)

// WarehouseStorehouseStoreForm 仓库-库房表单
func (receiver WarehouseStorehouseStoreForm) ShouldBind(ctx *gin.Context) WarehouseStorehouseStoreForm {
	if err := ctx.ShouldBind(&receiver); err != nil {
		wrongs.ThrowValidate("表单绑定失败：%s", err.Error())
	}
	if receiver.Name == "" {
		wrongs.ThrowValidate("库房名称必填")
	}
	if receiver.OrganizationWorkshopUuid == "" {
		wrongs.ThrowValidate("组织结构-车间必选")
	} else {
		ret := models.NewOrganizationWorkshopMdl().GetDb("").Where("uuid = ?", receiver.OrganizationWorkshopUuid).First(&receiver.organizationWorkshop)
		wrongs.ThrowWhenEmpty(ret, "组织结构-车间")
	}
	if receiver.OrganizationWorkAreaUuid != "" {
		ret := models.NewOrganizationWorkAreaMdl().GetDb("").Where("uuid = ?", receiver.OrganizationWorkAreaUuid).First(&receiver.organizationWorkArea)
		wrongs.ThrowWhenEmpty(ret, "组织结构-工区")
	}
	return receiver
}

// WarehouseStorehouseDestroyManyForm 仓库-库房批量删除表单
func (receiver WarehouseStorehouseDestroyManyForm) ShouldBind(ctx *gin.Context) WarehouseStorehouseDestroyManyForm {
	if err := ctx.ShouldBind(&receiver); err != nil {
		wrongs.ThrowValidate("表单绑定失败：%s", err.Error())
	}
	if len(receiver.Uuids) == 0 {
		wrongs.ThrowValidate("仓库-库房编号必填")
	}
	return receiver
}

// WarehouseAreaStoreForm 仓库-库区表单
func (receiver WarehouseAreaStoreForm) ShouldBind(ctx *gin.Context) WarehouseAreaStoreForm {
	if err := ctx.ShouldBind(&receiver); err != nil {
		wrongs.ThrowValidate("表单绑定失败：%s", err.Error())
	}
	if receiver.Name == "" {
		wrongs.ThrowValidate("仓库-库区名称必填")
	}
	if receiver.WarehouseStorehouseUuid == "" {
		wrongs.ThrowValidate("仓库-库房必选")
	} else {
		ret := models.NewWarehouseStorehouseMdl().GetDb("").Where("uuid = ?", receiver.WarehouseStorehouseUuid).First(&receiver.warehouseStorehouse)
		wrongs.ThrowWhenEmpty(ret, "仓库-库房")
	}
	return receiver
}

// WarehouseAreaDestroyManyForm 仓库-库区批量删除表单
func (receiver WarehouseAreaDestroyManyForm) ShouldBind(ctx *gin.Context) WarehouseAreaDestroyManyForm {
	if err := ctx.ShouldBind(&receiver); err != nil {
		wrongs.ThrowValidate("表单绑定失败：%s", err.Error())
	}
	if len(receiver.Uuids) == 0 {
		wrongs.ThrowValidate("仓库-库区编号必填")
	}
	return receiver
}

// WarehousePlatoonStoreForm 仓库-排表单
func (receiver WarehousePlatoonStoreForm) ShouldBind(ctx *gin.Context) WarehousePlatoonStoreForm {
	if err := ctx.ShouldBind(&receiver); err != nil {
		wrongs.ThrowValidate("表单绑定失败：%s", err.Error())
	}
	if receiver.Name == "" {
		wrongs.ThrowValidate("仓库-排名称必填")
	}
	if receiver.WarehouseAreaUuid == "" {
		wrongs.ThrowValidate("仓库-库区必选")
	} else {
		ret := models.NewWarehouseAreaMdl().GetDb("").Where("uuid = ?", receiver.WarehouseAreaUuid).First(&receiver.warehouseArea)
		wrongs.ThrowWhenEmpty(ret, "仓库-库区")
	}
	if !tools.InString(receiver.TypeCode, models.WarehousePlatoonMdl{}.GetTypeCodes()) {
		wrongs.ThrowValidate("仓库-排类型错误")
	}
	return receiver
}

// WarehousePlatoonDestroyManyForm 仓库-排批量删除表单
func (receiver WarehousePlatoonDestroyManyForm) ShouldBind(ctx *gin.Context) WarehousePlatoonDestroyManyForm {
	if err := ctx.ShouldBind(&receiver); err != nil {
		wrongs.ThrowValidate("表单绑定失败：%s", err.Error())
	}
	if len(receiver.Uuids) == 0 {
		wrongs.ThrowValidate("仓库-排编号必填")
	}
	return receiver
}

// WarehouseShelfStoreForm 仓库-架表单
func (receiver WarehouseShelfStoreForm) ShouldBind(ctx *gin.Context) WarehouseShelfStoreForm {
	if err := ctx.ShouldBind(&receiver); err != nil {
		wrongs.ThrowValidate("表单绑定失败：%s", err.Error())
	}
	if receiver.Name == "" {
		wrongs.ThrowValidate("仓库-架名称必填")
	}
	if receiver.WarehousePlatoonUuid == "" {
		wrongs.ThrowValidate("仓库-排必选")
	} else {
		ret := models.NewWarehousePlatoonMdl().GetDb("").Where("uuid = ?", receiver.WarehousePlatoonUuid).First(&receiver.warehousePlatoon)
		wrongs.ThrowWhenEmpty(ret, "仓库-排")
	}
	return receiver
}

// WarehouseShelfDestroyManyForm 仓库-架批量删除表单
func (receiver WarehouseShelfDestroyManyForm) ShouldBind(ctx *gin.Context) WarehouseShelfDestroyManyForm {
	if err := ctx.ShouldBind(&receiver); err != nil {
		wrongs.ThrowValidate("表单绑定失败：%s", err.Error())
	}
	if len(receiver.Uuids) == 0 {
		wrongs.ThrowValidate("仓库-架编号必填")
	}
	return receiver
}

// WarehouseTierStoreForm 仓库-层表单
func (receiver WarehouseTierStoreForm) ShouldBind(ctx *gin.Context) WarehouseTierStoreForm {
	if err := ctx.ShouldBind(&receiver); err != nil {
		wrongs.ThrowValidate("表单绑定失败：%s", err.Error())
	}
	if receiver.Name == "" {
		wrongs.ThrowValidate("仓库-层名称必填")
	}
	if receiver.WarehouseShelfUuid == "" {
		wrongs.ThrowValidate("仓库-架必选")
	} else {
		ret := models.NewWarehouseShelfMdl().GetDb("").Where("uuid = ?", receiver.WarehouseShelfUuid).First(&receiver.warehouseShelf)
		wrongs.ThrowWhenEmpty(ret, "仓库-架")
	}
	return receiver
}

// WarehouseTierStoreManyForm 仓库-层批量新建表单
func (receiver WarehouseTierStoreManyForm) ShouldBind(ctx *gin.Context) WarehouseTierStoreManyForm {
	if err := ctx.ShouldBind(&receiver); err != nil {
		wrongs.ThrowValidate("表单绑定失败：%s", err.Error())
	}
	if receiver.Number < 1 {
		wrongs.ThrowValidate("增加的层数不能小于1")
	}
	if receiver.WarehouseShelfUuid == "" {
		wrongs.ThrowValidate("仓库-架必选")
	} else {
		if ret := models.NewWarehouseShelfMdl().GetDb("").Where("uuid = ?", receiver.WarehouseShelfUuid).First(&receiver.warehouseShelf); ret.Error != nil {
			wrongs.ThrowValidate("仓库-架不存在")
		}
	}
	return receiver
}

// WarehouseTierDestroyManyForm 仓库-层批量删除表单
func (receiver WarehouseTierDestroyManyForm) ShouldBind(ctx *gin.Context) WarehouseTierDestroyManyForm {
	if err := ctx.ShouldBind(&receiver); err != nil {
		wrongs.ThrowValidate("表单绑定失败：%s", err.Error())
	}
	if len(receiver.Uuids) == 0 {
		wrongs.ThrowValidate("仓库-层编号必填")
	}
	return receiver
}

// WarehouseCellStoreForm 仓库-位表单
func (receiver WarehouseCellStoreForm) ShouldBind(ctx *gin.Context) WarehouseCellStoreForm {
	if err := ctx.ShouldBind(&receiver); err != nil {
		wrongs.ThrowValidate("表单绑定失败：%s", err.Error())
	}
	if receiver.Name == "" {
		wrongs.ThrowValidate("仓库-位名称必填")
	}
	if receiver.WarehouseTierUuid == "" {
		wrongs.ThrowValidate("仓库-层必选")
	} else {
		ret := models.NewWarehouseTierMdl().GetDb("").Where("uuid = ?", receiver.WarehouseTierUuid).First(&receiver.warehouseTier)
		wrongs.ThrowWhenEmpty(ret, "仓库-层")
	}
	return receiver
}

// WarehouseCellStoreManyForm 仓库-位批量新建表单
func (receiver WarehouseCellStoreManyForm) ShouldBind(ctx *gin.Context) WarehouseCellStoreManyForm {
	if err := ctx.ShouldBind(&receiver); err != nil {
		wrongs.ThrowValidate("表单绑定失败：%s", err.Error())
	}
	if receiver.Number < 1 {
		wrongs.ThrowValidate("增加的位数不能小于1")
	}
	if receiver.WarehouseTierUuid == "" {
		wrongs.ThrowValidate("仓库-层必选")
	} else {
		if ret := models.NewWarehouseTierMdl().GetDb("").Where("uuid = ?", receiver.WarehouseTierUuid).First(&receiver.warehouseTier); ret.Error != nil {
			wrongs.ThrowValidate("仓库-层不存在")
		}
	}
	return receiver
}

// WarehouseCellDestroyManyForm 仓库-位批量删除表单
func (receiver WarehouseCellDestroyManyForm) ShouldBind(ctx *gin.Context) WarehouseCellDestroyManyForm {
	if err := ctx.ShouldBind(&receiver); err != nil {
		wrongs.ThrowValidate("表单绑定失败：%s", err.Error())
	}
	if len(receiver.Uuids) == 0 {
		wrongs.ThrowValidate("仓库-位编号必填")
	}
	return receiver
}

// NewWarehouseStorehouseCtrl 构造函数
func NewWarehouseStorehouseCtrl() *WarehouseStorehouseCtrl {
	return &WarehouseStorehouseCtrl{}
}

// Store 新建
func (WarehouseStorehouseCtrl) Store(ctx *gin.Context) {
	var (
		ret    *gorm.DB
		repeat *models.WarehouseStorehouseMdl
	)

	// 表单
	form := WarehouseStorehouseStoreForm{}.ShouldBind(ctx)

	// 查重
	ret = models.NewWarehouseStorehouseMdl().
		GetDb("").
		Where("name = ?", form.Name).
		First(&repeat)
	wrongs.ThrowWhenNotEmpty(ret, "仓库名称")

	// 新建
	warehouseStorehouse := &models.WarehouseStorehouseMdl{
		Name:                     form.Name,
		OrganizationWorkshopUuid: form.organizationWorkshop.Uuid,
		OrganizationWorkAreaUuid: form.OrganizationWorkAreaUuid,
	}
	if ret = models.NewWarehouseStorehouseMdl().
		GetDb("").
		Create(&warehouseStorehouse); ret.Error != nil {
		wrongs.ThrowForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Created(map[string]any{"warehouse_storehouse": warehouseStorehouse}).ToGinResponse())
}

// Destroy 删除
func (WarehouseStorehouseCtrl) Destroy(ctx *gin.Context) {
	var (
		ret                 *gorm.DB
		warehouseStorehouse *models.WarehouseStorehouseMdl
	)

	// 查询
	ret = models.NewWarehouseStorehouseMdl().
		GetDb("").
		Where("uuid = ?", ctx.Param("uuid")).
		First(&warehouseStorehouse)
	wrongs.ThrowWhenEmpty(ret, "仓库")

	// 删除
	if ret := models.NewWarehouseStorehouseMdl().
		GetDb("").
		Where("uuid = ?", ctx.Param("uuid")).
		Delete(&warehouseStorehouse); ret.Error != nil {
		wrongs.ThrowForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Deleted().ToGinResponse())
}

// DestroyMany 批量删除
func (WarehouseStorehouseCtrl) DestroyMany(ctx *gin.Context) {
	form := WarehouseStorehouseDestroyManyForm{}.ShouldBind(ctx)
	if ret := models.NewWarehouseStorehouseMdl().GetDb("").Where("uuid in ?", form.Uuids).Delete(nil); ret.Error != nil {
		wrongs.ThrowForbidden("批量删除失败：%s", ret.Error.Error())
	}
	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Deleted().ToGinResponse())
}

// Update 编辑
func (WarehouseStorehouseCtrl) Update(ctx *gin.Context) {
	var (
		ret                         *gorm.DB
		warehouseStorehouse, repeat *models.WarehouseStorehouseMdl
	)

	// 表单
	form := WarehouseStorehouseStoreForm{}.ShouldBind(ctx)

	// 查重
	ret = models.NewWarehouseStorehouseMdl().
		GetDb("").
		Where("name = ? and uuid <> ?", form.Name, ctx.Param("uuid")).
		First(&repeat)
	wrongs.ThrowWhenNotEmpty(ret, "仓库名称")

	// 查询
	ret = models.NewWarehouseStorehouseMdl().
		GetDb("").
		Where("uuid = ?", ctx.Param("uuid")).
		First(&warehouseStorehouse)
	wrongs.ThrowWhenEmpty(ret, "仓库")

	// 编辑
	warehouseStorehouse.Name = form.Name
	warehouseStorehouse.OrganizationWorkshopUuid = form.organizationWorkshop.Uuid
	warehouseStorehouse.OrganizationWorkAreaUuid = form.OrganizationWorkAreaUuid
	if ret = models.NewWarehouseStorehouseMdl().
		GetDb("").
		Where("uuid = ?", ctx.Param("uuid")).
		Save(&warehouseStorehouse); ret.Error != nil {
		wrongs.ThrowForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Updated(map[string]any{"warehouse_storehouse": warehouseStorehouse}).ToGinResponse())
}

// Detail 详情
func (WarehouseStorehouseCtrl) Detail(ctx *gin.Context) {
	var (
		ret                 *gorm.DB
		warehouseStorehouse *models.WarehouseStorehouseMdl
	)
	ret = models.NewWarehouseStorehouseMdl().
		SetCtx(ctx).
		GetDbUseQuery("").
		Where("uuid = ?", ctx.Param("uuid")).
		First(&warehouseStorehouse)
	wrongs.ThrowWhenEmpty(ret, "仓库")

	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Datum(map[string]any{"warehouse_storehouse": warehouseStorehouse}).ToGinResponse())
}

// List 列表
func (receiver WarehouseStorehouseCtrl) List(ctx *gin.Context) {
	var warehouseStorehouses []*models.WarehouseStorehouseMdl

	ctx.JSON(
		tools.NewCorrectWithGinContext("", ctx).
			DataForPager(
				models.WarehouseStorehouseMdl{}.GetListByQuery(ctx),
				func(db *gorm.DB) map[string]any {
					db.Find(&warehouseStorehouses)
					return map[string]any{"warehouse_storehouses": warehouseStorehouses}
				},
			).
			ToGinResponse(),
	)
}

// ListJdt jquery-dataTable后端分页数据
func (receiver WarehouseStorehouseCtrl) ListJdt(ctx *gin.Context) {
	var warehouseStorehouses []*models.WarehouseStorehouseMdl

	ctx.JSON(
		tools.NewCorrectWithGinContext("", ctx).
			DataForJqueryDataTable(
				models.WarehouseStorehouseMdl{}.GetListByQuery(ctx),
				func(db *gorm.DB) map[string]any {
					db.Find(&warehouseStorehouses)
					return map[string]any{"warehouse_storehouses": warehouseStorehouses}
				},
			).
			ToGinResponse(),
	)
}

// NewWarehouseAreaCtrl 构造函数
func NewWarehouseAreaCtrl() *WarehouseAreaCtrl {
	return &WarehouseAreaCtrl{}
}

// Store 新建
func (WarehouseAreaCtrl) Store(ctx *gin.Context) {
	var (
		ret    *gorm.DB
		repeat *models.WarehouseAreaMdl
	)

	// 表单
	form := WarehouseAreaStoreForm{}.ShouldBind(ctx)

	// 查重
	ret = models.NewWarehouseAreaMdl().
		GetDb("").
		Where("name = ?", form.Name).
		First(&repeat)
	wrongs.ThrowWhenNotEmpty(ret, "仓库-库区名称")

	// 新建
	warehouseArea := &models.WarehouseAreaMdl{
		Name:                    form.Name,
		WarehouseStorehouseUuid: form.warehouseStorehouse.Uuid,
	}
	if ret = models.NewWarehouseAreaMdl().
		GetDb("").
		Create(&warehouseArea); ret.Error != nil {
		wrongs.ThrowForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Created(map[string]any{"warehouse_area": warehouseArea}).ToGinResponse())
}

// Destroy 删除
func (WarehouseAreaCtrl) Destroy(ctx *gin.Context) {
	var (
		ret           *gorm.DB
		warehouseArea *models.WarehouseAreaMdl
	)

	// 查询
	ret = models.NewWarehouseAreaMdl().
		GetDb("").
		Where("uuid = ?", ctx.Param("uuid")).
		First(&warehouseArea)
	wrongs.ThrowWhenEmpty(ret, "仓库-库区")

	// 删除
	if ret := models.NewWarehouseAreaMdl().
		GetDb("").
		Where("uuid = ?", ctx.Param("uuid")).
		Delete(&warehouseArea); ret.Error != nil {
		wrongs.ThrowForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Deleted().ToGinResponse())
}

// DestroyMany 批量删除
func (WarehouseAreaCtrl) DestroyMany(ctx *gin.Context) {
	form := WarehouseAreaDestroyManyForm{}.ShouldBind(ctx)
	if ret := models.NewWarehouseAreaMdl().GetDb("").Where("uuid in ?", form.Uuids).Delete(nil); ret.Error != nil {
		wrongs.ThrowForbidden("批量删除失败：%s", ret.Error.Error())
	}
	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Deleted().ToGinResponse())
}

// Update 编辑
func (WarehouseAreaCtrl) Update(ctx *gin.Context) {
	var (
		ret                   *gorm.DB
		warehouseArea, repeat *models.WarehouseAreaMdl
	)

	// 表单
	form := WarehouseAreaStoreForm{}.ShouldBind(ctx)

	// 查重
	ret = models.NewWarehouseAreaMdl().
		GetDb("").
		Where("name = ? and uuid <> ?", form.Name, ctx.Param("uuid")).
		First(&repeat)
	wrongs.ThrowWhenNotEmpty(ret, "仓库-库区名称")

	// 查询
	ret = models.NewWarehouseAreaMdl().
		GetDb("").
		Where("uuid = ?", ctx.Param("uuid")).
		First(&warehouseArea)
	wrongs.ThrowWhenEmpty(ret, "仓库-库区")

	// 编辑
	warehouseArea.Name = form.Name
	warehouseArea.WarehouseStorehouseUuid = form.warehouseStorehouse.Uuid
	if ret = models.NewWarehouseAreaMdl().
		GetDb("").
		Where("uuid = ?", ctx.Param("uuid")).
		Save(&warehouseArea); ret.Error != nil {
		wrongs.ThrowForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Updated(map[string]any{"warehouse_area": warehouseArea}).ToGinResponse())
}

// Detail 详情
func (WarehouseAreaCtrl) Detail(ctx *gin.Context) {
	var (
		ret           *gorm.DB
		warehouseArea *models.WarehouseAreaMdl
	)
	ret = models.NewWarehouseAreaMdl().
		SetCtx(ctx).
		GetDbUseQuery("").
		Where("uuid = ?", ctx.Param("uuid")).
		First(&warehouseArea)
	wrongs.ThrowWhenEmpty(ret, "仓库-库区")

	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Datum(map[string]any{"warehouse_area": warehouseArea}).ToGinResponse())
}

// List 列表
func (receiver WarehouseAreaCtrl) List(ctx *gin.Context) {
	var warehouseAreas []*models.WarehouseAreaMdl

	ctx.JSON(
		tools.NewCorrectWithGinContext("", ctx).
			DataForPager(
				models.WarehouseAreaMdl{}.GetListByQuery(ctx),
				func(db *gorm.DB) map[string]any {
					db.Find(&warehouseAreas)
					return map[string]any{"warehouse_areas": warehouseAreas}
				},
			).
			ToGinResponse(),
	)
}

// ListJdt jquery-dataTable后端分页数据
func (receiver WarehouseAreaCtrl) ListJdt(ctx *gin.Context) {
	var warehouseAreas []*models.WarehouseAreaMdl

	ctx.JSON(
		tools.NewCorrectWithGinContext("", ctx).
			DataForJqueryDataTable(
				models.WarehouseAreaMdl{}.GetListByQuery(ctx),
				func(db *gorm.DB) map[string]any {
					db.Find(&warehouseAreas)
					return map[string]any{"warehouse_areas": warehouseAreas}
				},
			).
			ToGinResponse(),
	)
}

// NewWarehousePlatoonCtrl 构造函数
func NewWarehousePlatoonCtrl() *WarehousePlatoonCtrl {
	return &WarehousePlatoonCtrl{}
}

// Store 新建
func (WarehousePlatoonCtrl) Store(ctx *gin.Context) {
	var (
		ret    *gorm.DB
		repeat *models.WarehousePlatoonMdl
	)

	// 表单
	form := WarehousePlatoonStoreForm{}.ShouldBind(ctx)

	// 查重
	ret = models.NewWarehousePlatoonMdl().
		GetDb("").
		Where("name = ?", form.Name).
		First(&repeat)
	wrongs.ThrowWhenNotEmpty(ret, "仓库-排名称")

	// 新建
	warehousePlatoon := &models.WarehousePlatoonMdl{
		Name:              form.Name,
		WarehouseAreaUuid: form.warehouseArea.Uuid,
		TypeCode:          form.TypeCode,
	}
	if ret = models.NewWarehousePlatoonMdl().
		GetDb("").
		Create(&warehousePlatoon); ret.Error != nil {
		wrongs.ThrowForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Created(map[string]any{"warehouse_platoon": warehousePlatoon}).ToGinResponse())
}

// Destroy 删除
func (WarehousePlatoonCtrl) Destroy(ctx *gin.Context) {
	var (
		ret              *gorm.DB
		warehousePlatoon *models.WarehousePlatoonMdl
	)

	// 查询
	ret = models.NewWarehousePlatoonMdl().
		GetDb("").
		Where("uuid = ?", ctx.Param("uuid")).
		First(&warehousePlatoon)
	wrongs.ThrowWhenEmpty(ret, "仓库-排")

	// 删除
	if ret := models.NewWarehousePlatoonMdl().
		GetDb("").
		Where("uuid = ?", ctx.Param("uuid")).
		Delete(&warehousePlatoon); ret.Error != nil {
		wrongs.ThrowForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Deleted().ToGinResponse())
}

// DestroyMany 批量删除
func (WarehousePlatoonCtrl) DestroyMany(ctx *gin.Context) {
	form := WarehousePlatoonDestroyManyForm{}.ShouldBind(ctx)
	if ret := models.NewWarehousePlatoonMdl().GetDb("").Where("uuid in ?", form.Uuids).Delete(nil); ret.Error != nil {
		wrongs.ThrowForbidden("批量删除失败：%s", ret.Error.Error())
	}
	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Deleted().ToGinResponse())
}

// Update 编辑
func (WarehousePlatoonCtrl) Update(ctx *gin.Context) {
	var (
		ret                      *gorm.DB
		warehousePlatoon, repeat *models.WarehousePlatoonMdl
	)

	// 表单
	form := WarehousePlatoonStoreForm{}.ShouldBind(ctx)

	// 查重
	ret = models.NewWarehousePlatoonMdl().
		GetDb("").
		Where("name = ? and uuid <> ?", form.Name, ctx.Param("uuid")).
		First(&repeat)
	wrongs.ThrowWhenNotEmpty(ret, "仓库-排名称")

	// 查询
	ret = models.NewWarehousePlatoonMdl().
		GetDb("").
		Where("uuid = ?", ctx.Param("uuid")).
		First(&warehousePlatoon)
	wrongs.ThrowWhenEmpty(ret, "仓库-排")

	// 编辑
	warehousePlatoon.Name = form.Name
	warehousePlatoon.WarehouseAreaUuid = form.warehouseArea.Uuid
	warehousePlatoon.TypeCode = form.TypeCode
	if ret = models.NewWarehousePlatoonMdl().
		GetDb("").
		Where("uuid = ?", ctx.Param("uuid")).
		Save(&warehousePlatoon); ret.Error != nil {
		wrongs.ThrowForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Updated(map[string]any{"warehouse_platoon": warehousePlatoon}).ToGinResponse())
}

// Detail 详情
func (WarehousePlatoonCtrl) Detail(ctx *gin.Context) {
	var (
		ret              *gorm.DB
		warehousePlatoon *models.WarehousePlatoonMdl
	)
	ret = models.NewWarehousePlatoonMdl().
		SetCtx(ctx).
		GetDbUseQuery("").
		Where("uuid = ?", ctx.Param("uuid")).
		First(&warehousePlatoon)
	wrongs.ThrowWhenEmpty(ret, "仓库-排")

	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Datum(map[string]any{"warehouse_platoon": warehousePlatoon}).ToGinResponse())
}

// List 列表
func (receiver WarehousePlatoonCtrl) List(ctx *gin.Context) {
	var warehousePlatoons []*models.WarehousePlatoonMdl

	ctx.JSON(
		tools.NewCorrectWithGinContext("", ctx).
			DataForPager(
				models.WarehousePlatoonMdl{}.GetListByQuery(ctx),
				func(db *gorm.DB) map[string]any {
					db.Find(&warehousePlatoons)
					return map[string]any{"warehouse_platoons": warehousePlatoons}
				},
			).
			ToGinResponse(),
	)
}

// ListJdt jquery-dataTable后端分页数据
func (receiver WarehousePlatoonCtrl) ListJdt(ctx *gin.Context) {
	var warehousePlatoons []*models.WarehousePlatoonMdl

	ctx.JSON(
		tools.NewCorrectWithGinContext("", ctx).
			DataForJqueryDataTable(
				models.WarehousePlatoonMdl{}.GetListByQuery(ctx),
				func(db *gorm.DB) map[string]any {
					db.Find(&warehousePlatoons)
					return map[string]any{"warehouse_platoons": warehousePlatoons}
				},
			).
			ToGinResponse(),
	)
}

// NewWarehouseShelfCtrl 构造函数
func NewWarehouseShelfCtrl() *WarehouseShelfCtrl {
	return &WarehouseShelfCtrl{}
}

// Store 新建
func (WarehouseShelfCtrl) Store(ctx *gin.Context) {
	var (
		ret    *gorm.DB
		repeat *models.WarehouseShelfMdl
	)

	// 表单
	form := WarehouseShelfStoreForm{}.ShouldBind(ctx)

	// 查重
	ret = models.NewWarehouseShelfMdl().
		GetDb("").
		Where("name = ?", form.Name).
		First(&repeat)
	wrongs.ThrowWhenNotEmpty(ret, "仓库-架名称")

	// 新建
	warehouseShelf := &models.WarehouseShelfMdl{
		Name:                 form.Name,
		WarehousePlatoonUuid: form.warehousePlatoon.Uuid,
	}
	if ret = models.NewWarehouseShelfMdl().
		GetDb("").
		Create(&warehouseShelf); ret.Error != nil {
		wrongs.ThrowForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Created(map[string]any{"warehouse_shelf": warehouseShelf}).ToGinResponse())
}

// Destroy 删除
func (WarehouseShelfCtrl) Destroy(ctx *gin.Context) {
	var (
		ret            *gorm.DB
		warehouseShelf *models.WarehouseShelfMdl
	)

	// 查询
	ret = models.NewWarehouseShelfMdl().
		GetDb("").
		Where("uuid = ?", ctx.Param("uuid")).
		First(&warehouseShelf)
	wrongs.ThrowWhenEmpty(ret, "仓库-架")

	// 删除
	if ret := models.NewWarehouseShelfMdl().
		GetDb("").
		Where("uuid = ?", ctx.Param("uuid")).
		Delete(&warehouseShelf); ret.Error != nil {
		wrongs.ThrowForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Deleted().ToGinResponse())
}

// DestroyMany 批量删除
func (WarehouseShelfCtrl) DestroyMany(ctx *gin.Context) {
	form := WarehouseShelfDestroyManyForm{}.ShouldBind(ctx)
	if ret := models.NewWarehouseShelfMdl().GetDb("").Where("uuid in ?", form.Uuids).Delete(nil); ret.Error != nil {
		wrongs.ThrowForbidden("批量删除失败：%s", ret.Error.Error())
	}
	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Deleted().ToGinResponse())
}

// Update 编辑
func (WarehouseShelfCtrl) Update(ctx *gin.Context) {
	var (
		ret                    *gorm.DB
		warehouseShelf, repeat *models.WarehouseShelfMdl
	)

	// 表单
	form := WarehouseShelfStoreForm{}.ShouldBind(ctx)

	// 查重
	ret = models.NewWarehouseShelfMdl().
		GetDb("").
		Where("name = ? and uuid <> ?", form.Name, ctx.Param("uuid")).
		First(&repeat)
	wrongs.ThrowWhenNotEmpty(ret, "仓库-架名称")

	// 查询
	ret = models.NewWarehouseShelfMdl().
		GetDb("").
		Where("uuid = ?", ctx.Param("uuid")).
		First(&warehouseShelf)
	wrongs.ThrowWhenEmpty(ret, "仓库-架")

	// 编辑
	warehouseShelf.Name = form.Name
	warehouseShelf.WarehousePlatoonUuid = form.warehousePlatoon.Uuid
	if ret = models.NewWarehouseShelfMdl().
		GetDb("").
		Where("uuid = ?", ctx.Param("uuid")).
		Save(&warehouseShelf); ret.Error != nil {
		wrongs.ThrowForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Updated(map[string]any{"warehouse_shelf": warehouseShelf}).ToGinResponse())
}

// Detail 详情
func (WarehouseShelfCtrl) Detail(ctx *gin.Context) {
	var (
		ret            *gorm.DB
		warehouseShelf *models.WarehouseShelfMdl
	)
	ret = models.NewWarehouseShelfMdl().
		SetCtx(ctx).
		GetDbUseQuery("").
		Where("uuid = ?", ctx.Param("uuid")).
		First(&warehouseShelf)
	wrongs.ThrowWhenEmpty(ret, "仓库-架")

	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Datum(map[string]any{"warehouse_shelf": warehouseShelf}).ToGinResponse())
}

// List 列表
func (receiver WarehouseShelfCtrl) List(ctx *gin.Context) {
	var warehouseShelves []*models.WarehouseShelfMdl

	ctx.JSON(
		tools.NewCorrectWithGinContext("", ctx).
			DataForPager(
				models.WarehouseShelfMdl{}.GetListByQuery(ctx),
				func(db *gorm.DB) map[string]any {
					db.Find(&warehouseShelves)
					return map[string]any{"warehouse_shelves": warehouseShelves}
				},
			).
			ToGinResponse(),
	)
}

// ListJdt jquery-dataTable后端分页数据
func (receiver WarehouseShelfCtrl) ListJdt(ctx *gin.Context) {
	var warehouseShelves []*models.WarehouseShelfMdl

	ctx.JSON(
		tools.NewCorrectWithGinContext("", ctx).
			DataForJqueryDataTable(
				models.WarehouseShelfMdl{}.GetListByQuery(ctx),
				func(db *gorm.DB) map[string]any {
					db.Find(&warehouseShelves)
					return map[string]any{"warehouse_shelves": warehouseShelves}
				},
			).
			ToGinResponse(),
	)
}

// NewWarehouseTierCtrl 构造函数
func NewWarehouseTierCtrl() *WarehouseTierCtrl {
	return &WarehouseTierCtrl{}
}

// Store 新建
func (WarehouseTierCtrl) Store(ctx *gin.Context) {
	var (
		ret    *gorm.DB
		repeat *models.WarehouseTierMdl
	)

	// 表单
	form := WarehouseTierStoreForm{}.ShouldBind(ctx)

	// 查重
	ret = models.NewWarehouseTierMdl().
		GetDb("").
		Where("name = ?", form.Name).
		First(&repeat)
	wrongs.ThrowWhenNotEmpty(ret, "仓库-层名称")

	// 新建
	warehouseTier := &models.WarehouseTierMdl{
		Name:               form.Name,
		WarehouseShelfUuid: form.warehouseShelf.Uuid,
	}
	if ret = models.NewWarehouseTierMdl().
		GetDb("").
		Create(&warehouseTier); ret.Error != nil {
		wrongs.ThrowForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Created(map[string]any{"warehouse_tier": warehouseTier}).ToGinResponse())
}

// StoreMany 批量新建
func (WarehouseTierCtrl) StoreMany(ctx *gin.Context) {
	var (
		lastWarehouseTierNameInt int
		err                      error
		newWarehouseTiers        []*models.WarehouseTierMdl
	)
	form := WarehouseTierStoreManyForm{}.ShouldBind(ctx)
	lastWarehouseTier := models.WarehouseTierMdl{}.GetLastByWarehouseShelf(form.warehouseShelf)

	lastWarehouseTierNameInt, err = strconv.Atoi(lastWarehouseTier.Name)
	if err != nil {
		lastWarehouseTierNameInt = 0
	}
	for i := 0; i < int(form.Number); i++ {
		newWarehouseTiers = append(newWarehouseTiers, &models.WarehouseTierMdl{
			Name:               strconv.Itoa(lastWarehouseTierNameInt + i + 1),
			WarehouseShelfUuid: form.warehouseShelf.Uuid,
		})
	}

	models.NewWarehouseTierMdl().GetDb("").Create(&newWarehouseTiers)

	ctx.JSON(tools.NewCorrectWithGinContext(fmt.Sprintf("成功新建：%d层", len(newWarehouseTiers)), ctx).Created(map[string]any{"warehouse_tiers": newWarehouseTiers}).ToGinResponse())
}

// Destroy 删除
func (WarehouseTierCtrl) Destroy(ctx *gin.Context) {
	var (
		ret           *gorm.DB
		warehouseTier *models.WarehouseTierMdl
	)

	// 查询
	ret = models.NewWarehouseTierMdl().
		GetDb("").
		Where("uuid = ?", ctx.Param("uuid")).
		First(&warehouseTier)
	wrongs.ThrowWhenEmpty(ret, "仓库-层")

	// 删除
	if ret := models.NewWarehouseTierMdl().
		GetDb("").
		Where("uuid = ?", ctx.Param("uuid")).
		Delete(&warehouseTier); ret.Error != nil {
		wrongs.ThrowForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Deleted().ToGinResponse())
}

// DestroyMany 批量删除
func (WarehouseTierCtrl) DestroyMany(ctx *gin.Context) {
	form := WarehouseTierDestroyManyForm{}.ShouldBind(ctx)
	if ret := models.NewWarehouseTierMdl().GetDb("").Where("uuid in ?", form.Uuids).Delete(nil); ret.Error != nil {
		wrongs.ThrowForbidden("批量删除失败：%s", ret.Error.Error())
	}
	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Deleted().ToGinResponse())
}

// Update 编辑
func (WarehouseTierCtrl) Update(ctx *gin.Context) {
	var (
		ret                   *gorm.DB
		warehouseTier, repeat *models.WarehouseTierMdl
	)

	// 表单
	form := WarehouseTierStoreForm{}.ShouldBind(ctx)

	// 查重
	ret = models.NewWarehouseTierMdl().
		GetDb("").
		Where("name = ? and uuid <> ?", form.Name, ctx.Param("uuid")).
		First(&repeat)
	wrongs.ThrowWhenNotEmpty(ret, "仓库-层名称")

	// 查询
	ret = models.NewWarehouseTierMdl().
		GetDb("").
		Where("uuid = ?", ctx.Param("uuid")).
		First(&warehouseTier)
	wrongs.ThrowWhenEmpty(ret, "仓库-层")

	// 编辑
	warehouseTier.Name = form.Name
	warehouseTier.WarehouseShelfUuid = form.warehouseShelf.Uuid
	if ret = models.NewWarehouseTierMdl().
		GetDb("").
		Where("uuid = ?", ctx.Param("uuid")).
		Save(&warehouseTier); ret.Error != nil {
		wrongs.ThrowForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Updated(map[string]any{"warehouse_tier": warehouseTier}).ToGinResponse())
}

// Detail 详情
func (WarehouseTierCtrl) Detail(ctx *gin.Context) {
	var (
		ret           *gorm.DB
		warehouseTier *models.WarehouseTierMdl
	)
	ret = models.NewWarehouseTierMdl().
		SetCtx(ctx).
		GetDbUseQuery("").
		Where("uuid = ?", ctx.Param("uuid")).
		First(&warehouseTier)
	wrongs.ThrowWhenEmpty(ret, "仓库-层")

	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Datum(map[string]any{"warehouse_tier": warehouseTier}).ToGinResponse())
}

// List 列表
func (receiver WarehouseTierCtrl) List(ctx *gin.Context) {
	var warehouseTiers []*models.WarehouseTierMdl

	ctx.JSON(
		tools.NewCorrectWithGinContext("", ctx).
			DataForPager(
				models.WarehouseTierMdl{}.GetListByQuery(ctx),
				func(db *gorm.DB) map[string]any {
					db.Find(&warehouseTiers)
					return map[string]any{"warehouse_tiers": warehouseTiers}
				},
			).
			ToGinResponse(),
	)
}

// ListJdt jquery-dataTable后端分页数据
func (receiver WarehouseTierCtrl) ListJdt(ctx *gin.Context) {
	var warehouseTiers []*models.WarehouseTierMdl

	ctx.JSON(
		tools.NewCorrectWithGinContext("", ctx).
			DataForJqueryDataTable(
				models.WarehouseTierMdl{}.GetListByQuery(ctx),
				func(db *gorm.DB) map[string]any {
					db.Find(&warehouseTiers)
					return map[string]any{"warehouse_tiers": warehouseTiers}
				},
			).
			ToGinResponse(),
	)
}

// NewWarehouseCellCtrl 构造函数
func NewWarehouseCellCtrl() *WarehouseCellCtrl {
	return &WarehouseCellCtrl{}
}

// Store 新建
func (WarehouseCellCtrl) Store(ctx *gin.Context) {
	var (
		ret    *gorm.DB
		repeat *models.WarehouseCellMdl
	)

	// 表单
	form := WarehouseCellStoreForm{}.ShouldBind(ctx)

	// 查重
	ret = models.NewWarehouseCellMdl().
		GetDb("").
		Where("name = ?", form.Name).
		First(&repeat)
	wrongs.ThrowWhenNotEmpty(ret, "仓库-位名称")

	// 新建
	warehouseCell := &models.WarehouseCellMdl{
		Name:              form.Name,
		WarehouseTierUuid: form.warehouseTier.Uuid,
	}
	if ret = models.NewWarehouseCellMdl().
		GetDb("").
		Create(&warehouseCell); ret.Error != nil {
		wrongs.ThrowForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Created(map[string]any{"warehouse_cell": warehouseCell}).ToGinResponse())
}

// StoreMany 批量新建
func (WarehouseCellCtrl) StoreMany(ctx *gin.Context) {
	var (
		lastWarehouseCellNameInt int
		err                      error
		newWarehouseCells        []*models.WarehouseCellMdl
	)
	form := WarehouseCellStoreManyForm{}.ShouldBind(ctx)
	lastWarehouseCell := models.WarehouseCellMdl{}.GetLastByWarehouseTier(form.warehouseTier)
	lastWarehouseCellNameInt, err = strconv.Atoi(lastWarehouseCell.Name)
	if err != nil {
		lastWarehouseCellNameInt = 0
	}
	for i := 0; i < int(form.Number); i++ {
		newWarehouseCells = append(newWarehouseCells, &models.WarehouseCellMdl{
			Name:              strconv.Itoa(lastWarehouseCellNameInt + i + 1),
			WarehouseTierUuid: form.warehouseTier.Uuid,
		})
	}

	models.NewWarehouseCellMdl().GetDb("").Create(&newWarehouseCells)

	ctx.JSON(tools.NewCorrectWithGinContext(fmt.Sprintf("成功新建：%d位", len(newWarehouseCells)), ctx).Created(map[string]any{"warehouse_cells": newWarehouseCells}).ToGinResponse())
}

// Destroy 删除
func (WarehouseCellCtrl) Destroy(ctx *gin.Context) {
	var (
		ret           *gorm.DB
		warehouseCell *models.WarehouseCellMdl
	)

	// 查询
	ret = models.NewWarehouseCellMdl().
		GetDb("").
		Where("uuid = ?", ctx.Param("uuid")).
		First(&warehouseCell)
	wrongs.ThrowWhenEmpty(ret, "仓库-位")

	// 删除
	if ret := models.NewWarehouseCellMdl().
		GetDb("").
		Where("uuid = ?", ctx.Param("uuid")).
		Delete(&warehouseCell); ret.Error != nil {
		wrongs.ThrowForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Deleted().ToGinResponse())
}

// DestroyMany 批量删除
func (WarehouseCellCtrl) DestroyMany(ctx *gin.Context) {
	form := WarehouseCellDestroyManyForm{}.ShouldBind(ctx)
	if ret := models.NewWarehouseCellMdl().GetDb("").Where("uuid in ?", form.Uuids).Delete(nil); ret.Error != nil {
		wrongs.ThrowForbidden("批量删除失败：%s", ret.Error.Error())
	}
	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Deleted().ToGinResponse())
}

// Update 编辑
func (WarehouseCellCtrl) Update(ctx *gin.Context) {
	var (
		ret                   *gorm.DB
		warehouseCell, repeat *models.WarehouseCellMdl
	)

	// 表单
	form := WarehouseCellStoreForm{}.ShouldBind(ctx)

	// 查重
	ret = models.NewWarehouseCellMdl().
		GetDb("").
		Where("name = ? and uuid <> ?", form.Name, ctx.Param("uuid")).
		First(&repeat)
	wrongs.ThrowWhenNotEmpty(ret, "仓库-位名称")

	// 查询
	ret = models.NewWarehouseCellMdl().
		GetDb("").
		Where("uuid = ?", ctx.Param("uuid")).
		First(&warehouseCell)
	wrongs.ThrowWhenEmpty(ret, "仓库-位")

	// 编辑
	warehouseCell.Name = form.Name
	warehouseCell.WarehouseTierUuid = form.warehouseTier.Uuid
	if ret = models.NewWarehouseCellMdl().
		GetDb("").
		Where("uuid = ?", ctx.Param("uuid")).
		Save(&warehouseCell); ret.Error != nil {
		wrongs.ThrowForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Updated(map[string]any{"warehouse_cell": warehouseCell}).ToGinResponse())
}

// Detail 详情
func (WarehouseCellCtrl) Detail(ctx *gin.Context) {
	var (
		ret           *gorm.DB
		warehouseCell *models.WarehouseCellMdl
	)
	ret = models.NewWarehouseCellMdl().
		SetCtx(ctx).
		GetDbUseQuery("").
		Where("uuid = ?", ctx.Param("uuid")).
		First(&warehouseCell)
	wrongs.ThrowWhenEmpty(ret, "仓库-位")

	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Datum(map[string]any{"warehouse_cell": warehouseCell}).ToGinResponse())
}

// List 列表
func (receiver WarehouseCellCtrl) List(ctx *gin.Context) {
	var warehouseCells []*models.WarehouseCellMdl

	ctx.JSON(
		tools.NewCorrectWithGinContext("", ctx).
			DataForPager(
				models.WarehouseCellMdl{}.GetListByQuery(ctx),
				func(db *gorm.DB) map[string]any {
					db.Find(&warehouseCells)
					return map[string]any{"warehouse_cells": warehouseCells}
				},
			).
			ToGinResponse(),
	)
}

// ListJdt jquery-dataTable后端分页数据
func (receiver WarehouseCellCtrl) ListJdt(ctx *gin.Context) {
	var warehouseCells []*models.WarehouseCellMdl

	ctx.JSON(
		tools.NewCorrectWithGinContext("", ctx).
			DataForJqueryDataTable(
				models.WarehouseCellMdl{}.GetListByQuery(ctx),
				func(db *gorm.DB) map[string]any {
					db.Find(&warehouseCells)
					return map[string]any{"warehouse_cells": warehouseCells}
				},
			).
			ToGinResponse(),
	)
}
