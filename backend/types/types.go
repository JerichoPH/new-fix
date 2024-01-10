package types

import (
	"reflect"
	"strings"
)

// 用户信息
type AccountInfo struct {
	Uuid       string `json:"uuid"`
	Username   string `json:"username"`
	Nickname   string `json:"nickname"`
	BeAdmin    bool   `json:"be_admin"`
	AvatarUuid string `json:"avatar_uuid"`
	Avatar     struct {
		SavePath          string `json:"save_path"`
		Filename          string `json:"filename"`
		OriginalFilename  string `json:"original_filename"`
		OriginalExtension string `json:"original_extension"`
		Url               string `json:"url"`
	}
	RbacRoles []struct {
		Uuid            string `json:"uuid"`
		Name            string `json:"name"`
		RbacPermissions []struct {
			Uuid   string `json:"uuid"`
			Name   string `json:"name"`
			Uri    string `json:"uri"`
			Method string `json:"method"`
		}
	}
}

// IsString 判断是否是字符串类型
func IsString(value any) bool {
	return reflect.TypeOf(value).Kind() == reflect.String
}

// IsNumber 判断是否是数字类型
func IsNumber(value any) bool {
	if IsInt(value) || IsInt8(value) || IsInt16(value) || IsInt32(value) || IsInt64(value) || IsUint(value) || IsUint8(value) || IsUint16(value) || IsUint32(value) || IsUint64(value) {
		return true
	}
	return false
}

// IsInt 判断是否是整数类型
func IsInt(value any) bool {
	return reflect.TypeOf(value).Kind() == reflect.Int
}

// IsInt8 判断是否是int8类型
func IsInt8(value any) bool {
	return reflect.TypeOf(value).Kind() == reflect.Int8
}

// IsInt16 判断是否是int8类型
func IsInt16(value any) bool {
	return reflect.TypeOf(value).Kind() == reflect.Int16
}

// IsInt32 判断是否是int32类型
func IsInt32(value any) bool {
	return reflect.TypeOf(value).Kind() == reflect.Int32
}

// IsInt64 判断是否是int64类型
func IsInt64(value any) bool {
	return reflect.TypeOf(value).Kind() == reflect.Int64
}

// IsUint 判断是否是uint类型
func IsUint(value any) bool {
	return reflect.TypeOf(value).Kind() == reflect.Int
}

// IsUint8 判断是否是uint8类型
func IsUint8(value any) bool {
	return reflect.TypeOf(value).Kind() == reflect.Uint8
}

func IsUint16(value any) bool {
	return reflect.TypeOf(value).Kind() == reflect.Uint16
}

// IsUint32 判断是否是uint32类型
func IsUint32(value any) bool {
	return reflect.TypeOf(value).Kind() == reflect.Uint32
}

// IsUint64 判断是否是uint64类型
func IsUint64(value any) bool {
	return reflect.TypeOf(value).Kind() == reflect.Uint64
}

// IsFloat 判断是否是浮点类型
func IsFloat(value any) bool {
	if IsFloat32(value) || IsFloat64(value) {
		return true
	}
	return false
}

// IsFloat32 判断是否是float32
func IsFloat32(value any) bool {
	return reflect.TypeOf(value).Kind() == reflect.Float32
}

// IsFloat64 判断是否是float64
func IsFloat64(value any) bool {
	return reflect.TypeOf(value).Kind() == reflect.Float64
}

// IsEmptyStruct 判断是否为空结构
func IsEmptyStruct(value, typeMode any) any {
	return reflect.Zero(reflect.TypeOf(value)).Interface()
}

// MapString 字符串字典处理
type MapString struct{}

// GetKeys 获取字符串字典所有key
func (receiver *MapString) GetKeys(value map[string]any) []string {
	var keys = make([]string, 0)

	for key := range value {
		keys = append(keys, key)
	}

	return keys
}

// IsTrue 判断字符串是否代表true
func IsTrue(value string) bool {
	switch strings.ToLower(value) {
	case "1":
	case "true":
	case "on":
	case "yes":
		return true
	}

	return false
}

// IsFalse 判断字符串是否代表false
func IsFalse(value string) bool {
	switch strings.ToLower(value) {
	case "0":
	case "false":
	case "off":
	case "no":
		return true
	}
	return false
}
