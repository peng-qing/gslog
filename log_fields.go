package gslog

import (
	"encoding"
	"encoding/json"
	"time"

	"gslog/internal/bufferPool"
)

var (
	// 检查 LogField 实现 encoding.TextMarshaler
	_ encoding.TextMarshaler = (*LogField)(nil)
	// 检查 LogField 实现 json.Marshaler
	_ json.Marshaler = (*LogField)(nil)
)

type LogField struct {
	Key   string
	Value LogFieldValue
}

func String[T ~string | ~[]byte | ~[]rune](key string, val T) LogField {
	return LogField{
		Key:   key,
		Value: StringFieldValue(string(val)),
	}
}

func Int[T ~int | ~int8 | ~int16 | ~int32 | ~int64](key string, val T) LogField {
	return LogField{
		Key:   key,
		Value: Int64FieldValue(int64(val)),
	}
}

func Uint[T ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64](key string, val T) LogField {
	return LogField{
		Key:   key,
		Value: Uint64FieldValue(uint64(val)),
	}
}

func Bool[T ~bool](key string, val T) LogField {
	return LogField{
		Key:   key,
		Value: BoolFieldValue(bool(val)),
	}
}

func Float[T ~float32 | ~float64](key string, val T) LogField {
	return LogField{
		Key:   key,
		Value: Float64FieldValue(float64(val)),
	}
}

func Errors(key string, val ...error) LogField {
	return LogField{
		Key:   key,
		Value: ErrorFieldValue(val...),
	}
}

func Fields(key string, val ...LogField) LogField {
	fields := LogField{Key: key}
	if len(val) <= 1 {
		fields.Value = FieldFieldValue(val[0])
		return fields
	}
	fields.Value = FieldArrayFieldValue(val...)
	return fields
}

func Time(key string, val time.Time) LogField {
	return LogField{
		Key:   key,
		Value: TimeFieldValue(val),
	}
}

func Duration(key string, val time.Duration) LogField {
	return LogField{
		Key:   key,
		Value: DurationFieldValue(val),
	}
}

func Any(key string, val any) LogField {
	return LogField{
		Key:   key,
		Value: AnyFieldValue(val),
	}
}

// MarshalText 实现 encoding.TextMarshaler
// 序列化文本格式 key=value
func (l LogField) MarshalText() ([]byte, error) {
	buffer := bufferPool.Get()
	defer buffer.Free()

	buffer.AppendString(l.Key)
	if l.Value.Kind() == LogFieldValueField {
		// key.k=v
		buffer.AppendByte('.')
		data, err := l.Value.Field().MarshalText()
		if err != nil {
			return nil, err
		}
		buffer.AppendBytes(data)
		return buffer.Bytes(), nil
	}
	buffer.AppendByte('=')
	if l.Value.Kind() == LogFieldValueAny {
		// 值是 any 类型调用 尝试调用 encoding.TextMarshaler
		if vv, ok := l.Value.Any().(encoding.TextMarshaler); ok {
			data, err := vv.MarshalText()
			if err != nil {
				return nil, err
			}
			buffer.AppendBytes(data)
			return buffer.Bytes(), nil
		}
	}
	// default
	buffer.AppendString(l.Value.String())

	return buffer.Bytes(), nil
}

// MarshalJSON 实现 json.Marshaler
// 序列化Json格式 {"key":value}
func (l LogField) MarshalJSON() ([]byte, error) {
	buffer := bufferPool.Get()
	defer buffer.Free()

	// TODO 这里是否需要对空指针处理...

	allFieldKV := make(map[string]any)
	allFieldKV[l.Key] = l.Value.Any()

	data, err := json.Marshal(allFieldKV)
	if err != nil {
		return nil, err
	}
	buffer.AppendBytes(data)

	return buffer.Bytes(), nil
}
