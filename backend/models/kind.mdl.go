package models

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// EquipmentKindCategoryMdl 器材种类模型
type EquipmentKindCategoryMdl struct {
	MySqlMdl
	UniqueCode         string                  `gorm:"unique;char(3);not null;comment:器材种类代码;" json:"unique_code"`
	Name               string                  `gorm:"unique;varchar(128);not null;comment:器材种类名称;" json:"nickname"`
	EquipmentKindTypes []*EquipmentKindTypeMdl `gorm:"foreignKey:equipment_kind_category_uuid;references:uuid;comment:相关器材类型;" json:"equipment_kind_types"`
}

// TableName 器材种类表名称
func (EquipmentKindCategoryMdl) TableName() string {
	return "equipment_kind_categories"
}

// NewEquipmentKindCategoryMdl 新建器材种类模型
func NewEquipmentKindCategoryMdl() *MySqlMdl {
	return NewMySqlMdl().SetModel(EquipmentKindCategoryMdl{})
}

// GetListByQuery 根据Query获取器材种类列表
func (receiver EquipmentKindCategoryMdl) GetListByQuery(ctx *gin.Context) *gorm.DB {
	return NewEquipmentKindCategoryMdl().
		SetWheresEqual().
		SetWheresFuzzy(map[string]string{
			"name": "ekc.name like ?",
		}).
		SetWheresDateBetween("ekc.created_at", "ekc.updated_at", "ekc.deleted_at").
		SetWheresExtraHasValue(map[string]func(string, *gorm.DB) *gorm.DB{}).
		SetWheresExtraHasValues(map[string]func([]string, *gorm.DB) *gorm.DB{}).
		SetCtx(ctx).
		GetDbUseQuery("").
		Table("equipment_kind_categories as ekc")
}

// EquipmentKindTypeMdl 器材类型模型
type EquipmentKindTypeMdl struct {
	MySqlMdl
	UniqueCode                string                    `gorm:"unique;char(5);not null;comment:器材类型代码;" json:"unique_code"`
	Name                      string                    `gorm:"unique;varchar(128);not null;comment:器材类型名称;" json:"nickname"`
	EquipmentKindCategoryUuid string                    `gorm:"type:char(36);not null;comment:器材种类UUID;" json:"equipment_kind_category_uuid"`
	EquipmentKindCategoryMdl  *EquipmentKindCategoryMdl `gorm:"foreignKey:equipment_kind_category_uuid;references:uuid;comment:所属器材种类;" json:"equipment_kind_category_mdl"`
	EquipmentKindModel        []*EquipmentKindModelMdl  `gorm:"foreignKey:equipment_kind_type_uuid;references:uuid;comment:相关器材型号;" json:"equipment_kind_models"`
}

// TableName 器材类型表名称
func (EquipmentKindTypeMdl) TableName() string {
	return "equipment_kind_types"
}

// NewEquipmentKindTypeMdl 新建器材类型模型
func NewEquipmentKindTypeMdl() *MySqlMdl {
	return NewMySqlMdl().SetModel(EquipmentKindTypeMdl{})
}

// GetListByQuery 根据Query获取器材类型列表
func (receiver EquipmentKindTypeMdl) GetListByQuery(ctx *gin.Context) *gorm.DB {
	return NewEquipmentKindTypeMdl().
		SetWheresEqual().
		SetWheresFuzzy(map[string]string{
			"name": "ekt.name like ?",
		}).
		SetWheresDateBetween("ekt.created_at", "ekt.updated_at", "ekt.deleted_at").
		SetWheresExtraHasValue(map[string]func(string, *gorm.DB) *gorm.DB{}).
		SetWheresExtraHasValues(map[string]func([]string, *gorm.DB) *gorm.DB{}).
		SetCtx(ctx).
		GetDbUseQuery("").
		Table("equipment_kind_types as ekt")
}

// EquipmentKindModelMdl 器材型号模型
type EquipmentKindModelMdl struct {
	MySqlMdl
	UniqueCode            string                `gorm:"unique;char(7);not null;comment:器材型号代码;" json:"unique_code"`
	Name                  string                `gorm:"unique;varchar(128);not null;comment:器材型号名称;" json:"nickname"`
	EquipmentKindTypeUuid string                `gorm:"type:char(36);not null;comment:器材类型UUID;" json:"equipment_kind_type_uuid"`
	EquipmentKindTypeMdl  *EquipmentKindTypeMdl `gorm:"foreignKey:equipment_kind_type_uuid;references:uuid;comment:所属器材类型;" json:"equipment_kind_type_mdl"`
}

// TableName 器材型号表名称
func (EquipmentKindModelMdl) TableName() string {
	return "equipment_kind_models"
}

// NewEquipmentKindModelMdl 新建器材型号模型
func NewEquipmentKindModelMdl() *MySqlMdl {
	return NewMySqlMdl().SetModel(EquipmentKindModelMdl{})
}

// GetListByQuery 根据Query获取器材型号列表
func (receiver EquipmentKindModelMdl) GetListByQuery(ctx *gin.Context) *gorm.DB {
	return NewEquipmentKindModelMdl().
		SetWheresEqual().
		SetWheresFuzzy(map[string]string{
			"name": "ekm.name like ?",
		}).
		SetWheresDateBetween("ekm.created_at", "ekm.updated_at", "ekm.deleted_at").
		SetWheresExtraHasValue(map[string]func(string, *gorm.DB) *gorm.DB{}).
		SetWheresExtraHasValues(map[string]func([]string, *gorm.DB) *gorm.DB{}).
		SetCtx(ctx).
		GetDbUseQuery("").
		Table("equipment_kind_models as ekm")
}
