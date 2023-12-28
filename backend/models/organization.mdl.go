package models

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type (
	// OrganizationRailwayMdl 组织结构-路局模型
	OrganizationRailwayMdl struct {
		MySqlMdl
		UniqueCode             string                      `gorm:"index;type:char(3);not null;comment:路局代码;" json:"unique_code"`
		Name                   string                      `gorm:"index;type:varchar(128);not null;comment:路局名称;" json:"name"`
		ShortName              string                      `grom:"index;type:varchar(64);not null;comment:路局简称;" json:"short_name"`
		OrganizationParagraphs []*OrganizationParagraphMdl `gorm:"foreignKey:organization_railway_uuid;references:uuid;comment:相关站段;" json:"organization_paragraphs"`
	}
	// OrganizationParagraphMdl 组织结构-站段模型
	OrganizationParagraphMdl struct {
		MySqlMdl
		UniqueCode              string                     `gorm:"index;type:char(4);not null;comment:站段代码;" json:"unique_code"`
		Name                    string                     `gorm:"index;type:varchar(128);not null;comment:站段名称;" json:"name"`
		OrganizationRailwayUuid string                     `gorm:"type:char(36);not null;" json:"organization_railway_uuid"`
		OrganizationRailway     *OrganizationRailwayMdl    `gorm:"foreignKey:organization_railway_uuid;references:uuid;comment:所属路局;" json:"organization_railway"`
		OrganizationWorkshops   []*OrganizationWorkshopMdl `gorm:"foreignKey:organization_paragraph_uuid;references:uuid;comment:相关车间;" json:"organization_workshops"`
		OrganizationLines       []*OrganizationLineMdl     `gorm:"foreignKey:organization_paragraph_uuid;references:uuid;comment:相关线别;" json:"organization_lines"`
	}
	// OrganizationWorkshopMdl 组织结构-车间模型
	OrganizationWorkshopMdl struct {
		MySqlMdl
		UniqueCode                string                      `gorm:"index;type:char(7);not null;comment:车间代码;" json:"unique_code"`
		Name                      string                      `gorm:"index;type:varchar(128);not null;comment:车间名称;" json:"name"`
		TypeCode                  string                      `gorm:"enum('','SCENE','FIX','ELECTRIC','HUMP','ON-BOARD');not null;default:'';comment:车间类型;" json:"type_code"`
		TypeText                  string                      `gorm:"->;type:varchar(128) as ((case type_code when 'SCENE' then '现场车间' when 'FIX' then '检修车间' when 'ELECTRIC' then '电子车间' when 'HUMP' then '驼峰车间' when 'ON-BOARD' then '车载车间' else '无' end))" json:"type_text"`
		OrganizationParagraphUuid string                      `gorm:"type:char(36);not null;" json:"organization_paragraph_uuid"`
		OrganizationParagraph     *OrganizationParagraphMdl   `gorm:"foreignKey:organization_paragraph_uuid;references:uuid;comment:所属站段;" json:"organization_paragraph"`
		OrganizationStations      []*OrganizationStationMdl   `gorm:"foreignKey:organization_workshop_uuid;references:uuid;comment:相关车站;" json:"organization_stations"`
		OrganizationCrossroads    []*OrganizationCrossroadMdl `gorm:"foreignKey:organization_workshop_uuid;references:uuid;comment:相关道口;" json:"organization_crossroads"`
		OrganizationCenters       []*OrganizationCenterMdl    `gorm:"foreignKey:organization_workshop_uuid;references:uuid;comment:相关中心;" json:"organization_centers"`
		OrganizationWorkAreas     []*OrganizationWorkAreaMdl  `gorm:"foreignKey:organization_workshop_uuid;references:uuid;comment:相关工区;" json:"organization_work_areas"`
		EquipmentsByUseBelong     []*EquipmentMdl             `gorm:"foreignKey:use_belong_organizaiton_workshop_uuid;references:uuid;comment:相关器材（使用归属车间）;" json:"equipments_by_use_belong"`
	}
	// OrganizationStationMdl 组织结构-站场模型
	OrganizationStationMdl struct {
		MySqlMdl
		UniqueCode               string                   `gorm:"index;type:char(6);not null;comment:车站代码;" json:"unique_code"`
		Name                     string                   `gorm:"index;type:varchar(128);not null;comment:车站名称;" json:"name"`
		OrganizationWorkshopUuid string                   `gorm:"type:char(36);not null;" json:"organization_workshop_uuid"`
		OrganizationWorkshop     *OrganizationWorkshopMdl `gorm:"foreignKey:organization_workshop_uuid;references:uuid;comment:所属车间;" json:"organization_workshop"`
		OrganizationLineUuid     string                   `gorm:"type:char(36);not null;" json:"organization_line_uuid"`
		OrganizationLine         *OrganizationLineMdl     `gorm:"foreignKey:organization_line_uuid;references:uuid;comment:所属线别;" json:"organization_line"`
		EquipmentsByUseBelong    []*EquipmentMdl          `gorm:"foreignKey:use_belong_organization_station_uuid;references:uuid;comment:相关器材（使用归属站场）;" json:"equipments_by_use_belong"`
	}
	// OrganizationCrossroadMdl 组织结构-道口模型
	OrganizationCrossroadMdl struct {
		MySqlMdl
		UniqueCode               string                   `gorm:"index;type:char(5);not null;comment:道口代码;" json:"unique_code"` // I1642
		Name                     string                   `gorm:"index;type:varchar(128);not null;comment:道口名称;" json:"name"`
		OrganizationWorkshopUuid string                   `gorm:"type:char(36);not null;" json:"organization_workshop_uuid"`
		OrganizationWorkshop     *OrganizationWorkshopMdl `gorm:"foreignKey:organization_workshop_uuid;references:uuid;comment:所属车间;" json:"organization_workshop"`
		OrganizationLineUuid     string                   `gorm:"type:char(36);not null;" json:"organization_line_uuid"`
		OrganizationLine         *OrganizationLineMdl     `gorm:"foreignKey:organization_line_uuid;references:uuid;comment:所属线别;" json:"organization_line"`
		EquipmentsByUseBelong    []*EquipmentMdl          `gorm:"foreignKey:use_belong_organization_crossroad_uuid;references:uuid;comment:相关器材（使用归属道口）;" json:"equipments_by_use_belong"`
	}
	// OrganizationCenterMdl 组织结构-中心模型
	OrganizationCenterMdl struct {
		MySqlMdl
		UniqueCode               string                   `gorm:"index;type:char(6);not null;comment:中心代码;" json:"unique_code"` // A12F34
		Name                     string                   `gorm:"index;type:varchar(128);not null;comment:中心名称;" json:"name"`
		OrganizationWorkshopUuid string                   `gorm:"type:char(36);not null;" json:"organization_workshop_uuid"`
		OrganizationWorkshop     *OrganizationWorkshopMdl `gorm:"foreignKey:organization_workshop_uuid;references:uuid;comment:所属车间;" json:"organization_workshop"`
		OrganizationLineUuid     string                   `gorm:"type:char(36);not null;" json:"organization_line_uuid"`
		OrganizationLine         *OrganizationLineMdl     `gorm:"foreignKey:organization_line_uuid;references:uuid;comment:所属线别;" json:"organization_line"`
		EquipmentsByUseBelong    []*EquipmentMdl          `gorm:"foreignKey:use_belong_organization_center_uuid;references:uuid;comment:相关器材（使用归属中心）;" json:"equipments_by_use_belong"`
	}
	// OrganizationWorkAreaMdl 组织结构-工区模型
	OrganizationWorkAreaMdl struct {
		MySqlMdl
		UniqueCode                 string                   `gorm:"index;type:char(8);not null;comment:工区代码;" json:"unique_code"` // B048D001
		Name                       string                   `gorm:"index;type:varchar(128);not null;comment:工区名称;" json:"name"`
		TypeCode                   string                   `gorm:"type:enum('','SCENE','POINT-SWITCH','REPLY','SYNTHESIZE', 'POWER-SUPPLY-PANEL');not null;default:'';comment:工区类型;" json:"type_code"`
		TypeText                   string                   `gorm:"->;type:varchar(128) as ((case type_code when 'SCENE' then '现场工区' when 'POINT-SWITCH' then '转辙机工区' when 'REPLY' then '继电器工区' when 'SYNTHESIZE' then '综合工区' when 'POWER-SUPPLY-PANEL' then '电源屏工区' else '无' end))" json:"type_text"`
		OrganizationWorkshopUuid   string                   `gorm:"type:char(36);not null;" json:"organization_workshop_uuid"`
		OrganizationWorkshop       *OrganizationWorkshopMdl `gorm:"foreignKey:organization_workshop_uuid;references:uuid;comment:所属车间;" json:"organization_workshop"`
		EquipmentsByPropertyBelong []*EquipmentMdl          `gorm:"foreignKey:property_belong_organization_work_area_uuid;references:uuid;comment:相关器材（资产归属工区）;" json:"equipments_by_property_belong"`
		EquipmentsByUseBelong      []*EquipmentMdl          `gorm:"foreignKey:use_belong_organization_work_area_uuid;references:uuid;comment:相关器材（使用归属工区）;" json:"equipments_by_use_belong"`
	}
	// OrganizationLineMdl 组织结构-线别模型
	OrganizationLineMdl struct {
		MySqlMdl
		UniqueCode                string                      `gorm:"index;type:char(5);not null;comment:线别代码;" json:"unique_code"` // E0001
		Name                      string                      `gorm:"index;type:varchar(128);not null;comment:线别名称;" json:"name"`
		OrganizationParagraphUuid string                      `gorm:"type:char(36);not null;" json:"organization_paragraph_uuid"`
		OrganizationParagraph     *OrganizationParagraphMdl   `gorm:"foreignKey:organization_paragraph_uuid;references:uuid;comment:所属站段;" json:"organization_paragraph"`
		OrganizationStations      []*OrganizationStationMdl   `gorm:"foreignKey:organization_line_uuid;references:uuid;comment:相关站场;" json:"organization_stations"`
		OrganizationCrossroads    []*OrganizationCrossroadMdl `gorm:"foreignKey:organization_line_uuid;references:uuid;comment:相关道口;" json:"organization_crossroads"`
		OrganizationCenters       []*OrganizationCenterMdl    `gorm:"foreignKey:organization_line_uuid;references:uuid;comment:相关中心;" json:"organization_centers"`
		EquipmentsByUseBelong     []*EquipmentMdl             `gorm:"foreignKey:use_belong_organization_line_uuid;references:uuid;comment:相关器材（使用归属线别）;" json:"equipments_by_use_belong"`
	}
)

// TableName 组织结构-路局表名称
func (OrganizationRailwayMdl) TableName() string {
	return "organization_railways"
}

// NewOrganizationRailwayMdl 新建组织结构-路局模型
func NewOrganizationRailwayMdl() *MySqlMdl {
	return NewMySqlMdl().SetModel(OrganizationRailwayMdl{})
}

// GetListByQuery 根据Query获取组织结构-路局列表
func (receiver OrganizationRailwayMdl) GetListByQuery(ctx *gin.Context) *gorm.DB {
	return NewOrganizationRailwayMdl().
		SetWheresEqual().
		SetWheresFuzzy(map[string]string{
			"name": "ora.name like ?",
		}).
		SetWheresDateBetween("ora.created_at", "ora.updated_at", "ora.deleted_at").
		SetWheresExtraHasValue(map[string]func(string, *gorm.DB) *gorm.DB{
			"unique_code": func(value string, db *gorm.DB) *gorm.DB {
				return db.Where("ora.unique_code = ?", value)
			},
		}).
		SetWheresExtraHasValues(map[string]func([]string, *gorm.DB) *gorm.DB{}).
		SetCtx(ctx).
		GetDbUseQuery("").
		Table("organization_railways as ora")
}

// TableName 组织结构-站段表名称
func (OrganizationParagraphMdl) TableName() string {
	return "organization_paragraphs"
}

// NewOrganizationParagraphMdl 新建组织结构-站段模型
func NewOrganizationParagraphMdl() *MySqlMdl {
	return NewMySqlMdl().SetModel(OrganizationParagraphMdl{})
}

// GetListByQuery 根据Query获取组织结构-站段列表
func (receiver OrganizationParagraphMdl) GetListByQuery(ctx *gin.Context) *gorm.DB {
	return NewOrganizationParagraphMdl().
		SetWheresEqual().
		SetWheresFuzzy(map[string]string{
			"name": "op.name like ?",
		}).
		SetWheresDateBetween("op.created_at", "op.updated_at", "op.deleted_at").
		SetWheresExtraHasValue(map[string]func(string, *gorm.DB) *gorm.DB{
			"unique_code": func(value string, db *gorm.DB) *gorm.DB {
				return db.Where("op.unique_code = ?", value)
			},
			"organization_railway_uuid": func(value string, db *gorm.DB) *gorm.DB {
				return db.Where("ora.uuid = ?", value)
			},
		}).
		SetWheresExtraHasValues(map[string]func([]string, *gorm.DB) *gorm.DB{}).
		SetCtx(ctx).
		GetDbUseQuery("").
		Table("organization_paragraphs as op").
		Distinct("op.*").
		Joins("join organization_railways ora on op.organization_railway_uuid = ora.uuid")
}

// TableName 组织结构-车间表名称
func (OrganizationWorkshopMdl) TableName() string {
	return "organization_workshops"
}

// NewOrganizationWorkshopMdl 新建组织结构-车间模型
func NewOrganizationWorkshopMdl() *MySqlMdl {
	return NewMySqlMdl().SetModel(OrganizationWorkshopMdl{})
}

// GetListByQuery 根据Query获取组织结构-车间列表
func (receiver OrganizationWorkshopMdl) GetListByQuery(ctx *gin.Context) *gorm.DB {
	return NewOrganizationWorkshopMdl().
		SetWheresEqual().
		SetWheresFuzzy(map[string]string{
			"name": "ow.name like ?",
		}).
		SetWheresDateBetween("ow.created_at", "ow.updated_at", "ow.deleted_at").
		SetWheresExtraHasValue(map[string]func(string, *gorm.DB) *gorm.DB{
			"unique_code": func(value string, db *gorm.DB) *gorm.DB {
				return db.Where("ow.unique_code = ?", value)
			},
			"organization_paragraph_uuid": func(value string, db *gorm.DB) *gorm.DB {
				return db.Where("op.uuid = ?", value)
			},
			"organization_railway_uuid": func(value string, db *gorm.DB) *gorm.DB {
				return db.Where("ora.uuid = ?", value)
			},
			"type_code": func(value string, db *gorm.DB) *gorm.DB {
				return db.Where("ow.type_code = ?", value)
			},
		}).
		SetWheresExtraHasValues(map[string]func([]string, *gorm.DB) *gorm.DB{}).
		SetCtx(ctx).
		GetDbUseQuery("").
		Table("organization_workshops as ow").
		Distinct("ow.*").
		Joins("join organization_paragraphs op on ow.organization_paragraph_uuid = op.uuid").
		Joins("join organization_railways ora on op.organization_railway_uuid = ora.uuid")
}

// GetTypeCodes 获取车间类型代码列表
func (OrganizationWorkshopMdl) GetTypeCodes() []string {
	return []string{
		"SCENE",
		"FIX",
		"ELECTRIC",
		"HUMP",
		"ON-BOARD",
	}
}

// GetTypeCodesMap 获取车间类型代码映射
func (OrganizationWorkshopMdl) GetTypeCodesMap() []map[string]string {
	return []map[string]string{
		{"code": "SCENE", "text": "现场车间"},
		{"code": "FIX", "text": "检修车间"},
		{"code": "ELECTRIC", "text": "电子车间"},
		{"code": "HUMP", "text": "驼峰车间"},
		{"code": "ON-BOARD", "text": "车载车间"},
	}
}

// TableName 组织结构-站场表名称
func (OrganizationStationMdl) TableName() string {
	return "organization_stations"
}

// NewOrganizationStationMdl 新建组织结构-站场模型
func NewOrganizationStationMdl() *MySqlMdl {
	return NewMySqlMdl().SetModel(OrganizationStationMdl{})
}

// GetListByQuery 根据Query获取组织结构-站场列表
func (receiver OrganizationStationMdl) GetListByQuery(ctx *gin.Context) *gorm.DB {
	return NewOrganizationStationMdl().
		SetWheresEqual().
		SetWheresFuzzy(map[string]string{
			"name": "os.name like ?",
		}).
		SetWheresDateBetween("os.created_at", "os.updated_at", "os.deleted_at").
		SetWheresExtraHasValue(map[string]func(string, *gorm.DB) *gorm.DB{
			"unique_code": func(value string, db *gorm.DB) *gorm.DB {
				return db.Where("os.unique_code = ?", value)
			},
			"organization_workshop_uuid": func(value string, db *gorm.DB) *gorm.DB {
				return db.Where("ow.uuid = ?", value)
			},
			"organization_paragraph_uuid": func(value string, db *gorm.DB) *gorm.DB {
				return db.Where("op.uuid = ?", value)
			},
			"organization_railway_uuid": func(value string, db *gorm.DB) *gorm.DB {
				return db.Where("ora.uuid = ?", value)
			},
			"organization_line_uuid": func(value string, db *gorm.DB) *gorm.DB {
				return db.Where("ol.uuid = ?", value)
			},
		}).
		SetWheresExtraHasValues(map[string]func([]string, *gorm.DB) *gorm.DB{}).
		SetCtx(ctx).
		GetDbUseQuery("").
		Table("organization_stations as os").
		Distinct("os.*").
		Joins("left join organization_lines ol on os.organization_line_uuid = ol.uuid").
		Joins("join organization_workshops ow on os.organization_workshop_uuid = ow.uuid").
		Joins("join organization_paragraphs op on ow.organization_paragraph_uuid = op.uuid").
		Joins("join organization_railways ora on op.organization_railway_uuid = ora.uuid")
}

// TableName 组织结构-道口表名称
func (OrganizationCrossroadMdl) TableName() string {
	return "organization_crossroads"
}

// NewOrganizationCrossroadMdl 新建组织结构-道口模型
func NewOrganizationCrossroadMdl() *MySqlMdl {
	return NewMySqlMdl().SetModel(OrganizationCrossroadMdl{})
}

// GetListByQuery 根据Query获取组织结构-道口列表
func (receiver OrganizationCrossroadMdl) GetListByQuery(ctx *gin.Context) *gorm.DB {
	return NewOrganizationCrossroadMdl().
		SetWheresEqual().
		SetWheresFuzzy(map[string]string{
			"name": "ocr.name like ?",
		}).
		SetWheresDateBetween("ocr.created_at", "ocr.updated_at", "ocr.deleted_at").
		SetWheresExtraHasValue(map[string]func(string, *gorm.DB) *gorm.DB{
			"unique_code": func(value string, db *gorm.DB) *gorm.DB {
				return db.Where("ocr.unique_code = ?", value)
			},
			"organization_workshop_uuid": func(value string, db *gorm.DB) *gorm.DB {
				return db.Where("ow.uuid = ?", value)
			},
			"organization_paragraph_uuid": func(value string, db *gorm.DB) *gorm.DB {
				return db.Where("op.uuid = ?", value)
			},
			"organization_railway_uuid": func(value string, db *gorm.DB) *gorm.DB {
				return db.Where("ora.uuid = ?", value)
			},
			"organization_line_uuid": func(value string, db *gorm.DB) *gorm.DB {
				return db.Where("ol.uuid = ?", value)
			},
		}).
		SetWheresExtraHasValues(map[string]func([]string, *gorm.DB) *gorm.DB{}).
		SetCtx(ctx).
		GetDbUseQuery("").
		Table("organization_crossroads as ocr").
		Distinct("ocr.*").
		Joins("left join organization_lines ol on ocr.organization_line_uuid = ol.uuid").
		Joins("join organization_workshops ow on ocr.organization_workshop_uuid = ow.uuid").
		Joins("join organization_paragraphs op on ow.organization_paragraph_uuid = op.uuid").
		Joins("join organization_railways ora on op.organization_railway_uuid = ora.uuid")
}

// TableName 组织结构-中心表名称
func (OrganizationCenterMdl) TableName() string {
	return "organization_centers"
}

// NewOrganizationCenterMdl 新建组织结构-中心模型
func NewOrganizationCenterMdl() *MySqlMdl {
	return NewMySqlMdl().SetModel(OrganizationCenterMdl{})
}

// GetListByQuery 根据Query获取组织结构-中心列表
func (receiver OrganizationCenterMdl) GetListByQuery(ctx *gin.Context) *gorm.DB {
	return NewOrganizationCenterMdl().
		SetWheresEqual().
		SetWheresFuzzy(map[string]string{
			"name": "oct.name like ?",
		}).
		SetWheresDateBetween("oct.created_at", "oct.updated_at", "oct.deleted_at").
		SetWheresExtraHasValue(map[string]func(string, *gorm.DB) *gorm.DB{
			"unique_code": func(value string, db *gorm.DB) *gorm.DB {
				return db.Where("oct.unique_code = ?", value)
			},
			"organization_workshop_uuid": func(value string, db *gorm.DB) *gorm.DB {
				return db.Where("ow.uuid = ?", value)
			},
			"organization_paragraph_uuid": func(value string, db *gorm.DB) *gorm.DB {
				return db.Where("op.uuid = ?", value)
			},
			"organization_line_uuid": func(value string, db *gorm.DB) *gorm.DB {
				return db.Where("ol.uuid = ?", value)
			},
		}).
		SetWheresExtraHasValues(map[string]func([]string, *gorm.DB) *gorm.DB{}).
		SetCtx(ctx).
		GetDbUseQuery("").
		Table("organization_centers as oct").
		Distinct("oct.*").
		Joins("left join organization_lines ol on oct.organization_line_uuid = ol.uuid").
		Joins("join organization_workshops ow on oct.organization_workshop_uuid = ow.uuid").
		Joins("join organization_paragraphs op on ow.organization_paragraph_uuid = op.uuid").
		Joins("join organization_railways ora on op.organization_railway_uuid = ora.uuid")
}

// TableName 组织结构-工区表名称
func (OrganizationWorkAreaMdl) TableName() string {
	return "organization_work_areas"
}

// NewOrganizationWorkAreaMdl 新建组织结构-工区模型
func NewOrganizationWorkAreaMdl() *MySqlMdl {
	return NewMySqlMdl().SetModel(OrganizationWorkAreaMdl{})
}

// GetListByQuery 根据Query获取组织结构-工区列表
func (receiver OrganizationWorkAreaMdl) GetListByQuery(ctx *gin.Context) *gorm.DB {
	return NewOrganizationWorkAreaMdl().
		SetWheresEqual().
		SetWheresFuzzy(map[string]string{
			"name": "owa.name like ?",
		}).
		SetWheresDateBetween("owa.created_at", "owa.updated_at", "owa.deleted_at").
		SetWheresExtraHasValue(map[string]func(string, *gorm.DB) *gorm.DB{
			"unique_code": func(value string, db *gorm.DB) *gorm.DB {
				return db.Where("owa.unique_code = ?", value)
			},
			"organization_workshop_uuid": func(value string, db *gorm.DB) *gorm.DB {
				return db.Where("ow.uuid = ?", value)
			},
			"organization_paragraph_uuid": func(value string, db *gorm.DB) *gorm.DB {
				return db.Where("op.uuid = ?", value)
			},
			"organization_railway_uuid": func(value string, db *gorm.DB) *gorm.DB {
				return db.Where("ora.uuid = ?", value)
			},
			"type_code": func(value string, d *gorm.DB) *gorm.DB {
				return d.Where("owa.type_code = ?", value)
			},
		}).
		SetWheresExtraHasValues(map[string]func([]string, *gorm.DB) *gorm.DB{}).
		SetCtx(ctx).
		GetDbUseQuery("").
		Table("organization_work_areas as owa").
		Distinct("owa.*").
		Joins("join organization_workshops ow on owa.organization_workshop_uuid = ow.uuid").
		Joins("join organization_paragraphs op on ow.organization_paragraph_uuid = op.uuid").
		Joins("join organization_railways ora on op.organization_railway_uuid = ora.uuid")
}

// GetTypeCodes 获取工区类型代码列表
func (OrganizationWorkAreaMdl) GetTypeCodes() []string {
	return []string{
		"SCENE",
		"POINT-SWITCH",
		"REPLY",
		"SYNTHESIZE",
		"POWER-SUPPLY-PANEL",
	}
}

// GetTypeCodesMap 获取工区类型代码映射
func (OrganizationWorkAreaMdl) GetTypeCodesMap() []map[string]string {
	return []map[string]string{
		{"code": "SCENE", "text": "现场工区"},
		{"code": "POINT-SWITCH", "text": "转辙机工区"},
		{"code": "REPLY", "text": "继电器工区"},
		{"code": "SYNTHESIZE", "text": "综合工区"},
		{"code": "POWER-SUPPLY-PANEL", "text": "电源屏工区"},
	}
}

// TableName 组织结构-线别表名称
func (OrganizationLineMdl) TableName() string {
	return "organization_lines"
}

// NewOrganizationLineMdl 新建组织结构-线别模型
func NewOrganizationLineMdl() *MySqlMdl {
	return NewMySqlMdl().SetModel(OrganizationLineMdl{})
}

// GetListByQuery 根据Query获取组织结构-线别列表
func (receiver OrganizationLineMdl) GetListByQuery(ctx *gin.Context) *gorm.DB {
	return NewOrganizationLineMdl().
		SetWheresEqual().
		SetWheresFuzzy(map[string]string{
			"name": "ol.name like ?",
		}).
		SetWheresDateBetween("ol.created_at", "ol.updated_at", "ol.deleted_at").
		SetWheresExtraHasValue(map[string]func(string, *gorm.DB) *gorm.DB{
			"unique_code": func(value string, db *gorm.DB) *gorm.DB {
				return db.Where("ol.unique_code = ?", value)
			},
			"organization_station_uuid": func(value string, db *gorm.DB) *gorm.DB {
				return db.Where("os.uuid = ?", value)
			},
			"organization_crossroad_uuid": func(value string, db *gorm.DB) *gorm.DB {
				return db.Where("ocr.uuid = ?", value)
			},
			"organization_center_uuid": func(value string, db *gorm.DB) *gorm.DB {
				return db.Where("oct.uuid = ?", value)
			},
			"organization_paragraph_uuid": func(value string, db *gorm.DB) *gorm.DB {
				return db.Where("op.uuid = ?", value)
			},
		}).
		SetWheresExtraHasValues(map[string]func([]string, *gorm.DB) *gorm.DB{}).
		SetCtx(ctx).
		GetDbUseQuery("").
		Table("organization_lines as ol").
		Distinct("ol.*").
		Joins("left join organization_stations os on ol.uuid = os.organization_line_uuid").
		Joins("left join organization_crossroads ocr on ol.uuid = ocr.organization_line_uuid").
		Joins("left join organization_centers oct on ol.uuid = oct.organization_line_uuid").
		Joins("left join organization_paragraphs op on ol.organization_paragraph_uuid = op.uuid")
}

// 线别绑定站场
func (receiver OrganizationLineMdl) BindOrganizationStations(organizationStations []*OrganizationStationMdl) {
	// 解绑所有已绑定的站场
	NewOrganizationStationMdl().GetDb("").Where("organization_line_uuid = ?", receiver.Uuid).Update("organization_line_uuid", "")

	if len(organizationStations) > 0 {
		for _, organizationStation := range organizationStations {
			organizationStation.OrganizationLineUuid = receiver.Uuid
			NewOrganizationStationMdl().GetDb("").Where("uuid = ?", organizationStation.Uuid).Update("organization_line_uuid", receiver.Uuid)
		}
	}
}

// 线别绑定站场
func (receiver OrganizationLineMdl) BindOrganizationCrossroads(organizationCrossroads []*OrganizationCrossroadMdl) {
	// 解绑所有已绑定的站场
	NewOrganizationCrossroadMdl().GetDb("").Where("organization_line_uuid = ?", receiver.Uuid).Update("organization_line_uuid", "")

	if len(organizationCrossroads) > 0 {
		for _, organizationCrossroad := range organizationCrossroads {
			organizationCrossroad.OrganizationLineUuid = receiver.Uuid
			NewOrganizationCrossroadMdl().GetDb("").Where("uuid = ?", organizationCrossroad.Uuid).Update("organization_line_uuid", receiver.Uuid)
		}
	}
}

// 线别绑定中心
func (receiver OrganizationLineMdl) BindOrganizationCenters(organizationCenters []*OrganizationCenterMdl) {
	// 解绑所有已绑定的中心
	NewOrganizationCenterMdl().GetDb("").Where("organization_line_uuid = ?", receiver.Uuid).Update("organization_line_uuid", "")

	if len(organizationCenters) > 0 {
		for _, organizationCenter := range organizationCenters {
			organizationCenter.OrganizationLineUuid = receiver.Uuid
			NewOrganizationCenterMdl().GetDb("").Where("uuid = ?", organizationCenter.Uuid).Update("organization_line_uuid", receiver.Uuid)
		}
	}
}
