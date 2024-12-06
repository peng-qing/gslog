package gslog

var (
	// 检查 optionFunc 是否实现 Options 接口
	_ Options = (*optionFunc)(nil)
)

type LogOptions struct {
	// 输出日志等级
	Level LogLevel `json:"level"`
	// 日期输出格式
	Layout string `json:"layout"`

	// 文本格式输出前缀
	TextPrefix string `json:"text_prefix"`
	// 日志输出格式标志
	TextFlag LTextFlag `json:"text_flag"`

	// Json格式一些默认字段的key
	TimeEncodeKey    string `json:"time_encode_key"`
	SourceEncodeKey  string `json:"source_encode_key"`
	LevelEncodeKey   string `json:"level_encode_key"`
	MessageEncodeKey string `json:"message_encode_key"`
	FieldEncodeKey   string `json:"field_encode_key"`
}

// Options Option模式接口
type Options interface {
	apply(logOptions *LogOptions)
}

// 自定义 Options func 格式
type optionFunc func(logOptions *LogOptions)

// 实现 Options 接口
func (optFunc optionFunc) apply(logOptions *LogOptions) {
	optFunc(logOptions)
}

// WithLevel 设置日志级别
func WithLevel(level LogLevel) Options {
	return optionFunc(func(logOptions *LogOptions) {
		logOptions.Level = level
	})
}

// WithPrefix 设置文本日志前缀
func WithPrefix(prefix string) Options {
	return optionFunc(func(logOptions *LogOptions) {
		logOptions.TextPrefix = prefix
	})
}

// WithTextFlag 设置文本日志相关配置
func WithTextFlag(flag LTextFlag) Options {
	return optionFunc(func(logOptions *LogOptions) {
		logOptions.TextFlag = flag
	})
}

// WithLayout 设置时间日期格式
func WithLayout(layout string) Options {
	return optionFunc(func(logOptions *LogOptions) {
		logOptions.Layout = layout
	})
}

// WithTimeEncodeKey 设置Json key
func WithTimeEncodeKey(key string) Options {
	return optionFunc(func(logOptions *LogOptions) {
		logOptions.TimeEncodeKey = key
	})
}

// WithSourceEncodeKey 设置Json key
func WithSourceEncodeKey(key string) Options {
	return optionFunc(func(logOptions *LogOptions) {
		logOptions.SourceEncodeKey = key
	})
}

// WithLevelEncodeKey 设置Json key
func WithLevelEncodeKey(key string) Options {
	return optionFunc(func(logOptions *LogOptions) {
		logOptions.LevelEncodeKey = key
	})
}

// WithMessageEncodeKey 设置Json key
func WithMessageEncodeKey(key string) Options {
	return optionFunc(func(logOptions *LogOptions) {
		logOptions.MessageEncodeKey = key
	})
}

// WithFieldEncodeKey 设置Json key
func WithFieldEncodeKey(key string) Options {
	return optionFunc(func(logOptions *LogOptions) {
		logOptions.FieldEncodeKey = key
	})
}
