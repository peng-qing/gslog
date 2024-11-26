package gslog

import (
	"errors"
	"fmt"
	"time"
)

type LogFieldValueKind int

const (
	LogFieldValueAny LogFieldValueKind = iota
	LogFieldValueInt64
	LogFieldValueInt64s
	LogFieldValueUint64
	LogFieldValueUint64s
	LogFieldValueFloat64
	LogFieldValueFloat64s
	LogFieldValueString
	LogFieldValueStrings
	LogFieldValueBool
	LogFieldValueBools
	LogFieldValueTime
	LogFieldValueDuration
	LogFieldValueField
	LogFieldValueFields
	LogFieldValueError
)

var (
	kindStrings = map[LogFieldValueKind]string{
		LogFieldValueAny:      "Any",
		LogFieldValueInt64:    "Int64",
		LogFieldValueInt64s:   "Int64s",
		LogFieldValueUint64:   "Uint64",
		LogFieldValueUint64s:  "Uint64s",
		LogFieldValueFloat64:  "Float64",
		LogFieldValueFloat64s: "Floats",
		LogFieldValueString:   "String",
		LogFieldValueStrings:  "Strings",
		LogFieldValueBool:     "Bool",
		LogFieldValueBools:    "Bools",
		LogFieldValueTime:     "Time",
		LogFieldValueDuration: "Duration",
		LogFieldValueField:    "Field",
		LogFieldValueFields:   "Fields",
		LogFieldValueError:    "Error",
	}
)

// String 获取类型对应字符串
func (l LogFieldValueKind) String() string {
	if l >= 0 && l < LogFieldValueKind(len(kindStrings)) {
		return kindStrings[l]
	}

	return fmt.Sprintf("LogFieldValueKind(%d)", l)
}

// LogFieldValue 日志字段值类型
type LogFieldValue struct {
	// 禁止比较运算符 ==
	_ [0]func()
	// 实际类型说明
	kind LogFieldValueKind
	// int/int8/int16/int32/int64 => int64
	// uint/uint8/uint16/uint32/uint64/byte => uint64
	// float32/float64 => float64
	value any
}

// IntFieldValue Int
func IntFieldValue(val int) LogFieldValue {
	return LogFieldValue{kind: LogFieldValueInt64, value: int64(val)}
}

// Int64FieldValue Int64
func Int64FieldValue(val int64) LogFieldValue {
	return LogFieldValue{kind: LogFieldValueInt64, value: val}
}

// Int64ArrayFieldValue int64 array
func Int64ArrayFieldValue(val ...int64) LogFieldValue {
	return LogFieldValue{kind: LogFieldValueInt64, value: val}
}

// Uint64FieldValue uint64
func Uint64FieldValue(val uint64) LogFieldValue {
	return LogFieldValue{kind: LogFieldValueUint64, value: val}
}

// Uint64ArrayFieldValue uint64 array
func Uint64ArrayFieldValue(val ...uint64) LogFieldValue {
	return LogFieldValue{kind: LogFieldValueUint64s, value: val}
}

// Float64FieldValue float64
func Float64FieldValue(val float64) LogFieldValue {
	return LogFieldValue{kind: LogFieldValueFloat64, value: val}
}

// Float64ArrayFieldValue float64 array
func Float64ArrayFieldValue(val ...float64) LogFieldValue {
	return LogFieldValue{kind: LogFieldValueFloat64s, value: val}
}

// StringFieldValue string
func StringFieldValue(val string) LogFieldValue {
	return LogFieldValue{kind: LogFieldValueString, value: val}
}

// StringArrayFieldValue string array
func StringArrayFieldValue(val ...string) LogFieldValue {
	return LogFieldValue{kind: LogFieldValueStrings, value: val}
}

// BoolFieldValue bool
func BoolFieldValue(val bool) LogFieldValue {
	return LogFieldValue{kind: LogFieldValueBool, value: val}
}

// BoolArrayFieldValue bool array
func BoolArrayFieldValue(val ...bool) LogFieldValue {
	return LogFieldValue{kind: LogFieldValueBools, value: val}
}

// TimeFieldValue time.Time
func TimeFieldValue(val time.Time) LogFieldValue {
	return LogFieldValue{kind: LogFieldValueTime, value: val}
}

// DurationFieldValue time.Duration
func DurationFieldValue(val time.Duration) LogFieldValue {
	return LogFieldValue{kind: LogFieldValueDuration, value: val}
}

// FieldFieldValue LogField
func FieldFieldValue(val LogField) LogFieldValue {
	return LogFieldValue{kind: LogFieldValueField, value: val}
}

// FieldArrayFieldValue LogField array
func FieldArrayFieldValue(val ...LogField) LogFieldValue {
	return LogFieldValue{kind: LogFieldValueFields, value: val}
}

// ErrorFieldValue error
func ErrorFieldValue(val ...error) LogFieldValue {
	return LogFieldValue{kind: LogFieldValueError, value: errors.Join(val...)}
}

// AnyFieldValue any
func AnyFieldValue(val any) LogFieldValue {
	switch vv := val.(type) {
	case int:
		return IntFieldValue(vv)
	case []int:
		return Int64ArrayFieldValue()
	case int8:
		return Int64FieldValue(int64(vv))
	case []int8:
		return Int64ArrayFieldValue(utils.ConvertIntSliceToInt64s[int8](vv)...)
	case int16:
		return Int64FieldValue(int64(vv))
	case []int16:
		return Int64ArrayFieldValue(utils.ConvertIntSliceToInt64s[int16](vv)...)
	case int32:
		return Int64FieldValue(int64(vv))
	case []int32:
		return Int64ArrayFieldValue(utils.ConvertIntSliceToInt64s[int32](vv)...)
	case int64:
		return Int64FieldValue(vv)
	case []int64:
		return Int64ArrayFieldValue(vv...)
	case uint:
		return Uint64FieldValue(uint64(vv))
	case []uint:
		return Uint64ArrayFieldValue(utils.ConvertUintSliceToUint64s[uint](vv)...)
	case uint8:
		return Uint64FieldValue(uint64(vv))
	case []uint8:
		return Uint64ArrayFieldValue(utils.ConvertUintSliceToUint64s[uint8](vv)...)
	case uint16:
		return Uint64FieldValue(uint64(vv))
	case []uint16:
		return Uint64ArrayFieldValue(utils.ConvertUintSliceToUint64s[uint16](vv)...)
	case uint32:
		return Uint64FieldValue(uint64(vv))
	case []uint32:
		return Uint64ArrayFieldValue(utils.ConvertUintSliceToUint64s[uint32](vv)...)
	case uint64:
		return Uint64FieldValue(vv)
	case []uint64:
		return Uint64ArrayFieldValue(vv...)
	case float32:
		return Float64FieldValue(float64(vv))
	case []float32:
		return Float64ArrayFieldValue(utils.ConvertFloatSliceToFloat64s(vv)...)
	case float64:
		return Float64FieldValue(vv)
	case []float64:
		return Float64ArrayFieldValue(vv...)
	case string:
		return StringFieldValue(vv)
	case []string:
		return StringArrayFieldValue(vv...)
	case bool:
		return BoolFieldValue(vv)
	case []bool:
		return BoolArrayFieldValue(vv...)
	case time.Time:
		return TimeFieldValue(vv)
	case time.Duration:
		return DurationFieldValue(vv)
	case LogField:
		return FieldFieldValue(vv)
	case []LogField:
		return FieldArrayFieldValue(vv...)
	case error:
		return ErrorFieldValue(vv)
	default:
		return LogFieldValue{kind: LogFieldValueAny, value: val}
	}
}
