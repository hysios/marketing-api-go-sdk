package model

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// LoginType 是一个可以同时支持字符串和整数的类型
type LoginType struct {
	strVal string
	intVal int64
	isInt  bool
}

// NewLoginTypeFromString 从字符串创建 LoginType
func NewLoginTypeFromString(s string) *LoginType {
	return &LoginType{
		strVal: s,
		isInt:  false,
	}
}

// NewLoginTypeFromInt 从整数创建 LoginType
func NewLoginTypeFromInt(i int64) *LoginType {
	return &LoginType{
		intVal: i,
		isInt:  true,
	}
}

// IsInt 返回值是否为整数
func (lt *LoginType) IsInt() bool {
	return lt.isInt
}

// IsString 返回值是否为字符串
func (lt *LoginType) IsString() bool {
	return !lt.isInt
}

// String 获取字符串值
func (lt *LoginType) String() (string, bool) {
	if lt.isInt {
		return strconv.FormatInt(lt.intVal, 10), true
	}
	return lt.strVal, true
}

// Int 获取整数值
func (lt *LoginType) Int() (int64, bool) {
	if lt.isInt {
		return lt.intVal, true
	}
	// 尝试将字符串转换为整数
	if i, err := strconv.ParseInt(lt.strVal, 10, 64); err == nil {
		return i, true
	}
	return 0, false
}

// MarshalJSON 实现 json.Marshaler 接口
func (lt *LoginType) MarshalJSON() ([]byte, error) {
	if lt == nil {
		return []byte("null"), nil
	}
	if lt.isInt {
		return json.Marshal(lt.intVal)
	}
	return json.Marshal(lt.strVal)
}

// UnmarshalJSON 实现 json.Unmarshaler 接口
func (lt *LoginType) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}

	// 尝试解析为整数
	var intVal int64
	if err := json.Unmarshal(data, &intVal); err == nil {
		lt.intVal = intVal
		lt.isInt = true
		return nil
	}

	// 尝试解析为字符串
	var strVal string
	if err := json.Unmarshal(data, &strVal); err == nil {
		lt.strVal = strVal
		lt.isInt = false
		return nil
	}

	return fmt.Errorf("value must be either string or integer")
}

// Value 返回实际值的接口类型
func (lt *LoginType) Value() interface{} {
	if lt.isInt {
		return lt.intVal
	}
	return lt.strVal
}
