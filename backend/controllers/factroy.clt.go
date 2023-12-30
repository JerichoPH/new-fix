package controllers

import (
	"new-fix/models"
	"new-fix/tools"
	"new-fix/wrongs"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type (
	// 厂家控制器
	FactoryCtrl struct{}
	// 厂家表单
	FactoryStoreForm struct {
		UniqueCode string `json:"unique_code"`
		Name       string `json:"name"`
	}
	// 厂家批量删除表单
	FactoryDestroyForm struct {
		Uuids []string `json:"uuids"`
	}
)

// ShouldBind 表单绑定（厂家）
func (receiver FactoryStoreForm) ShouldBind(ctx *gin.Context) FactoryStoreForm {
	if err := ctx.ShouldBind(&receiver); err != nil {
		wrongs.ThrowValidate("表单绑定失败：%s", err.Error())
	}
	if receiver.UniqueCode == "" {
		wrongs.ThrowValidate("厂家代码必填")
	} else {
		if len(receiver.UniqueCode) != 5 {
			wrongs.ThrowValidate("厂家代码长度必须为5位")
		}
	}
	if receiver.Name == "" {
		wrongs.ThrowValidate("厂家名称必填")
	}
	return receiver
}

// ShouldBind 表单绑定（厂家批量删除）
func (receiver FactoryDestroyForm) ShouldBind(ctx *gin.Context) FactoryDestroyForm {
	if err := ctx.ShouldBind(&receiver); err != nil {
		wrongs.ThrowValidate("表单绑定失败：%s", err.Error())
	}
	if len(receiver.Uuids) == 0 {
		wrongs.ThrowValidate("uuid必填")
	}
	return receiver
}

// NewFactoryCtrl 构造函数
func NewFactoryCtrl() *FactoryCtrl {
	return &FactoryCtrl{}
}

// Store 新建
func (FactoryCtrl) Store(ctx *gin.Context) {
	var (
		ret    *gorm.DB
		repeat *models.FactoryMdl
	)

	// 表单
	form := FactoryStoreForm{}.ShouldBind(ctx)

	// 查重
	ret = models.NewFactoryMdl().
		GetDb("").
		Where("unique_code = ?", form.UniqueCode).
		First(&repeat)
	wrongs.ThrowWhenNotEmpty(ret, "厂家代码")
	ret = models.NewFactoryMdl().
		GetDb("").
		Where("name = ?", form.Name).
		First(&repeat)
	wrongs.ThrowWhenNotEmpty(ret, "厂家名称")

	// 新建
	factory := &models.FactoryMdl{
		UniqueCode: form.UniqueCode,
		Name:       form.Name,
	}
	if ret = models.NewFactoryMdl().
		GetDb("").
		Create(&factory); ret.Error != nil {
		wrongs.ThrowForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Created(map[string]any{"factory": factory}).ToGinResponse())
}

// Destroy 删除
func (FactoryCtrl) Destroy(ctx *gin.Context) {
	var (
		ret     *gorm.DB
		factory *models.FactoryMdl
	)

	// 查询
	ret = models.NewFactoryMdl().
		GetDb("").
		Where("uuid = ?", ctx.Param("uuid")).
		First(&factory)
	wrongs.ThrowWhenEmpty(ret, "厂家")

	// 删除
	if ret := models.NewFactoryMdl().
		GetDb("").
		Where("uuid = ?", ctx.Param("uuid")).
		Delete(&factory); ret.Error != nil {
		wrongs.ThrowForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Deleted().ToGinResponse())
}

// DestroyMany 批量删除
func (FactoryCtrl) DestroyMany(ctx *gin.Context) {
	// 表单
	form := FactoryDestroyForm{}.ShouldBind(ctx)
	models.NewFactoryMdl().GetDb("").Where("uuid in ?", form.Uuids).Delete(nil)
}

// Update 编辑
func (FactoryCtrl) Update(ctx *gin.Context) {
	var (
		ret             *gorm.DB
		factory, repeat *models.FactoryMdl
	)

	// 表单
	form := FactoryStoreForm{}.ShouldBind(ctx)

	// 查重
	ret = models.NewFactoryMdl().
		GetDb("").
		Where("unique_code = ? and uuid <> ?", form.UniqueCode, ctx.Param("uuid")).
		First(&repeat)
	wrongs.ThrowWhenNotEmpty(ret, "厂家代码")
	ret = models.NewFactoryMdl().
		GetDb("").
		Where("name = ? and uuid <> ?", form.Name, ctx.Param("uuid")).
		First(&repeat)
	wrongs.ThrowWhenNotEmpty(ret, "厂家名称")

	// 查询
	ret = models.NewFactoryMdl().
		GetDb("").
		Where("uuid = ?", ctx.Param("uuid")).
		First(&factory)
	wrongs.ThrowWhenEmpty(ret, "厂家")

	// 编辑
	factory.UniqueCode = form.UniqueCode
	factory.Name = form.Name
	if ret = models.NewFactoryMdl().
		GetDb("").
		Where("uuid = ?", ctx.Param("uuid")).
		Save(&factory); ret.Error != nil {
		wrongs.ThrowForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Updated(map[string]any{"factory": factory}).ToGinResponse())
}

// Detail 详情
func (FactoryCtrl) Detail(ctx *gin.Context) {
	var (
		ret     *gorm.DB
		factory *models.FactoryMdl
	)
	ret = models.NewFactoryMdl().
		SetCtx(ctx).
		GetDbUseQuery("").
		Where("uuid = ?", ctx.Param("uuid")).
		First(&factory)
	wrongs.ThrowWhenEmpty(ret, "厂家")

	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Datum(map[string]any{"factory": factory}).ToGinResponse())
}

// List 列表
func (receiver FactoryCtrl) List(ctx *gin.Context) {
	var factories []*models.FactoryMdl

	ctx.JSON(
		tools.NewCorrectWithGinContext("", ctx).
			DataForPager(
				models.FactoryMdl{}.GetListByQuery(ctx),
				func(db *gorm.DB) map[string]any {
					db.Find(&factories)
					return map[string]any{"factories": factories}
				},
			).
			ToGinResponse(),
	)
}

// ListJdt jquery-dataTable后端分页数据
func (receiver FactoryCtrl) ListJdt(ctx *gin.Context) {
	var factories []*models.FactoryMdl

	ctx.JSON(
		tools.NewCorrectWithGinContext("", ctx).
			DataForJqueryDataTable(
				models.FactoryMdl{}.GetListByQuery(ctx),
				func(db *gorm.DB) map[string]any {
					db.Find(&factories)
					return map[string]any{"factories": factories}
				},
			).
			ToGinResponse(),
	)
}
