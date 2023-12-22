package models

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type (
	// EquipmentKindCategoryMdl 器材种类模型
	EquipmentKindCategoryMdl struct {
		MySqlMdl
		UniqueCode         string                  `gorm:"unique;char(3);not null;comment:器材种类代码;" json:"unique_code"`
		Name               string                  `gorm:"unique;varchar(128);not null;comment:器材种类名称;" json:"nickname"`
		EquipmentKindTypes []*EquipmentKindTypeMdl `gorm:"foreignKey:equipment_kind_category_uuid;references:uuid;comment:相关器材类型;" json:"equipment_kind_types"`
	}
	// EquipmentKindTypeMdl 器材类型模型
	EquipmentKindTypeMdl struct {
		MySqlMdl
		UniqueCode                string                    `gorm:"unique;char(5);not null;comment:器材类型代码;" json:"unique_code"`
		Name                      string                    `gorm:"unique;varchar(128);not null;comment:器材类型名称;" json:"nickname"`
		EquipmentKindCategoryUuid string                    `gorm:"type:char(36);not null;comment:器材种类UUID;" json:"equipment_kind_category_uuid"`
		EquipmentKindCategory     *EquipmentKindCategoryMdl `gorm:"foreignKey:equipment_kind_category_uuid;references:uuid;comment:所属器材种类;" json:"equipment_kind_category"`
		EquipmentKindModel        []*EquipmentKindModelMdl  `gorm:"foreignKey:equipment_kind_type_uuid;references:uuid;comment:相关器材型号;" json:"equipment_kind_models"`
		FullName                  string                    `gorm:"-" json:"full_name"`
	}
	// EquipmentKindModelMdl 器材型号模型
	EquipmentKindModelMdl struct {
		MySqlMdl
		UniqueCode            string                `gorm:"unique;char(7);not null;comment:器材型号代码;" json:"unique_code"`
		Name                  string                `gorm:"unique;varchar(128);not null;comment:器材型号名称;" json:"nickname"`
		EquipmentKindTypeUuid string                `gorm:"type:char(36);not null;comment:器材类型UUID;" json:"equipment_kind_type_uuid"`
		EquipmentKindType     *EquipmentKindTypeMdl `gorm:"foreignKey:equipment_kind_type_uuid;references:uuid;comment:所属器材类型;" json:"equipment_kind_type"`
		FullName              string                `gorm:"-" json:"full_name"`
	}
)

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

// AfterFind 获取类型全名
func (receiver EquipmentKindTypeMdl) AfterFind(db *gorm.DB) (err error) {
	receiver.FullName = fmt.Sprintf(
		"%s %s",
		receiver.EquipmentKindCategory.Name,
		receiver.Name,
	)
	return
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

// AfterFind 获取型号全名
func (receiver EquipmentKindModelMdl) AfterFind(db *gorm.DB) (err error) {
	receiver.FullName = fmt.Sprintf(
		"%s %s %s",
		receiver.EquipmentKindType.EquipmentKindCategory.Name,
		receiver.EquipmentKindType.Name,
		receiver.Name,
	)
	return
}
