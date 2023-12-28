package models

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// IndoorInstallRoomMdl 室内上道位置-机房模型
type IndoorInstallRoomMdl struct {
	MySqlMdl
}

// TableName 室内上道位置-机房表名称
func (IndoorInstallRoomMdl) TableName() string {
	return "indoor_install_rooms"
}

// NewIndoorInstallRoomMdl 新建室内上道位置-机房模型
func NewIndoorInstallRoomMdl() *MySqlMdl {
	return NewMySqlMdl().SetModel(IndoorInstallRoomMdl{})
}

// GetListByQuery 根据Query获取室内上道位置-机房列表
func (receiver IndoorInstallRoomMdl) GetListByQuery(ctx *gin.Context) *gorm.DB {
	return NewIndoorInstallRoomMdl().
		SetWheresEqual().
		SetWheresFuzzy(map[string]string{
			"name": "iir.name like ?",
		}).
		SetWheresDateBetween("iir.created_at", "iir.updated_at", "iir.deleted_at").
		SetWheresExtraHasValue(map[string]func(string, *gorm.DB) *gorm.DB{}).
		SetWheresExtraHasValues(map[string]func([]string, *gorm.DB) *gorm.DB{}).
		SetCtx(ctx).
		GetDbUseQuery("").
		Table("indoor_install_rooms as iir")
}

// IndoorInstallPlatoonMdl 室内上道位置-排模型
type IndoorInstallPlatoonMdl struct {
	MySqlMdl
}

// TableName 室内上道位置-排表名称
func (IndoorInstallPlatoonMdl) TableName() string {
	return "indoor_install_platoons"
}

// NewIndoorInstallPlatoonMdl 新建室内上道位置-排模型
func NewIndoorInstallPlatoonMdl() *MySqlMdl {
	return NewMySqlMdl().SetModel(IndoorInstallPlatoonMdl{})
}

// GetListByQuery 根据Query获取室内上道位置-排列表
func (receiver IndoorInstallPlatoonMdl) GetListByQuery(ctx *gin.Context) *gorm.DB {
	return NewIndoorInstallPlatoonMdl().
		SetWheresEqual().
		SetWheresFuzzy(map[string]string{
			"name": "iip.name like ?",
		}).
		SetWheresDateBetween("iip.created_at", "iip.updated_at", "iip.deleted_at").
		SetWheresExtraHasValue(map[string]func(string, *gorm.DB) *gorm.DB{}).
		SetWheresExtraHasValues(map[string]func([]string, *gorm.DB) *gorm.DB{}).
		SetCtx(ctx).
		GetDbUseQuery("").
		Table("indoor_install_platoons as iip")
}

// IndoorInstallShelfMdl 室内上道位置-机柜模型
type IndoorInstallShelfMdl struct {
	MySqlMdl
}

// TableName 室内上道位置-机柜表名称
func (IndoorInstallShelfMdl) TableName() string {
	return "indoor_install_shelves"
}

// NewIndoorInstallCabinetMdl 新建室内上道位置-机柜模型
func NewIndoorInstallCabinetMdl() *MySqlMdl {
	return NewMySqlMdl().SetModel(IndoorInstallShelfMdl{})
}

// GetListByQuery 根据Query获取室内上道位置-机柜列表
func (receiver IndoorInstallShelfMdl) GetListByQuery(ctx *gin.Context) *gorm.DB {
	return NewIndoorInstallCabinetMdl().
		SetWheresEqual().
		SetWheresFuzzy(map[string]string{
			"name": "iis.name like ?",
		}).
		SetWheresDateBetween("iis.created_at", "iis.updated_at", "iis.deleted_at").
		SetWheresExtraHasValue(map[string]func(string, *gorm.DB) *gorm.DB{}).
		SetWheresExtraHasValues(map[string]func([]string, *gorm.DB) *gorm.DB{}).
		SetCtx(ctx).
		GetDbUseQuery("").
		Table("indoor_install_shelves as iis")
}

// IndoorInstallTierMdl 室内上道位置-层模型
type IndoorInstallTierMdl struct {
	MySqlMdl
}

// TableName 室内上道位置-层表名称
func (IndoorInstallTierMdl) TableName() string {
	return "indoor_install_tiers"
}

// NewIndoorInstallTierMdl 新建室内上道位置-层模型
func NewIndoorInstallTierMdl() *MySqlMdl {
	return NewMySqlMdl().SetModel(IndoorInstallTierMdl{})
}

// GetListByQuery 根据Query获取室内上道位置-层列表
func (receiver IndoorInstallTierMdl) GetListByQuery(ctx *gin.Context) *gorm.DB {
	return NewIndoorInstallTierMdl().
		SetWheresEqual().
		SetWheresFuzzy(map[string]string{
			"name": "iit.name like ?",
		}).
		SetWheresDateBetween("iit.created_at", "iit.updated_at", "iit.deleted_at").
		SetWheresExtraHasValue(map[string]func(string, *gorm.DB) *gorm.DB{}).
		SetWheresExtraHasValues(map[string]func([]string, *gorm.DB) *gorm.DB{}).
		SetCtx(ctx).
		GetDbUseQuery("").
		Table("indoor_install_tiers as iit")
}

// IndoorInstallCellMdl 室内上道位置-格位模型
type IndoorInstallCellMdl struct {
	MySqlMdl
}

// TableName 室内上道位置-格位表名称
func (IndoorInstallCellMdl) TableName() string {
	return "indoor_install_cells"
}

// NewIndoorInstallCellMdl 新建室内上道位置-格位模型
func NewIndoorInstallCellMdl() *MySqlMdl {
	return NewMySqlMdl().SetModel(IndoorInstallCellMdl{})
}

// GetListByQuery 根据Query获取室内上道位置-格位列表
func (receiver IndoorInstallCellMdl) GetListByQuery(ctx *gin.Context) *gorm.DB {
	return NewIndoorInstallCellMdl().
		SetWheresEqual().
		SetWheresFuzzy(map[string]string{
			"name": "iic.name like ?",
		}).
		SetWheresDateBetween("iic.created_at", "iic.updated_at", "iic.deleted_at").
		SetWheresExtraHasValue(map[string]func(string, *gorm.DB) *gorm.DB{}).
		SetWheresExtraHasValues(map[string]func([]string, *gorm.DB) *gorm.DB{}).
		SetCtx(ctx).
		GetDbUseQuery("").
		Table("indoor_install_cells as iic")
}
