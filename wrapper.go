package gslog

import (
	"context"
	"sync"
	"sync/atomic"
)

var (
	defaultLogger atomic.Pointer[Logger]
	once          sync.Once
)

func init() {
	//once.Do(func() {
	//	defaultLogger.Store()
	//})
}

// Default 默认全局日志器
func Default() *Logger {
	return defaultLogger.Load()
}

// SetDefault 设置全局默认日志器
func SetDefault(logger *Logger) {
	defaultLogger.Store(logger)
}

// Trace 格式化输出 TraceLevel 级别日志
func Trace(msg string, args ...any) {
	Default().log(context.Background(), TraceLevel, msg, args...)
}

// Debug 格式化输出 DebugLevel 级别日志
func Debug(msg string, args ...any) {
	Default().log(context.Background(), DebugLevel, msg, args...)
}

// Info 格式化输出 InfoLevel 级别日志
func Info(msg string, args ...any) {
	Default().log(context.Background(), InfoLevel, msg, args...)
}

// Warn 格式化输出 WarnLevel 级别日志
func Warn(msg string, args ...any) {
	Default().log(context.Background(), WarnLevel, msg, args...)
}

// Error 格式化输出 ErrorLevel 级别日志
func Error(msg string, args ...any) {
	Default().log(context.Background(), ErrorLevel, msg, args...)
}

// Panic 格式化输出 PanicLevel 级别日志
func Panic(msg string, args ...any) {
	Default().log(context.Background(), PanicLevel, msg, args...)
}

// Fatal 格式化输出 FatalLevel 级别日志
func Fatal(msg string, args ...any) {
	Default().log(context.Background(), FatalLevel, msg, args...)
}

// TraceContext 格式化输出 TraceLevel 级别日志
func TraceContext(ctx context.Context, msg string, args ...any) {
	Default().log(ctx, TraceLevel, msg, args...)
}

// DebugContext 格式化输出 DebugLevel 级别日志
func DebugContext(ctx context.Context, msg string, args ...any) {
	Default().log(ctx, DebugLevel, msg, args...)
}

// InfoContext 格式化输出 InfoLevel 级别日志
func InfoContext(ctx context.Context, msg string, args ...any) {
	Default().log(ctx, InfoLevel, msg, args...)
}

// WarnContext 格式化输出 WarnLevel 级别日志
func WarnContext(ctx context.Context, msg string, args ...any) {
	Default().log(ctx, WarnLevel, msg, args...)
}

// ErrorContext 格式化输出 ErrorLevel 级别日志
func ErrorContext(ctx context.Context, msg string, args ...any) {
	Default().log(ctx, ErrorLevel, msg, args...)
}

// PanicContext 格式化输出 PanicLevel 级别日志
func PanicContext(ctx context.Context, msg string, args ...any) {
	Default().log(ctx, PanicLevel, msg, args...)
}

// FatalContext 格式化输出 FatalLevel 级别日志
func FatalContext(ctx context.Context, msg string, args ...any) {
	Default().log(ctx, FatalLevel, msg, args...)
}

func TraceFields(msg string, args ...LogField) {
	Default().logFields(context.Background(), TraceLevel, msg, args...)
}

func DebugFields(msg string, args ...LogField) {
	Default().logFields(context.Background(), DebugLevel, msg, args...)
}

func InfoFields(msg string, args ...LogField) {
	Default().logFields(context.Background(), InfoLevel, msg, args...)
}

func WarnFields(msg string, args ...LogField) {
	Default().logFields(context.Background(), WarnLevel, msg, args...)
}

func ErrorFields(msg string, args ...LogField) {
	Default().logFields(context.Background(), ErrorLevel, msg, args...)
}

func PanicFields(msg string, args ...LogField) {
	Default().logFields(context.Background(), PanicLevel, msg, args...)
}

func FatalFields(msg string, args ...LogField) {
	Default().logFields(context.Background(), FatalLevel, msg, args...)
}

func TraceFieldsContext(ctx context.Context, msg string, args ...LogField) {
	Default().logFields(ctx, TraceLevel, msg, args...)
}

func DebugFieldsContext(ctx context.Context, msg string, args ...LogField) {
	Default().logFields(ctx, DebugLevel, msg, args...)
}

func InfoFieldsContext(ctx context.Context, msg string, args ...LogField) {
	Default().logFields(ctx, InfoLevel, msg, args...)
}

func WarnFieldsContext(ctx context.Context, msg string, args ...LogField) {
	Default().logFields(ctx, WarnLevel, msg, args...)
}

func ErrorFieldsContext(ctx context.Context, msg string, args ...LogField) {
	Default().logFields(ctx, ErrorLevel, msg, args...)
}

func PanicFieldsContext(ctx context.Context, msg string, args ...LogField) {
	Default().logFields(ctx, PanicLevel, msg, args...)
}

func FatalFieldsContext(ctx context.Context, msg string, args ...LogField) {
	Default().logFields(ctx, FatalLevel, msg, args...)
}
