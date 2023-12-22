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
