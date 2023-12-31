package models

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type (
	// SourceTypeMdl 来源类型模型
	SourceTypeMdl struct {
		MySqlMdl
		UniqueCode     string              `gorm:"type:char(2);not null;comment:来源类型代码;" json:"unique_code"`
		Name           string              `gorm:"unique;type:varchar(128);not null;comment:来源类型名称" json:"name"`
		SourceProjects []*SourceProjectMdl `gorm:"foreignKey:source_type_uuid;references:uuid;comment:相关来源项目" json:"source_projects"`
		Equipments     []*EquipmentMdl     `gorm:"foreignKey:source_type_uuid;references:uuid;comment:相关器材" json:"equipments"`
	}
	// SourceProjectMdl 来源项目模型
	SourceProjectMdl struct {
		MySqlMdl
		SourceTypeUuid string          `gorm:"type:char(36);not null;comment:来源类型uuid;" json:"source_type_uuid"`
		SourceType     *SourceTypeMdl  `gorm:"foreignKey:source_type_uuid;references:uuid;comment:来源类型" json:"source_type"`
		Name           string          `gorm:"unique;type:varchar(128);not null;comment:来源项目名称" json:"name"`
		Equipments     []*EquipmentMdl `gorm:"foreignKey:source_project_uuid;references:uuid;comment:相关器材" json:"equipments"`
	}
)

// TableName 来源类型表名称
func (SourceTypeMdl) TableName() string {
	return "source_types"
}

// NewSourceTypeMdl 新建来源类型模型
func NewSourceTypeMdl() *MySqlMdl {
	return NewMySqlMdl().SetModel(SourceTypeMdl{})
}

// GetListByQuery 根据Query获取来源类型列表
func (receiver SourceTypeMdl) GetListByQuery(ctx *gin.Context) *gorm.DB {
	return NewSourceTypeMdl().
		SetWheresEqual().
		SetWheresFuzzy(map[string]string{
			"name": "st.name like ?",
		}).
		SetWheresDateBetween("st.created_at", "st.updated_at", "st.deleted_at").
		SetWheresExtraHasValue(map[string]func(string, *gorm.DB) *gorm.DB{}).
		SetWheresExtraHasValues(map[string]func([]string, *gorm.DB) *gorm.DB{}).
		SetCtx(ctx).
		GetDbUseQuery("").
		Table("source_types as st")
}

// TableName 来源项目表名称
func (SourceProjectMdl) TableName() string {
	return "source_projects"
}

// NewSourceProjectMdl 新建来源项目模型
func NewSourceProjectMdl() *MySqlMdl {
	return NewMySqlMdl().SetModel(SourceProjectMdl{})
}

// GetListByQuery 根据Query获取来源项目列表
func (receiver SourceProjectMdl) GetListByQuery(ctx *gin.Context) *gorm.DB {
	return NewSourceProjectMdl().
		SetWheresEqual().
		SetWheresFuzzy(map[string]string{
			"name": "sp.name like ?",
		}).
		SetWheresDateBetween("sp.created_at", "sp.updated_at", "sp.deleted_at").
		SetWheresExtraHasValue(map[string]func(string, *gorm.DB) *gorm.DB{}).
		SetWheresExtraHasValues(map[string]func([]string, *gorm.DB) *gorm.DB{}).
		SetCtx(ctx).
		GetDbUseQuery("").
		Table("source_projects as sp")
}
