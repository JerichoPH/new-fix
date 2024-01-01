package models

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type (
	// WarehouseStorehouseMdl 仓库模型
	WarehouseStorehouseMdl struct {
		MySqlMdl
		Name              string              `gorm:"type:varchar(128);not null;default:'';comment:名称;" json:"name"`
		WarehouseSections []*WarehouseAreaMdl `gorm:"foreignKey:warehouse_storehouse_uuid;references:uuid;" json:"warehouse_areas"`
	}

	// WarehouseAreaMdl 仓库区模型
	WarehouseAreaMdl struct {
		MySqlMdl
		Name                    string                  `gorm:"type:varchar(128);not null;default:'';comment:名称;" json:"name"`
		WarehouseStorehouseUuid string                  `gorm:"index;type:char(36);not null;default:'';comment:仓库-库房uuid;" json:"warehouse_storehouse_uuid"`
		WarehouseStorehouse     *WarehouseStorehouseMdl `gorm:"foreignKey:warehouse_storehouse_uuid;references:uuid;" json:"warehouse_storehouse"`
		WarehousePlatoons       []*WarehousePlatoonMdl  `gorm:"foreignKey:warehouse_area_uuid;references:uuid;" json:"warehouse_platoons"`
	}

	// WarehousePlatoonMdl 仓库-排模型
	WarehousePlatoonMdl struct {
		MySqlMdl
		Name              string               `gorm:"type:varchar(128);not null;default:'';comment:名称;" json:"name"`
		WarehouseAreaUuid string               `gorm:"index;type:char(36);not null;default:'';comment:仓库-区uuid;" json:"warehouse_area_uuid"`
		WarehouseArea     *WarehouseAreaMdl    `gorm:"foreignKey:warehouse_area_uuid;references:uuid;" json:"warehouse_area"`
		WarehouseShelves  []*WarehouseShelfMdl `gorm:"foreignKey:warehouse_platoon_uuid;references:uuid;" json:"warehouse_shelves"`
	}

	// WarehouseShelfMdl 仓库-柜架模型
	WarehouseShelfMdl struct {
		MySqlMdl
		Name                 string               `gorm:"type:varchar(128);not null;default:'';comment:名称;" json:"name"`
		WarehousePlatoonUuid string               `gorm:"index;type:char(36);not null;default:'';comment:仓库-排uuid;" json:"warehouse_platoon_uuid"`
		WarehousePlatoon     *WarehousePlatoonMdl `gorm:"foreignKey:warehouse_platoon_uuid;references:uuid;" json:"warehouse_platoon"`
		WarehouseTiers       []*WarehouseTierMdl  `gorm:"foreignKey:warehouse_shelf_uuid;references:uuid;" json:"warehouse_tiers"`
	}

	// WarehouseTierMdl 仓库-层模型
	WarehouseTierMdl struct {
		MySqlMdl
		Name               string              `gorm:"type:varchar(128);not null;default:'';comment:名称;" json:"name"`
		WarehouseShelfUuid string              `gorm:"index;type:char(36);not null;default:'';comment:仓库-柜架uuid;" json:"warehouse_shelf_uuid"`
		WarehouseShelf     *WarehouseShelfMdl  `gorm:"foreignKey:warehouse_shelf_uuid;references:uuid;" json:"warehouse_shelf"`
		WarehouseCells     []*WarehouseCellMdl `gorm:"foreignKey:warehouse_tier_uuid;references:uuid;" json:"warehouse_cells"`
	}

	// WarehouseCellMdl 仓库-格位模型
	WarehouseCellMdl struct {
		MySqlMdl
		Name              string            `gorm:"type:varchar(128);not null;default:'';comment:名称;" json:"name"`
		WarehouseTierUuid string            `gorm:"index;type:char(36);not null;default:'';comment:仓库-层uuid;" json:"warehouse_tier_uuid"`
		WarehouseTier     *WarehouseTierMdl `gorm:"foreignKey:warehouse_tier_uuid;references:uuid;" json:"warehouse_tier"`
		Equipments        []*EquipmentMdl   `gorm:"foreignKey:warehouse_cell_uuid;references:uuid;comment:相关器材;" json:"equipments"`
	}

	// WarehouseScanItemMdl 出入库扫码记录模型
	WarehouseScanItemMdl struct {
		MySqlMdl
		EquipmentUuid string        `gorm:"unique;type:char(36);not null;comment:所属器材uuid;" json:"equipment_uuid"`
		Equipment     *EquipmentMdl `gorm:"foreignKey:equipment_uuid;references:uuid;" json:"equipment"`
		ProcessorUuid string        `gorm:"type:char(36);not null;comment:经办人uuid;" json:"processor_uuid"`
		Processor     *AccountMdl   `gorm:"foreignKey:processor_uuid;references:uuid;" json:"processor"`
		DirectionCode string        `gorm:"enum('IN','OUT');not null;comment:出入库方向;" json:"direction_code"`
		DirectionText string        `gorm:"-" json:"direction_text"`
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
func (WarehouseAreaMdl) TableName() string {
	return "warehouse_areas"
}

// NewWarehouseSectionMdl 新建仓库区模型
func NewWarehouseSectionMdl() *MySqlMdl {
	return NewMySqlMdl().SetModel(WarehouseAreaMdl{})
}

// GetListByQuery 根据Query获取仓库区列表
func (receiver WarehouseAreaMdl) GetListByQuery(ctx *gin.Context) *gorm.DB {
	return NewWarehouseSectionMdl().
		SetWheresEqual().
		SetWheresFuzzy(map[string]string{
			"name": "wa.name like ?",
		}).
		SetWheresDateBetween("wa.created_at", "wa.updated_at", "wa.deleted_at").
		SetWheresExtraHasValue(map[string]func(string, *gorm.DB) *gorm.DB{}).
		SetWheresExtraHasValues(map[string]func([]string, *gorm.DB) *gorm.DB{}).
		SetCtx(ctx).
		GetDbUseQuery("").
		Table("warehouse as wa")
}

// TableName 仓库-排表名称
func (WarehousePlatoonMdl) TableName() string {
	return "warehouse_platoons"
}

// NewWarehousePlatoonMdl 新建仓库-排模型
func NewWarehousePlatoonMdl() *MySqlMdl {
	return NewMySqlMdl().SetModel(WarehousePlatoonMdl{})
}

// GetListByQuery 根据Query获取仓库-排列表
func (receiver WarehousePlatoonMdl) GetListByQuery(ctx *gin.Context) *gorm.DB {
	return NewWarehousePlatoonMdl().
		SetWheresEqual().
		SetWheresFuzzy(map[string]string{
			"name": "wp.name like ?",
		}).
		SetWheresDateBetween("wp.created_at", "wp.updated_at", "wp.deleted_at").
		SetWheresExtraHasValue(map[string]func(string, *gorm.DB) *gorm.DB{}).
		SetWheresExtraHasValues(map[string]func([]string, *gorm.DB) *gorm.DB{}).
		SetCtx(ctx).
		GetDbUseQuery("").
		Table("warehouse_platoons as wp")
}

// TableName 仓库-柜架表名称
func (WarehouseShelfMdl) TableName() string {
	return "warehouse_cabinets"
}

// NewWarehouseShelfMdl 新建仓库-柜架模型
func NewWarehouseShelfMdl() *MySqlMdl {
	return NewMySqlMdl().SetModel(WarehouseShelfMdl{})
}

// GetListByQuery 根据Query获取仓库-柜架列表
func (receiver WarehouseShelfMdl) GetListByQuery(ctx *gin.Context) *gorm.DB {
	return NewWarehouseShelfMdl().
		SetWheresEqual().
		SetWheresFuzzy(map[string]string{
			"name": "ws.name like ?",
		}).
		SetWheresDateBetween("ws.created_at", "ws.updated_at", "ws.deleted_at").
		SetWheresExtraHasValue(map[string]func(string, *gorm.DB) *gorm.DB{}).
		SetWheresExtraHasValues(map[string]func([]string, *gorm.DB) *gorm.DB{}).
		SetCtx(ctx).
		GetDbUseQuery("").
		Table("warehouse_shelves as ws")
}

// TableName 仓库-层表名称
func (WarehouseTierMdl) TableName() string {
	return "warehouse_tiers"
}

// NewWarehouseTierMdl 新建仓库-层模型
func NewWarehouseTierMdl() *MySqlMdl {
	return NewMySqlMdl().SetModel(WarehouseTierMdl{})
}

// GetListByQuery 根据Query获取仓库-层列表
func (receiver WarehouseTierMdl) GetListByQuery(ctx *gin.Context) *gorm.DB {
	return NewWarehouseTierMdl().
		SetWheresEqual().
		SetWheresFuzzy(map[string]string{
			"name": "wt.name like ?",
		}).
		SetWheresDateBetween("wt.created_at", "wt.updated_at", "wt.deleted_at").
		SetWheresExtraHasValue(map[string]func(string, *gorm.DB) *gorm.DB{}).
		SetWheresExtraHasValues(map[string]func([]string, *gorm.DB) *gorm.DB{}).
		SetCtx(ctx).
		GetDbUseQuery("").
		Table("warehouse_tiers as wt")
}

// TableName 仓库-格位表名称
func (WarehouseCellMdl) TableName() string {
	return "warehouse_cells"
}

// NewWarehouseCellMdl 新建仓库-格位模型
func NewWarehouseCellMdl() *MySqlMdl {
	return NewMySqlMdl().SetModel(WarehouseCellMdl{})
}

// GetListByQuery 根据Query获取仓库-格位列表
func (receiver WarehouseCellMdl) GetListByQuery(ctx *gin.Context) *gorm.DB {
	return NewWarehouseCellMdl().
		SetWheresEqual().
		SetWheresFuzzy(map[string]string{
			"name": "wc.name like ?",
		}).
		SetWheresDateBetween("wc.created_at", "wc.updated_at", "wc.deleted_at").
		SetWheresExtraHasValue(map[string]func(string, *gorm.DB) *gorm.DB{}).
		SetWheresExtraHasValues(map[string]func([]string, *gorm.DB) *gorm.DB{}).
		SetCtx(ctx).
		GetDbUseQuery("").
		Table("warehouse_cells as wc")
}

// TableName 出入库扫码记录表名称
func (WarehouseScanItemMdl) TableName() string {
	return "warehouse_scan_items"
}

// NewWarehouseScanItemMdl 新建出入库扫码记录模型
func NewWarehouseScanItemMdl() *MySqlMdl {
	return NewMySqlMdl().SetModel(WarehouseScanItemMdl{})
}

// GetListByQuery 根据Query获取出入库扫码记录列表
func (receiver WarehouseScanItemMdl) GetListByQuery(ctx *gin.Context) *gorm.DB {
	return NewWarehouseScanItemMdl().
		SetWheresEqual().
		SetWheresFuzzy(map[string]string{
			"name": "wsi.name like ?",
		}).
		SetWheresDateBetween("wsi.created_at", "wsi.updated_at", "wsi.deleted_at").
		SetWheresExtraHasValue(map[string]func(string, *gorm.DB) *gorm.DB{}).
		SetWheresExtraHasValues(map[string]func([]string, *gorm.DB) *gorm.DB{}).
		SetCtx(ctx).
		GetDbUseQuery("").
		Table("warehouse_scan_items as wsi")
}

// GetDirectionCodes 获取出入库方向代码
func (WarehouseScanItemMdl) GetDirectionCodes() []string {
	return []string{"IN", "OUT"}
}

// GetDirectionCodesMap 获取出入库方向代码映射
func (WarehouseScanItemMdl) GetDirectionCodesMap() []map[string]string {
	return []map[string]string{
		{"code": "IN", "text": "入库"},
		{"code": "OUT", "text": "出库"},
	}
}

// GetDirectionTexts 获取出入库方向描述
func (receiver *WarehouseScanItemMdl) AfterFind(db *gorm.DB) (err error) {
	var directionText string

	for _, item := range receiver.GetDirectionCodesMap() {
		if item["code"] == receiver.DirectionCode {
			directionText = item["text"]
		}
	}

	receiver.DirectionText = directionText
	return
}
