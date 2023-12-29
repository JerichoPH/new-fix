package models

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type (
	// BreakdownTypeMdl 故障类型模型
	BreakdownTypeMdl struct {
		MySqlMdl
		Name                      string                    `gorm:"type:varchar(128);not null;comment:故障类型名称;" json:"name"`
		EquipmentKindCategoryUuid string                    `gorm:"type:char(36);not null;comment:所属器材型号编号;" json:"equipment_kind_category_uuid"`
		EquipmentKindCategory     *EquipmentKindCategoryMdl `gorm:"foreignKey:equipment_kind_category_uuid;references:uuid;comment:所属器材型号;" json:"equipment_kind_category"`
		BreakdownLogs             []*BreakdownLogMdl        `gorm:"foreignKey:breakdown_type_uuid;references:uuid;comment:故障日志;" json:"breakdown_logs"`
	}

	// BreakdownLogMdl 故障日志模型
	BreakdownLogMdl struct {
		MySqlMdl
		EquipmentUuid     string            `gorm:"type:char(36);not null;comment:器材编号;" json:"equipment_uuid"`
		Equipment         *EquipmentMdl     `gorm:"foreignKey:equipment_uuid;references:uuid;comment:所属器材;" json:"equipment"`
		BreakdownTypeUuid string            `gorm:"type:char(36);not null;comment:故障类型编号;" json:"breakdown_type_uuid"`
		BreakdownType     *BreakdownTypeMdl `gorm:"foreignKey:breakdown_type_uuid;references:uuid;comment:所属故障类型;" json:"breakdown_type"`
		ReportPersonUuid  string            `gorm:"type:char(36);not null;comment:报修人编号;" json:"report_person_uuid"`
		ReportPerson      *AccountMdl       `gorm:"foreignKey:report_person_uuid;references:uuid;comment:报修人;" json:"report_person"`
		TypeCode          string            `gorm:"type:enum('WORKSHOP', 'SCENE');not null;comment:故障日志类型;" json:"type_code"`
		TypeText          string            `gorm:"-" json:"type_text"`
	}
)

// TableName 故障类型表名称
func (BreakdownTypeMdl) TableName() string {
	return "breakdown_types"
}

// NewBreakdownTypeMdl 新建故障类型模型
func NewBreakdownTypeMdl() *MySqlMdl {
	return NewMySqlMdl().SetModel(BreakdownTypeMdl{})
}

// GetListByQuery 根据Query获取故障类型列表
func (receiver BreakdownTypeMdl) GetListByQuery(ctx *gin.Context) *gorm.DB {
	return NewBreakdownTypeMdl().
		SetWheresEqual().
		SetWheresFuzzy(map[string]string{
			"name": "bt.name like ?",
		}).
		SetWheresDateBetween("bt.created_at", "bt.updated_at", "bt.deleted_at").
		SetWheresExtraHasValue(map[string]func(string, *gorm.DB) *gorm.DB{
			"equipment_kind_category_uuid": func(value string, db *gorm.DB) *gorm.DB {
				return db.Where("ekc.uuid = ?", value)
			},
		}).
		SetWheresExtraHasValues(map[string]func([]string, *gorm.DB) *gorm.DB{}).
		SetCtx(ctx).
		GetDbUseQuery("").
		Table("breakdown_types as bt").
		Joins("join equipment_kind_categories ekc on bt.equipment_kind_category_uuid = ekc.uuid").Debug()
}

// TableName 故障日志表名称
func (BreakdownLogMdl) TableName() string {
	return "breakdown_logs"
}

// NewBreakdownLogMdl 新建故障日志模型
func NewBreakdownLogMdl() *MySqlMdl {
	return NewMySqlMdl().SetModel(BreakdownLogMdl{})
}

// GetListByQuery 根据Query获取故障日志列表
func (receiver BreakdownLogMdl) GetListByQuery(ctx *gin.Context) *gorm.DB {
	return NewBreakdownLogMdl().
		SetWheresEqual().
		SetWheresFuzzy(map[string]string{
			"name": "bl.name like ?",
		}).
		SetWheresDateBetween("bl.created_at", "bl.updated_at", "bl.deleted_at").
		SetWheresExtraHasValue(map[string]func(string, *gorm.DB) *gorm.DB{
			"breakdown_type_uuid": func(value string, db *gorm.DB) *gorm.DB {
				return db.Where("bt.uuid = ?", value)
			},
			"equipment_uuid": func(value string, db *gorm.DB) *gorm.DB {
				return db.Where("e.uuid = ?", value)
			},
			"report_person_uuid": func(value string, db *gorm.DB) *gorm.DB {
				return db.Where("a.uuid = ?", value)
			},
		}).
		SetWheresExtraHasValues(map[string]func([]string, *gorm.DB) *gorm.DB{}).
		SetCtx(ctx).
		GetDbUseQuery("").
		Table("breakdown_logs as bl").
		Joins("join breakdown_types bt on bl.breakdown_type_uuid = bt.uuid").
		Joins("join equipments e on bl.equipment_uuid = e.uuid").
		Joins("join accounts a on bl.report_person_uuid = a.uuid")
}

// GetTypeCodes 获取故障日志类型列表
func (BreakdownLogMdl) GetTypeCodes() []string {
	return []string{"WORKSHOP", "SCENE"}
}

// GetTypeCodesMap 获取故障日志类型映射
func (BreakdownLogMdl) GetTypeCodesMap() []map[string]string {
	return []map[string]string{
		{"code": "WORKSHOP", "text": "车间"},
		{"code": "SCENE", "text": "现场"},
	}
}

// getTypeText 获取故障日志类型描述
func (receiver *BreakdownLogMdl) getTypeText() {
	var typeText string
	for _, item := range receiver.GetTypeCodesMap() {
		if item["code"] == receiver.TypeCode {
			typeText = item["text"]
		}
	}
	receiver.TypeText = typeText
}

// AfterFind 查询后置钩子
func (receiver *BreakdownLogMdl) AfterFind(db *gorm.DB) (err error) {
	receiver.getTypeText()
	return
}
