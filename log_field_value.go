package gslog

import (
	"errors"
	"fmt"
	"time"

	"gslog/internal/utils"
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
		return Int64ArrayFieldValue(utils.IntSliceToInt64[int](vv)...)
	case int8:
		return Int64FieldValue(int64(vv))
	case []int8:
		return Int64ArrayFieldValue(utils.IntSliceToInt64[int8](vv)...)
	case int16:
		return Int64FieldValue(int64(vv))
	case []int16:
		return Int64ArrayFieldValue(utils.IntSliceToInt64[int16](vv)...)
	case int32:
		return Int64FieldValue(int64(vv))
	case []int32:
		return Int64ArrayFieldValue(utils.IntSliceToInt64[int32](vv)...)
	case int64:
		return Int64FieldValue(vv)
	case []int64:
		return Int64ArrayFieldValue(vv...)
	case uint:
		return Uint64FieldValue(uint64(vv))
	case []uint:
		return Uint64ArrayFieldValue(utils.UintSliceToUint64[uint](vv)...)
	case uint8:
		return Uint64FieldValue(uint64(vv))
	case []uint8:
		return Uint64ArrayFieldValue(utils.UintSliceToUint64[uint8](vv)...)
	case uint16:
		return Uint64FieldValue(uint64(vv))
	case []uint16:
		return Uint64ArrayFieldValue(utils.UintSliceToUint64[uint16](vv)...)
	case uint32:
		return Uint64FieldValue(uint64(vv))
	case []uint32:
		return Uint64ArrayFieldValue(utils.UintSliceToUint64[uint32](vv)...)
	case uint64:
		return Uint64FieldValue(vv)
	case []uint64:
		return Uint64ArrayFieldValue(vv...)
	case float32:
		return Float64FieldValue(float64(vv))
	case []float32:
		return Float64ArrayFieldValue(utils.FloatToFloat64(vv)...)
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
	case fmt.Stringer:
		return StringFieldValue(vv.String())
	default:
		return LogFieldValue{kind: LogFieldValueAny, value: val}
	}
}

// Kind 获取存储的具体值类型
func (l LogFieldValue) Kind() LogFieldValueKind {
	return l.kind
}

func (l LogFieldValue) Int64() int64 {
	if current, target := l.Kind(), LogFieldValueInt64; current != target {
		panic(fmt.Sprintf("current FieldValueKind is %s, not %s", kindStrings[current], kindStrings[target]))
	}

	return l.value.(int64)
}

func (l LogFieldValue) Int64s() []int64 {
	if current, target := l.Kind(), LogFieldValueInt64s; current != target {
		panic(fmt.Sprintf("current FieldValueKind is %s, not %s", kindStrings[current], kindStrings[target]))
	}
	return l.value.([]int64)
}

func (l LogFieldValue) Uint64() uint64 {
	if current, target := l.Kind(), LogFieldValueUint64; current != target {
		panic(fmt.Sprintf("current FieldValueKind is %s, not %s", kindStrings[current], kindStrings[target]))
	}
	return l.value.(uint64)
}

func (l LogFieldValue) Uint64s() []uint64 {
	if current, target := l.Kind(), LogFieldValueUint64s; current != target {
		panic(fmt.Sprintf("current FieldValueKind is %s, not %s", kindStrings[current], kindStrings[target]))
	}
	return l.value.([]uint64)
}

func (l LogFieldValue) Float64() float64 {
	if current, target := l.Kind(), LogFieldValueFloat64; current != target {
		panic(fmt.Sprintf("current FieldValueKind is %s, not %s", kindStrings[current], kindStrings[target]))
	}
	return l.value.(float64)
}

func (l LogFieldValue) Float64s() []float64 {
	if current, target := l.Kind(), LogFieldValueFloat64s; current != target {
		panic(fmt.Sprintf("current FieldValueKind is %s, not %s", kindStrings[current], kindStrings[target]))
	}
	return l.value.([]float64)
}

// String never panic
func (l LogFieldValue) String() string {
	if current, target := l.Kind(), LogFieldValueString; current == target {
		return l.value.(string)
	}

	buffer := make([]byte, 0)
	return string(l.appendFieldValue(buffer))
}

func (l LogFieldValue) Strings() []string {
	if current, target := l.Kind(), LogFieldValueStrings; current != target {
		panic(fmt.Sprintf("current FieldValueKind is %s, not %s", kindStrings[current], kindStrings[target]))
	}
	return l.value.([]string)
}

func (l LogFieldValue) Bool() bool {
	if current, target := l.Kind(), LogFieldValueBool; current != target {
		panic(fmt.Sprintf("current FieldValueKind is %s, not %s", kindStrings[current], kindStrings[target]))
	}
	return l.value.(bool)
}

func (l LogFieldValue) Bools() []bool {
	if current, target := l.Kind(), LogFieldValueBools; current != target {
		panic(fmt.Sprintf("current FieldValueKind is %s, not %s", kindStrings[current], kindStrings[target]))
	}
	return l.value.([]bool)
}

func (l LogFieldValue) Time() time.Time {
	if current, target := l.Kind(), LogFieldValueTime; current != target {
		panic(fmt.Sprintf("current FieldValueKind is %s, not %s", kindStrings[current], kindStrings[target]))
	}
	return l.value.(time.Time)
}

func (l LogFieldValue) Duration() time.Duration {
	if current, target := l.Kind(), LogFieldValueDuration; current != target {
		panic(fmt.Sprintf("current FieldValueKind is %s, not %s", kindStrings[current], kindStrings[target]))
	}
	return l.value.(time.Duration)
}

func (l LogFieldValue) Field() LogField {
	if current, target := l.Kind(), LogFieldValueField; current != target {
		panic(fmt.Sprintf("current FieldValueKind is %s, not %s", kindStrings[current], kindStrings[target]))
	}
	return l.value.(LogField)
}

func (l LogFieldValue) Fields() []LogField {
	if current, target := l.Kind(), LogFieldValueFields; current != target {
		panic(fmt.Sprintf("current FieldValueKind is %s, not %s", kindStrings[current], kindStrings[target]))
	}
	return l.value.([]LogField)
}

func (l LogFieldValue) Error() error {
	if current, target := l.Kind(), LogFieldValueError; current != target {
		panic(fmt.Sprintf("current FieldValueKind is %s, not %s", kindStrings[current], kindStrings[target]))
	}
	return l.value.(error)
}

func (l LogFieldValue) Any() any {
	switch l.Kind() {
	case LogFieldValueAny:
		return l.value
	case LogFieldValueInt64:
		return l.Int64()
	case LogFieldValueInt64s:
		return l.Int64s()
	case LogFieldValueUint64:
		return l.Uint64()
	case LogFieldValueUint64s:
		return l.Uint64s()
	case LogFieldValueFloat64:
		return l.Float64()
	case LogFieldValueFloat64s:
		return l.Float64s()
	case LogFieldValueString:
		return l.String()
	case LogFieldValueStrings:
		return l.Strings()
	case LogFieldValueBool:
		return l.Bool()
	case LogFieldValueBools:
		return l.Bools()
	case LogFieldValueTime:
		return l.Time()
	case LogFieldValueDuration:
		return l.Duration()
	case LogFieldValueField:
		return l.Field()
	case LogFieldValueFields:
		return l.Fields()
	case LogFieldValueError:
		return l.Error()
	default:
		panic(fmt.Sprintf("unknown kind %s", l.Kind()))
	}
}

func (l LogFieldValue) appendFieldValue(dst []byte) []byte {
	return nil
}
