package models

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type (
	// InstallIndoorRoomTypeMdl 室内上道位置-机房类型模型
	InstallIndoorRoomTypeMdl struct {
		MySqlMdl
		UniqueCode         string                  `gorm:"unique;type:char(2);not null;comment:代码;" json:"unique_code"`
		Name               string                  `gorm:"unique;type:varchar(64);not null;comment:名称;" json:"name"`
		InstallIndoorRooms []*InstallIndoorRoomMdl `gorm:"foreignKey:install_indoor_room_type_uuid;references:uuid;comment:所属室内上道位置-机房类型;" json:"install_indoor_rooms"`
	}

	// InstallIndoorRoomMdl 室内上道位置-机房模型
	InstallIndoorRoomMdl struct {
		MySqlMdl
		Name                      string                     `gorm:"type:varchar(64);not null;comment:名称;" json:"name"`
		InstallIndoorRoomTypeUuid string                     `gorm:"index;char(36);not null;default:;comment:所属室内上道位置-机房类型uuid;" json:"install_indoor_room_type_uuid"`
		InstallIndoorRoomType     *InstallIndoorRoomMdl      `gorm:"foreignKey:install_indoor_room_type_uuid;references:uuid;comment:所属室内上道位置-机房类型;" json:"install_indoor_room_type"`
		OrganizationStationUuid   string                     `gorm:"index;type:char(36);not null;default:;comment:所属组织结构-站场uuid;" json:"organization_station_uuid"`
		OrganizationStation       *OrganizationStationMdl    `gorm:"foreignKey:organization_station_uuid;references:uuid;comment:所属组织结构-站场;" json:"organization_station"`
		OrganizationCrossroadUuid string                     `gorm:"index;type:char(36);not null;default:;comment:所属组织结构-道口uuid;" json:"organization_crossroad_uuid"`
		OrganizationCrossroad     *OrganizationCrossroadMdl  `gorm:"foreignKey:organization_crossroad_uuid;references:uuid;comment:所属组织结构-道口;" json:"organization_crossroad"`
		OrganizationCenterUuid    string                     `gorm:"index;type:char(36);not null;default:;comment:所属组织结构-中心uuid;" json:"organization_center_uuid"`
		OrganizationCenter        *OrganizationCenterMdl     `gorm:"foreignKey:organization_center_uuid;references:uuid;comment:所属组织结构-中心;" json:"organization_center"`
		InstallIndoorPlatoons     []*InstallIndoorPlatoonMdl `gorm:"foreignKey:install_indoor_room_uuid;references:uuid;comment:相关室内上道位置-排;" json:"install_indoor_platoons"`
	}

	// InstallIndoorPlatoonMdl 室内上道位置-排模型
	InstallIndoorPlatoonMdl struct {
		MySqlMdl
		Name                  string                   `gorm:"type:varchar(64);not null;comment:名称;" json:"name"`
		InstallIndoorRoomUuid string                   `gorm:"index;type:char(36);not null;default:;comment:所属室内上道位置-机房uuid;" json:"install_indoor_room_uuid"`
		InstallIndoorRoom     *InstallIndoorRoomMdl    `gorm:"foreignKey:install_indoor_room_uuid;references:uuid;comment:所属室内上道位置-机房;" json:"install_indoor_room"`
		InstallIndoorShelves  []*InstallIndoorShelfMdl `gorm:"foreignKey:install_indoor_platoon_uuid;references:uuid;comment:相关室内上道位置-架;" json:"install_indoor_shelves"`
	}

	// InstallIndoorShelfMdl 室内上道位置-架模型
	InstallIndoorShelfMdl struct {
		MySqlMdl
		Name                     string                   `gorm:"type:varchar(64);not null;comment:名称;" json:"name"`
		InstallIndoorPlatoonUuid string                   `gorm:"index;type:char(36);not null;default:;comment:所属室内上道位置-排uuid;" json:"install_indoor_platoon_uuid"`
		InstallIndoorPlatoon     *InstallIndoorPlatoonMdl `gorm:"foreignKey:install_indoor_platoon_uuid;references:uuid;comment:所属室内上道位置-排;" json:"install_indoor_platoon"`
		InstallIndoorTiers       []*InstallIndoorTierMdl  `gorm:"foreignKey:install_indoor_shelf_uuid;references:uuid;comment:相关室内上道位置-层;" json:"install_indoor_tiers"`
	}

	// InstallIndoorTierMdl 室内上道位置-层模型
	InstallIndoorTierMdl struct {
		MySqlMdl
		Name                   string                  `gorm:"type:varchar(64);not null;comment:名称;" json:"name"`
		InstallIndoorShelfUuid string                  `gorm:"index;type:char(36);not null;default:;comment:所属室内上道位置-架uuid;" json:"install_indoor_shelf_uuid"`
		InstallIndoorShelf     *InstallIndoorShelfMdl  `gorm:"foreignKey:install_indoor_shelf_uuid;references:uuid;comment:所属室内上道位置-架;" json:"install_indoor_shelf"`
		InstallIndoorCells     []*InstallIndoorCellMdl `gorm:"foreignKey:install_indoor_tier_uuid;references:uuid;comment:相关室内上道位置-位;" json:"install_indoor_cells"`
	}

	// InstallIndoorCellMdl 室内上道位置-位模型
	InstallIndoorCellMdl struct {
		MySqlMdl
		Name                  string                `gorm:"type:varchar(64);not null;comment:名称;" json:"name"`
		InstallIndoorTierUuid string                `gorm:"index;type:char(36);not null;default:;comment:所属室内上道位置-层uuid;" json:"install_indoor_tier_uuid"`
		InstallIndoorTier     *InstallIndoorTierMdl `gorm:"foreignKey:install_indoor_tier_uuid;references:uuid;comment:所属室内上道位置-层;" json:"install_indoor_tier"`
		Volumn                uint64                `gorm:"type:bigint(20);not null;default:1;comment:容量;" json:"volumn"`
		Equipments            []*EquipmentMdl       `gorm:"foreignKey:install_indoor_cell_uuid;references:uuid;comment:相关器材;" json:"equipments"`
	}
)

// TableName 室内上道位置-机房类型表名称
func (InstallIndoorRoomTypeMdl) TableName() string {
	return "install_indoor_room_types"
}

// NewInstallIndoorRoomTypeMdl 新建室内上道位置-机房类型模型
func NewInstallIndoorRoomTypeMdl() *MySqlMdl {
	return NewMySqlMdl().SetModel(InstallIndoorRoomTypeMdl{})
}

// GetListByQuery 根据Query获取室内上道位置-机房类型列表
func (receiver InstallIndoorRoomTypeMdl) GetListByQuery(ctx *gin.Context) *gorm.DB {
	return NewInstallIndoorRoomTypeMdl().
		SetWheresEqual().
		SetWheresFuzzy(map[string]string{
			"name": "iirt.name like ?",
		}).
		SetWheresDateBetween("iirt.created_at", "iirt.updated_at", "iirt.deleted_at").
		SetWheresExtraHasValue(map[string]func(string, *gorm.DB) *gorm.DB{}).
		SetWheresExtraHasValues(map[string]func([]string, *gorm.DB) *gorm.DB{}).
		SetCtx(ctx).
		GetDbUseQuery("").
		Table("install_indoor_room_types as iirt")
}

// TableName 室内上道位置-机房表名称
func (InstallIndoorRoomMdl) TableName() string {
	return "install_indoor_rooms"
}

// NewInstallIndoorRoomMdl 新建室内上道位置-机房模型
func NewInstallIndoorRoomMdl() *MySqlMdl {
	return NewMySqlMdl().SetModel(InstallIndoorRoomMdl{})
}

// GetListByQuery 根据Query获取室内上道位置-机房列表
func (receiver InstallIndoorRoomMdl) GetListByQuery(ctx *gin.Context) *gorm.DB {
	return NewInstallIndoorRoomMdl().
		SetWheresEqual().
		SetWheresFuzzy(map[string]string{
			"name": "iir.name like ?",
		}).
		SetWheresDateBetween("iir.created_at", "iir.updated_at", "iir.deleted_at").
		SetWheresExtraHasValue(map[string]func(string, *gorm.DB) *gorm.DB{
			"organization_workshop_uuid": func(value string, db *gorm.DB) *gorm.DB {
				return db.Where("ows.uuid = ? or owcr.uuid = ? or owc.uuid = ?", value)
			},
			"organization_station_uuid": func(value string, db *gorm.DB) *gorm.DB {
				return db.Where("os.uuid = ?", value)
			},
			"organization_crossroad_uuid": func(value string, db *gorm.DB) *gorm.DB {
				return db.Where("ocr.uuid = ?", value)
			},
			"organization_center_uuid": func(value string, db *gorm.DB) *gorm.DB {
				return db.Where("oc.uuid = ?", value)
			},
			"install_indoor_room_type_uuid": func(value string, db *gorm.DB) *gorm.DB {
				return db.Where("iirt.uuid = ?", value)
			},
		}).
		SetWheresExtraHasValues(map[string]func([]string, *gorm.DB) *gorm.DB{}).
		SetCtx(ctx).
		GetDbUseQuery("").
		Table("install_indoor_rooms as iir").
		Joins("join install_indoor_room_types iirt on iirt.uuid = iir.install_indoor_room_type_uuid").
		Joins("join organization_stations os on os.uuid = iir.organization_station_uuid").
		Joins("join organization_crossroads ocr on ocr.uuid = iir.organization_crossroad_uuid").
		Joins("join organization_centers oc on oc.uuid = iir.organization_center_uuid").
		Joins("left join organization_workshops ows on ows.uuid = os.organization_workshop_uuid").
		Joins("left join organization_workshops owcr on owcr.uuid = ocr.organization_workshop_uuid").
		Joins("left join organization_workshops owc on owc.uuid = oc.organization_workshop_uuid")
}

// TableName 室内上道位置-排表名称
func (InstallIndoorPlatoonMdl) TableName() string {
	return "install_indoor_platoons"
}

// NewInstallIndoorPlatoonMdl 新建室内上道位置-排模型
func NewInstallIndoorPlatoonMdl() *MySqlMdl {
	return NewMySqlMdl().SetModel(InstallIndoorPlatoonMdl{})
}

// GetListByQuery 根据Query获取室内上道位置-排列表
func (receiver InstallIndoorPlatoonMdl) GetListByQuery(ctx *gin.Context) *gorm.DB {
	return NewInstallIndoorPlatoonMdl().
		SetWheresEqual().
		SetWheresFuzzy(map[string]string{
			"name": "iip.name like ?",
		}).
		SetWheresDateBetween("iip.created_at", "iip.updated_at", "iip.deleted_at").
		SetWheresExtraHasValue(map[string]func(string, *gorm.DB) *gorm.DB{
			"organization_workshop_uuid": func(value string, db *gorm.DB) *gorm.DB {
				return db.Where("ows.uuid = ? or owcr.uuid = ? or owc.uuid = ?", value)
			},
			"organization_station_uuid": func(value string, db *gorm.DB) *gorm.DB {
				return db.Where("os.uuid = ?", value)
			},
			"organization_crossroad_uuid": func(value string, db *gorm.DB) *gorm.DB {
				return db.Where("ocr.uuid = ?", value)
			},
			"organization_center_uuid": func(value string, db *gorm.DB) *gorm.DB {
				return db.Where("oc.uuid = ?", value)
			},
			"install_indoor_room_type_uuid": func(value string, db *gorm.DB) *gorm.DB {
				return db.Where("iirt.uuid = ?", value)
			},
			"install_indoor_room_uuid": func(value string, db *gorm.DB) *gorm.DB {
				return db.Where("iir.uuid = ?", value)
			},
		}).
		SetWheresExtraHasValues(map[string]func([]string, *gorm.DB) *gorm.DB{}).
		SetCtx(ctx).
		GetDbUseQuery("").
		Table("install_indoor_platoons as iip").
		Joins("join install_indoor_rooms iir on iir.uuid = iip.install_indoor_room_uuid").
		Joins("join install_indoor_room_types iirt on iirt.uuid = iir.install_indoor_room_type_uuid").
		Joins("join organization_stations os on os.uuid = iir.organization_station_uuid").
		Joins("join organization_crossroads ocr on ocr.uuid = iir.organization_crossroad_uuid").
		Joins("join organization_centers oc on oc.uuid = iir.organization_center_uuid").
		Joins("left join organization_workshops ows on ows.uuid = os.organization_workshop_uuid").
		Joins("left join organization_workshops owcr on owcr.uuid = ocr.organization_workshop_uuid").
		Joins("left join organization_workshops owc on owc.uuid = oc.organization_workshop_uuid")
}

// TableName 室内上道位置-架表名称
func (InstallIndoorShelfMdl) TableName() string {
	return "install_indoor_shelves"
}

// NewInstallIndoorShelfMdl 新建室内上道位置-架模型
func NewInstallIndoorShelfMdl() *MySqlMdl {
	return NewMySqlMdl().SetModel(InstallIndoorShelfMdl{})
}

// GetListByQuery 根据Query获取室内上道位置-架列表
func (receiver InstallIndoorShelfMdl) GetListByQuery(ctx *gin.Context) *gorm.DB {
	return NewInstallIndoorShelfMdl().
		SetWheresEqual().
		SetWheresFuzzy(map[string]string{
			"name": "iis.name like ?",
		}).
		SetWheresDateBetween("iis.created_at", "iis.updated_at", "iis.deleted_at").
		SetWheresExtraHasValue(map[string]func(string, *gorm.DB) *gorm.DB{
			"organization_workshop_uuid": func(value string, db *gorm.DB) *gorm.DB {
				return db.Where("ows.uuid = ? or owcr.uuid = ? or owc.uuid = ?", value)
			},
			"organization_station_uuid": func(value string, db *gorm.DB) *gorm.DB {
				return db.Where("os.uuid = ?", value)
			},
			"organization_crossroad_uuid": func(value string, db *gorm.DB) *gorm.DB {
				return db.Where("ocr.uuid = ?", value)
			},
			"organization_center_uuid": func(value string, db *gorm.DB) *gorm.DB {
				return db.Where("oc.uuid = ?", value)
			},
			"install_indoor_room_type_uuid": func(value string, db *gorm.DB) *gorm.DB {
				return db.Where("iirt.uuid = ?", value)
			},
			"install_indoor_room_uuid": func(value string, db *gorm.DB) *gorm.DB {
				return db.Where("iir.uuid = ?", value)
			},
			"install_indoor_platoon_uuid": func(value string, db *gorm.DB) *gorm.DB {
				return db.Where("iip.uuid = ?", value)
			},
		}).
		SetWheresExtraHasValues(map[string]func([]string, *gorm.DB) *gorm.DB{}).
		SetCtx(ctx).
		GetDbUseQuery("").
		Table("install_indoor_shelves as iis").
		Joins("join install_indoor_platoons iip on iip.uuid = iis.install_indoor_platoon_uuid").
		Joins("join install_indoor_rooms iir on iir.uuid = iip.install_indoor_room_uuid").
		Joins("join install_indoor_room_types iirt on iirt.uuid = iir.install_indoor_room_type_uuid").
		Joins("join organization_stations os on os.uuid = iir.organization_station_uuid").
		Joins("join organization_crossroads ocr on ocr.uuid = iir.organization_crossroad_uuid").
		Joins("join organization_centers oc on oc.uuid = iir.organization_center_uuid").
		Joins("left join organization_workshops ows on ows.uuid = os.organization_workshop_uuid").
		Joins("left join organization_workshops owcr on owcr.uuid = ocr.organization_workshop_uuid").
		Joins("left join organization_workshops owc on owc.uuid = oc.organization_workshop_uuid")
}

// TableName 室内上道位置-层表名称
func (InstallIndoorTierMdl) TableName() string {
	return "install_indoor_tiers"
}

// NewInstallIndoorTierMdl 新建室内上道位置-层模型
func NewInstallIndoorTierMdl() *MySqlMdl {
	return NewMySqlMdl().SetModel(InstallIndoorTierMdl{})
}

// GetListByQuery 根据Query获取室内上道位置-层列表
func (receiver InstallIndoorTierMdl) GetListByQuery(ctx *gin.Context) *gorm.DB {
	return NewInstallIndoorTierMdl().
		SetWheresEqual().
		SetWheresFuzzy(map[string]string{
			"name": "iit.name like ?",
		}).
		SetWheresDateBetween("iit.created_at", "iit.updated_at", "iit.deleted_at").
		SetWheresExtraHasValue(map[string]func(string, *gorm.DB) *gorm.DB{
			"organization_workshop_uuid": func(value string, db *gorm.DB) *gorm.DB {
				return db.Where("ows.uuid = ? or owcr.uuid = ? or owc.uuid = ?", value)
			},
			"organization_station_uuid": func(value string, db *gorm.DB) *gorm.DB {
				return db.Where("os.uuid = ?", value)
			},
			"organization_crossroad_uuid": func(value string, db *gorm.DB) *gorm.DB {
				return db.Where("ocr.uuid = ?", value)
			},
			"organization_center_uuid": func(value string, db *gorm.DB) *gorm.DB {
				return db.Where("oc.uuid = ?", value)
			},
			"install_indoor_room_type_uuid": func(value string, db *gorm.DB) *gorm.DB {
				return db.Where("iirt.uuid = ?", value)
			},
			"install_indoor_room_uuid": func(value string, db *gorm.DB) *gorm.DB {
				return db.Where("iir.uuid = ?", value)
			},
			"install_indoor_platoon_uuid": func(value string, db *gorm.DB) *gorm.DB {
				return db.Where("iip.uuid = ?", value)
			},
			"install_indoor_shelf_uuid": func(value string, db *gorm.DB) *gorm.DB {
				return db.Where("iis.uuid = ?", value)
			},
		}).
		SetWheresExtraHasValues(map[string]func([]string, *gorm.DB) *gorm.DB{}).
		SetCtx(ctx).
		GetDbUseQuery("").
		Table("install_indoor_tiers as iit").
		Joins("join install_indoor_shelves iis on iis.uuid = iit.install_indoor_shelf_uuid").
		Joins("join install_indoor_platoons iip on iip.uuid = iis.install_indoor_platoon_uuid").
		Joins("join install_indoor_rooms iir on iir.uuid = iip.install_indoor_room_uuid").
		Joins("join install_indoor_room_types iirt on iirt.uuid = iir.install_indoor_room_type_uuid").
		Joins("join organization_stations os on os.uuid = iir.organization_station_uuid").
		Joins("join organization_crossroads ocr on ocr.uuid = iir.organization_crossroad_uuid").
		Joins("join organization_centers oc on oc.uuid = iir.organization_center_uuid").
		Joins("left join organization_workshops ows on ows.uuid = os.organization_workshop_uuid").
		Joins("left join organization_workshops owcr on owcr.uuid = ocr.organization_workshop_uuid").
		Joins("left join organization_workshops owc on owc.uuid = oc.organization_workshop_uuid")
}

// TableName 室内上道位置-位表名称
func (InstallIndoorCellMdl) TableName() string {
	return "install_indoor_cells"
}

// NewInstallIndoorCellMdl 新建室内上道位置-位模型
func NewInstallIndoorCellMdl() *MySqlMdl {
	return NewMySqlMdl().SetModel(InstallIndoorCellMdl{})
}

// GetListByQuery 根据Query获取室内上道位置-位列表
func (receiver InstallIndoorCellMdl) GetListByQuery(ctx *gin.Context) *gorm.DB {
	return NewInstallIndoorCellMdl().
		SetWheresEqual().
		SetWheresFuzzy(map[string]string{
			"name": "iic.name like ?",
		}).
		SetWheresDateBetween("iic.created_at", "iic.updated_at", "iic.deleted_at").
		SetWheresExtraHasValue(map[string]func(string, *gorm.DB) *gorm.DB{
			"organization_workshop_uuid": func(value string, db *gorm.DB) *gorm.DB {
				return db.Where("ows.uuid = ? or owcr.uuid = ? or owc.uuid = ?", value)
			},
			"organization_station_uuid": func(value string, db *gorm.DB) *gorm.DB {
				return db.Where("os.uuid = ?", value)
			},
			"organization_crossroad_uuid": func(value string, db *gorm.DB) *gorm.DB {
				return db.Where("ocr.uuid = ?", value)
			},
			"organization_center_uuid": func(value string, db *gorm.DB) *gorm.DB {
				return db.Where("oc.uuid = ?", value)
			},
			"install_indoor_room_type_uuid": func(value string, db *gorm.DB) *gorm.DB {
				return db.Where("iirt.uuid = ?", value)
			},
			"install_indoor_room_uuid": func(value string, db *gorm.DB) *gorm.DB {
				return db.Where("iir.uuid = ?", value)
			},
			"install_indoor_platoon_uuid": func(value string, db *gorm.DB) *gorm.DB {
				return db.Where("iip.uuid = ?", value)
			},
			"install_indoor_shelf_uuid": func(value string, db *gorm.DB) *gorm.DB {
				return db.Where("iis.uuid = ?", value)
			},
			"install_indoor_tier_uuid": func(value string, db *gorm.DB) *gorm.DB {
				return db.Where("iit.uuid = ?", value)
			},
		}).
		SetWheresExtraHasValues(map[string]func([]string, *gorm.DB) *gorm.DB{}).
		SetCtx(ctx).
		GetDbUseQuery("").
		Table("install_indoor_cells as iic").
		Joins("join install_indoor_tiers iit on iit.uuid = iic.install_indoor_tier_uuid").
		Joins("join install_indoor_shelves iis on iis.uuid = iit.install_indoor_shelf_uuid").
		Joins("join install_indoor_platoons iip on iip.uuid = iis.install_indoor_platoon_uuid").
		Joins("join install_indoor_rooms iir on iir.uuid = iip.install_indoor_room_uuid").
		Joins("join install_indoor_room_types iirt on iirt.uuid = iir.install_indoor_room_type_uuid").
		Joins("join organization_stations os on os.uuid = iir.organization_station_uuid").
		Joins("join organization_crossroads ocr on ocr.uuid = iir.organization_crossroad_uuid").
		Joins("join organization_centers oc on oc.uuid = iir.organization_center_uuid").
		Joins("left join organization_workshops ows on ows.uuid = os.organization_workshop_uuid").
		Joins("left join organization_workshops owcr on owcr.uuid = ocr.organization_workshop_uuid").
		Joins("left join organization_workshops owc on owc.uuid = oc.organization_workshop_uuid")
}
