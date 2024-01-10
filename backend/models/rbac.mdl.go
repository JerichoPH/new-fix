package models

import (
	"fmt"
	"new-fix/database"
	"new-fix/wrongs"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type (
	// RbacRoleMdl 角色模型
	RbacRoleMdl struct {
		MySqlMdl
		Name            string               `gorm:"type:varchar(64);not null;comment:角色名称;" json:"name"`
		Accounts        []*AccountMdl        `gorm:"many2many:pivot_rbac_roles__accounts;foreignKey:uuid;joinForeignKey:rbac_role_uuid;references:uuid;joinReferences:accountUuid;" json:"accounts"`
		RbacPermissions []*RbacPermissionMdl `gorm:"many2many:pivot_rbac_roles__rbac_permissions;foreignKey:uuid;joinForeignKey:rbac_role_uuid;references:uuid;joinReferences:rbacPermissionUuid;" json:"rbac_permissions"`
		RbacMenus       []*RbacMenuMdl       `gorm:"many2many:pivot_rbac_roles__rbac_menus;foreignKey:uuid;joinForeignKey:rbac_role_uuid;references:uuid;joinReferences:rbacMenuUuid;" json:"rbac_menus"`
	}

	// RbacPermissionMdl 权限模型
	RbacPermissionMdl struct {
		MySqlMdl
		Name        string         `gorm:"type:varchar(64);not null;comment:权限名称;" json:"name"`
		Description *string        `gorm:"type:text;comment:权限描述;" json:"description"`
		Uri         string         `gorm:"type:varchar(255);not null;default:'';comment:权限所属路由;" json:"uri"`
		Method      string         `gorm:"type:varchar(32);not null;comment:请求方法;" json:"method"`
		RbacRoles   []*RbacRoleMdl `gorm:"many2many:pivot_rbac_roles__rbac_permissions;foreignKey:uuid;joinForeignKey:rbac_permission_uuid;references:uuid;joinReferences:rbacRoleUuid;" json:"rbac_roles"`
	}

	// RbacMenuMdl 菜单模型
	RbacMenuMdl struct {
		MySqlMdl
		Name          string         `gorm:"type:varchar(64);not null;comment:菜单名称" json:"name"`
		SubTitle      string         `gorm:"type:varchar(128);not null;default:'';comment:菜单副标题" json:"sub_title"`
		Description   *string        `gorm:"type:text;comment:菜单描述" json:"description"`
		Uri           string         `gorm:"type:varchar(128);not null;default:'';comment:菜单所属路由" json:"uri"`
		Icon          string         `gorm:"type:varchar(64);not null;default:'';comment:菜单图标" json:"icon"`
		PageRouteName string         `gorm:"type:varchar(64);not null;default:'';comment:页面路由名称;" json:"page_route_name"`
		RbacRoles     []*RbacRoleMdl `gorm:"many2many:pivot_rbac_roles__rbac_menus;foreignKey:uuid;joinForeignKey:rbac_menu_uuid;references:uuid;joinReferences:rbacRoleUuid;" json:"rbac_roles"`
		ParentUuid    string         `gorm:"type:char(36);not null;default:'';comment:父级uuid;" json:"parent_uuid"`
		Parent        *RbacMenuMdl   `gorm:"foreignKey:parent_uuid;references:uuid;comment:所属父级;" json:"parent"`
		Subs          []*RbacMenuMdl `gorm:"foreignKey:parent_uuid;references:uuid;comment:相关子集;" json:"subs"`
	}

	PivotRbacRoleAccountMdl struct {
		RbacRoleUuid string `gorm:"type:char(36);not null;default:'';comment:角色uuid" json:"rbac_role_uuid"`
		AccountUuid  string `gorm:"type:char(36);not null;default:'';comment:用户uuid" json:"account_uuid"`
	}

	PivotRbacRoleRbacPermissionMdl struct {
		RbacRoleUuid       string `gorm:"type:char(36);not null;default:'';comment:角色uuid" json:"rbac_role_uuid"`
		RbacPermissionUuid string `gorm:"type:char(36);not null;default:'';comment:权限uuid" json:"rbac_permission_uuid"`
	}

	PivotRbacRoleRbacMenuMdl struct {
		RbacRoleUuid string `gorm:"type:char(36);not null;default:'';comment:角色uuid" json:"rbac_role_uuid"`
		RbacMenuUuid string `gorm:"type:char(36);not null;default:'';comment:菜单uuid" json:"rbac_menu_uuid"`
	}
)

// NewRbacRoleMdl 创建一个新的 RBAC 角色模型
func NewRbacRoleMdl() *MySqlMdl {
	return NewMySqlMdl().SetModel(&RbacRoleMdl{})
}

// TableName 角色表名称
func (RbacRoleMdl) TableName() string {
	return "rbac_roles"
}

// GetListByQuery 根据Query获取角色列表
func (receiver RbacRoleMdl) GetListByQuery(ctx *gin.Context) *gorm.DB {
	return NewRbacRoleMdl().
		SetWheresEqual("rr.be_enable").
		SetWheresDateBetween("rr.created_at", "rr.updated_at", "rr.deleted_at").
		SetWheresExtraHasValue(map[string]func(string, *gorm.DB) *gorm.DB{
			"name": func(value string, db *gorm.DB) *gorm.DB {
				return db.Where(fmt.Sprintf("rr.name like '%%%s%%'", value))
			},
		}).
		SetWheresExtraHasValues(map[string]func([]string, *gorm.DB) *gorm.DB{
			"names[]": func(values []string, db *gorm.DB) *gorm.DB {
				return db.Where("rr.name in (?)", values)
			},
		}).
		SetCtx(ctx).
		GetDbUseQuery("").
		Table("rbac_roles as rr")
}

// NewRbacPermissionMdl 返回一个新的 RbacPermissionModel 模型实例化的指针
func NewRbacPermissionMdl() *MySqlMdl {
	return NewMySqlMdl().SetModel(&RbacPermissionMdl{})
}

// TableName 权限表名称
func (RbacPermissionMdl) TableName() string {
	return "rbac_permissions"
}

// GetListByQuery 根据Query获取权限列表
func (receiver RbacPermissionMdl) GetListByQuery(ctx *gin.Context) *gorm.DB {
	return NewRbacPermissionMdl().
		SetWheresEqual("rp.be_enable").
		SetWheresDateBetween("rp.created_at", "rp.updated_at", "rp.deleted_at").
		SetWheresExtraHasValue(map[string]func(string, *gorm.DB) *gorm.DB{
			"name": func(value string, db *gorm.DB) *gorm.DB {
				return db.Where(fmt.Sprintf("rp.name like '%%%s%%'", value))
			},
			"rbac_role_uuid": func(value string, db *gorm.DB) *gorm.DB {
				return db.
					Joins("left join pivot_rbac_roles__rbac_permissions prrrp on rp.uuid = prrrp.rbac_permission_uuid").
					Joins("left join rbac_roles rr on prrrp.rbac_role_uuid = rr.uuid").
					Where("rr.uuid =?", value)
			},
		}).
		SetWheresExtraHasValues(map[string]func([]string, *gorm.DB) *gorm.DB{
			"names[]": func(values []string, db *gorm.DB) *gorm.DB {
				return db.Where("rp.name in (?)", values)
			},
			"rbac_role_uuids[]": func(values []string, db *gorm.DB) *gorm.DB {
				return db.
					Joins("left join pivot_rbac_roles__rbac_permissions prrrp on rp.uuid = prrrp.rbac_permission_uuid").
					Joins("left join rbac_roles rr on prrrp.rbac_role_uuid = rr.uuid").
					Where("rr.uuid in (?)", values)
			},
		}).
		SetCtx(ctx).
		GetDbUseQuery("").
		Table("rbac_permissions as rp")
}

// NewRbacMenuMdl 返回一个新的 RbacMenuModel 模型实例指针
func NewRbacMenuMdl() *MySqlMdl {
	return NewMySqlMdl().SetModel(&RbacMenuMdl{})
}

// TableName 菜单表名称
func (RbacMenuMdl) TableName() string {
	return "rbac_menus"
}

// GetListByQuery 根据Query获取菜单列表
func (receiver RbacMenuMdl) GetListByQuery(ctx *gin.Context) *gorm.DB {
	var (
		notHasSubs = ctx.Query("not_has_subs")
		subs       = make(map[string]map[string]string)
		subUuids   []string
	)
	subUuids = make([]string, 0)
	if notHasSubs != "" {
		subs = receiver.GetSubUuidsByParentUuid(notHasSubs)
		for _, sub := range subs {
			subUuids = append(subUuids, sub["uuid"])
		}
	}

	return NewRbacMenuMdl().
		SetWheresEqual("rm.be_enable").
		SetWheresDateBetween("rm.created_at", "rm.updated_at", "rm.deleted_at").
		SetWheresExtraHasValue(map[string]func(string, *gorm.DB) *gorm.DB{
			"name": func(value string, db *gorm.DB) *gorm.DB {
				return db.Where(fmt.Sprintf("rm.name like '%%%s%%'", value))
			},
			"uri": func(value string, db *gorm.DB) *gorm.DB {
				return db.Where("rm.uri = ?", value)
			},
			"not_has_subs": func(value string, db *gorm.DB) *gorm.DB {
				if len(subUuids) == 0 {
					return db
				}
				return db.Where("rm.uuid not in ?", subUuids)
			},
		}).
		SetWheresExtraHasValues(map[string]func([]string, *gorm.DB) *gorm.DB{
			"names[]": func(values []string, db *gorm.DB) *gorm.DB {
				return db.Where("rm. name in (?)", values)
			},
			"uris[]": func(values []string, db *gorm.DB) *gorm.DB {
				return db.Where("rm. uri in (?)", values)
			},
		}).
		SetCtx(ctx).
		GetDbUseQuery("").
		Table("rbac_menus as rm")
}

// GetSubUuidsByParentUuid 根据父级uuid获取所有子集uuid
func (receiver RbacMenuMdl) GetSubUuidsByParentUuid(parentUuid string) map[string]map[string]string {
	if rows := database.ExecSql(strings.Join([]string{
		"WITH RECURSIVE cte AS (",
		"SELECT uuid, name, NULL AS parent_uuid",
		"FROM rbac_menus",
		"WHERE parent_uuid = ?",
		"AND deleted_at IS NULL",
		"UNION ALL",
		"SELECT m.uuid, m.name, c.parent_uuid",
		"FROM rbac_menus m INNER JOIN cte c ON m.parent_uuid = c.uuid",
		"WHERE m.deleted_at IS NULL",
		")",
		"SELECT uuid, name FROM cte",
	}, "\r\n"), parentUuid); rows != nil {
		var (
			subs = make(map[string]map[string]string)
			err  error
		)
		for rows.Next() {
			var (
				uuid string
				name string
			)
			if err = rows.Scan(&uuid, &name); err != nil {
				wrongs.ThrowForbidden(err.Error())
			}
			subs[uuid] = map[string]string{
				"uuid": uuid,
				"name": name,
			}
		}
		return subs
	}
	return map[string]map[string]string{}
}

// NewPivotRbacRoleAccountMdl 返回一个新的 PivotRbacRoleAccountModel 模型实例
func NewPivotRbacRoleAccountMdl() *MySqlMdl {
	return NewMySqlMdl().SetModel(&PivotRbacRoleAccountMdl{})
}

// TableName 角色与用户对应关系表名称
func (PivotRbacRoleAccountMdl) TableName() string {
	return "pivot_rbac_roles__accounts"
}

// BindRbacRoles 绑定角色与用户
func (PivotRbacRoleAccountMdl) BindRbacRoles(rbacRole *RbacRoleMdl, accounts []*AccountMdl) {
	database.NewGormLauncher().GetConn("").Table("pivot_rbac_roles__accounts").Where("rbac_role_uuid", rbacRole.Uuid).Delete(nil)
	if len(accounts) > 0 {
		for _, account := range accounts {
			NewPivotRbacRoleAccountMdl().
				GetDb("").
				Create(&PivotRbacRoleAccountMdl{
					RbacRoleUuid: rbacRole.Uuid,
					AccountUuid:  account.Uuid,
				})
		}
	}
}

// NewPivotRbacRoleRbacPermissionMdl 返回一个新的 PivotRbacRoleRbacPermissionModel 模型的实例。
func NewPivotRbacRoleRbacPermissionMdl() *MySqlMdl {
	return NewMySqlMdl().SetModel(&PivotRbacRoleRbacPermissionMdl{})
}

// TableName 角色与权限对应关系表名称
func (PivotRbacRoleRbacPermissionMdl) TableName() string {
	return "pivot_rbac_roles__rbac_permissions"
}

// BindRbacRoles 绑定角色与权限
func (PivotRbacRoleRbacPermissionMdl) BindRbacRoles(rbacPermission *RbacPermissionMdl, rbacRoles []*RbacRoleMdl) {
	database.NewGormLauncher().GetConn("").Table("pivot_rbac_roles__rbac_permissions").Where("rbac_permission_uuid", rbacPermission.Uuid).Delete(nil)
	if len(rbacRoles) > 0 {
		for _, rbacRole := range rbacRoles {
			NewPivotRbacRoleRbacPermissionMdl().
				GetDb("").
				Create(&PivotRbacRoleRbacPermissionMdl{
					RbacRoleUuid:       rbacRole.Uuid,
					RbacPermissionUuid: rbacPermission.Uuid,
				})
		}
	}
}

// NewPivotRbacRoleRbacMenuMdl 返回一个新的 PivotRbacRoleRbacMenuModel 模型的实例。
func NewPivotRbacRoleRbacMenuMdl() *MySqlMdl {
	return NewMySqlMdl().SetModel(&PivotRbacRoleRbacMenuMdl{})
}

// TableName 角色与菜单对应关系表名称
func (PivotRbacRoleRbacMenuMdl) TableName() string {
	return "pivot_rbac_roles__rbac_menus"
}

// BindRbacRoles 绑定角色与菜单
func (PivotRbacRoleRbacMenuMdl) BindRbacRoles(rbacMenu *RbacMenuMdl, rbacRoles []*RbacRoleMdl) {
	var pivotRbacRoleRbacMenus = make([]*PivotRbacRoleRbacMenuMdl, 0)

	database.NewGormLauncher().GetConn("").Table("pivot_rbac_roles__rbac_menus").Where("rbac_menu_uuid = ?", rbacMenu.Uuid).Delete(nil)

	if len(rbacRoles) > 0 {
		for _, rbacRole := range rbacRoles {
			pivotRbacRoleRbacMenus = append(pivotRbacRoleRbacMenus, &PivotRbacRoleRbacMenuMdl{
				RbacRoleUuid: rbacRole.Uuid,
				RbacMenuUuid: rbacMenu.Uuid,
			})
		}
	}

	NewPivotRbacRoleRbacMenuMdl().GetDb("").Create(&pivotRbacRoleRbacMenus)
}
