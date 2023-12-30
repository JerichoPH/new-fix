package models

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// FactoryMdl 厂家模型
type FactoryMdl struct {
	MySqlMdl
	UniqueCode string          `gorm:"index;type:char(5);not null;comment:代码;" json:"unique_code"`
	Name       string          `gorm:"index;type:varchar(64);not null;comment:名称;" json:"name"`
	Equipments []*EquipmentMdl `gorm:"foreignKey:factory_uuid;references:uuid;comment:厂家;" json:"equipments"`
}

// TableName 厂家表名称
func (FactoryMdl) TableName() string {
	return "factories"
}

// NewFactoryMdl 新建厂家模型
func NewFactoryMdl() *MySqlMdl {
	return NewMySqlMdl().SetModel(FactoryMdl{})
}

// GetListByQuery 根据Query获取厂家列表
func (receiver FactoryMdl) GetListByQuery(ctx *gin.Context) *gorm.DB {
	return NewFactoryMdl().
		SetWheresEqual().
		SetWheresFuzzy(map[string]string{
			"name": "f.name like ?",
		}).
		SetWheresDateBetween("f.created_at", "f.updated_at", "f.deleted_at").
		SetWheresExtraHasValue(map[string]func(string, *gorm.DB) *gorm.DB{
			"unique_code": func(value string, db *gorm.DB) *gorm.DB {
				return db.Where("f.unique_code = ?", value)
			},
		}).
		SetWheresExtraHasValues(map[string]func([]string, *gorm.DB) *gorm.DB{}).
		SetCtx(ctx).
		GetDbUseQuery("").
		Table("factories as f")
}
