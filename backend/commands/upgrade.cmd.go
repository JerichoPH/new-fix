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
			{Name: "故障类型列表", Method: "get", Uri: "/api/breakdown/type"},
			{Name: "故障类型详情", Method: "get", Uri: "/api/breakdown/type/:uuid"},
			{Name: "新建故障类型", Method: "post", Uri: "/api/breakdown/type"},
			{Name: "编辑故障类型", Method: "put", Uri: "/api/breakdown/type/:uuid"},
			{Name: "删除故障类型", Method: "delete", Uri: "/api/breakdown/type/:uuid"},
			{Name: "故障日志列表", Method: "get", Uri: "/api/breakdown/log"},
			{Name: "故障日志详情", Method: "get", Uri: "/api/breakdown/log/:uuid"},
			{Name: "厂家列表", Method: "get", Uri: "/api/factory"},
			{Name: "厂家详情", Method: "get", Uri: "/api/factory/:uuid"},
			{Name: "新建厂家", Method: "post", Uri: "/api/factory"},
			{Name: "编辑厂家", Method: "put", Uri: "/api/factory/:uuid"},
			{Name: "删除厂家", Method: "delete", Uri: "/api/factory/:uuid"},
			{Name: "来源类型列表", Method: "get", Uri: "/api/source/type"},
			{Name: "来源类型详情", Method: "get", Uri: "/api/source/type/:uuid"},
			{Name: "新建来源类型", Method: "post", Uri: "/api/source/type"},
			{Name: "编辑来源类型", Method: "put", Uri: "/api/source/type/:uuid"},
			{Name: "删除来源类型", Method: "delete", Uri: "/api/source/type/:uuid"},
			{Name: "来源项目列表", Method: "get", Uri: "/api/source/project"},
			{Name: "来源项目详情", Method: "get", Uri: "/api/source/project/:uuid"},
			{Name: "新建来源项目", Method: "post", Uri: "/api/source/project"},
			{Name: "编辑来源项目", Method: "put", Uri: "/api/source/project/:uuid"},
			{Name: "删除来源项目", Method: "delete", Uri: "/api/source/project/:uuid"},
			{Name: "仓库-库房列表", Method: "get", Uri: "/api/warehouse/storehouse"},
			{Name: "仓库-库房详情", Method: "get", Uri: "/api/warehouse/storehouse/:uuid"},
			{Name: "新建仓库-库房", Method: "post", Uri: "/api/warehouse/storehouse"},
			{Name: "编辑仓库-库房", Method: "put", Uri: "/api/warehouse/storehouse/:uuid"},
			{Name: "删除仓库-库房", Method: "delete", Uri: "/api/warehouse/storehouse/:uuid"},
			{Name: "仓库-库区列表", Method: "get", Uri: "/api/warehouse/area"},
			{Name: "仓库-库区详情", Method: "get", Uri: "/api/warehouse/area/:uuid"},
			{Name: "新建仓库-库区", Method: "post", Uri: "/api/warehouse/area"},
			{Name: "编辑仓库-库区", Method: "put", Uri: "/api/warehouse/area/:uuid"},
			{Name: "删除仓库-库区", Method: "delete", Uri: "/api/warehouse/area/:uuid"},
			{Name: "仓库-排列表", Method: "get", Uri: "/api/warehouse/platoon"},
			{Name: "仓库-排详情", Method: "get", Uri: "/api/warehouse/platoon/:uuid"},
			{Name: "新建仓库-排", Method: "post", Uri: "/api/warehouse/platoon"},
			{Name: "编辑仓库-排", Method: "put", Uri: "/api/warehouse/platoon/:uuid"},
			{Name: "删除仓库-排", Method: "delete", Uri: "/api/warehouse/platoon/:uuid"},
			{Name: "仓库-架列表", Method: "get", Uri: "/api/warehouse/shelf"},
			{Name: "仓库-架详情", Method: "get", Uri: "/api/warehouse/shelf/:uuid"},
			{Name: "新建仓库-架", Method: "post", Uri: "/api/warehouse/shelf"},
			{Name: "编辑仓库-架", Method: "put", Uri: "/api/warehouse/shelf/:uuid"},
			{Name: "仓库-架删除", Method: "delete", Uri: "/api/warehouse/shelf/:uuid"},
			{Name: "仓库-层列表", Method: "get", Uri: "/api/warehouse/tier"},
			{Name: "详情仓库-层", Method: "get", Uri: "/api/warehouse/tier/:uuid"},
			{Name: "新建仓库-层", Method: "post", Uri: "/api/warehouse/tier"},
			{Name: "编辑仓库-层", Method: "put", Uri: "/api/warehouse/tier/:uuid"},
			{Name: "删除仓库-层", Method: "delete", Uri: "/api/warehouse/tier/:uuid"},
			{Name: "仓库-位列表", Method: "get", Uri: "/api/warehouse/cell"},
			{Name: "仓库-位详情", Method: "get", Uri: "/api/warehouse/cell/:uuid"},
			{Name: "新建仓库-位", Method: "post", Uri: "/api/warehouse/cell"},
			{Name: "编辑仓库-位", Method: "put", Uri: "/api/warehouse/cell/:uuid"},
			{Name: "删除仓库-位", Method: "delete", Uri: "/api/warehouse/cell/:uuid"},
			{Name: "室内上道位置-机房类型列表", Method: "get", Uri: "/api/install/indoorRoomType"},
			{Name: "室内上道位置-机房类型详情", Method: "get", Uri: "/api/install/indoorRoomType/:uuid"},
			{Name: "新建室内上道位置-机房类型", Method: "post", Uri: "/api/install/indoorRoomType"},
			{Name: "编辑室内上道位置-机房类型", Method: "put", Uri: "/api/install/indoorRoomType/:uuid"},
			{Name: "删除室内上道位置-机房类型", Method: "delete", Uri: "/api/install/indoorRoomType/:uuid"},
			{Name: "室内上道位置-机房列表", Method: "get", Uri: "/api/install/indoorRoom"},
			{Name: "室内上道位置-机房详情", Method: "get", Uri: "/api/install/indoorRoom/:uuid"},
			{Name: "新建室内上道位置-机房", Method: "post", Uri: "/api/install/indoorRoom"},
			{Name: "编辑室内上道位置-机房", Method: "put", Uri: "/api/install/indoorRoom/:uuid"},
			{Name: "删除室内上道位置-机房", Method: "delete", Uri: "/api/install/indoorRoom/:uuid"},
			{Name: "室内上道位置-排列表", Method: "get", Uri: "/api/install/indoorPlatoon"},
			{Name: "室内上道位置-排详情", Method: "get", Uri: "/api/install/indoorPlatoon/:uuid"},
			{Name: "新建室内上道位置-排", Method: "post", Uri: "/api/install/indoorPlatoon"},
			{Name: "编辑室内上道位置-排", Method: "put", Uri: "/api/install/indoorPlatoon/:uuid"},
			{Name: "删除室内上道位置-排", Method: "delete", Uri: "/api/install/indoorPlatoon/:uuid"},
			{Name: "室内上道位置-架列表", Method: "get", Uri: "/api/install/indoorShelf"},
			{Name: "室内上道位置-架详情", Method: "get", Uri: "/api/install/indoorShelf/:uuid"},
			{Name: "新建室内上道位置-架", Method: "post", Uri: "/api/install/indoorShelf"},
			{Name: "编辑室内上道位置-架", Method: "put", Uri: "/api/install/indoorShelf/:uuid"},
			{Name: "室内上道位置-架删除", Method: "delete", Uri: "/api/install/indoorShelf/:uuid"},
			{Name: "室内上道位置-层列表", Method: "get", Uri: "/api/install/indoorTier"},
			{Name: "详情室内上道位置-层", Method: "get", Uri: "/api/install/indoorTier/:uuid"},
			{Name: "新建室内上道位置-层", Method: "post", Uri: "/api/install/indoorTier"},
			{Name: "编辑室内上道位置-层", Method: "put", Uri: "/api/install/indoorTier/:uuid"},
			{Name: "删除室内上道位置-层", Method: "delete", Uri: "/api/install/indoorTier/:uuid"},
			{Name: "室内上道位置-位列表", Method: "get", Uri: "/api/install/indoorCell"},
			{Name: "室内上道位置-位详情", Method: "get", Uri: "/api/install/indoorCell/:uuid"},
			{Name: "新建室内上道位置-位", Method: "post", Uri: "/api/install/indoorCell"},
			{Name: "编辑室内上道位置-位", Method: "put", Uri: "/api/install/indoorCell/:uuid"},
			{Name: "删除室内上道位置-位", Method: "delete", Uri: "/api/install/indoorCell/:uuid"},
		})
	if ret = models.NewRbacPermissionMdl().GetDb("").Create(&rbacPermissions); ret.Error != nil {
		std.EchoLineWrong(fmt.Sprintf("错误：%v", ret.Error.Error()))
	}
	std.EchoLineSuccess("成功")

	std.EchoLineDebug("初始化菜单")
	database.ExecSql("truncate table rbac_menus")
	std.EchoLineSuccess("截断菜单表成功")
	rbacMenu = &models.RbacMenuMdl{Name: "基础信息", Icon: "fa fa-cogs"}
	models.NewRbacMenuMdl().GetDb("").Create(rbacMenu)
	models.NewRbacMenuMdl().GetDb("").Create([]*models.RbacMenuMdl{
		{Name: "器材种类管理", Uri: "/equipmentKind/category", Icon: "fa fa-list", PageRouteName: "equipmentKindCategory:index", ParentUuid: rbacMenu.Uuid},
		{Name: "器材类型管理", Uri: "/equipmentKind/type", Icon: "fa fa-list", PageRouteName: "equipmentKindType:index", ParentUuid: rbacMenu.Uuid},
		{Name: "器材型号管理", Uri: "/equipmentKind/model", Icon: "fa fa-list", PageRouteName: "equipmentKindModel:index", ParentUuid: rbacMenu.Uuid},

		{Name: "路局管理", Uri: "/organization/railway", Icon: "fa fa-list", PageRouteName: "organizationRailway:index", ParentUuid: rbacMenu.Uuid},
		{Name: "站段管理", Uri: "/organization/paragraph", Icon: "fa fa-list", PageRouteName: "organizationParagraph:index", ParentUuid: rbacMenu.Uuid},
		{Name: "车间管理", Uri: "/organization/workshop", Icon: "fa fa-list", PageRouteName: "organizationWorkshop:index", ParentUuid: rbacMenu.Uuid},
		{Name: "站场管理", Uri: "/organization/station", Icon: "fa fa-list", PageRouteName: "organizationStation:index", ParentUuid: rbacMenu.Uuid},
		{Name: "道口管理", Uri: "/organization/crossroad", Icon: "fa fa-list", PageRouteName: "organizationCrossroad:index", ParentUuid: rbacMenu.Uuid},
		{Name: "中心管理", Uri: "/organization/center", Icon: "fa fa-list", PageRouteName: "organizationCenter:index", ParentUuid: rbacMenu.Uuid},
		{Name: "工区管理", Uri: "/organization/workArea", Icon: "fa fa-list", PageRouteName: "organizationWorkArea:index", ParentUuid: rbacMenu.Uuid},
		{Name: "线别管理", Uri: "/organization/line", Icon: "fa fa-list", PageRouteName: "organizationLine:index", ParentUuid: rbacMenu.Uuid},
		{Name: "故障类型管理", Uri: "/breakdown/type", Icon: "fa fa-list", PageRouteName: "breakdownType:index", ParentUuid: rbacMenu.Uuid},
		{Name: "厂家管理", Uri: "/factory", Icon: "fa fa-list", PageRouteName: "factory:index", ParentUuid: rbacMenu.Uuid},
		{Name: "来源类型管理", Uri: "/source/type", Icon: "fa fa-list", PageRouteName: "sourceType:index", ParentUuid: rbacMenu.Uuid},
		{Name: "来源项目管理", Uri: "/source/project", Icon: "fa fa-list", PageRouteName: "sourceProject:index", ParentUuid: rbacMenu.Uuid},
		{Name: "仓库-库房管理", Uri: "/warehouse/storehouse", Icon: "fa fa-list", PageRouteName: "warehouseStorehouse:index", ParentUuid: rbacMenu.Uuid},
		{Name: "仓库-库区管理", Uri: "/warehouse/area", Icon: "fa fa-list", PageRouteName: "warehouseArea:index", ParentUuid: rbacMenu.Uuid},
		{Name: "仓库-排管理", Uri: "/warehouse/platoon", Icon: "fa fa-list", PageRouteName: "warehousePlatoon:index", ParentUuid: rbacMenu.Uuid},
		{Name: "仓库-架管理", Uri: "/warehouse/shelf", Icon: "fa fa-list", PageRouteName: "warehouseShelf:index", ParentUuid: rbacMenu.Uuid},
		{Name: "室内上道位置-机房类型管理", Uri: "/install/indoorRoomType", Icon: "fa fa-list", PageRouteName: "installIndoorRoomType:index", ParentUuid: rbacMenu.Uuid},
		{Name: "室内上道位置-机房管理", Uri: "/install/indoorRoom", Icon: "fa fa-list", PageRouteName: "installIndoorRoom:index", ParentUuid: rbacMenu.Uuid},
		{Name: "室内上道位置-排管理", Uri: "/install/indoorPlatoon", Icon: "fa fa-list", PageRouteName: "installIndoorPlatoon:index", ParentUuid: rbacMenu.Uuid},
		{Name: "室内上道位置-架管理", Uri: "/install/indoorShelf", Icon: "fa fa-list", PageRouteName: "installIndoorShelf:index", ParentUuid: rbacMenu.Uuid},
	})
	rbacMenu = &models.RbacMenuMdl{Name: "系统管理", Icon: "fa fa-cogs"}
	models.NewRbacMenuMdl().GetDb("").Create(rbacMenu)
	models.NewRbacMenuMdl().GetDb("").Create([]*models.RbacMenuMdl{
		{Name: "用户管理", Uri: "/account", Icon: "fa fa-list", PageRouteName: "acocunt:index", ParentUuid: rbacMenu.Uuid},
		{Name: "角色管理", Uri: "/rbac/role", Icon: "fa fa-list", PageRouteName: "rbacRole:index", ParentUuid: rbacMenu.Uuid},
		{Name: "权限管理", Uri: "/rbac/permission", Icon: "fa fa-list", PageRouteName: "rbacPermission:index", ParentUuid: rbacMenu.Uuid},
		{Name: "菜单管理", Uri: "/rbac/menu", Icon: "fa fa-list", PageRouteName: "rbacMenu:index", ParentUuid: rbacMenu.Uuid},
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
