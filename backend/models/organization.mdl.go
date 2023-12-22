package models

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// OrganizationRailwayMdl 组织结构-路局模型
type OrganizationRailwayMdl struct {
	MySqlMdl
	UniqueCode             string                      `gorm:"unique;char(3);not null;comment:路局代码;" json:"unique_code"`
	Name                   string                      `gorm:"unique;varchar(128);not null;comment:路局名称;" json:"name"`
	OrganizationParagraphs []*OrganizationParagraphMdl `gorm:"foreignKey:organization_railway_uuid;references:uuid;comment:相关站段;" json:"organization_paragraphs"`
}

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
			"name": "or.name like ?",
		}).
		SetWheresDateBetween("or.created_at", "or.updated_at", "or.deleted_at").
		SetWheresExtraHasValue(map[string]func(string, *gorm.DB) *gorm.DB{}).
		SetWheresExtraHasValues(map[string]func([]string, *gorm.DB) *gorm.DB{}).
		SetCtx(ctx).
		GetDbUseQuery("").
		Table("organization_railways as or")
}

// OrganizationParagraphMdl 组织结构-站段模型
type OrganizationParagraphMdl struct {
	MySqlMdl
	UniqueCode              string                     `gorm:"unique;char(4);not null;comment:站段代码;" json:"unique_code"`
	Name                    string                     `gorm:"unique;varchar(128);not null;comment:站段名称;" json:"name"`
	OrganizationRailwayUuid string                     `gorm:"type:char(36);not null;" json:"organization_railway_uuid"`
	OrganizationRailway     *OrganizationRailwayMdl    `gorm:"foreignKey:organization_railway_uuid;references:uuid;comment:所属路局;" json:"organization_railway"`
	OrganizationWorkshops   []*OrganizationWorkshopMdl `gorm:"foreignKey:organization_paragraph_uuid;references:uuid;comment:相关车间;" json:"organization_workshops"`
	OrganizationLines       []*OrganizationLineMdl     `gorm:"foreignKey:organization_paragraph_uuid;references:uuid;comment:相关线别;" json:"organization_lines"`
}

// TableName 组织结构-站段表名称
func (OrganizationParagraphMdl) TableName() string {
	return "oraganizaiton_paragraphs"
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
		SetWheresExtraHasValue(map[string]func(string, *gorm.DB) *gorm.DB{}).
		SetWheresExtraHasValues(map[string]func([]string, *gorm.DB) *gorm.DB{}).
		SetCtx(ctx).
		GetDbUseQuery("").
		Table("oraganizaiton_paragraphs as op")
}

// OrganizationWorkshopMdl 组织结构-车间模型
type OrganizationWorkshopMdl struct {
	MySqlMdl
	UniqueCode                string                      `gorm:"unique;char(7);not null;comment:车间代码;" json:"unique_code"`
	Name                      string                      `gorm:"unique;varchar(128);not null;comment:车间名称;" json:"name"`
	OrganizationParagraphUuid string                      `gorm:"type:char(36);not null;" json:"organization_paragraph_uuid"`
	OrganizationParagraph     *OrganizationParagraphMdl   `gorm:"foreignKey:organization_paragraph_uuid;references:uuid;comment:所属站段;" json:"organization_paragraph"`
	OrganizationStations      []*OrganizationStationMdl   `gorm:"foreignKey:organization_workshop_uuid;references:uuid;comment:相关车站;" json:"organization_stations"`
	OrganizationCrossroads    []*OrganizationCrossroadMdl `gorm:"foreignKey:organization_workshop_uuid;references:uuid;comment:相关道口;" json:"organization_crossroads"`
	OrganizationCenters       []*OrganizationCenterMdl    `gorm:"foreignKey:organization_workshop_uuid;references:uuid;comment:相关中心;" json:"organization_centers"`
	OrganizationWorkAreas     []*OrganizationWorkAreaMdl  `gorm:"foreignKey:organization_workshop_uuid;references:uuid;comment:相关工区;" json:"organization_work_areas"`
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
		SetWheresExtraHasValue(map[string]func(string, *gorm.DB) *gorm.DB{}).
		SetWheresExtraHasValues(map[string]func([]string, *gorm.DB) *gorm.DB{}).
		SetCtx(ctx).
		GetDbUseQuery("").
		Table("organization_workshops as ow")
}

// OrganizationStationMdl 组织结构-站场模型
type OrganizationStationMdl struct {
	MySqlMdl
	UniqueCode               string                   `gorm:"unique;char(8);not null;comment:车站代码;" json:"unique_code"`
	Name                     string                   `gorm:"unique;varchar(128);not null;comment:车站名称;" json:"name"`
	OrganizationWorkshopUuid string                   `gorm:"type:char(36);not null;" json:"organization_workshop_uuid"`
	OrganizationWorkshop     *OrganizationWorkshopMdl `gorm:"foreignKey:organization_workshop_uuid;references:uuid;comment:所属车间;" json:"organization_workshop"`
	OrganizationLineUuid     string                   `gorm:"type:char(36);not null;" json:"organization_line_uuid"`
	OrganizationLine         *OrganizationLineMdl     `gorm:"foreignKey:organization_line_uuid;references:uuid;comment:所属线别;" json:"organization_line"`
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
		SetWheresExtraHasValue(map[string]func(string, *gorm.DB) *gorm.DB{}).
		SetWheresExtraHasValues(map[string]func([]string, *gorm.DB) *gorm.DB{}).
		SetCtx(ctx).
		GetDbUseQuery("").
		Table("organization_stations as os")
}

// OrganizationCrossroadMdl 组织结构-道口模型
type OrganizationCrossroadMdl struct {
	MySqlMdl
	UniqueCode               string                   `gorm:"unique;char(5);not null;comment:道口代码;" json:"unique_code"` // I1642
	Name                     string                   `gorm:"unique;varchar(128);not null;comment:道口名称;" json:"name"`
	OrganizationWorkshopUuid string                   `gorm:"type:char(36);not null;" json:"organization_workshop_uuid"`
	OrganizationWorkshop     *OrganizationWorkshopMdl `gorm:"foreignKey:organization_workshop_uuid;references:uuid;comment:所属车间;" json:"organization_workshop"`
	OrganizationLineUuid     string                   `gorm:"type:char(36);not null;" json:"organization_line_uuid"`
	OrganizationLine         *OrganizationLineMdl     `gorm:"foreignKey:organization_line_uuid;references:uuid;comment:所属线别;" json:"organization_line"`
}

// TableName 组织结构-道口表名称
func (OrganizationCrossroadMdl) TableName() string {
	return "oraganization_crossroads"
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
		SetWheresExtraHasValue(map[string]func(string, *gorm.DB) *gorm.DB{}).
		SetWheresExtraHasValues(map[string]func([]string, *gorm.DB) *gorm.DB{}).
		SetCtx(ctx).
		GetDbUseQuery("").
		Table("oraganization_crossroads as ocr")
}

// OrganizationCenterMdl 组织结构-中心模型
type OrganizationCenterMdl struct {
	MySqlMdl
	UniqueCode               string                   `gorm:"unique;char(6);not null;comment:中心代码;" json:"unique_code"` // A12F34
	Name                     string                   `gorm:"unique;varchar(128);not null;comment:中心名称;" json:"name"`
	OrganizationWorkshopUuid string                   `gorm:"type:char(36);not null;" json:"organization_workshop_uuid"`
	OrganizationWorkshop     *OrganizationWorkshopMdl `gorm:"foreignKey:organization_workshop_uuid;references:uuid;comment:所属车间;" json:"organization_workshop"`
	OrganizationLineUuid     string                   `gorm:"type:char(36);not null;" json:"organization_line_uuid"`
	OrganizationLine         *OrganizationLineMdl     `gorm:"foreignKey:organization_line_uuid;references:uuid;comment:所属线别;" json:"organization_line"`
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
		SetWheresExtraHasValue(map[string]func(string, *gorm.DB) *gorm.DB{}).
		SetWheresExtraHasValues(map[string]func([]string, *gorm.DB) *gorm.DB{}).
		SetCtx(ctx).
		GetDbUseQuery("").
		Table("organization_centers as oct")
}

// OrganizationWorkAreaMdl 组织结构-工区模型
type OrganizationWorkAreaMdl struct {
	MySqlMdl
	UniqueCode               string                   `gorm:"unique;char(8);not null;comment:工区代码;" json:"unique_code"` // B048D001
	Name                     string                   `gorm:"unique;varchar(128);not null;comment:工区名称;" json:"name"`
	Category                 string                   `gorm:"enum('','SCENE','POINT-SWITCH','REPLY','SYNTHESIZE', 'POWER-SUPPLY-PANEL');not null;default:'';comment:工区类型;" json:"category"`
	OrganizationWorkshopUuid string                   `gorm:"type:char(36);not null;" json:"organization_workshop_uuid"`
	OrganizationWorkshop     *OrganizationWorkshopMdl `gorm:"foreignKey:organization_workshop_uuid;references:uuid;comment:所属车间;" json:"organization_workshop"`
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
		SetWheresExtraHasValue(map[string]func(string, *gorm.DB) *gorm.DB{}).
		SetWheresExtraHasValues(map[string]func([]string, *gorm.DB) *gorm.DB{}).
		SetCtx(ctx).
		GetDbUseQuery("").
		Table("organization_work_areas as owa")
}

// OrganizationLineMdl 组织结构-线别模型
type OrganizationLineMdl struct {
	MySqlMdl
	UniqueCode                string                      `gorm:"unique;char(5);not null;comment:线别代码;" json:"unique_code"` // E0001
	Name                      string                      `gorm:"unique;varchar(128);not null;comment:线别名称;" json:"name"`
	OrganizationParagraphUuid string                      `gorm:"type:char(36);not null;" json:"organization_paragraph_uuid"`
	OrganizationParagraph     *OrganizationParagraphMdl   `gorm:"foreignKey:organization_paragraph_uuid;references:uuid;comment:所属站段;" json:"organization_paragraph"`
	OrganizationStations      []*OrganizationStationMdl   `gorm:"foreignKey:organization_line_uuid;references:uuid;comment:相关站场;" json:"organization_stations"`
	OrganizationCrossroads    []*OrganizationCrossroadMdl `gorm:"foreignKey:organization_line_uuid;references:uuid;comment:相关道口;" json:"organization_crossroads"`
	OrganizationCenters       []*OrganizationCenterMdl    `gorm:"foreignKey:organization_line_uuid;references:uuid;comment:相关中心;" json:"organization_centers"`
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
		SetWheresExtraHasValue(map[string]func(string, *gorm.DB) *gorm.DB{}).
		SetWheresExtraHasValues(map[string]func([]string, *gorm.DB) *gorm.DB{}).
		SetCtx(ctx).
		GetDbUseQuery("").
		Table("organization_lines as ol")
}
