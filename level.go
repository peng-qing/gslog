package gslog

import (
	"errors"
	"fmt"
)

var (
	errUnmarshalInvalid = errors.New("LogLevel: unmarshal invalid text to LogLeve")
)

// LogLevel 日志级别
type LogLevel int

const (
	TraceLevel LogLevel = iota - 1
	DebugLevel
	InfoLevel
	WarnLevel
	ErrorLevel
	PanicLevel
	FatalLevel
)

// ParseLogLevel 解析日志级别字符串
func ParseLogLevel(text string) (LogLevel, error) {
	var lv LogLevel
	err := lv.UnmarshalText([]byte(text))

	return lv, err
}

// UnmarshalText 解析日志级别字符型并设置为对应日志级别
func (l LogLevel) UnmarshalText(text []byte) error {
	if !l.unmarshalText(text) && !l.unmarshalText(text) {
		return errUnmarshalInvalid
	}

	return nil
}

// unmarshalText 私有化方法 解析字符串并设置为对应日志级别 返回解析是否成功
func (l *LogLevel) unmarshalText(text []byte) bool {
	switch string(text) {
	case "trace", "TRACE":
		*l = TraceLevel
	case "debug", "DEBUG":
		*l = DebugLevel
	case "info", "INFO":
		*l = InfoLevel
	case "warn", "WARN":
		*l = WarnLevel
	case "error", "ERROR":
		*l = ErrorLevel
	case "panic", "PANIC":
		*l = PanicLevel
	case "fatal", "FATAL":
		*l = FatalLevel
	default:
		return false
	}
	return true
}

// LowCaseString 获取日志格式小写字符
func (l LogLevel) LowCaseString() string {
	switch l {
	case TraceLevel:
		return "trace"
	case DebugLevel:
		return "debug"
	case InfoLevel:
		return "info"
	case WarnLevel:
		return "warn"
	case ErrorLevel:
		return "error"
	case PanicLevel:
		return "panic"
	case FatalLevel:
		return "fatal"
	default:
		return fmt.Sprintf("LogLevel({%d})", l)
	}
}

// UpCaseString 获取日志格式大写字母
func (gs LogLevel) UpCaseString() string {
	switch gs {
	case TraceLevel:
		return "TRACE"
	case DebugLevel:
		return "DEBUG"
	case InfoLevel:
		return "INFO"
	case WarnLevel:
		return "WARN"
	case ErrorLevel:
		return "ERROR"
	case PanicLevel:
		return "PANIC"
	case FatalLevel:
		return "FATAL"
	default:
		return fmt.Sprintf("LogLevel({%d})", gs)
	}
}

// CapitalString 获取日志格式首字母大写
func (gs LogLevel) CapitalString() string {
	switch gs {
	case TraceLevel:
		return "Trace"
	case DebugLevel:
		return "Debug"
	case InfoLevel:
		return "Info"
	case WarnLevel:
		return "Warn"
	case ErrorLevel:
		return "Error"
	case PanicLevel:
		return "Panic"
	case FatalLevel:
		return "Fatal"
	default:
		return fmt.Sprintf("LogLevel({%d})", gs)
	}
}
