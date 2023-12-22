package commands

import (
	"fmt"
	"new-fix/database"
	"new-fix/models"
	"new-fix/tools"

	"gorm.io/gorm"
)

// UpgradeCmd 程序升级
type UpgradeCmd struct{}

// NewUpgradeCmd 构造函数
func NewUpgradeCmd() UpgradeCmd {
	return UpgradeCmd{}
}

// 初始化项目
func (UpgradeCmd) init() []string {
	var (
		rbacPermissions = make([]*models.RbacPermissionMdl, 0)
		rbacMenus       = make([]*models.RbacMenuMdl, 0)
		rbacMenu        *models.RbacMenuMdl
		ret             *gorm.DB
	)

	std := tools.NewStdoutHelper("初始化项目")

	std.EchoLineDebug("初始化管理员")
	database.ExecSql("truncate table accounts")
	std.EchoLineSuccess("截断管理员表成功")
	// 创建管理员
	models.
		NewAccountMdl().
		GetDb("").
		Create(&models.AccountMdl{
			Username: "admin",
			Nickname: "管理员",
			Password: tools.GeneratePassword("123123"),
			BeAdmin:  true,
		})

	database.ExecSql("truncate table rbac_roles")
	std.EchoLineSuccess("截断角色表成功")

	std.EchoLineDebug("初始化权限")
	database.ExecSql("truncate table rbac_permissions")
	std.EchoLineSuccess("截断权限表成功")
	// 创建权限
	models.
		NewRbacPermissionMdl().
		GetDb("").
		Create([]*models.RbacPermissionMdl{
			{Name: "用户列表", Method: "get", Uri: "/api/account"},
			{Name: "用户详情", Method: "get", Uri: "/api/account/:uuid"},
			{Name: "新建用户", Method: "post", Uri: "/api/account"},
			{Name: "编辑用户", Method: "put", Uri: "/api/account/:uuid"},
			{Name: "删除用户", Method: "delete", Uri: "/api/account/:uuid"},
			{Name: "角色列表", Method: "get", Uri: "/api/account"},
			{Name: "角色详情", Method: "get", Uri: "/api/account/:uuid"},
			{Name: "新建角色", Method: "post", Uri: "/api/account"},
			{Name: "编辑角色", Method: "put", Uri: "/api/account/:uuid"},
			{Name: "删除角色", Method: "delete", Uri: "/api/account/:uuid"},
			{Name: "权限列表", Method: "get", Uri: "/api/account"},
			{Name: "权限详情", Method: "get", Uri: "/api/account/:uuid"},
			{Name: "新建权限", Method: "post", Uri: "/api/account"},
			{Name: "编辑权限", Method: "put", Uri: "/api/account/:uuid"},
			{Name: "删除权限", Method: "delete", Uri: "/api/account/:uuid"},
			{Name: "菜单列表", Method: "get", Uri: "/api/account"},
			{Name: "菜单详情", Method: "get", Uri: "/api/account/:uuid"},
			{Name: "新建菜单", Method: "post", Uri: "/api/account"},
			{Name: "编辑菜单", Method: "put", Uri: "/api/account/:uuid"},
			{Name: "删除菜单", Method: "delete", Uri: "/api/account/:uuid"},
			{Name: "器材列表", Method: "get", Uri: "/api/equipment"},
			{Name: "器材种类列表", Method: "get", Uri: "/api/equipmentKind/category"},
			{Name: "器材种类详情", Method: "get", Uri: "/api/equipmentKind/category/:uuid"},
			{Name: "新建器材种类", Method: "post", Uri: "/api/equipmentKind/category"},
			{Name: "编辑器材种类", Method: "put", Uri: "/api/equipmentKind/category/:uuid"},
			{Name: "删除器材种类", Method: "delete", Uri: "/api/equipmentKind/category/:uuid"},
			{Name: "器材类型列表", Method: "get", Uri: "/api/equipmentKind/type"},
			{Name: "器材类型详情", Method: "get", Uri: "/api/equipmentKind/type/:uuid"},
			{Name: "新建器材类型", Method: "post", Uri: "/api/equipmentKind/type"},
			{Name: "编辑器材类型", Method: "put", Uri: "/api/equipmentKind/type/:uuid"},
			{Name: "删除器材类型", Method: "delete", Uri: "/api/equipmentKind/type/:uuid"},
			{Name: "器材型号列表", Method: "get", Uri: "/api/equipmentKind/model"},
			{Name: "器材型号详情", Method: "get", Uri: "/api/equipmentKind/model/:uuid"},
			{Name: "新建器材型号", Method: "post", Uri: "/api/equipmentKind/model"},
			{Name: "编辑器材型号", Method: "put", Uri: "/api/equipmentKind/model/:uuid"},
			{Name: "删除器材型号", Method: "delete", Uri: "/api/equipmentKind/model/:uuid"},
			{Name: "路局列表", Method: "get", Uri: "/api/organization/railway"},
			{Name: "路局详情", Method: "get", Uri: "/api/organization/railway/:uuid"},
			{Name: "新建路局", Method: "post", Uri: "/api/organization/railway"},
			{Name: "编辑路局", Method: "put", Uri: "/api/organization/railway/:uuid"},
			{Name: "删除路局", Method: "delete", Uri: "/api/organization/railway/:uuid"},
			{Name: "站段列表", Method: "get", Uri: "/api/organization/paragraph"},
			{Name: "站段详情", Method: "get", Uri: "/api/organization/paragraph/:uuid"},
			{Name: "新建站段", Method: "post", Uri: "/api/organization/paragraph"},
			{Name: "编辑站段", Method: "put", Uri: "/api/organization/paragraph/:uuid"},
			{Name: "删除站段", Method: "delete", Uri: "/api/organization/paragraph/:uuid"},
			{Name: "车间列表", Method: "get", Uri: "/api/organization/workshop"},
			{Name: "车间详情", Method: "get", Uri: "/api/organization/workshop/:uuid"},
			{Name: "新建车间", Method: "post", Uri: "/api/organization/workshop"},
			{Name: "编辑车间", Method: "put", Uri: "/api/organization/workshop/:uuid"},
			{Name: "删除车间", Method: "delete", Uri: "/api/organization/workshop/:uuid"},
			{Name: "站场列表", Method: "get", Uri: "/api/organization/station"},
			{Name: "站场详情", Method: "get", Uri: "/api/organization/station/:uuid"},
			{Name: "新建站场", Method: "post", Uri: "/api/organization/station"},
			{Name: "编辑站场", Method: "put", Uri: "/api/organization/station/:uuid"},
			{Name: "删除站场", Method: "delete", Uri: "/api/organization/station/:uuid"},
			{Name: "道口列表", Method: "get", Uri: "/api/organization/crossroad"},
			{Name: "道口详情", Method: "get", Uri: "/api/organization/crossroad/:uuid"},
			{Name: "新建道口", Method: "post", Uri: "/api/organization/crossroad"},
			{Name: "编辑道口", Method: "put", Uri: "/api/organization/crossroad/:uuid"},
			{Name: "删除道口", Method: "delete", Uri: "/api/organization/crossroad/:uuid"},
			{Name: "中心列表", Method: "get", Uri: "/api/organization/center"},
			{Name: "中心详情", Method: "get", Uri: "/api/organization/center/:uuid"},
			{Name: "新建中心", Method: "post", Uri: "/api/organization/center"},
			{Name: "编辑中心", Method: "put", Uri: "/api/organization/center/:uuid"},
			{Name: "删除中心", Method: "delete", Uri: "/api/organization/center/:uuid"},
			{Name: "工区列表", Method: "get", Uri: "/api/organization/workArea"},
			{Name: "工区详情", Method: "get", Uri: "/api/organization/workArea/:uuid"},
			{Name: "新建工区", Method: "post", Uri: "/api/organization/workArea"},
			{Name: "编辑工区", Method: "put", Uri: "/api/organization/workArea/:uuid"},
			{Name: "删除工区", Method: "delete", Uri: "/api/organization/workArea/:uuid"},
			{Name: "线别列表", Method: "get", Uri: "/api/organization/line"},
			{Name: "线别详情", Method: "get", Uri: "/api/organization/line/:uuid"},
			{Name: "新建线别", Method: "post", Uri: "/api/organization/line"},
			{Name: "编辑线别", Method: "put", Uri: "/api/organization/line/:uuid"},
			{Name: "删除线别", Method: "delete", Uri: "/api/organization/line/:uuid"},
		})
	if ret = models.NewRbacPermissionMdl().GetDb("").Create(&rbacPermissions); ret.Error != nil {
		std.EchoLineWrong(fmt.Sprintf("错误：%v", ret.Error.Error()))
	}
	std.EchoLineSuccess("成功")

	std.EchoLineDebug("初始化菜单")
	database.ExecSql("truncate table rbac_menus")
	std.EchoLineSuccess("截断菜单表成功")
	rbacMenu = &models.RbacMenuMdl{Name: "系统管理", Icon: "fa fa-cogs"}
	models.NewRbacMenuMdl().GetDb("").Create(rbacMenu)
	models.NewRbacMenuMdl().GetDb("").Create([]*models.RbacMenuMdl{
		{Name: "用户列表", Uri: "/account", Icon: "fa fa-list", PageRouteName: "acocunt:index", ParentUuid: rbacMenu.Uuid},
		{Name: "角色列表", Uri: "/rbac/role", Icon: "fa fa-list", PageRouteName: "rbacRole:index", ParentUuid: rbacMenu.Uuid},
		{Name: "权限列表", Uri: "/rbac/permission", Icon: "fa fa-list", PageRouteName: "rbacPermission:index", ParentUuid: rbacMenu.Uuid},
		{Name: "菜单列表", Uri: "/rbac/menu", Icon: "fa fa-list", PageRouteName: "rbacMenu:index", ParentUuid: rbacMenu.Uuid},
	})
	if ret = models.NewRbacMenuMdl().GetDb("").Create(&rbacMenus); ret.Error != nil {
		std.EchoLineWrong(fmt.Sprintf("错误：%v", ret.Error.Error()))
	}
	std.EchoLineSuccess("成功")

	return []string{}
}

// Do 执行命令
func (receiver UpgradeCmd) Do(params []string) []string {
	switch params[0] {
	case "init":
		return receiver.init()
	}

	return []string{"执行完成"}
}
