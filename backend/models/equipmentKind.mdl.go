package models

import (
	"fmt"
	"new-fix/tools"
	"new-fix/wrongs"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type (
	// EquipmentKindCategoryMdl 器材种类模型
	EquipmentKindCategoryMdl struct {
		MySqlMdl
		UniqueCode         string                  `gorm:"unique;char(3);not null;comment:器材种类代码;" json:"unique_code"`
		Name               string                  `gorm:"unique;varchar(128);not null;comment:器材种类名称;" json:"name"`
		EquipmentKindTypes []*EquipmentKindTypeMdl `gorm:"foreignKey:equipment_kind_category_uuid;references:uuid;comment:相关器材类型;" json:"equipment_kind_types"`
		BreakdownTypes     []*BreakdownTypeMdl     `gorm:"foreignKey:equipment_kind_category_uuid;references:uuid;comment:相关故障类型;" json:"breakdown_types"`
	}
	// EquipmentKindTypeMdl 器材类型模型
	EquipmentKindTypeMdl struct {
		MySqlMdl
		UniqueCode                string                    `gorm:"unique;char(5);not null;comment:器材类型代码;" json:"unique_code"`
		Name                      string                    `gorm:"unique;varchar(128);not null;comment:器材类型名称;" json:"name"`
		EquipmentKindCategoryUuid string                    `gorm:"type:char(36);not null;comment:器材种类UUID;" json:"equipment_kind_category_uuid"`
		EquipmentKindCategory     *EquipmentKindCategoryMdl `gorm:"foreignKey:equipment_kind_category_uuid;references:uuid;comment:所属器材种类;" json:"equipment_kind_category"`
		EquipmentKindModel        []*EquipmentKindModelMdl  `gorm:"foreignKey:equipment_kind_type_uuid;references:uuid;comment:相关器材型号;" json:"equipment_kind_models"`
	}
	// EquipmentKindModelMdl 器材型号模型
	EquipmentKindModelMdl struct {
		MySqlMdl
		UniqueCode            string                `gorm:"unique;char(7);not null;comment:器材型号代码;" json:"unique_code"`
		Name                  string                `gorm:"unique;varchar(128);not null;comment:器材型号名称;" json:"name"`
		EquipmentKindTypeUuid string                `gorm:"type:char(36);not null;comment:器材类型UUID;" json:"equipment_kind_type_uuid"`
		EquipmentKindType     *EquipmentKindTypeMdl `gorm:"foreignKey:equipment_kind_type_uuid;references:uuid;comment:所属器材类型;" json:"equipment_kind_type"`
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

// GetLastUniqueCode 获取最后一个UniqueCode
func (receiver EquipmentKindCategoryMdl) GetLastUniqueCode() string {
	ret := NewEquipmentKindCategoryMdl().GetDb("").Order("unique_code desc").First(&receiver)

	if wrongs.ThrowWhenEmpty(ret, "") {
		return ""
	}

	return receiver.UniqueCode
}

// GetNewUniqueCode 获取新的UniqueCode
func (receiver EquipmentKindCategoryMdl) GetNewUniqueCode() string {
	var (
		lastUniqueCode, newUniqueCode           string
		lastUniqueCodeInt64, newUniqueCodeInt64 int64
		err                                     error
	)

	lastUniqueCode = receiver.GetLastUniqueCode()
	if lastUniqueCode != "" {
		lastUniqueCodeInt64, err = tools.NewMath().Base36ToDec(lastUniqueCode[len(lastUniqueCode)-2:])
		if err != nil {
			wrongs.ThrowForbidden("器材种类代码36 -> 10失败：%s", err.Error())
		}
		newUniqueCodeInt64 = lastUniqueCodeInt64 + 1
		newUniqueCode, err = tools.NewMath().DecToBase36(newUniqueCodeInt64)
		if err != nil {
			wrongs.ThrowForbidden("器材种类代码10 -> 36失败：%s", err.Error())
		}
		return fmt.Sprintf("Q%s", tools.NewMath().PadLeftZeros(newUniqueCode, 2))
	}

	return "Q01"
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
		SetWheresExtraHasValue(map[string]func(string, *gorm.DB) *gorm.DB{
			"equipment_kind_category_uuid": func(value string, db *gorm.DB) *gorm.DB {
				return db.Where("ekc.uuid = ?", value)
			},
		}).
		SetWheresExtraHasValues(map[string]func([]string, *gorm.DB) *gorm.DB{}).
		SetCtx(ctx).
		GetDbUseQuery("").
		Table("equipment_kind_types as ekt").
		Joins("join equipment_kind_categories as ekc on ekc.uuid = ekt.equipment_kind_category_uuid")
}

// GetLastUniqueCode 获取最后一个UniqueCode
func (receiver EquipmentKindTypeMdl) GetLastUniqueCode(equipmentKindCategory *EquipmentKindCategoryMdl) string {
	ret := NewEquipmentKindTypeMdl().GetDb("").Where("equipment_kind_category_uuid = ?", equipmentKindCategory.Uuid).Order("unique_code desc").First(&receiver)

	if wrongs.ThrowWhenEmpty(ret, "") {
		return ""
	}

	return receiver.UniqueCode
}

// GetNewUniqueCode 获取新的UniqueCode
func (receiver EquipmentKindTypeMdl) GetNewUniqueCode(equipmentKindCategory *EquipmentKindCategoryMdl) string {
	var (
		lastUniqueCode, newUniqueCode           string
		lastUniqueCodeInt64, newUniqueCodeInt64 int64
		err                                     error
	)

	lastUniqueCode = receiver.GetLastUniqueCode(equipmentKindCategory)
	if lastUniqueCode != "" {
		lastUniqueCodeInt64, err = tools.NewMath().Base36ToDec(lastUniqueCode[len(lastUniqueCode)-4:])
		if err != nil {
			wrongs.ThrowForbidden("器材类型代码36 -> 10失败：%s", err.Error())
		}
		newUniqueCodeInt64 = lastUniqueCodeInt64 + 1
		newUniqueCode, err = tools.NewMath().DecToBase36(newUniqueCodeInt64)
		if err != nil {
			wrongs.ThrowForbidden("器材类型代码10 -> 36失败：%s", err.Error())
		}
		return fmt.Sprintf("%s%s", equipmentKindCategory.UniqueCode, tools.NewMath().PadLeftZeros(newUniqueCode, 2))
	}

	return fmt.Sprintf("%s01", equipmentKindCategory.UniqueCode)
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
		SetWheresExtraHasValue(map[string]func(string, *gorm.DB) *gorm.DB{
			"equipment_kind_type_uuid": func(value string, db *gorm.DB) *gorm.DB {
				return db.Where("ekt.uuid = ?", value)
			},
			"equipment_kind_category_uuid": func(value string, db *gorm.DB) *gorm.DB {
				return db.Where("ekc.uuid = ?", value)
			},
		}).
		SetWheresExtraHasValues(map[string]func([]string, *gorm.DB) *gorm.DB{}).
		SetCtx(ctx).
		GetDbUseQuery("").
		Table("equipment_kind_models as ekm").
		Joins("join equipment_kind_types as ekt on ekt.uuid = ekm.equipment_kind_type_uuid").
		Joins("join equipment_kind_categories as ekc on ekc.uuid = ekt.equipment_kind_category_uuid")
}

// GetLastUniqueCode 获取最后一个UniqueCode
func (receiver EquipmentKindModelMdl) GetLastUniqueCode(equipmentKindType *EquipmentKindTypeMdl) string {
	ret := NewEquipmentKindTypeMdl().GetDb("").Where("equipment_kind_type_uuid = ?", equipmentKindType.Uuid).Order("unique_code desc").First(&receiver)

	if wrongs.ThrowWhenEmpty(ret, "") {
		return ""
	}

	return receiver.UniqueCode
}

// GetNewUniqueCode 获取新的UniqueCode
func (receiver EquipmentKindModelMdl) GetNewUniqueCode(equipmentKindType *EquipmentKindTypeMdl) string {
	var (
		lastUniqueCode, newUniqueCode           string
		lastUniqueCodeInt64, newUniqueCodeInt64 int64
		err                                     error
	)

	lastUniqueCode = receiver.GetLastUniqueCode(equipmentKindType)
	if lastUniqueCode != "" {
		lastUniqueCodeInt64, err = tools.NewMath().Base36ToDec(lastUniqueCode[len(lastUniqueCode)-6:])
		if err != nil {
			wrongs.ThrowForbidden("器材类型代码36 -> 10失败：%s", err.Error())
		}
		newUniqueCodeInt64 = lastUniqueCodeInt64 + 1
		newUniqueCode, err = tools.NewMath().DecToBase36(newUniqueCodeInt64)
		if err != nil {
			wrongs.ThrowForbidden("器材类型代码10 -> 36失败：%s", err.Error())
		}
		return fmt.Sprintf("%s%s", equipmentKindType.UniqueCode, tools.NewMath().PadLeftZeros(newUniqueCode, 2))
	}

	return fmt.Sprintf("%s01", equipmentKindType.UniqueCode)
}
