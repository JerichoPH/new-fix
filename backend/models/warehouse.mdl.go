package models

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type (
	// WarehouseStorehouseMdl 仓库-库房模型
	WarehouseStorehouseMdl struct {
		MySqlMdl
		Name                     string                   `gorm:"type:varchar(128);not null;default:'';comment:名称;" json:"name"`
		OrganizationWorkshopUuid string                   `gorm:"index;type:char(36);not null;default:'';comment:所属组织-车间uuid;" json:"organization_workshop_uuid"`
		OrganizationWorkshop     *OrganizationWorkshopMdl `gorm:"foreignKey:organization_workshop_uuid;references:uuid;" json:"organization_workshop"`
		OrganizationWorkAreaUuid string                   `gorm:"index;type:char(36);not null;default:'';comment:所属组织-工区uuid;" json:"organization_work_area_uuid"`
		OrganizationWorkArea     *OrganizationWorkAreaMdl `gorm:"foreignKey:organization_work_area_uuid;references:uuid;" json:"organization_work_area"`
		WarehouseAreas           []*WarehouseAreaMdl      `gorm:"foreignKey:warehouse_storehouse_uuid;references:uuid;" json:"warehouse_areas"`
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
		TypeCode          string               `gorm:"type:enum('FIXING','FIXED','EMERGENCY');comment:类型代码;" json:"type_code"`
		TypeText          string               `gorm:"-" json:"type_text"`
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
		EquipmentUuid                    string          `gorm:"unique;type:char(36);not null;comment:所属器材uuid;" json:"equipment_uuid"`
		Equipment                        *EquipmentMdl   `gorm:"foreignKey:equipment_uuid;references:uuid;comment:所属器材;" json:"equipment"`
		EquipmentsByWarehouseInScanItem  []*EquipmentMdl `gorm:"foreignKey:warehouse_in_scan_item_uuid;references:uuid;comment:相关器材（入库扫码）;" json:"equipments_by_warehouse_in_scan_item"`
		EquipmentsByWarehouseOutScanItem []*EquipmentMdl `gorm:"foreignKey:warehouse_out_scan_item_uuid;references:uuid;comment:相关器材（出库扫码）;" json:"equipments_by_warehouse_out_scan_item"`
		ProcessorUuid                    string          `gorm:"type:char(36);not null;comment:经办人uuid;" json:"processor_uuid"`
		Processor                        *AccountMdl     `gorm:"foreignKey:processor_uuid;references:uuid;" json:"processor"`
		DirectionCode                    string          `gorm:"enum('IN','OUT');not null;comment:出入库方向;" json:"direction_code"`
		DirectionText                    string          `gorm:"-" json:"direction_text"`
	}
	// WarehouseOrderMdl 仓库-出入库单模型
	WarehouseOrderMdl struct {
		MySqlMdl
		EquipmentUuid                 string          `gorm:"type:char(36);not null;comment:所属器材uuid;" json:"equipment_uuid"`
		Equipment                     *EquipmentMdl   `gorm:"foreignKey:equipment_uuid;references:uuid;comment:所属器材;" json:"equipment"`
		EquipmentsByWarehouseInOrder  []*EquipmentMdl `gorm:"foreignKey:warehouse_in_order_uuid;references:uuid;comment:相关器材（最后入库单）;" json:"equipments_by_warehouse_in_order"`
		EquipmentsByWarehouseOutOrder []*EquipmentMdl `gorm:"foreignKey:warehouse_out_order_uuid;references:uuid;comment:相关器材（最后出库单）;" json:"equipments_by_warehouse_out_order"`
		ProcessorUuid                 string          `gorm:"type:char(36);not null;comment:经办人uuid;" json:"processor_uuid"`
		Processor                     *AccountMdl     `gorm:"foreignKey:processor_uuid;references:uuid;" json:"processor"`
		DirectionCode                 string          `gorm:"enum('IN','OUT');not null;comment:出入库方向;" json:"direction_code"`
		DirectionText                 string          `gorm:"-" json:"direction_text"`
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
		SetWheresExtraHasValue(map[string]func(string, *gorm.DB) *gorm.DB{
			"organization_railway_uuid": func(value string, db *gorm.DB) *gorm.DB {
				return db.Where("oa.uuid = ?", value)
			},
			"organization_paragraph_uuid": func(value string, db *gorm.DB) *gorm.DB {
				return db.Where("op.uuid = ?", value)
			},
			"organization_workshop_uuid": func(value string, db *gorm.DB) *gorm.DB {
				return db.Where("ow.uuid = ?", value)
			},
			"organization_work_area_uuid": func(value string, db *gorm.DB) *gorm.DB {
				return db.Where("owa.uuid = ?", value)
			},
		}).
		SetWheresExtraHasValues(map[string]func([]string, *gorm.DB) *gorm.DB{}).
		SetCtx(ctx).
		GetDbUseQuery("").
		Table("warehouse_storehouses as ws").
		Joins("join organization_workshops ow on ws.organization_workshop_uuid = ow.uuid").
		Joins("left join organization_work_areas owa on ws.organization_work_area_uuid = owa.uuid").
		Joins("join organization_paragraphs op on ow.organization_paragraph_uuid = op.uuid").
		Joins("join organization_railways oa on op.organization_railway_uuid = oa.uuid")
}

// TableName 仓库区表名称
func (WarehouseAreaMdl) TableName() string {
	return "warehouse_areas"
}

// NewWarehouseAreaMdl 新建仓库区模型
func NewWarehouseAreaMdl() *MySqlMdl {
	return NewMySqlMdl().SetModel(WarehouseAreaMdl{})
}

// GetListByQuery 根据Query获取仓库区列表
func (receiver WarehouseAreaMdl) GetListByQuery(ctx *gin.Context) *gorm.DB {
	return NewWarehouseAreaMdl().
		SetWheresEqual().
		SetWheresFuzzy(map[string]string{
			"name": "wa.name like ?",
		}).
		SetWheresDateBetween("wa.created_at", "wa.updated_at", "wa.deleted_at").
		SetWheresExtraHasValue(map[string]func(string, *gorm.DB) *gorm.DB{
			"organization_railway_uuid": func(value string, db *gorm.DB) *gorm.DB {
				return db.Where("oa.uuid = ?", value)
			},
			"organization_paragraph_uuid": func(value string, db *gorm.DB) *gorm.DB {
				return db.Where("op.uuid = ?", value)
			},
			"organization_workshop_uuid": func(value string, db *gorm.DB) *gorm.DB {
				return db.Where("ow.uuid = ?", value)
			},
			"organization_work_area_uuid": func(value string, db *gorm.DB) *gorm.DB {
				return db.Where("owa.uuid = ?", value)
			},
			"warehouse_storehouse_uuid": func(value string, db *gorm.DB) *gorm.DB {
				return db.Where("ws.uuid = ?", value)
			},
		}).
		SetWheresExtraHasValues(map[string]func([]string, *gorm.DB) *gorm.DB{}).
		SetCtx(ctx).
		GetDbUseQuery("").
		Table("warehouse_areas as wa").
		Joins("join warehouse_storehouses ws on wa.warehouse_storehouse_uuid = ws.uuid").
		Joins("join organization_workshops ow on ws.organization_workshop_uuid = ow.uuid").
		Joins("left join organization_work_areas owa on ws.organization_work_area_uuid = owa.uuid").
		Joins("join organization_paragraphs op on ow.organization_paragraph_uuid = op.uuid").
		Joins("join organization_railways oa on op.organization_railway_uuid = oa.uuid")
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
		SetWheresExtraHasValue(map[string]func(string, *gorm.DB) *gorm.DB{
			"organization_railway_uuid": func(value string, db *gorm.DB) *gorm.DB {
				return db.Where("oa.uuid = ?", value)
			},
			"organization_paragraph_uuid": func(value string, db *gorm.DB) *gorm.DB {
				return db.Where("op.uuid = ?", value)
			},
			"organization_workshop_uuid": func(value string, db *gorm.DB) *gorm.DB {
				return db.Where("ow.uuid = ?", value)
			},
			"organization_work_area_uuid": func(value string, db *gorm.DB) *gorm.DB {
				return db.Where("owa.uuid = ?", value)
			},
			"warehouse_storehouse_uuid": func(value string, db *gorm.DB) *gorm.DB {
				return db.Where("ws.uuid = ?", value)
			},
			"warehouse_area_uuid": func(value string, db *gorm.DB) *gorm.DB {
				return db.Where("wa.uuid = ?", value)
			},
			"type_code": func(value string, db *gorm.DB) *gorm.DB {
				return db.Where("wp.type_code = ?", value)
			},
		}).
		SetWheresExtraHasValues(map[string]func([]string, *gorm.DB) *gorm.DB{}).
		SetCtx(ctx).
		GetDbUseQuery("").
		Table("warehouse_platoons as wp").
		Joins("join warehouse_areas wa on wp.warehouse_area_uuid = wa.uuid").
		Joins("join warehouse_storehouses ws on wa.warehouse_storehouse_uuid = ws.uuid").
		Joins("join organization_workshops ow on ws.organization_workshop_uuid = ow.uuid").
		Joins("left join organization_work_areas owa on ws.organization_work_area_uuid = owa.uuid").
		Joins("join organization_paragraphs op on ow.organization_paragraph_uuid = op.uuid").
		Joins("join organization_railways oa on op.organization_railway_uuid = oa.uuid")
}

// GetTypeCodes 仓库-排类型代码列表
func (WarehousePlatoonMdl) GetTypeCodes() []string {
	return []string{"FIXING", "FIXED", "EMERGENCY"}
}

// GetTypeCodesMap 仓库-排类型代码映射
func (WarehousePlatoonMdl) GetTypeCodesMap() []map[string]string {
	return []map[string]string{
		{"code": "FIXING", "text": "待修"},
		{"code": "FIXED", "text": "成品"},
		{"code": "EMERGENCY", "text": "应急备品"},
	}
}

// getTypeText 获取仓库-排类型描述
func (*WarehousePlatoonMdl) getTypeText(warehousePlatoon *WarehousePlatoonMdl) {
	for _, item := range warehousePlatoon.GetTypeCodesMap() {
		if item["code"] == warehousePlatoon.TypeCode {
			warehousePlatoon.TypeText = item["text"]
			return
		}
	}
}

func (receiver *WarehousePlatoonMdl) AfterFind(db *gorm.DB) (err error) {
	receiver.getTypeText(receiver)
	return
}

// TableName 仓库-柜架表名称
func (WarehouseShelfMdl) TableName() string {
	return "warehouse_shelves"
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
			"name": "wh.name like ?",
		}).
		SetWheresDateBetween("wsh.created_at", "wsh.updated_at", "wsh.deleted_at").
		SetWheresExtraHasValue(map[string]func(string, *gorm.DB) *gorm.DB{
			"organization_railway_uuid": func(value string, db *gorm.DB) *gorm.DB {
				return db.Where("oa.uuid = ?", value)
			},
			"organization_paragraph_uuid": func(value string, db *gorm.DB) *gorm.DB {
				return db.Where("op.uuid = ?", value)
			},
			"organization_workshop_uuid": func(value string, db *gorm.DB) *gorm.DB {
				return db.Where("ow.uuid = ?", value)
			},
			"organization_work_area_uuid": func(value string, db *gorm.DB) *gorm.DB {
				return db.Where("owa.uuid = ?", value)
			},
			"warehouse_storehouse_uuid": func(value string, db *gorm.DB) *gorm.DB {
				return db.Where("ws.uuid = ?", value)
			},
			"warehouse_area_uuid": func(value string, db *gorm.DB) *gorm.DB {
				return db.Where("wa.uuid = ?", value)
			},
			"warehouse_platoon_uuid": func(value string, db *gorm.DB) *gorm.DB {
				return db.Where("wp.uuid = ?", value)
			},
		}).
		SetWheresExtraHasValues(map[string]func([]string, *gorm.DB) *gorm.DB{}).
		SetCtx(ctx).
		GetDbUseQuery("").
		Table("warehouse_shelves as wsh").
		Joins("join warehouse_platoons wp on wsh.warehouse_platoon_uuid = wp.uuid").
		Joins("join warehouse_areas wa on wp.warehouse_area_uuid = wa.uuid").
		Joins("join warehouse_storehouses ws on wa.warehouse_storehouse_uuid = ws.uuid").
		Joins("join organization_workshops ow on ws.organization_workshop_uuid = ow.uuid").
		Joins("left join organization_work_areas owa on ws.organization_work_area_uuid = owa.uuid").
		Joins("join organization_paragraphs op on ow.organization_paragraph_uuid = op.uuid").
		Joins("join organization_railways oa on op.organization_railway_uuid = oa.uuid")
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
		SetWheresExtraHasValue(map[string]func(string, *gorm.DB) *gorm.DB{
			"organization_railway_uuid": func(value string, db *gorm.DB) *gorm.DB {
				return db.Where("oa.uuid = ?", value)
			},
			"organization_paragraph_uuid": func(value string, db *gorm.DB) *gorm.DB {
				return db.Where("op.uuid = ?", value)
			},
			"organization_workshop_uuid": func(value string, db *gorm.DB) *gorm.DB {
				return db.Where("ow.uuid = ?", value)
			},
			"organization_work_area_uuid": func(value string, db *gorm.DB) *gorm.DB {
				return db.Where("owa.uuid = ?", value)
			},
			"warehouse_storehouse_uuid": func(value string, db *gorm.DB) *gorm.DB {
				return db.Where("ws.uuid = ?", value)
			},
			"warehouse_area_uuid": func(value string, db *gorm.DB) *gorm.DB {
				return db.Where("wa.uuid = ?", value)
			},
			"warehouse_platoon_uuid": func(value string, db *gorm.DB) *gorm.DB {
				return db.Where("wp.uuid = ?", value)
			},
			"warehouse_shelf_uuid": func(value string, db *gorm.DB) *gorm.DB {
				return db.Where("wsh.uuid = ?", value)
			},
		}).
		SetWheresExtraHasValues(map[string]func([]string, *gorm.DB) *gorm.DB{}).
		SetCtx(ctx).
		GetDbUseQuery("").
		Table("warehouse_tiers as wt").
		Joins("join warehouse_shelves wsh on wt.warehouse_shelf_uuid = wsh.uuid").
		Joins("join warehouse_platoons wp on wsh.warehouse_platoon_uuid = wp.uuid").
		Joins("join warehouse_areas wa on wp.warehouse_area_uuid = wa.uuid").
		Joins("join warehouse_storehouses ws on wa.warehouse_storehouse_uuid = ws.uuid").
		Joins("join organization_workshops ow on ws.organization_workshop_uuid = ow.uuid").
		Joins("left join organization_work_areas owa on ws.organization_work_area_uuid = owa.uuid").
		Joins("join organization_paragraphs op on ow.organization_paragraph_uuid = op.uuid").
		Joins("join organization_railways oa on op.organization_railway_uuid = oa.uuid")
}

// 根据仓库-架获取最后一层
func (WarehouseTierMdl) GetLastByWarehouseShelf(warehouseShelf *WarehouseShelfMdl) (warehouseTier *WarehouseTierMdl) {
	NewWarehouseTierMdl().
		GetDb("").
		Where("warehouse_shelf_uuid = ?", warehouseShelf.Uuid).
		Order("name desc").
		Order("id desc").
		First(&warehouseTier)
	return
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
		SetWheresExtraHasValue(map[string]func(string, *gorm.DB) *gorm.DB{
			"organization_railway_uuid": func(value string, db *gorm.DB) *gorm.DB {
				return db.Where("oa.uuid = ?", value)
			},
			"organization_paragraph_uuid": func(value string, db *gorm.DB) *gorm.DB {
				return db.Where("op.uuid = ?", value)
			},
			"organization_workshop_uuid": func(value string, db *gorm.DB) *gorm.DB {
				return db.Where("ow.uuid = ?", value)
			},
			"organization_work_area_uuid": func(value string, db *gorm.DB) *gorm.DB {
				return db.Where("owa.uuid = ?", value)
			},
			"warehouse_storehouse_uuid": func(value string, db *gorm.DB) *gorm.DB {
				return db.Where("ws.uuid = ?", value)
			},
			"warehouse_area_uuid": func(value string, db *gorm.DB) *gorm.DB {
				return db.Where("wa.uuid = ?", value)
			},
			"warehouse_platoon_uuid": func(value string, db *gorm.DB) *gorm.DB {
				return db.Where("wp.uuid = ?", value)
			},
			"warehouse_shelf_uuid": func(value string, db *gorm.DB) *gorm.DB {
				return db.Where("wsh.uuid = ?", value)
			},
			"warehouse_tier_uuid": func(value string, db *gorm.DB) *gorm.DB {
				return db.Where("wt.uuid = ?", value)
			},
		}).
		SetWheresExtraHasValues(map[string]func([]string, *gorm.DB) *gorm.DB{}).
		SetCtx(ctx).
		GetDbUseQuery("").
		Table("warehouse_cells as wc").
		Joins("join warehouse_tiers wt on wc.warehouse_tier_uuid = wt.uuid").
		Joins("join warehouse_shelves wsh on wt.warehouse_shelf_uuid = wsh.uuid").
		Joins("join warehouse_platoons wp on wsh.warehouse_platoon_uuid = wp.uuid").
		Joins("join warehouse_areas wa on wp.warehouse_area_uuid = wa.uuid").
		Joins("join warehouse_storehouses ws on wa.warehouse_storehouse_uuid = ws.uuid").
		Joins("join organization_workshops ow on ws.organization_workshop_uuid = ow.uuid").
		Joins("left join organization_work_areas owa on ws.organization_work_area_uuid = owa.uuid").
		Joins("join organization_paragraphs op on ow.organization_paragraph_uuid = op.uuid").
		Joins("join organization_railways oa on op.organization_railway_uuid = oa.uuid")
}

// 根据仓库-层获取最后一位
func (WarehouseCellMdl) GetLastByWarehouseTier(warehouseTier *WarehouseTierMdl) (warehouseCell *WarehouseCellMdl) {
	NewWarehouseCellMdl().
		GetDb("").
		Where("warehouse_tier_uuid = ?", warehouseTier.Uuid).
		Order("name desc").
		Order("id desc").
		First(&warehouseCell)
	return
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
		SetWheresEqual("direction_code").
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

// getDirectionText 获取出入库方向描述
func (*WarehouseScanItemMdl) getDirectionText(warehouseScanItem *WarehouseScanItemMdl) (typeText string) {
	for _, item := range warehouseScanItem.GetDirectionCodesMap() {
		if item["code"] == warehouseScanItem.DirectionCode {
			typeText = item["text"]
		}
	}
	return
}

// AfterFind 查询回调
func (receiver *WarehouseScanItemMdl) AfterFind(db *gorm.DB) (err error) {
	receiver.DirectionText = receiver.getDirectionText(receiver)
	return
}

// TableName 仓库-出入库单表名称
func (WarehouseOrderMdl) TableName() string {
	return "warehouse_orders"
}

// NewWarehouseOrderMdl 新建仓库-出入库单模型
func NewWarehouseOrderMdl() *MySqlMdl {
	return NewMySqlMdl().SetModel(WarehouseOrderMdl{})
}

// GetListByQuery 根据Query获取仓库-出入库单列表
func (receiver WarehouseOrderMdl) GetListByQuery(ctx *gin.Context) *gorm.DB {
	return NewWarehouseOrderMdl().
		SetWheresEqual("direction_code").
		SetWheresFuzzy(map[string]string{
			"name": "wo.name like ?",
		}).
		SetWheresDateBetween("wo.created_at", "wo.updated_at", "wo.deleted_at").
		SetWheresExtraHasValue(map[string]func(string, *gorm.DB) *gorm.DB{}).
		SetWheresExtraHasValues(map[string]func([]string, *gorm.DB) *gorm.DB{}).
		SetCtx(ctx).
		GetDbUseQuery("").
		Table("warehouse_orders as wo")
}

// GetDirectionCodes 获取出入库方向代码
func (WarehouseOrderMdl) GetDirectionCodes() []string {
	return []string{"IN", "OUT"}
}

// GetDirectionCodesMap 获取出入库方向代码映射
func (WarehouseOrderMdl) GetDirectionCodesMap() []map[string]string {
	return []map[string]string{
		{"code": "IN", "text": "入库"},
		{"code": "OUT", "text": "出库"},
	}
}

// getDirectionText 获取出入库方向描述
func (*WarehouseOrderMdl) getDirectionText(warehouseOrder *WarehouseOrderMdl) (typeText string) {
	for _, item := range warehouseOrder.GetDirectionCodesMap() {
		if item["code"] == warehouseOrder.DirectionCode {
			typeText = item["text"]
		}
	}
	return
}

// AfterFind 查询回调
func (receiver *WarehouseOrderMdl) AfterFind(db *gorm.DB) (err error) {
	receiver.DirectionText = receiver.getDirectionText(receiver)
	return
}
