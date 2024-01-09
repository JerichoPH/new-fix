package controllers

import (
	"fmt"
	"new-fix/models"
	"new-fix/utils"
	"new-fix/wrongs"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type (
	// RbacRoleCtrl 角色控制器
	RbacRoleCtrl struct{}
	// RbacRoleStoreForm 角色表单
	RbacRoleStoreForm struct {
		Name string `json:"name"`
	}
	// RbacRoleDestroyManyForm 批量删除角色表单
	RbacRoleDestroyManyForm struct {
		Uuids     []string `json:"uuids"`
		rbacRoles []*models.RbacRoleMdl
	}

	// RbacPermissionCtrl 权限控制器
	RbacPermissionCtrl struct{}
	// RbacPermissionStoreForm 权限表单
	RbacPermissionStoreForm struct {
		Name          string   `json:"name"`
		Description   string   `json:"description"`
		Uri           string   `json:"uri"`
		Icon          string   `json:"icon"`
		RbacRoleUuids []string `json:"rbac_role_uuids"`
		rbacRoles     []*models.RbacRoleMdl
	}
	// RbacPermissionDestroyMany 批量删除权限表单
	RbacPermissionDestroyMany struct {
		Uuids           []string `json:"uuids"`
		rbacPermissions []*models.RbacPermissionMdl
	}

	// RbacMenuCtrl 菜单控制器
	RbacMenuCtrl struct{}
	// RbacMenuStoreForm 菜单表单
	RbacMenuStoreForm struct {
		Name          string `json:"name"`
		SubTitle      string `json:"sub_title"`
		Description   string `json:"description"`
		Uri           string `json:"uri"`
		Icon          string `json:"icon"`
		PageRouteName string `json:"page_route_name"`
		ParentUuid    string `json:"parent_uuid"`
		parentMenu    *models.RbacMenuMdl
		RbacRoleUuids []string `json:"rbac_role_uuids"`
		rbacRoles     []*models.RbacRoleMdl
	}
	// RbacMenuDestroyMany 批量删除菜单表单
	RbacMenuDestroyMany struct {
		Uuids     []string `json:"uuids"`
		rbacMenus []*models.RbacMenuMdl
	}
)

// ShouldBind 角色表单绑定
func (receiver RbacRoleStoreForm) ShouldBind(ctx *gin.Context) RbacRoleStoreForm {
	var err error

	if ctx.Request.Method == "PUT" && ctx.Param("uuid") == "" {
		wrongs.ThrowValidate("角色编号不能为空")
	}
	if err = ctx.ShouldBind(&receiver); err != nil {
		if len(receiver.Name) == 0 {
			wrongs.ThrowValidate("角色名称不能为空")
		}
	}

	return receiver
}

// ShouldBind 批量删除角色表单绑定
func (receiver RbacRoleDestroyManyForm) ShouldBind(ctx *gin.Context) RbacRoleDestroyManyForm {
	if err := ctx.ShouldBind(&receiver); err != nil {
		wrongs.ThrowValidate("表单绑定失败：%s", err.Error())
	}

	if len(receiver.Uuids) == 0 {
		wrongs.ThrowValidate("角色编号不能为空")
	}

	models.NewRbacRoleMdl().GetDb("").Where("uuid in ?", receiver.Uuids).Find(&receiver.rbacRoles)

	return receiver
}

// ShouldBind 权限表单绑定
func (receiver RbacPermissionStoreForm) ShouldBind(ctx *gin.Context) RbacPermissionStoreForm {
	if ctx.Request.Method == "PUT" && ctx.Param("uuid") == "" {
		wrongs.ThrowValidate("权限编号不能为空")
	}
	if err := ctx.ShouldBind(&receiver); err != nil {
		wrongs.ThrowValidate(err.Error())
	}
	if receiver.Name == "" {
		wrongs.ThrowValidate("权限名称必填")
	}
	if receiver.Uri == "" {
		wrongs.ThrowValidate("权限路由必填")
	}
	if len(receiver.RbacRoleUuids) > 0 {
		models.NewRbacRoleMdl().GetDb("").Where("uuid in ?", receiver.RbacRoleUuids).Find(&receiver.rbacRoles)
	}

	return receiver
}

// ShouldBind 批量删除权限表单绑定
func (receiver RbacPermissionDestroyMany) ShouldBind(ctx *gin.Context) RbacPermissionDestroyMany {
	if err := ctx.ShouldBind(&receiver); err != nil {
		wrongs.ThrowValidate("表单绑定失败：%s", err.Error())
	}

	if len(receiver.Uuids) == 0 {
		wrongs.ThrowValidate("权限编号不能为空")
	}

	models.NewRbacPermissionMdl().GetDb("").Where("uuid in ?", receiver.Uuids).Find(&receiver.rbacPermissions)

	return receiver
}

// ShouldBind 菜单表单绑定
func (receiver RbacMenuStoreForm) ShouldBind(ctx *gin.Context) RbacMenuStoreForm {
	var ret *gorm.DB

	if ctx.Request.Method == "PUT" && ctx.Param("uuid") == "" {
		wrongs.ThrowValidate("菜单编号不能为空")
	}
	if err := ctx.ShouldBind(&receiver); err != nil {
		wrongs.ThrowValidate(err.Error())
	}
	if receiver.Name == "" {
		wrongs.ThrowValidate("菜单名称必填")
	}
	if receiver.ParentUuid != "" {
		ret = models.NewRbacMenuMdl().GetDb("").Where("uuid = ?", receiver.ParentUuid).First(&receiver.parentMenu)
		wrongs.ThrowWhenEmpty(ret, fmt.Sprintf("父级菜单（%s）", receiver.ParentUuid))
	}
	if len(receiver.RbacRoleUuids) > 0 {
		models.NewRbacRoleMdl().GetDb("").Where("uuid in ?", receiver.RbacRoleUuids).Find(&receiver.rbacRoles)
	}

	return receiver
}

// ShouldBind 批量删除菜单表单绑定
func (receiver RbacMenuDestroyMany) ShouldBind(ctx *gin.Context) RbacMenuDestroyMany {
	if err := ctx.ShouldBind(&receiver); err != nil {
		wrongs.ThrowValidate("表单绑定失败：%s", err.Error())
	}

	if len(receiver.Uuids) == 0 {
		wrongs.ThrowValidate("菜单编号不能为空")
	}

	models.NewRbacMenuMdl().GetDb("").Where("uuid in ?", receiver.Uuids).Find(&receiver.rbacMenus)

	return receiver
}

// NewRbacRoleCtrl 构造函数
func NewRbacRoleCtrl() *RbacRoleCtrl {
	return &RbacRoleCtrl{}
}

// Store 新建
func (RbacRoleCtrl) Store(ctx *gin.Context) {
	var (
		ret    *gorm.DB
		repeat models.RbacRoleMdl
	)

	// 表单
	form := (&RbacRoleStoreForm{}).ShouldBind(ctx)

	// 查重
	ret = models.NewRbacRoleMdl().
		GetDb("").
		Where("name = ?", form.Name).
		First(&repeat)
	wrongs.ThrowWhenNotEmpty(ret, "角色名称")

	// 新建
	rbacRole := &models.RbacRoleMdl{
		MySqlMdl: models.MySqlMdl{Uuid: uuid.NewV4().String()},
		Name:     form.Name,
	}
	if ret = models.NewRbacRoleMdl().
		GetDb("").
		Create(&rbacRole); ret.Error != nil {
		wrongs.ThrowForbidden(ret.Error.Error())
	}

	ctx.JSON(utils.NewCorrectWithGinContext("", ctx).Created(map[string]any{"rbac_role": rbacRole}).ToGinResponse())
}

// Destory 删除
func (RbacRoleCtrl) Destory(ctx *gin.Context) {
	var (
		ret      *gorm.DB
		rbacRole models.RbacRoleMdl
	)

	// 查询
	ret = models.NewRbacRoleMdl().
		GetDb("").
		Where("uuid = ?", ctx.Param("uuid")).
		First(&rbacRole)
	wrongs.ThrowWhenEmpty(ret, "角色")

	// 删除
	if ret := models.NewRbacRoleMdl().
		GetDb("").
		Where("uuid = ?", ctx.Param("uuid")).
		Delete(&rbacRole); ret.Error != nil {
		wrongs.ThrowForbidden(ret.Error.Error())
	}

	ctx.JSON(utils.NewCorrectWithGinContext("", ctx).Deleted().ToGinResponse())
}

// DestroyMany 批量删除
func (receiver RbacRoleCtrl) DestroyMany(ctx *gin.Context) {
	form := RbacRoleDestroyManyForm{}.ShouldBind(ctx)

	if len(form.rbacRoles) > 0 {
		models.NewRbacRoleMdl().GetDb("").Where("uuid = ?", form.Uuids).Delete(&form.rbacRoles)
	}

	ctx.JSON(utils.NewCorrectWithGinContext("", ctx).Deleted().ToGinResponse())
}

// Update 编辑
func (RbacRoleCtrl) Update(ctx *gin.Context) {
	var (
		ret              *gorm.DB
		rbacRole, repeat models.RbacRoleMdl
	)

	// 表单
	form := (&RbacRoleStoreForm{}).ShouldBind(ctx)

	// 查重
	ret = models.NewRbacRoleMdl().
		GetDb("").
		Where("name = ? and uuid <> ?", form.Name, ctx.Param("uuid")).
		First(&repeat)
	wrongs.ThrowWhenNotEmpty(ret, "角色名称")

	// 查询
	ret = models.NewRbacRoleMdl().
		GetDb("").
		Where("uuid = ?", ctx.Param("uuid")).
		First(&rbacRole)
	wrongs.ThrowWhenEmpty(ret, "角色")

	// 编辑
	rbacRole.Name = form.Name
	if ret = models.NewRbacRoleMdl().
		GetDb("").
		Where("uuid = ?", ctx.Param("uuid")).
		Save(&rbacRole); ret.Error != nil {
		wrongs.ThrowForbidden(ret.Error.Error())
	}

	ctx.JSON(utils.NewCorrectWithGinContext("", ctx).Updated(map[string]any{"rbac_role": rbacRole}).ToGinResponse())
}

// Detail 详情
func (RbacRoleCtrl) Detail(ctx *gin.Context) {
	var (
		ret      *gorm.DB
		rbacRole models.RbacRoleMdl
	)
	ret = models.NewRbacRoleMdl().
		SetCtx(ctx).
		GetDbUseQuery("").
		Where("uuid = ?", ctx.Param("uuid")).
		First(&rbacRole)
	wrongs.ThrowWhenEmpty(ret, "角色")

	ctx.JSON(utils.NewCorrectWithGinContext("", ctx).Datum(map[string]any{"rbac_role": rbacRole}).ToGinResponse())
}

// List 列表
func (receiver RbacRoleCtrl) List(ctx *gin.Context) {
	var rbacRoles []*models.RbacRoleMdl

	ctx.JSON(
		utils.NewCorrectWithGinContext("", ctx).
			DataForPager(
				models.RbacRoleMdl{}.GetListByQuery(ctx),
				func(db *gorm.DB) map[string]interface{} {
					db.Find(&rbacRoles)
					return map[string]any{"rbac_roles": rbacRoles}
				},
			).
			ToGinResponse(),
	)
}

// ListJdt jquery-dataTable后端分页数据
func (receiver RbacRoleCtrl) ListJdt(ctx *gin.Context) {
	var rbacRoles []*models.RbacRoleMdl

	ctx.JSON(
		utils.NewCorrectWithGinContext("", ctx).
			DataForJqueryDataTable(
				models.RbacRoleMdl{}.GetListByQuery(ctx),
				func(db *gorm.DB) map[string]interface{} {
					db.Find(&rbacRoles)
					return map[string]any{"rbac_roles": rbacRoles}
				},
			).
			ToGinResponse(),
	)
}

// NewRbacPermissionCtrl 构造函数
func NewRbacPermissionCtrl() *RbacPermissionCtrl {
	return &RbacPermissionCtrl{}
}

// Store 新建
func (RbacPermissionCtrl) Store(ctx *gin.Context) {
	var (
		ret    *gorm.DB
		repeat models.RbacPermissionMdl
	)

	// 表单
	form := RbacPermissionStoreForm{}.ShouldBind(ctx)

	// 查重
	ret = models.NewRbacPermissionMdl().
		GetDb("").
		Where("name = ?", form.Name).
		First(&repeat)
	wrongs.ThrowWhenNotEmpty(ret, "权限名称")

	// 新建
	rbacPermission := &models.RbacPermissionMdl{
		MySqlMdl:    models.MySqlMdl{Uuid: uuid.NewV4().String()},
		Name:        form.Name,
		Uri:         form.Uri,
		Description: &form.Description,
	}
	if ret = models.NewRbacPermissionMdl().
		GetDb("").
		Create(&rbacPermission); ret.Error != nil {
		wrongs.ThrowForbidden(ret.Error.Error())
	}

	// 绑定角色与权限
	models.PivotRbacRoleRbacPermissionMdl{}.BindRbacRoles(rbacPermission, form.rbacRoles)

	ctx.JSON(utils.NewCorrectWithGinContext("", ctx).Created(map[string]any{"rbac_permission": rbacPermission}).ToGinResponse())
}

// Destroy 删除
func (RbacPermissionCtrl) Destroy(ctx *gin.Context) {
	var (
		ret            *gorm.DB
		rbacPermission models.RbacPermissionMdl
	)

	// 查询
	ret = models.NewRbacPermissionMdl().
		GetDb("").
		Where("uuid = ?", ctx.Param("uuid")).
		First(&rbacPermission)
	wrongs.ThrowWhenEmpty(ret, "权限")

	// 删除
	if ret := models.NewRbacPermissionMdl().
		GetDb("").
		Where("uuid = ?", ctx.Param("uuid")).
		Delete(&rbacPermission); ret.Error != nil {
		wrongs.ThrowForbidden(ret.Error.Error())
	}

	ctx.JSON(utils.NewCorrectWithGinContext("", ctx).Deleted().ToGinResponse())
}

// DestroyMany 批量删除
func (receiver RbacPermissionCtrl) DestroyMany(ctx *gin.Context) {
	form := RbacPermissionDestroyMany{}.ShouldBind(ctx)

	if len(form.rbacPermissions) > 0 {
		models.NewRbacPermissionMdl().GetDb("").Where("uuid in ?", form.Uuids).Delete(&form.rbacPermissions)
	}

	ctx.JSON(utils.NewCorrectWithGinContext("", ctx).Deleted().ToGinResponse())
}

// Update 编辑
func (RbacPermissionCtrl) Update(ctx *gin.Context) {
	var (
		ret                    *gorm.DB
		rbacPermission, repeat *models.RbacPermissionMdl
	)

	// 表单
	form := RbacPermissionStoreForm{}.ShouldBind(ctx)

	// 查重
	ret = models.NewRbacPermissionMdl().
		GetDb("").
		Where("name = ? and uuid <> ?", form.Name, ctx.Param("uuid")).
		First(&repeat)
	wrongs.ThrowWhenNotEmpty(ret, fmt.Sprintf("权限名称 %s %s", ctx.Param("uuid"), repeat.Uuid))

	// 查询
	ret = models.NewRbacPermissionMdl().
		GetDb("").
		Where("uuid = ?", ctx.Param("uuid")).
		First(&rbacPermission)
	wrongs.ThrowWhenEmpty(ret, "权限")

	// 编辑
	rbacPermission.Name = form.Name
	rbacPermission.Description = &form.Description
	rbacPermission.Uri = form.Uri
	if ret = models.NewRbacPermissionMdl().
		GetDb("").
		Where("uuid = ?", ctx.Param("uuid")).
		Save(&rbacPermission); ret.Error != nil {
		wrongs.ThrowForbidden(ret.Error.Error())
	}

	// 绑定角色与权限
	models.PivotRbacRoleRbacPermissionMdl{}.BindRbacRoles(rbacPermission, form.rbacRoles)

	ctx.JSON(utils.NewCorrectWithGinContext("", ctx).Updated(map[string]any{"rbac_permission": rbacPermission}).ToGinResponse())
}

// Detail 详情
func (RbacPermissionCtrl) Detail(ctx *gin.Context) {
	var (
		ret            *gorm.DB
		rbacPermission models.RbacPermissionMdl
	)
	ret = models.NewRbacPermissionMdl().
		SetCtx(ctx).
		GetDbUseQuery("").
		Where("uuid = ?", ctx.Param("uuid")).
		First(&rbacPermission)
	wrongs.ThrowWhenEmpty(ret, "权限")

	ctx.JSON(utils.NewCorrectWithGinContext("", ctx).Datum(map[string]any{"rbac_permission": rbacPermission}).ToGinResponse())
}

// List 列表
func (receiver RbacPermissionCtrl) List(ctx *gin.Context) {
	var rbacPermissions []*models.RbacPermissionMdl

	ctx.JSON(
		utils.NewCorrectWithGinContext("", ctx).
			DataForPager(
				models.RbacPermissionMdl{}.GetListByQuery(ctx),
				func(db *gorm.DB) map[string]interface{} {
					db.Find(&rbacPermissions)
					return map[string]any{"rbac_permissions": rbacPermissions}
				},
			).
			ToGinResponse(),
	)
}

// ListJdt jquery-dataTable后端分页数据
func (receiver RbacPermissionCtrl) ListJdt(ctx *gin.Context) {
	var rbacPermissions []*models.RbacPermissionMdl

	ctx.JSON(
		utils.NewCorrectWithGinContext("", ctx).
			DataForJqueryDataTable(
				models.RbacPermissionMdl{}.GetListByQuery(ctx),
				func(db *gorm.DB) map[string]interface{} {
					db.Find(&rbacPermissions)
					return map[string]any{"rbac_permissions": rbacPermissions}
				},
			).
			ToGinResponse(),
	)
}

// NewRbacMenuCtrl 构造函数
func NewRbacMenuCtrl() *RbacMenuCtrl {
	return &RbacMenuCtrl{}
}

// Store 新建
func (RbacMenuCtrl) Store(ctx *gin.Context) {
	var (
		ret    *gorm.DB
		repeat *models.RbacMenuMdl
	)

	// 表单
	form := RbacMenuStoreForm{}.ShouldBind(ctx)

	// 查重
	ret = models.NewRbacMenuMdl().
		GetDb("").
		Where("name = ?", form.Name).
		Where("parent_uuid = ?", form.ParentUuid).
		First(&repeat)
	wrongs.ThrowWhenNotEmpty(ret, "菜单名称")

	// 新建
	rbacMenu := &models.RbacMenuMdl{
		MySqlMdl:      models.MySqlMdl{Uuid: uuid.NewV4().String()},
		Name:          form.Name,
		SubTitle:      form.SubTitle,
		Description:   &form.Description,
		Uri:           form.Uri,
		Icon:          form.Icon,
		PageRouteName: form.PageRouteName,
		ParentUuid:    form.ParentUuid,
	}
	if ret = models.NewRbacMenuMdl().
		GetDb("").
		Create(&rbacMenu); ret.Error != nil {
		wrongs.ThrowForbidden(ret.Error.Error())
	}

	// 绑定角色与菜单
	models.PivotRbacRoleRbacMenuMdl{}.BindRbacRoles(rbacMenu, form.rbacRoles)

	ctx.JSON(utils.NewCorrectWithGinContext("", ctx).Created(map[string]any{"rbac_menu": rbacMenu}).ToGinResponse())
}

// Destroy 删除
func (RbacMenuCtrl) Destroy(ctx *gin.Context) {
	var (
		ret      *gorm.DB
		rbacMenu *models.RbacMenuMdl
		subs     = make(map[string]map[string]string)
	)

	// 查询
	ret = models.NewRbacMenuMdl().
		GetDb("").
		Where("uuid = ?", ctx.Param("uuid")).
		First(&rbacMenu)
	wrongs.ThrowWhenEmpty(ret, "菜单")
	// 查询该菜单下是否存在子集
	subs = rbacMenu.GetSubUuidsByParentUuid(rbacMenu.Uuid)
	if len(subs) > 0 {
		wrongs.ThrowForbidden("该菜单下存在子集，请先删除子集")
	}

	// 删除
	if ret := models.NewRbacMenuMdl().
		GetDb("").
		Where("uuid = ?", ctx.Param("uuid")).
		Delete(&rbacMenu); ret.Error != nil {
		wrongs.ThrowForbidden(ret.Error.Error())
	}

	ctx.JSON(utils.NewCorrectWithGinContext("", ctx).Deleted().ToGinResponse())
}

// DestroyMany 批量删除
func (receiver RbacMenuCtrl) DestroyMany(ctx *gin.Context) {
	form := RbacMenuDestroyMany{}.ShouldBind(ctx)

	if len(form.rbacMenus) > 0 {
		models.NewRbacMenuMdl().GetDb("").Where("uuid in ?", form.Uuids).Delete(&form.rbacMenus)
	}

	ctx.JSON(utils.NewCorrectWithGinContext("", ctx).Deleted().ToGinResponse())
}

// Update 编辑
func (RbacMenuCtrl) Update(ctx *gin.Context) {
	var (
		ret              *gorm.DB
		rbacMenu, repeat *models.RbacMenuMdl
		subs             map[string]map[string]string
	)

	// 表单
	form := RbacMenuStoreForm{}.ShouldBind(ctx)

	// 查重
	ret = models.NewRbacMenuMdl().
		GetDb("").
		Where("name = ? and uuid <> ?", form.Name, ctx.Param("uuid")).
		Where("uuid <> ?", ctx.Param("uuid")).
		Where("parent_uuid <> ?", form.ParentUuid).Debug().
		First(&repeat)
	wrongs.ThrowWhenNotEmpty(ret, "菜单名称")

	// 查询
	ret = models.NewRbacMenuMdl().
		GetDb("").
		Where("uuid = ?", ctx.Param("uuid")).
		First(&rbacMenu)
	wrongs.ThrowWhenEmpty(ret, "菜单")
	if form.ParentUuid != "" {
		if rbacMenu.ParentUuid == form.ParentUuid {
			wrongs.ThrowValidate("父级菜单不能是自己")
		}
	}
	// 查询所有子菜单
	subs = models.RbacMenuMdl{}.GetSubUuidsByParentUuid(rbacMenu.Uuid)
	if len(subs) > 0 && form.ParentUuid != "" {
		if sub, exist := subs[form.ParentUuid]; exist {
			wrongs.ThrowValidate("「%s」是「%s」的父级，不能将子集绑定为子集的父级", form.Name, sub["name"])
		}
	}

	// 编辑
	rbacMenu.Name = form.Name
	rbacMenu.SubTitle = form.SubTitle
	rbacMenu.Description = &form.Description
	rbacMenu.Icon = form.Icon
	rbacMenu.Uri = form.Uri
	rbacMenu.PageRouteName = form.PageRouteName
	rbacMenu.ParentUuid = form.ParentUuid
	if ret = models.NewRbacMenuMdl().
		GetDb("").
		Where("uuid = ?", ctx.Param("uuid")).
		Save(&rbacMenu); ret.Error != nil {
		wrongs.ThrowForbidden(ret.Error.Error())
	}

	// 绑定角色与菜单
	models.PivotRbacRoleRbacMenuMdl{}.BindRbacRoles(rbacMenu, form.rbacRoles)

	ctx.JSON(utils.NewCorrectWithGinContext("", ctx).Updated(map[string]any{"rbac_menu": rbacMenu}).ToGinResponse())
}

// Detail 详情
func (RbacMenuCtrl) Detail(ctx *gin.Context) {
	var (
		ret      *gorm.DB
		rbacMenu models.RbacMenuMdl
	)
	ret = models.NewRbacMenuMdl().
		SetCtx(ctx).
		GetDbUseQuery("").
		Where("uuid = ?", ctx.Param("uuid")).
		First(&rbacMenu)
	wrongs.ThrowWhenEmpty(ret, "菜单")

	ctx.JSON(utils.NewCorrectWithGinContext("", ctx).Datum(map[string]any{"rbac_menu": rbacMenu}).ToGinResponse())
}

// List 列表
func (receiver RbacMenuCtrl) List(ctx *gin.Context) {
	var rbacMenus []*models.RbacMenuMdl

	ctx.JSON(
		utils.NewCorrectWithGinContext("", ctx).
			DataForPager(
				models.RbacMenuMdl{}.GetListByQuery(ctx),
				func(db *gorm.DB) map[string]interface{} {
					db.Find(&rbacMenus)
					return map[string]any{"rbac_menus": rbacMenus}
				},
			).
			ToGinResponse(),
	)
}

// ListJdt jquery-dataTable后端分页数据
func (receiver RbacMenuCtrl) ListJdt(ctx *gin.Context) {
	var rbacMenus []*models.RbacMenuMdl

	ctx.JSON(
		utils.NewCorrectWithGinContext("", ctx).
			DataForJqueryDataTable(
				models.RbacMenuMdl{}.GetListByQuery(ctx),
				func(db *gorm.DB) map[string]interface{} {
					db.Find(&rbacMenus)
					return map[string]any{"rbac_menus": rbacMenus}
				},
			).
			ToGinResponse(),
	)
}
