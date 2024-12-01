package gslog

import (
	"runtime"
	"time"
)

type LogEntry struct {
	Time   time.Time
	Level  LogLevel
	PC     uintptr
	Msg    string
	Fields []LogField
}

// NewLogEntry 实例化日志实体
func NewLogEntry(t time.Time, level LogLevel, msg string, pc uintptr) *LogEntry {
	return &LogEntry{
		Time:   t,
		Level:  level,
		PC:     pc,
		Msg:    msg,
		Fields: make([]LogField, 0),
	}
}

// AppendFields 添加日志字段
func (l *LogEntry) AppendFields(fields ...LogField) {
	l.Fields = append(l.Fields, fields...)
}

// Source 日志源
func (l *LogEntry) Source() (file string, line int, function string) {
	frames := runtime.CallersFrames([]uintptr{l.PC})
	frame, _ := frames.Next()

	return frame.File, frame.Line, frame.Function
}

// AddArgs 添加参数
func (l *LogEntry) AddArgs(args ...any) {
	var field LogField
	for len(args) > 0 {
		field, args = l.argsToLogFields(args...)
		l.AppendFields(field)
	}
}

func (l *LogEntry) argsToLogFields(args ...any) (LogField, []any) {
	switch vv := args[0].(type) {
	case LogField:
		return vv, args[1:]
	case string:
		if len(args) <= 1 {
			return String[string](badFieldsKey, vv), nil
		}
		return Any(vv, args[1:]), args[2:]
	default:
		return Any(badFieldsKey, vv), args[1:]
	}
}
