package utils

import (
	"math"
	"strconv"
)

// MathHelper 数学
type MathHelper struct{}

// NewMath 构造函数
func NewMath() *MathHelper {
	return &MathHelper{}
}

// Decimal 保留小数（四舍五入）
func (MathHelper) Decimal(value float64) float64 {
	return math.Trunc(value*1e2+0.5) * 1e-2
}

// DecToBase36 10 -> 36
func (MathHelper) DecToBase36(dec int64) (string, error) {
	return strconv.FormatInt(dec, 36), nil
}

// Base36ToDec 36 -> 10
func (MathHelper) Base36ToDec(base36 string) (int64, error) {
	return strconv.ParseInt(base36, 36, 64)
}

// PadLeftZeros 前置补零
func (MathHelper) PadLeftZeros(str string, length int) string {
	for len(str) < length {
		str = "0" + str
	}
	return str
}
