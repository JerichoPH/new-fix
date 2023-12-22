package models

import (
	"new-fix/database"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type (
	// AccountMdl 用户模型
	AccountMdl struct {
		MySqlMdl
		Username   string         `gorm:"unique;type:varchar(64);not null;comment:账号;" json:"username"`
		Password   string         `gorm:"type:varchar(255);not null;comment:密码;" json:"-"`
		Nickname   string         `gorm:"unique;type:varchar(64);not null;comment:昵称;" json:"nickname"`
		BeAdmin    bool           `gorm:"type:boolean;not null;default:0;comment:是否是管理员" json:"be_admin"`
		AvatarUuid string         `gorm:"index;type:char(36);not null;default:'';comment:用户头像文件uuid;" json:"avatar_uuid"`
		Avatar     *FileMdl       `gorm:"foreignKey:avatar_uuid;references:uuid" json:"avatar"`
		RbacRoles  []*RbacRoleMdl `gorm:"many2many:pivot_rbac_roles__accounts;foreignKey:uuid;joinForeignKey:account_uuid;references:uuid;joinReferences:rbac_role_uuid;" json:"rbac_roles"`
	}
)

func (AccountMdl) TableName() string {
	return "accounts"
}

func NewAccountMdl() *MySqlMdl {
	return NewMySqlMdl().SetModel(AccountMdl{})
}

// GetPermissionUuids 获取当前用户所有权限
func (receiver AccountMdl) GetPermissionUuids() (rbacPermissionUuids []string) {
	database.NewGormLauncher().
		GetConn("").
		Table("accounts as a").
		Select(strings.Join([]string{
			"DISTINCT rp.uuid",
		}, ",")).
		Joins("join pivot_rbac_roles__accounts prra on a.uuid = prra.account_uuid").
		Joins("join pivot_rbac_roles__rbac_permissions prrrp on prra.rbac_role_uuid = prrrp.rbac_role_uuid").
		Joins("join rbac_permissions rp on prrrp.rbac_permission_uuid = rp.uuid").
		Where("a.uuid = ?", receiver.Uuid).
		Pluck("uuid", &rbacPermissionUuids)

	return
}

// GetListByQuery 通过Query获取列表
func (receiver AccountMdl) GetListByQuery(ctx *gin.Context) *gorm.DB {
	return NewAccountMdl().
		SetWheresEqual("a.uuid", "a.work_area_unique_code", "a.rank").
		SetWheresFuzzy(map[string]string{
			"account":  "a.account like ?",
			"nickname": "a.nickname like ?",
		}).
		SetWheresDateBetween("a.created_at", "a.updated_at", "a.deleted_at").
		SetWheresExtraHasValue(map[string]func(string, *gorm.DB) *gorm.DB{}).
		SetWheresExtraHasValues(map[string]func([]string, *gorm.DB) *gorm.DB{}).
		SetCtx(ctx).
		GetDbUseQuery("").
		Table("accounts as a")
}

// BindRbacRoles 用户绑定角色
func (receiver AccountMdl) BindRbacRoles(rbacRoles []*RbacRoleMdl) {
	var pivotRbacRoleAccounts []*PivotRbacRoleAccountMdl
	database.
		NewGormLauncher().
		GetConn("").
		Table("pivot_rbac_roles__accounts").
		Where("account_uuid = ?", receiver.Uuid).
		Delete(nil)

	if len(rbacRoles) > 0 {
		for _, rbacRole := range rbacRoles {
			pivotRbacRoleAccounts = append(pivotRbacRoleAccounts, &PivotRbacRoleAccountMdl{
				AccountUuid:  receiver.Uuid,
				RbacRoleUuid: rbacRole.Uuid,
			})
		}
		NewPivotRbacRoleAccountMdl().
			GetDb("").
			Create(&pivotRbacRoleAccounts)
	}
}
