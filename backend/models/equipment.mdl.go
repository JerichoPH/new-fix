package models

import (
	"new-fix/tools"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// EquipmentMdl 器材模型
type EquipmentMdl struct {
	MySqlMdl
	UniqueCode                              string                    `gorm:"unique;char(19);not null;comment:唯一编号;" json:"unique_code"`
	EquipmentKindModelUuid                  string                    `gorm:"type:char(36);not null;comment:器材种类UUID;" json:"equipment_kind_model_uuid"`
	EquipmentKindModel                      *EquipmentKindModelMdl    `gorm:"foreignKey:equipment_kind_model_uuid;references:uuid;comment:所属器材型号;" json:"equipment_kind_model"`
	PropertyBelongOrganizationParagraphUuid string                    `gorm:"index;type:char(36);not null;default:'';comment:资产归属站段;" json:"property_belong_organization_paragraph_uuid"`
	PropertyBelongOrganizationParagraph     *OrganizationParagraphMdl `gorm:"foreignKey:property_belong_organization_paragraph_uuid;references:uuid;comment:资产归属站段;" json:"property_belong_organization_paragraph"`
	PropertyBelongOrganizationWorkshopUuid  string                    `gorm:"index;type:char(36);not null;default:'';comment:资产归属车间;" json:"property_belong_organization_workshop_uuid"`
	PropertyBelongOrganizationWorkshop      *OrganizationWorkshopMdl  `gorm:"foreignKey:property_belong_organization_workshop_uuid;references:uuid;comment:资产归属车间;" json:"property_belong_organization_workshop"`
	PropertyBelongOrganizationWorkAreaUuid  string                    `gorm:"index;type:char(36);not null;default:'';comment:资产归属;" json:"property_belong_organization_work_area_uuid"`
	PropertyBelongOrganiationWorkArea       *OrganizationWorkAreaMdl  `gorm:"foreignKey:property_belong_organization_work_area_uuid;references:uuid;comment:资产归属工区;" json:"property_belong_organization_work_area"`
	UseBelongOrganizaitonWorkshopUuid       string                    `gorm:"index;type:char(36);not null;default:'';comment:使用归属车间;" json:"use_belong_organizaiton_workshop_uuid"`
	UseBelongOrganizaitonWorkshop           *OrganizationWorkshopMdl  `gorm:"foreignKey:use_belong_organizaiton_workshop_uuid;references:uuid;comment:使用归属车间;" json:"use_belong_organizaiton_workshop"`
	UseBelongOrganizationStationUuid        string                    `gorm:"index;type:char(36);not null;default:'';comment:使用归属站场;" json:"use_belong_organization_station_uuid"`
	UseBelongOrganizationStation            *OrganizationStationMdl   `gorm:"foreignKey:use_belong_organization_station_uuid;references:uuid;comment:使用归属站场;" json:"use_belong_organization_station"`
	UseBelongOrganizationCrossroadUuid      string                    `gorm:"index;type:char(36);not null;default:'';comment:使用归属道口;" json:"use_belong_organization_crossroad_uuid"`
	UseBelongOrganizationCrossroad          *OrganizationCrossroadMdl `gorm:"foreignKey:use_belong_organization_crossroad_uuid;references:uuid;comment:使用归属道口;" json:"use_belong_organization_crossroad"`
	UseBelongOrganizationCenterUuid         string                    `gorm:"index;type:char(36);not null;default:'';comment:使用归属中心;" json:"use_belong_organization_center_uuid"`
	UseBelongOrganizationCenter             *OrganizationCenterMdl    `gorm:"foreignKey:use_belong_organization_center_uuid;references:uuid;comment:使用归属中心;" json:"use_belong_organization_center"`
	UseBelongOrganizationWorkAreaUuid       string                    `gorm:"index;type:char(36);not null;default:'';comment:使用归属工区;" json:"use_belong_organization_work_area_uuid"`
	UseBelongOrganizationWorkArea           *OrganizationWorkAreaMdl  `gorm:"foreignKey:use_belong_organization_work_area_uuid;references:uuid;comment:使用归属工区;" json:"use_belong_organization_work_area"`
	UseBelongOrganizationLineUuid           string                    `gorm:"index;type:char(36);not null;default:'';comment:使用归属线别;" json:"use_belong_organization_line_uuid"`
	UseBelongOrganizationLine               *OrganizationLineMdl      `gorm:"foreignKey:use_belong_organization_line_uuid;references:uuid;comment:使用归属线别;" json:"use_belong_organization_line"`
	StatusCode                              string                    `gorm:"type:enum ('BUY_IN', 'INSTALLING', 'INSTALLED', 'FIXING', 'FIXED', 'RETURN_FACTORY', 'FACTORY_RETURN', 'SCRAP', 'TRANSFER_OUT', 'TRANSFER_IN', 'UNINSTALLED', 'FRMLOSS', 'SEND_REPAIR', 'REPAIRING', 'UNINSTALLED_BREAKDOWN');comment:状态代码" json:"status_code"`
	StatusText                              string                    `gorm:"-"`
	FactoryUuid                             string                    `gorm:"index;type:char(36);not null;default:'';comment:厂家编号;" json:"factory_uuid"`
	Factory                                 *FactoryMdl               `gorm:"foreignKey:factory_uuid;references:uuid;comment:厂家;" json:"factory"`
	MadeAt                                  *time.Time                `gorm:"type:date;comment:出厂日期;" json:"made_at"`
	UsageWorkshopInOrderUuid                string                    `gorm:"index;type:char(36);not null;default:'';comment:入所单编号;" json:"in_workshop_order_uuid"`
	UsageWorkshopOutOrderUuid               string                    `gorm:"index;type:char(36);not null;default:'';comment:出所单编号;" json:"out_workshop_order_uuid"`
	UsageWarehouseInOrderUuid               string                    `gorm:"index;type:char(36);not null;default:'';comment:入库单编号;" json:"in_warehouse_order_uuid"`
	UsageWarehouseOutOrderUuid              string                    `gorm:"index;type:char(36);not null;default:'';comment:出库单编号;" json:"out_warehouse_order_uuid"`
	UsageRepairOrderUuid                    string                    `gorm:"index;type:char(36);not null;default:'';comment:检修单编号;" json:"repair_order_uuid"`
	UsageScrapOrderUuid                     string                    `gorm:"index;type:char(36);not null;default:'';comment:报废单编号;" json:"scrap_order_uuid"`
	UsageInstallOrderUuid                   string                    `gorm:"index;type:char(36);not null;default:'';comment:上道单编号;" json:"install_order_uuid"`
	UsageUninstallOrderUuid                 string                    `gorm:"index;type:char(36);not null;default:'';comment:下道单编号;" json:"uninstall_order_uuid"`
	UsageBeSpareOrderUuid                   string                    `gorm:"index;type:char(36);not null;default:'';comment:成为备品单编号;" json:"be_spare_order_uuid"`
	CanIWorkshopOut                         bool                      `gorm:"-" json:"can_i_workshop_out"`
	CanIInstall                             bool                      `gorm:"-" json:"can_i_install"`
	CanIUninstall                           bool                      `gorm:"-" json:"can_i_uninstall"`
	CanIBeSpare                             bool                      `gorm:"-" json:"can_i_be_spare"`
	BreakdonwLogs                           []*BreakdownLogMdl        `gorm:"foreignKey:equipment_uuid;references:uuid;comment:故障日志;" json:"breakdown_logs"`
	SourceTypeUuid                          string                    `gorm:"type:char(36);not null;default:'';comment:所属来源类型uuid;" json:"source_type_uuid"`
	SourceType                              *SourceTypeMdl            `gorm:"foreignKey:source_type_uuid;references:uuid;comment:所属来源类型;" json:"source_type"`
	SourceProjectUuid                       string                    `gorm:"type:char(36);not null;default:'';comment:所属来源项目uuid;" json:"source_project_uuid"`
	SourceProject                           *SourceProjectMdl         `gorm:"foreignKey:source_project_uuid;references:uuid;comment:所属来源项目;" json:"source_project"`
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

// GetStatusCodes 获取器材状态代码列表
func (EquipmentMdl) GetStatusCodes() []string {
	return []string{
		"FIXING",
		"FIXED",
		"TRANSFER_OUT",
		"INSTALLED",
		"INSTALLING",
		"TRANSFER_IN",
		"UNINSTALLED",
		"UNINSTALLED_BREAKDOWN",
		"SEND_REPAIR",
		"SCRAP",
	}
}

// GetStatusCodesMap 获取器材状态代码映射
func (EquipmentMdl) GetStatusCodesMap() []map[string]string {
	return []map[string]string{
		{"code": "FIXING", "text": "待修"},
		{"code": "FIXED", "text": "所内成品"},
		{"code": "TRANSFER_OUT", "text": "出所在途"},
		{"code": "INSTALLED", "text": "上道使用"},
		{"code": "INSTALLING", "text": "备品"},
		{"code": "TRANSFER_IN", "text": "入所在途"},
		{"code": "UNINSTALLED", "text": "下道停用"},
		{"code": "UNINSTALLED_BREAKDOWN", "text": "故障停用"},
		{"code": "SEND_REPAIR", "text": "送修中"},
		{"code": "SCRAP", "text": "报废"},
	}
}

// getStatusText 获取状态描述
func (receiver *EquipmentMdl) getStatusText() {
	var statusText string
	for _, item := range receiver.GetStatusCodesMap() {
		if item["code"] == receiver.StatusCode {
			statusText = item["text"]
			break
		}
	}
	receiver.StatusText = statusText
}

// canIWorkshopOut 获取是否可以出所标记
func (receiver *EquipmentMdl) getCanIWorkshopOut() {
	receiver.CanIWorkshopOut = receiver.StatusCode == "FIXED"
}

// getCanIInstall 获取是否可以上道标记
func (receiver *EquipmentMdl) getCanIInstall() {
	receiver.CanIInstall = tools.InString(receiver.StatusCode, []string{"TRANSFER_OUT", "UNINSTALLED"})
}

// getCanIUninstall 获取是否可以下道标记
func (receiver *EquipmentMdl) getCanIUninstall() {
	receiver.CanIUninstall = tools.InString(receiver.StatusCode, []string{"INSTALLED", "INSTALLING"})
}

// getCanIBeSpare 获取是否可以成为备品标记
func (receiver *EquipmentMdl) getCanIBeSpare() {
	receiver.CanIBeSpare = tools.InString(receiver.StatusCode, []string{"TRANSFER_OUT", "UNINSTALLED"})
}

// AfterFind 查询后钩子
func (receiver *EquipmentMdl) AfterFind(db *gorm.DB) (err error) {
	receiver.getStatusText()
	receiver.getCanIWorkshopOut()
	receiver.getCanIInstall()
	receiver.getCanIUninstall()
	receiver.getCanIBeSpare()
	return
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
