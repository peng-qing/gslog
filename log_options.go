package gslog

type LogOptions struct {
	Level LogLevel
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
