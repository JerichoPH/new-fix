package models

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// EquipmentMdl 器材模型
type EquipmentMdl struct {
	MySqlMdl
	UniqueCode                 string                   `gorm:"unique;char(19);not null;comment:唯一编号;" json:"unique_code"`
	EquipmentKindModelUuid     string                   `gorm:"type:char(36);not null;comment:器材种类UUID;" json:"equipment_kind_model_uuid"`
	EquipmentKindModel         *EquipmentKindModelMdl   `gorm:"foreignKey:equipment_kind_model_uuid;references:uuid;comment:所属器材型号;" json:"equipment_kind_model"`
	PropertyBelongWorkAreaUuid string                   `gorm:"index;type:char(36);not null;default:'';comment:资产归属;" json:"property_belong_work_area_uuid"`
	PropertyBelongWorkArea     *OrganizationWorkAreaMdl `gorm:"foreignKey:property_belong_work_area_uuid;references:uuid;comment:资产归属工区;" json:"property_belong_work_area"`
	UseBelongWorkshopUuid      string                   `gorm:"index;type:char(36);not null;default:'';comment:使用归属车间;" json:"use_belong_workshop_uuid"`
}

// TableName 器材表名称
func (EquipmentMdl) TableName() string {
	return "equipments"
}

// NewEquipmentMdl 新建器材模型
func NewEquipmentMdl() *MySqlMdl {
	return NewMySqlMdl().SetModel(EquipmentMdl{})
}

// GetListByQuery 根据Query获取器材列表
func (receiver EquipmentMdl) GetListByQuery(ctx *gin.Context) *gorm.DB {
	return NewEquipmentMdl().
		SetWheresEqual().
		SetWheresFuzzy(map[string]string{
			"name": "e.name like ?",
		}).
		SetWheresDateBetween("e.created_at", "e.updated_at", "e.deleted_at").
		SetWheresExtraHasValue(map[string]func(string, *gorm.DB) *gorm.DB{}).
		SetWheresExtraHasValues(map[string]func([]string, *gorm.DB) *gorm.DB{}).
		SetCtx(ctx).
		GetDbUseQuery("").
		Table("equipments as e")
}

// EquipmentLogMdl 器材日志模型
type EquipmentLogMdl struct {
	MySqlMdl
}

// TableName 器材日志表名称
func (EquipmentLogMdl) TableName() string {
	return "equipment_logs"
}

// NewEquipmentLogMdl 新建器材日志模型
func NewEquipmentLogMdl() *MySqlMdl {
	return NewMySqlMdl().SetModel(EquipmentLogMdl{})
}

// GetListByQuery 根据Query获取器材日志列表
func (receiver EquipmentLogMdl) GetListByQuery(ctx *gin.Context) *gorm.DB {
	return NewEquipmentLogMdl().
		SetWheresEqual().
		SetWheresFuzzy(map[string]string{
			"name": "el.name like ?",
		}).
		SetWheresDateBetween("el.created_at", "el.updated_at", "el.deleted_at").
		SetWheresExtraHasValue(map[string]func(string, *gorm.DB) *gorm.DB{}).
		SetWheresExtraHasValues(map[string]func([]string, *gorm.DB) *gorm.DB{}).
		SetCtx(ctx).
		GetDbUseQuery("").
		Table("equipment_logs as el")
}
