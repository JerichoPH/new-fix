package models

import (
	"fmt"
	"new-fix/database"
	"new-fix/wrongs"
	"strconv"
	"strings"
	"time"

	uuid "github.com/satori/go.uuid"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// MySqlMdl 基础模型
type MySqlMdl struct {
	Id                       uint64         `gorm:"primaryKey;autoIncrement;type:bigint unsigned auto_increment;" json:"id"`
	CreatedAt                time.Time      `gorm:"<-:create;type:datetime;default:CURRENT_TIMESTAMP;comment:创建时间;" json:"created_at"`
	UpdatedAt                time.Time      `gorm:"type:datetime;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;comment:更新时间;" json:"updated_at"`
	DeletedAt                gorm.DeletedAt `gorm:"index;type:datetime" json:"deleted_at"`
	Uuid                     string         `gorm:"unique;type:varchar(36);not null;comment:uuid;" json:"uuid"`
	Sort                     int64          `gorm:"type:bigint;default:0;comment:排序;" json:"sort"`
	BeEnable                 bool           `gorm:"type:boolean;not null;default:0;comment:是否禁用;" json:"be_enable"`
	ctx                      *gin.Context
	preloads                 []string
	selects                  []string
	omits                    []string
	whereFields              []string
	notWhereFields           []string
	orWhereFields            []string
	ignoreFields             []string
	distinctFieldNames       []string
	wheres                   map[string]any
	notWheres                map[string]any
	orWheres                 map[string]any
	wheresExtra              map[string]func(string, *gorm.DB) *gorm.DB
	wheresExtraExist         map[string]func(string, *gorm.DB) *gorm.DB
	wheresExtraExists        map[string]func([]string, *gorm.DB) *gorm.DB
	scopes                   []func(*gorm.DB) *gorm.DB
	whereDateBetween         []string
	whereDatetimeBetween     []string
	whereBetween             []string
	whereIntBetween          []string
	whereFloatBetween        []string
	whereIn                  map[string]string
	whereFuzzyQueryCondition map[string]string
	model                    interface{}
}

// JoinString 生成join语句
func JoinString(slaveTableName, slaveShortTableName, slaveFieldName, masterFieldName string) string {
	return fmt.Sprintf("join %s %s on %s = %s", slaveTableName, slaveShortTableName, slaveFieldName, masterFieldName)
}

// LeftJoinString 生成left join语句
func LeftJoinString(slaveTableName, slaveShortTableName, slaveFieldName, masterFieldName string) string {
	return "left " + JoinString(slaveTableName, slaveShortTableName, slaveFieldName, masterFieldName)
}

// RightJoinString 生成right join语句
func RightJoinString(slaveTableName, slaveShortTableName, slaveFieldName, masterFieldName string) string {
	return "right " + JoinString(slaveTableName, slaveShortTableName, slaveFieldName, masterFieldName)
}

// NewMySqlMdl 构造函数
func NewMySqlMdl() *MySqlMdl {
	return &MySqlMdl{}
}

// demoFindOne 获取单条数据演示
func (receiver *MySqlMdl) demoFindOne() {
	var b MySqlMdl
	ret := receiver.
		SetModel(MySqlMdl{}).
		SetWheres(map[string]any{}).
		SetNotWheres(map[string]any{}).
		GetDb("").
		First(b)
	wrongs.ThrowWhenEmpty(ret, "XX")
}

// demoFind 获取多条数据演示
func (receiver *MySqlMdl) demoFind() {
	var b MySqlMdl
	var ctx *gin.Context
	receiver.
		SetModel(MySqlMdl{}).
		SetWheresEqual("uuid").
		SetWheresHasKey(map[string]func(string, *gorm.DB) *gorm.DB{
			"name": func(fieldName string, db *gorm.DB) *gorm.DB {
				if queryValue, exist := ctx.GetQuery(fieldName); exist {
					db = db.Where(fieldName+" like ?", fmt.Sprintf("%%%s%%", queryValue))
				}
				return db
			},
		}).
		SetWheresDateBetween("created_at", "updated_at").
		SetCtx(ctx).
		GetDbUseQuery("连接池名称：空字符串为默认").
		Find(&b)
}

// ScopeBeEnableTrue 启用（查询域）
func (*MySqlMdl) ScopeBeEnableTrue(db *gorm.DB) *gorm.DB {
	return db.Where("be_enable is true")
}

// ScopeBeEnableFalse 不启用（查询域）
func (*MySqlMdl) ScopeBeEnableFalse(db *gorm.DB) *gorm.DB {
	return db.Where("be_enable is false")
}

// SetCtx 设置Context
func (receiver *MySqlMdl) SetCtx(ctx *gin.Context) *MySqlMdl {
	receiver.ctx = ctx
	return receiver
}

// SetModel 设置使用的模型
func (receiver *MySqlMdl) SetModel(model interface{}) *MySqlMdl {
	receiver.model = model
	return receiver
}

// SetDistinct 设置不重复字段
func (receiver *MySqlMdl) SetDistinct(distinctFieldNames ...string) *MySqlMdl {
	receiver.distinctFieldNames = distinctFieldNames

	return receiver
}

// SetPreloads 设置Preloads
func (receiver *MySqlMdl) SetPreloads(preloads ...string) *MySqlMdl {
	receiver.preloads = preloads
	return receiver
}

// SetPreloadsByDefault 设置Preloads为默认
func (receiver *MySqlMdl) SetPreloadsByDefault() *MySqlMdl {
	receiver.preloads = []string{clause.Associations}
	return receiver
}

// SetSelects 设置Selects
func (receiver *MySqlMdl) SetSelects(selects ...string) *MySqlMdl {
	receiver.selects = selects
	return receiver
}

// SetOmits 设置Omits
func (receiver *MySqlMdl) SetOmits(omits ...string) *MySqlMdl {
	receiver.omits = omits
	return receiver
}

// SetWheresEqual 设置WhereFields
func (receiver *MySqlMdl) SetWheresEqual(whereFields ...string) *MySqlMdl {
	receiver.whereFields = whereFields
	return receiver
}

// SetNotWhereFields 设置NotWhereFields
func (receiver *MySqlMdl) SetNotWhereFields(notWhereFields ...string) *MySqlMdl {
	receiver.notWhereFields = notWhereFields
	return receiver
}

// SetOrWhereFields 设置OrWhere字段
func (receiver *MySqlMdl) SetOrWhereFields(orWhereFields ...string) *MySqlMdl {
	receiver.orWhereFields = orWhereFields
	return receiver
}

// SetIgnoreFields 设置IgnoreFields
func (receiver *MySqlMdl) SetIgnoreFields(ignoreFields ...string) *MySqlMdl {
	receiver.ignoreFields = ignoreFields
	return receiver
}

// SetWheres 通过Map设置Wheres
func (receiver *MySqlMdl) SetWheres(wheres map[string]any) *MySqlMdl {
	receiver.wheres = wheres
	return receiver
}

// SetNotWheres 设置NotWheres
func (receiver *MySqlMdl) SetNotWheres(notWheres map[string]any) *MySqlMdl {
	receiver.notWheres = notWheres
	return receiver
}

// SetOrWheres 设置Or条件
func (receiver *MySqlMdl) SetOrWheres(orWheres map[string]any) *MySqlMdl {
	receiver.orWheres = orWheres
	return receiver
}

// SetScopes 设置Scopes
func (receiver *MySqlMdl) SetScopes(scopes ...func(*gorm.DB) *gorm.DB) *MySqlMdl {
	receiver.scopes = scopes
	return receiver
}

// SetWheresHasKey 设置额外搜索条件字段
func (receiver *MySqlMdl) SetWheresHasKey(wheresExtra map[string]func(string, *gorm.DB) *gorm.DB) *MySqlMdl {
	receiver.wheresExtra = wheresExtra

	return receiver
}

// SetWheresExtraHasValue 当Query参数存在时设置额外搜索条件（单个条件）
func (receiver *MySqlMdl) SetWheresExtraHasValue(wheresExtraExist map[string]func(string, *gorm.DB) *gorm.DB) *MySqlMdl {
	receiver.wheresExtraExist = wheresExtraExist

	return receiver
}

// SetWheresExtraHasValues 当Query参数存在时设置额外搜索条件（多个条件）
func (receiver *MySqlMdl) SetWheresExtraHasValues(wheresExtraExists map[string]func([]string, *gorm.DB) *gorm.DB) *MySqlMdl {
	receiver.wheresExtraExists = wheresExtraExists

	return receiver
}

// SetWheresDateBetween 设置需要检查的日期范围字段
func (receiver *MySqlMdl) SetWheresDateBetween(fieldNames ...string) *MySqlMdl {
	receiver.whereDateBetween = fieldNames

	return receiver
}

// SetWheresBetween 设置需要范围查询的条件
func (receiver *MySqlMdl) SetWheresBetween(fieldNames ...string) *MySqlMdl {
	receiver.whereBetween = fieldNames

	return receiver
}

// SetWheresIntBetween 设置需要范围查询的条件（数字）
func (receiver *MySqlMdl) SetWheresIntBetween(fieldNames ...string) *MySqlMdl {
	receiver.whereIntBetween = fieldNames

	return receiver
}

// SetWheresFloatBetween 设置需要范围查询的条件（浮点）
func (receiver *MySqlMdl) SetWheresFloatBetween(fieldNames ...string) *MySqlMdl {
	receiver.whereFloatBetween = fieldNames

	return receiver
}

// SetWheresDatetimeBetween 设置需要检查的日期时间范围字段
func (receiver *MySqlMdl) SetWheresDatetimeBetween(fieldNames ...string) *MySqlMdl {
	receiver.whereDatetimeBetween = fieldNames

	return receiver
}

// SetWheresIn 设置自定义查询条件(in)
func (receiver *MySqlMdl) SetWheresIn(condition map[string]string) *MySqlMdl {
	receiver.whereIn = condition

	return receiver
}

// SetWheresFuzzy 设置模糊查询条件
func (receiver *MySqlMdl) SetWheresFuzzy(condition map[string]string) *MySqlMdl {
	receiver.whereFuzzyQueryCondition = condition

	return receiver
}

// BeforeCreate 插入数据前
func (receiver *MySqlMdl) BeforeCreate(db *gorm.DB) (err error) {
	receiver.Uuid = uuid.NewV4().String()
	receiver.CreatedAt = time.Now()
	receiver.UpdatedAt = time.Now()
	return
}

// BeforeSave 修改数据前
func (receiver *MySqlMdl) BeforeSave(db *gorm.DB) (err error) {
	receiver.UpdatedAt = time.Now()
	return
}

// GetDb 初始化
func (receiver *MySqlMdl) GetDb(dbConnName string) (query *gorm.DB) {
	query = database.NewGormLauncher().GetConn(dbConnName)

	query = query.Where(receiver.wheres).Not(receiver.notWheres)

	if receiver.model != nil {
		query = query.Model(&receiver.model)
	}

	// 设置scopes
	if len(receiver.scopes) > 0 {
		query = query.Scopes(receiver.scopes...)
	}

	// 拼接preloads关系
	if len(receiver.preloads) > 0 {
		for _, v := range receiver.preloads {
			query = query.Preload(v)
		}
	}

	// 拼接distinct
	if len(receiver.distinctFieldNames) > 0 {
		query = query.Distinct(receiver.distinctFieldNames)
	}

	// 拼接selects字段
	if len(receiver.selects) > 0 {
		query = query.Select(receiver.selects)
	}

	// 拼接omits字段
	if len(receiver.omits) > 0 {
		query = query.Omit(receiver.omits...)
	}

	return query
}

// GetDbUseQuery 根据Query参数初始化
// 自助条件规则说明：
// :== 等于
// :==i 等于整数
// :==f 等于小数
// :==b 等于布尔
// :||&== 或者等于
// :||&==i 或者等于整数
// :||&==f 或者等于小数
// :||&==b 或者等于布尔
// :<>&== 不等于
// :<>&==i 不等于整数
// :<>&==f 不等于小数
// :<>&==b 不等于布尔
// :<<>> 模糊查询
// :<< 左模糊查询
// :>> 右模糊查询
// :== [] 等于空字符串
// :<> [] 不等于空字符串
// :==![] 等于空
// :<>![] 非空
// :< 小于
// :<i 小于整数
// :<f 小于小数
// :> 大于
// :>i 大于整数
// :>f 大于小数
// :<> 不等于
// :<>i 不等于整数
// :<>f 不等于小数
// :<>b 不等于布尔
// :<= 小于等于
// :<=i 小于等于整数
// :<=f 小于等于小数
// :>= 大于等于
// :>=i 大于等于整数
// :>=f 大于等于小数
// :[*] in
// :[i] in 整数
// :[f] in 小数
// :, between
// :,i between 整数
// :,f between 小数
// :~[] preloads
// :. order
// :… order field
func (receiver *MySqlMdl) GetDbUseQuery(dbConnName string) *gorm.DB {
	dbSession := receiver.GetDb(dbConnName)

	wheres := make(map[string]any)
	notWheres := make(map[string]any)
	orWheres := make(map[string]any)

	// 自动化处理⬇

	// 拼接需要跳过的字段
	ignoreFields := make(map[string]int32)
	if len(receiver.ignoreFields) > 0 {
		for _, v := range receiver.ignoreFields {
			ignoreFields[v] = 1
		}
	}

	// 拼接Where条件
	for _, whereField := range receiver.whereFields {
		if _, exist := ignoreFields[whereField]; !exist {
			if val, ok := receiver.ctx.GetQuery(whereField); ok {
				wheres[whereField] = val
			}
		}
	}

	// 拼接NotWhere条件
	for _, notWhereField := range receiver.notWhereFields {
		if _, exist := ignoreFields[notWhereField]; !exist {
			if notWhereValue, ok := receiver.ctx.GetQuery(notWhereField); ok {
				notWheres[notWhereField] = notWhereValue
			}
		}
	}

	// 拼接OrWhere条件
	for _, orWhereField := range receiver.orWhereFields {
		if _, exist := ignoreFields[orWhereField]; !exist {
			if val, ok := receiver.ctx.GetQuery(orWhereField); ok {
				orWheres[orWhereField] = val
			}
		}
	}

	// 拼接WhereMap条件 (string)
	if m, exist := receiver.ctx.GetQueryMap(":=="); exist {
		for whereMapField, whereMapValue := range m {
			if _, exist := ignoreFields[whereMapField]; !exist {
				if whereMapValue != "" {
					wheres[whereMapField] = whereMapValue
				}
			}
		}
	}

	// 拼接WhereMap条件 (int)
	if m, exist := receiver.ctx.GetQueryMap(":==i"); exist {
		for whereMapField, whereMapValue := range m {
			if _, exist := ignoreFields[whereMapField]; !exist {
				if whereMapValue != "" {
					if i, e := strconv.Atoi(whereMapValue); e != nil {
						wrongs.ThrowValidate("%s 必须是整数（当前值%s）", whereMapField, whereMapValue)
					} else {
						wheres[whereMapField] = i
					}
				}
			}
		}
	}

	// 拼接WhereMap条件 (float)
	if m, exist := receiver.ctx.GetQueryMap(":==f"); exist {
		for whereMapField, whereMapValue := range m {
			if _, exist := ignoreFields[whereMapField]; !exist {
				if whereMapValue != "" {
					if f, e := strconv.ParseFloat(whereMapValue, 64); e != nil {
						wrongs.ThrowValidate("%s 必须是小数（当前值%s）", whereMapField, whereMapValue)
					} else {
						wheres[whereMapField] = f
					}
				}
			}
		}
	}

	// 拼接WhereMap条件 (bool)
	if m, exist := receiver.ctx.GetQueryMap(":==b"); exist {
		for whereMapField, whereMapValue := range m {
			if _, exist := ignoreFields[whereMapField]; !exist {
				if whereMapValue != "" {
					if b, e := strconv.ParseBool(whereMapValue); e != nil {
						wrongs.ThrowValidate("%s 必须是布尔（当前值%s）", whereMapField, whereMapValue)
					} else {
						wheres[whereMapField] = b
					}
				}
			}
		}
	}

	// 拼接OrWhereMap条件 (string)
	if m, exist := receiver.ctx.GetQueryMap(":||&=="); exist {
		for orWhereMapField, orWhereMapValue := range m {
			if _, exist := ignoreFields[orWhereMapField]; !exist {
				if orWhereMapValue != "" {
					notWheres[orWhereMapField] = orWhereMapValue
				}
			}
		}
	}

	// 拼接OrWhereMap条件 (int)
	if m, exist := receiver.ctx.GetQueryMap(":||&==i"); exist {
		for orWhereMapField, orWhereMapValue := range m {
			if _, exist := ignoreFields[orWhereMapField]; !exist {
				if orWhereMapValue != "" {
					if i, e := strconv.Atoi(orWhereMapValue); e != nil {
						wrongs.ThrowValidate("%s 必须是整数（当前值%s）", orWhereMapField, orWhereMapValue)
					} else {
						orWheres[orWhereMapField] = i
					}
				}
			}
		}
	}

	// 拼接OrWhereMap条件 (float)
	if m, exist := receiver.ctx.GetQueryMap(":||&==f"); exist {
		for orWhereMapField, orWhereMapValue := range m {
			if _, exist := ignoreFields[orWhereMapField]; !exist {
				if orWhereMapValue != "" {
					if f, e := strconv.ParseFloat(orWhereMapValue, 64); e != nil {
						wrongs.ThrowValidate("%s 必须是小数（当前值%s）", orWhereMapField, orWhereMapValue)
					} else {
						orWheres[orWhereMapField] = f
					}
				}
			}
		}
	}

	// 拼接OrWhereMap条件 (bool)
	if m, exist := receiver.ctx.GetQueryMap(":||&==b"); exist {
		for orWhereMapField, orWhereMapValue := range m {
			if _, exist := ignoreFields[orWhereMapField]; !exist {
				if orWhereMapValue != "" {
					if b, e := strconv.ParseBool(orWhereMapValue); e != nil {
						wrongs.ThrowValidate("%s 必须是布尔（当前值%s）", orWhereMapField, orWhereMapValue)
					} else {
						orWheres[orWhereMapField] = b
					}
				}
			}
		}
	}

	// 拼接NotWhereMap条件 (string)
	if m, exist := receiver.ctx.GetQueryMap(":<>&=="); exist {
		for orWhereMapField, orWhereMapValue := range m {
			if _, exist := ignoreFields[orWhereMapField]; !exist {
				if orWhereMapValue != "" {
					notWheres[orWhereMapField] = orWhereMapValue
				}
			}
		}
	}

	// 拼接NotWhereMap条件 (int)
	if m, exist := receiver.ctx.GetQueryMap(":<>&==i"); exist {
		for orWhereMapField, notWhereMapValue := range m {
			if _, exist := ignoreFields[orWhereMapField]; !exist {
				if notWhereMapValue != "" {
					if i, e := strconv.Atoi(notWhereMapValue); e != nil {
						wrongs.ThrowValidate("%s 必须是整数（当前值%s）", orWhereMapField, notWhereMapValue)
					} else {
						orWheres[orWhereMapField] = i
					}
				}
			}
		}
	}

	// 拼接NotWhereMap条件 (float)
	if m, exist := receiver.ctx.GetQueryMap(":<>&==f"); exist {
		for orWhereMapField, notWhereMapValue := range m {
			if _, exist := ignoreFields[orWhereMapField]; !exist {
				if notWhereMapValue != "" {
					if f, e := strconv.ParseFloat(notWhereMapValue, 64); e != nil {
						wrongs.ThrowValidate("%s 必须是小数（当前值%s）", orWhereMapField, notWhereMapValue)
					} else {
						orWheres[orWhereMapField] = f
					}
				}
			}
		}
	}

	// 拼接NotWhereMap条件 (bool)
	if m, exist := receiver.ctx.GetQueryMap(":<>&==b"); exist {
		for orWhereMapField, notWhereMapValue := range m {
			if _, exist := ignoreFields[orWhereMapField]; !exist {
				if notWhereMapValue != "" {
					if b, e := strconv.ParseBool(notWhereMapValue); e != nil {
						wrongs.ThrowValidate("%s 必须是布尔（当前值%s）", orWhereMapField, notWhereMapValue)
					} else {
						orWheres[orWhereMapField] = b
					}
				}
			}
		}
	}
	dbSession = dbSession.Where(wheres).Not(notWheres).Or(orWheres)

	// 拼接自定义条件（like）
	if m, exist := receiver.ctx.GetQueryMap(":<<>>"); exist {
		for field, value := range m {
			if _, exist := ignoreFields[field]; !exist {
				if value != "" {
					dbSession = dbSession.Where(fmt.Sprintf("%s like ? %%%s%%", field, value))
				}
			}
		}
	}

	// 拼接自定义条件（左like）
	if m, exist := receiver.ctx.GetQueryMap(":<<"); exist {
		for field, value := range m {
			if _, exist := ignoreFields[field]; !exist {
				if value != "" {
					dbSession = dbSession.Where(fmt.Sprintf("%s like ? %%%s", field, value))
				}
			}
		}
	}

	// 拼接自定义条件（右like）
	if m, exist := receiver.ctx.GetQueryMap(":>>"); exist {
		for field, value := range m {
			if _, exist := ignoreFields[field]; !exist {
				if value != "" {
					dbSession = dbSession.Where(fmt.Sprintf("%s like ? %s%%", field, value))
				}
			}
		}
	}

	// 拼接自定义条件（等于空字符串）
	if fields, exist := receiver.ctx.GetQueryArray(":== []"); exist {
		for _, field := range fields {
			if _, exist := ignoreFields[field]; !exist {
				dbSession = dbSession.Where(fmt.Sprintf("%s = ''", field))
			}
		}
	}

	// 拼接自定义条件（不等于空字符串）
	if fields, exist := receiver.ctx.GetQueryArray(":<> []"); exist {
		for _, field := range fields {
			if _, exist := ignoreFields[field]; !exist {
				dbSession = dbSession.Where(fmt.Sprintf("%s <> ''", field))
			}
		}
	}

	// 拼接自定义条件（等于空）
	if fields, exist := receiver.ctx.GetQueryArray(":==![]"); exist {
		for _, field := range fields {
			if _, exist := ignoreFields[field]; !exist {
				dbSession = dbSession.Where(fmt.Sprintf("%s is null", field))
			}
		}
	}

	// 拼接自定义条件（非空）
	if fields, exist := receiver.ctx.GetQueryArray(":<>![]"); exist {
		for _, field := range fields {
			if _, exist := ignoreFields[field]; !exist {
				dbSession = dbSession.Where(fmt.Sprintf("%s is not null", field))
			}
		}
	}

	// 拼接自定义条件（小于）(string)
	if m, exist := receiver.ctx.GetQueryMap(":<"); exist {
		for field, value := range m {
			if _, exist := ignoreFields[field]; !exist {
				if value != "" {
					dbSession = dbSession.Where(fmt.Sprintf("%s < ?", field), value)
				}
			}
		}
	}

	// 拼接自定义条件（小于）(int)
	if m, exist := receiver.ctx.GetQueryMap(":<i"); exist {
		for field, value := range m {
			if _, exist := ignoreFields[field]; !exist {
				if value != "" {
					if i, e := strconv.Atoi(value); e != nil {
						wrongs.ThrowValidate("%s 必须是整数（当前值%s）", field, value)
					} else {
						dbSession = dbSession.Where(fmt.Sprintf("%s < ?", field), i)
					}
				}
			}
		}
	}

	// 拼接自定义条件（小于）(float)
	if m, exist := receiver.ctx.GetQueryMap(":<f"); exist {
		for field, value := range m {
			if _, exist := ignoreFields[field]; !exist {
				if value != "" {
					if f, e := strconv.ParseFloat(value, 64); e != nil {
						wrongs.ThrowValidate("%s 必须是小数（当前值%s）", field, value)
					} else {
						dbSession = dbSession.Where(fmt.Sprintf("%s < ?", field), f)
					}
				}
			}
		}
	}

	// 拼接自定义条件（大于）(string)
	if m, exist := receiver.ctx.GetQueryMap(":>"); exist {
		for field, value := range m {
			if _, exist := ignoreFields[field]; !exist {
				if value != "" {
					dbSession = dbSession.Where(fmt.Sprintf("%s > ?", field), value)
				}
			}
		}
	}

	// 拼接自定义条件（大于）(int)
	if m, exist := receiver.ctx.GetQueryMap(":>i"); exist {
		for field, value := range m {
			if _, exist := ignoreFields[field]; !exist {
				if value != "" {
					if i, e := strconv.Atoi(value); e != nil {
						wrongs.ThrowValidate("%s 必须是整数（当前值%s）", field, value)
					} else {
						dbSession = dbSession.Where(fmt.Sprintf("%s > ?", field), i)
					}
				}
			}
		}
	}

	// 拼接自定义条件（大于）(fkoat)
	if m, exist := receiver.ctx.GetQueryMap(":>f"); exist {
		for field, value := range m {
			if _, exist := ignoreFields[field]; !exist {
				if value != "" {
					if f, e := strconv.ParseFloat(value, 64); e != nil {
						wrongs.ThrowValidate("%s 必须是小数（当前值%s）", field, value)
					} else {
						dbSession = dbSession.Where(fmt.Sprintf("%s > ?", field), f)
					}
				}
			}
		}
	}

	// 拼接自定义条件（不等于）(string)
	if m, exist := receiver.ctx.GetQueryMap(":<>"); exist {
		for field, value := range m {
			if _, exist := ignoreFields[field]; !exist {
				if value != "" {
					dbSession = dbSession.Where(fmt.Sprintf("%s != ?", field), value)
				}
			}
		}
	}

	// 拼接自定义条件（不等于）(int)
	if m, exist := receiver.ctx.GetQueryMap(":<>i"); exist {
		for field, value := range m {
			if _, exist := ignoreFields[field]; !exist {
				if value != "" {
					if i, e := strconv.Atoi(value); e != nil {
						wrongs.ThrowValidate("%s 必须是整数（当前值%s）", field, value)
					} else {
						dbSession = dbSession.Where(fmt.Sprintf("%s != ?", field), i)
					}
				}
			}
		}
	}

	// 拼接自定义条件（不等于）(float)
	if m, exist := receiver.ctx.GetQueryMap(":<>f"); exist {
		for field, value := range m {
			if _, exist := ignoreFields[field]; !exist {
				if value != "" {
					if f, e := strconv.ParseFloat(value, 64); e != nil {
						wrongs.ThrowValidate("%s 必须是小数（当前值%s）", field, value)
					} else {
						dbSession = dbSession.Where(fmt.Sprintf("%s != ?", field), f)
					}
				}
			}
		}
	}

	// 拼接自定义条件（不等于）(bool)
	if m, exist := receiver.ctx.GetQueryMap(":<>b"); exist {
		for field, value := range m {
			if _, exist := ignoreFields[field]; !exist {
				if value != "" {
					if b, e := strconv.ParseBool(value); e != nil {
						wrongs.ThrowValidate("%s 必须是布尔（当前值%s）", field, value)
					} else {
						dbSession = dbSession.Where(fmt.Sprintf("%s != ?", field), b)
					}
				}
			}
		}
	}

	// 拼接自定义条件（小于等于）(string)
	if m, exist := receiver.ctx.GetQueryMap(":<="); exist {
		for field, value := range m {
			if _, exist := ignoreFields[field]; !exist {
				if value != "" {
					dbSession = dbSession.Where(fmt.Sprintf("%s <= ?", field), value)
				}
			}
		}
	}

	// 拼接自定义条件（小于等于）(int)
	if m, exist := receiver.ctx.GetQueryMap(":<=i"); exist {
		for field, value := range m {
			if _, exist := ignoreFields[field]; !exist {
				if value != "" {
					if i, e := strconv.Atoi(value); e != nil {
						wrongs.ThrowValidate("%s 必须是整数（当前值%s）", field, value)
					} else {
						dbSession = dbSession.Where(fmt.Sprintf("%s <= ?", field), i)
					}
				}
			}
		}
	}

	// 拼接自定义条件（小于等于）(float)
	if m, exist := receiver.ctx.GetQueryMap(":<=f"); exist {
		for field, value := range m {
			if _, exist := ignoreFields[field]; !exist {
				if value != "" {
					if f, e := strconv.ParseFloat(value, 64); e != nil {
						wrongs.ThrowValidate("%s 必须是小数（当前值%s）", field, value)
					} else {
						dbSession = dbSession.Where(fmt.Sprintf("%s <= ?", field), f)
					}
				}
			}
		}
	}

	// 拼接自定义条件（大于等于）(string)
	if m, exist := receiver.ctx.GetQueryMap(":>="); exist {
		for field, value := range m {
			if _, exist := ignoreFields[field]; !exist {
				if value != "" {
					dbSession = dbSession.Where(fmt.Sprintf("%s >= ?", field), value)
				}
			}
		}
	}

	// 拼接自定义条件（大于等于）(int)
	if m, exist := receiver.ctx.GetQueryMap(":>=i"); exist {
		for field, value := range m {
			if _, exist := ignoreFields[field]; !exist {
				if value != "" {
					if i, e := strconv.Atoi(value); e != nil {
						wrongs.ThrowValidate("%s 必须是整数（当前值%s）", field, value)
					} else {
						dbSession = dbSession.Where(fmt.Sprintf("%s >= ?", field), i)
					}
				}
			}
		}
	}

	// 拼接自定义条件（大于等于）(float)
	if m, exist := receiver.ctx.GetQueryMap(":>=f"); exist {
		for field, value := range m {
			if _, exist := ignoreFields[field]; !exist {
				if value != "" {
					if f, e := strconv.ParseFloat(value, 64); e != nil {
						wrongs.ThrowValidate("%s 必须是小数（当前值%s）", field, value)
					} else {
						dbSession = dbSession.Where(fmt.Sprintf("%s >= ?", field), f)
					}
				}
			}
		}
	}

	// 拼接自定义条件（in）(string)
	if m, exist := receiver.ctx.GetQueryMap(":[]"); exist {
		for field, value := range m {
			if _, exist := ignoreFields[field]; !exist {
				values := strings.Split(value, ",")
				if len(values) > 0 {
					dbSession = dbSession.Where(fmt.Sprintf("%s in ?", field), values)
				}
			}
		}
	}

	// 拼接自定义条件（in）(int)
	if m, exist := receiver.ctx.GetQueryMap(":[i]"); exist {
		for field, value := range m {
			if _, exist := ignoreFields[field]; !exist {
				values := strings.Split(value, ",")
				valuesI := make([]int, 0)
				if len(values) > 0 {
					for k, v := range values {
						if i, e := strconv.Atoi(v); e != nil {
							wrongs.ThrowValidate("in条件必须是整数（当前值%s）", v)
						} else {
							valuesI[k] = i
						}
					}
					dbSession = dbSession.Where(fmt.Sprintf("%s in ?", field), valuesI)
				}
			}
		}
	}

	// 拼接自定义条件（in）(float)
	if m, exist := receiver.ctx.GetQueryMap(":[f]"); exist {
		for field, value := range m {
			if _, exist := ignoreFields[field]; !exist {
				values := strings.Split(value, ",")
				valuesF := make([]float64, 0)
				if len(values) > 0 {
					for k, v := range values {
						if f, e := strconv.ParseFloat(v, 64); e != nil {
							wrongs.ThrowValidate("in条件必须是小数（当前值%s）", v)
						} else {
							valuesF[k] = f
						}
					}
					dbSession = dbSession.Where(fmt.Sprintf("%s in ?", field), valuesF)
				}
			}
		}
	}

	// 拼接自定义条件（between）
	if m, exist := receiver.ctx.GetQueryMap(":,"); exist {
		for field, values := range m {
			if _, exist := ignoreFields[field]; !exist {
				values := strings.Split(values, ",")
				if len(values) > 0 {
					dbSession = dbSession.Where(fmt.Sprintf("%s between ? and ?", field), values[0], values[1])
				}
			}
		}
	}

	// 自动拼接条件（between int）
	if m, exist := receiver.ctx.GetQueryMap(":,i"); exist {
		for field, values := range m {
			if _, exist := ignoreFields[field]; !exist {
				values := strings.Split(values, ",")
				int1, err := strconv.Atoi(values[0])
				if err != nil {
					wrongs.ThrowValidate("%s 必须是数字，（当前值：%v）", field, values[0])
				}
				int2, err := strconv.Atoi(values[1])
				if err != nil {
					wrongs.ThrowValidate("%s 必须是数字，（当前值：%v）", field, values[1])
				}
				if len(values) > 0 {
					dbSession = dbSession.Where(fmt.Sprintf("%s between ? and ?", field), int1, int2)
				}
			}
		}
	}

	// 自动拼接条件（between float）
	if m, exist := receiver.ctx.GetQueryMap(":,f"); exist {
		for field, values := range m {
			if _, exist := ignoreFields[field]; !exist {
				values := strings.Split(values, ",")
				float1, err := strconv.ParseFloat(values[0], 64)
				if err != nil {
					wrongs.ThrowValidate("%s 必须是小数，（当前值：%v）", field, values[0])
				}
				float2, err := strconv.ParseFloat(values[1], 64)
				if err != nil {
					wrongs.ThrowValidate("%s 必须是小数，（当前值：%v）", field, values[1])
				}
				if len(values) > 0 {
					dbSession = dbSession.Where(fmt.Sprintf("%s between ? and ?", field), float1, float2)
				}
			}
		}
	}

	// 自动拼接preload
	if preloads, exist := receiver.ctx.GetQueryArray(":~[]"); exist {
		for _, preload := range preloads {
			dbSession = dbSession.Preload(preload)
		}
	}

	// 手动处理⬇

	// 拼接额外搜索条件
	for fieldName, v := range receiver.wheresExtra {
		if _, exist := ignoreFields[fieldName]; !exist {
			if _, ok := receiver.ctx.GetQuery(fieldName); ok {
				dbSession = v(fieldName, dbSession)
			}
		}
	}

	// 拼接额外搜索条件（判断值是否存在，单个条件）
	for fieldName, v := range receiver.wheresExtraExist {
		if _, exist := ignoreFields[fieldName]; !exist {
			if value := receiver.ctx.Query(fieldName); value != "" {
				dbSession = v(value, dbSession)
			}
		}
	}

	// 拼接额外搜索条件（判断值是否存在，多个条件）
	for fieldName, v := range receiver.wheresExtraExists {
		if _, exist := ignoreFields[fieldName]; !exist {
			if values := receiver.ctx.QueryArray(fieldName); len(values) > 0 {
				dbSession = v(values, dbSession)
			}
		}
	}

	// 拼接日期范围查询
	if len(receiver.whereDateBetween) > 0 {
		for _, fieldName := range receiver.whereDateBetween {
			if value, exist := receiver.ctx.GetQuery(fieldName); exist {
				times := strings.Split(value, "~")
				originalAt, finishedAt := times[0], times[1]
				dbSession = dbSession.Where(fieldName+" between ? and ?", originalAt+" 00:00:00", finishedAt+" 23:59:59")
			}
		}
	}

	// 拼接日期时间范围查询
	if len(receiver.whereDatetimeBetween) > 0 {
		for _, fieldName := range receiver.whereDateBetween {
			if value, exist := receiver.ctx.GetQuery(fieldName); exist {
				times := strings.Split(value, "~")
				originalAt, finishedAt := times[0], times[1]
				dbSession = dbSession.Where(fieldName+" between ? and ?", originalAt, finishedAt)
			}
		}
	}

	// 拼接between范围查询条件
	if len(receiver.whereBetween) > 0 {
		for _, fieldName := range receiver.whereBetween {
			if values, exist := receiver.ctx.GetQuery(fieldName); exist {
				values := strings.Split(values, ",")
				if len(values) == 2 {
					dbSession = dbSession.Where(fmt.Sprintf("%s between ? and ?", fieldName), values[0], values[1])
				}
			}
		}
	}

	// 拼接int between范围查询条件
	if len(receiver.whereIntBetween) > 0 {
		for _, fieldName := range receiver.whereIntBetween {
			if values, exist := receiver.ctx.GetQuery(fieldName); exist {
				values := strings.Split(values, ",")
				if len(values) == 2 {
					int1, err := strconv.Atoi(values[0])
					if err != nil {
						wrongs.ThrowValidate("%s 必须是整数（当前值：%v）", fieldName, values[0])
					}
					int2, err := strconv.Atoi(values[1])
					if err != nil {
						wrongs.ThrowValidate("%s 必须是整数（当前值：%v）", fieldName, values[1])
					}

					dbSession = dbSession.Where(fmt.Sprintf("%s between ? and ?", fieldName), int1, int2)
				}
			}
		}
	}

	// 拼接flat between范围查询条件
	if len(receiver.whereFloatBetween) > 0 {
		for _, fieldName := range receiver.whereFloatBetween {
			if values, exist := receiver.ctx.GetQuery(fieldName); exist {
				values := strings.Split(values, ",")
				if len(values) == 2 {
					float1, err := strconv.ParseFloat(values[0], 64)
					if err != nil {
						wrongs.ThrowValidate("%s 必须是小数（当前值：%v）", fieldName, values[0])
					}
					float2, err := strconv.ParseFloat(values[0], 64)
					if err != nil {
						wrongs.ThrowValidate("%s 必须是小数（当前值：%v）", fieldName, values[1])
					}

					dbSession = dbSession.Where(fmt.Sprintf("%s between ? and ?", fieldName), float1, float2)
				}
			}
		}
	}

	// 拼接in查询条件
	if len(receiver.whereIn) > 0 {
		for field, condition := range receiver.whereIn {
			if values, exist := receiver.ctx.GetQueryArray(field); exist {
				if len(values) > 0 {
					dbSession = dbSession.Where(condition, values)
				}
			}
		}
	}

	// 拼接模糊查询条件
	if len(receiver.whereFuzzyQueryCondition) > 0 {
		for field, condition := range receiver.whereFuzzyQueryCondition {
			if value, exist := receiver.ctx.GetQuery(field); exist {
				if value != "" {
					dbSession = dbSession.Where(condition, "%"+value+"%")
				}
			}
		}
	}

	// 排序
	if order, exist := receiver.ctx.GetQuery(":."); exist {
		dbSession.Order(order)
	}

	// 指定排序
	if orderField, exist := receiver.ctx.GetQueryArray(":..."); exist {
		if len(orderField) > 2 {
			dbSession.Order(fmt.Sprintf("field (%s,'%s')", orderField[0], strings.Join(orderField[1:], "','")))
		}
	}

	return dbSession
}
