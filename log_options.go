package gslog

var (
	// 检查 optionFunc 是否实现 Options 接口
	_ Options = (*optionFunc)(nil)
)

type TextOptions struct {
	// 文本格式输出前缀
	TextPrefix string
	// 日志输出格式标志
	TextFlag LTextFlag
}

type JSONOptions struct {
	// 一些默认字段的key
	TimeEncodeKey    string
	SourceEncodeKey  string
	LevelEncodeKey   string
	MessageEncodeKey string
	FieldEncodeKey   string
}

type LogOptions struct {
	// 输出日志等级
	Level LogLevel
	// 日期输出格式
	Layout string
	// 文本日志控制
	TextConf *TextOptions
	// JSON日志配置
	JSONConf *JSONOptions
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
		logOptions.TextConf.TextPrefix = prefix
	})
}

// WithTextConf 设置文本日志相关配置
func WithTextConf(conf *TextOptions) Options {
	return optionFunc(func(logOptions *LogOptions) {
		logOptions.TextConf = conf
	})
}

// WithLayout 设置时间日期格式
func WithLayout(layout string) Options {
	return optionFunc(func(logOptions *LogOptions) {
		logOptions.Layout = layout
	})
}
