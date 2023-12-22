package tools

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"new-fix/types"
	"new-fix/wrongs"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// GetAuth 获取登陆信息
func GetAuth(ctx *gin.Context) any {
	authorization, exist := ctx.Get(string(types.ACCOUNT_AUTH))
	if !exist {
		wrongs.ThrowUnLogin("登陆失效")
	}

	return authorization
}

// 生成密码
func GeneratePassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes)
}

// 验证密码
func CheckPassword(password, target string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(target), []byte(password)); err != nil {
		return false
	}
	return true
}

// ToJson json序列化
func ToJson(v any) string {
	jsonBytes, _ := json.Marshal(v)
	return string(jsonBytes)
}

// ToJsonFormat json序列化 格式化
func ToJsonFormat(v any) string {
	jsonBytes, _ := json.MarshalIndent(v, "", "  ")
	return string(jsonBytes)
}

// FromJson json反序列化
func FromJson(s string) any {
	var v any
	json.Unmarshal([]byte(s), &v)
	return v
}

// DeepCopyByGob 深拷贝
func DeepCopyByGob(dst, src any) error {
	var buffer bytes.Buffer
	if err := gob.NewEncoder(&buffer).Encode(src); err != nil {
		return err
	}

	return gob.NewDecoder(&buffer).Decode(dst)
}

// JoinWithoutEmpty 去掉空值然后合并
func JoinWithoutEmpty(values []string, sep string) string {
	newValues := make([]string, 0)

	for _, value := range values {
		if value != "" {
			newValues = append(newValues, value)
		}
	}

	return strings.Join(newValues, sep)
}

// AddPrefix 给字符串增加前缀，否则返回默认值
func AddPrefix(value, prefix, defaultValue string) string {
	return Operation{}.Ternary(value != "", fmt.Sprintf("%s%s", prefix, value), defaultValue).(string)
}

// InString 判断字符串是否存在数组中
func InString(target string, strings []string) bool {
	for _, element := range strings {
		if target == element {
			return true
		}
	}
	return false
}

// InInt 判断int是否在数组中
func InInt(target int, values types.ListInt) bool {
	for _, element := range values {
		if target == element {
			return true
		}
	}
	return false
}

// InInt8 判断int8是否在数组中
func InInt8(target int8, values types.ListInt8) bool {
	for _, element := range values {
		if target == element {
			return true
		}
	}
	return false
}

// InInt16 判断int16是否在数组中
func InInt16(target int16, values types.ListInt16) bool {
	for _, element := range values {
		if target == element {
			return true
		}
	}
	return false
}

// InInt32 判断int32是否在数组中
func InInt32(target int32, values types.ListInt32) bool {
	for _, element := range values {
		if target == element {
			return true
		}
	}
	return false
}

// InInt64 判断int64是否在数组中
func InInt64(target int64, values types.ListInt64) bool {
	for _, element := range values {
		if target == element {
			return true
		}
	}
	return false
}

// InUint 判断uint是否在数组中
func InUint(target uint, values types.ListUint) bool {
	for _, element := range values {
		if target == element {
			return true
		}
	}
	return false
}

// InUint8 判断uint8是否在数组中
func InUint8(target uint8, values types.ListUint8) bool {
	for _, element := range values {
		if target == element {
			return true
		}
	}
	return false
}

// InUint16 判断uint16是否在数组中
func InUint16(target uint16, values types.ListUint16) bool {
	for _, element := range values {
		if target == element {
			return true
		}
	}
	return false
}

// InUint32 判断uint32是否在数组中
func InUint32(target uint32, values types.ListUint32) bool {
	for _, element := range values {
		if target == element {
			return true
		}
	}
	return false
}

// InUint64 判断uint64是否在数组中
func InUint64(target uint64, values types.ListUint64) bool {
	for _, element := range values {
		if target == element {
			return true
		}
	}
	return false
}

// InFloat32 判断float32是否在数组中
func InFloat32(target float32, strings types.ListFloat32) bool {
	for _, element := range strings {
		if target == element {
			return true
		}
	}
	return false
}

// InFloat64 判断float64是否在数组中
func InFloat64(target float64, strings types.ListFloat64) bool {
	for _, element := range strings {
		if target == element {
			return true
		}
	}
	return false
}
