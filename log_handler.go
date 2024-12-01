package gslog

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"reflect"
	"sync"
	"time"

	"gslog/internal/bufferPool"
	"gslog/pool"
)

var (
	// 检查 commonHandler 实现 LogHandler 接口
	_ LogHandler = (*commonHandler)(nil)
)

// LogHandler 日志处理
type LogHandler interface {
	//Enabled 针对每个 LogHandler 处理不同的日志级别
	Enabled(ctx context.Context, level LogLevel) bool
	// LogRecord 写入日志元数据
	LogRecord(ctx context.Context, entry *LogEntry) error
	// Sync 强制同步
	Sync() error
	// Closer 需要实现 io.Closer 接口
	io.Closer
}

// commonHandler 应该是 LogHandler 实现的一个基类
type commonHandler struct {
	mutex       sync.Mutex
	writeSyncer WriteSyncer
	options     *LogOptions
}

// newCommonHandler 实例化commonHandler方法 基类 不对外
func newCommonHandler(writeSyncer WriteSyncer, opts ...Options) *commonHandler {
	logOptions := &LogOptions{}

	for _, optFunc := range opts {
		optFunc.apply(logOptions)
	}

	return &commonHandler{
		writeSyncer: writeSyncer,
		options:     logOptions,
	}
}

// Enabled 判断日志是否需要输出
func (c *commonHandler) Enabled(ctx context.Context, level LogLevel) bool {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.options == nil {
		// 默认日志级别 InfoLevel
		return level >= InfoLevel
	}

	return c.options.Level >= level
}

// LogRecord 写入日志
func (c *commonHandler) LogRecord(ctx context.Context, entry *LogEntry) error {
	// TODO 由子类实现具体的效果
	return nil
}

// Close 关闭对应 Handler
func (c *commonHandler) Close() error {
	return c.writeSyncer.Close()
}

// WithOptions 修改日志配置参数
func (c *commonHandler) WithOptions(opts ...Options) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	for _, optFunc := range opts {
		optFunc.apply(c.options)
	}
}

// Sync 强制同步
func (c *commonHandler) Sync() error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	return c.writeSyncer.Sync()
}

// TextHandler 文本格式日志处理
type TextHandler struct {
	*commonHandler
}

// NewTextHandler 创建文本日志处理器
func NewTextHandler(writeSyncer WriteSyncer, opts ...Options) (*TextHandler, error) {
	instance := &TextHandler{
		commonHandler: newCommonHandler(writeSyncer, opts...),
	}

	return instance, nil
}

func (t *TextHandler) LogRecord(_ context.Context, entry *LogEntry) error {
	buffer := bufferPool.Get()
	defer buffer.Free()

	// 前缀
	if t.options.TextConf.TextPrefix != "" {
		buffer.AppendByte(serializePrefixBegin)
		buffer.AppendString(t.options.TextConf.TextPrefix)
		buffer.AppendByte(serializePrefixEnd)
		buffer.AppendByte(serializeSpaceSplit)
		// <prefix><space>
	}
	// 时间
	if !entry.Time.IsZero() && t.options.TextConf.TextFlag&LTextTime != 0 {
		layout := t.options.Layout
		if layout == "" {
			layout = DefaultTimeLayout
		}
		buffer.AppendString(entry.Time.Format(layout))
		buffer.AppendByte(serializeSpaceSplit)
		// <prefix> 2006/01/02 15:04:05.000000<space>
	}
	// LogLevel
	if t.options.TextConf.TextFlag&lCheckLogLevel != 0 {
		logLevel := entry.Level
		buffer.AppendByte(serializeArrayBegin)
		if t.options.TextConf.TextFlag&LTextLogLevel != 0 {
			buffer.AppendString(logLevel.CapitalString())
		} else if t.options.TextConf.TextFlag&LTextLogLevelUpCase != 0 {
			buffer.AppendString(logLevel.UpCaseString())
		} else {
			buffer.AppendString(logLevel.LowCaseString())
		}
		buffer.AppendByte(serializeArrayEnd)
		buffer.AppendByte(serializeSpaceSplit)
		// <prefix> 2006/01/02 15:04:05.000000 [Level]<space>
	}
	// File/Function
	if t.options.TextConf.TextFlag&lCheckShortFile != 0 {
		file, line, function := entry.Source()
		if t.options.TextConf.TextFlag&LTextFile != 0 {
			if file == "" {
				file = unknownFile
			}
			buffer.AppendString(file)
			buffer.AppendByte(serializeColonSplit)
			buffer.AppendInt(int64(line))
			buffer.AppendByte(serializeSpaceSplit)
			// <prefix> 2006/01/02 15:04:05.000000 [Level] file:line<space>
		}
		if t.options.TextConf.TextFlag&LTextFunction != 0 {
			buffer.AppendString(function)
			buffer.AppendByte(serializeSpaceSplit)
			// <prefix> 2006/01/02 15:04:05.000000 [Level] file:line function<space>
		}
	}
	// Message
	{
		buffer.AppendString(entry.Msg)
		buffer.AppendByte(serializeSpaceSplit)
		// <prefix> 2006/01/02 15:04:05.000000 [Level] file:line message<space>
	}
	// Fields
	for _, field := range entry.Fields {
		data, err := field.MarshalText()
		if err != nil {
			buffer.AppendString(fmt.Sprintf("%s=Error:%s", field.Key, err.Error()))
			continue
		}
		buffer.AppendBytes(data)
		buffer.AppendByte(serializeSpaceSplit)
		//  <prefix> 2024/06/11 10:00:00.000000 [Info] file:line function<space>message fieldKey=fieldValue...<space>
	}

	t.mutex.Lock()
	defer t.mutex.Unlock()
	_, err := t.writeSyncer.Write(buffer.Bytes())
	return err
}

// JsonHandler JSON格式日志处理
type JsonHandler struct {
	*commonHandler
}

// NewJsonHandler 实例化 JsonHandler
func NewJsonHandler(writeSyncer WriteSyncer, opts ...Options) *JsonHandler {
	instance := &JsonHandler{
		newCommonHandler(writeSyncer, opts...),
	}
	return instance
}

// LogRecord 记录日志
func (j *JsonHandler) LogRecord(_ context.Context, entry *LogEntry) error {
	buffer := bufferPool.Get()
	defer buffer.Free()

	j.mutex.Lock()
	defer j.mutex.Unlock()

	buffer.AppendByte(serializeJsonStart)
	// 时间
	if !entry.Time.IsZero() {
		layout := j.options.Layout
		if layout == "" {
			layout = DefaultTimeLayout
		}
		key := j.options.JSONConf.TimeEncodeKey
		if key == "" {
			key = defaultJsonTimeKey
		}
		j.appendJsonKey(buffer, key)
		j.appendJsonValue(buffer, entry.Time.Format(layout))
	}
	// source
	{
		key := j.options.JSONConf.SourceEncodeKey
		if key == "" {
			key = defaultJsonSourceKey
		}
		j.appendJsonKey(buffer, key)
		file, line, function := entry.Source()
		sourceStr := fmt.Sprintf("%s:%d %s", file, line, function)
		j.appendJsonValue(buffer, sourceStr)
	}
	// 日志级别
	{
		key := j.options.JSONConf.SourceEncodeKey
		if key == "" {
			key = defaultJsonLevelKey
		}
		j.appendJsonKey(buffer, key)
		j.appendJsonValue(buffer, entry.Level.LowCaseString())
	}
	// Message
	{
		key := j.options.JSONConf.SourceEncodeKey
		if key == "" {
			key = defaultJsonMessageKey
		}
		j.appendJsonKey(buffer, key)
		j.appendJsonValue(buffer, entry.Msg)
	}
	// fields...
	{
		key := j.options.JSONConf.FieldEncodeKey
		if key == "" {
			key = defaultJsonFieldsKey
		}
		j.appendJsonKey(buffer, key)
		j.appendJsonValue(buffer, entry.Fields)
	}
	buffer.AppendByte(serializeJsonEnd)
	buffer.AppendByte(serializeNewLine)

	_, err := j.writeSyncer.Write(buffer.Bytes())

	return err
}

// appendJsonKey 往buffer写入 "key":
func (j *JsonHandler) appendJsonKey(buffer *pool.Buffer, jsonKey string) {
	// json string has prefix '{'
	if buffer.Len() > 1 {
		buffer.AppendByte(serializeCommaStep)
	}
	// "key":
	buffer.AppendByte(serializeStringMarks)
	buffer.AppendString(jsonKey)
	buffer.AppendByte(serializeStringMarks)
	buffer.AppendByte(serializeColonSplit)
}

// appendJsonValue 往buffer写入Json value
func (j *JsonHandler) appendJsonValue(buffer *pool.Buffer, val any) {
	defer func() {
		if r := recover(); r != nil {
			if vv := reflect.ValueOf(val); vv.Kind() == reflect.Pointer && vv.IsNil() {
				buffer.AppendByte(serializeStringMarks)
				buffer.AppendString("<nil>")
				buffer.AppendByte(serializeStringMarks)
				return
			}
			buffer.AppendByte(serializeStringMarks)
			buffer.AppendString(fmt.Sprintf("Panic: %v", r))
			buffer.AppendByte(serializeStringMarks)
		}
	}()
	// 如果有定制
	switch vv := val.(type) {
	case time.Time:
		buffer.AppendByte(serializeStringMarks)
		layout := j.options.JSONConf.TimeEncodeKey
		if layout == "" {
			layout = DefaultTimeLayout
		}
		buffer.AppendString(vv.Format(layout))
		buffer.AppendByte(serializeStringMarks)
	case []LogField:
		// 其实这个可以不需要 LogField 已经实现了 json.Marshaler 接口
		// 在调用 jsonEncoder.Encode 的时候判断是否实现 json.Marshaler 然后会自动调用 MarshalJSON
		buffer.AppendByte(serializeArrayBegin)
		for idx, field := range vv {
			if idx > 0 {
				buffer.AppendByte(serializeCommaStep)
			}
			data, err := field.MarshalJSON()
			if err != nil {
				panic(err)
			}
			buffer.AppendBytes(data)
		}
		buffer.AppendByte(serializeArrayEnd)
	default:
		// 默认
		data, err := j.appendJsonMarshal(val)
		if err != nil {
			panic(err)
		}
		buffer.AppendBytes(data)
	}
}

func (j *JsonHandler) appendJsonMarshal(val any) ([]byte, error) {
	// 实现 json.Marshaler 接口
	if vv, ok := val.(json.Marshaler); ok {
		data, err := vv.MarshalJSON()
		if err != nil {
			return nil, err
		}
		return data, nil
	}

	buffer := bufferPool.Get()
	defer buffer.Free()

	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(val)
	if err != nil {
		return nil, err
	}
	buffer.TrimNewLine()
	return buffer.Bytes(), nil
}
