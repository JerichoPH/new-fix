package models

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type (
	// WarehouseStorehouseMdl 仓库模型
	WarehouseStorehouseMdl struct {
		MySqlMdl
		Name              string                 `json:"name" gorm:"column:name;type:varchar(255);not null;default:'';comment:仓库名称"`
		WarehouseSections []*WarehouseSectionMdl `gorm:"foreignKey:WarehouseStorehouseUuid;references:Uuid" json:"warehouse_sections"`
	}

	// WarehouseSectionMdl 仓库区模型
	WarehouseSectionMdl struct {
		MySqlMdl
		WarehouseStorehouseUuid string                  `json:"warehouse_storehouse_uuid" gorm:"column:warehouse_storehouse_uuid;type:char(36);not null;default:'';comment:仓库编号"`
		WarehouseStorehouse     *WarehouseStorehouseMdl `gorm:"foreignKey:WarehouseStorehouseUuid;references:Uuid" json:"warehouse_storehouse"`
	}
)

// TableName 仓库表名称
func (WarehouseStorehouseMdl) TableName() string {
	return "warehouse_storehouses"
}

// NewWarehouseStorehouseMdl 新建仓库模型
func NewWarehouseStorehouseMdl() *MySqlMdl {
	return NewMySqlMdl().SetModel(WarehouseStorehouseMdl{})
}

// GetListByQuery 根据Query获取仓库列表
func (receiver WarehouseStorehouseMdl) GetListByQuery(ctx *gin.Context) *gorm.DB {
	return NewWarehouseStorehouseMdl().
		SetWheresEqual().
		SetWheresFuzzy(map[string]string{
			"name": "ws.name like ?",
		}).
		SetWheresDateBetween("ws.created_at", "ws.updated_at", "ws.deleted_at").
		SetWheresExtraHasValue(map[string]func(string, *gorm.DB) *gorm.DB{}).
		SetWheresExtraHasValues(map[string]func([]string, *gorm.DB) *gorm.DB{}).
		SetCtx(ctx).
		GetDbUseQuery("").
		Table("warehouse_storehouses as ws")
}

// TableName 仓库区表名称
func (WarehouseSectionMdl) TableName() string {
	return "warehouse_sections"
}

// NewWarehouseSectionMdl 新建仓库区模型
func NewWarehouseSectionMdl() *MySqlMdl {
	return NewMySqlMdl().SetModel(WarehouseSectionMdl{})
}

// GetListByQuery 根据Query获取仓库区列表
func (receiver WarehouseSectionMdl) GetListByQuery(ctx *gin.Context) *gorm.DB {
	return NewWarehouseSectionMdl().
		SetWheresEqual().
		SetWheresFuzzy(map[string]string{
			"name": "wse.name like ?",
		}).
		SetWheresDateBetween("wse.created_at", "wse.updated_at", "wse.deleted_at").
		SetWheresExtraHasValue(map[string]func(string, *gorm.DB) *gorm.DB{}).
		SetWheresExtraHasValues(map[string]func([]string, *gorm.DB) *gorm.DB{}).
		SetCtx(ctx).
		GetDbUseQuery("").
		Table("warehouse as wse")
}
