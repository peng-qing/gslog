package gslog

import (
	"context"
)

// Logger 日志器
type Logger struct {
	handler LogHandler
}

// NewLogger 实例化日志器
func NewLogger(handler LogHandler) *Logger {
	return &Logger{
		handler: handler,
	}
}

// Trace 格式化输出 TraceLevel 级别日志
func (l *Logger) Trace(msg string, args ...any) {
	l.log(context.Background(), TraceLevel, msg, args...)
}

// Debug 格式化输出 DebugLevel 级别日志
func (l *Logger) Debug(msg string, args ...any) {
	l.log(context.Background(), DebugLevel, msg, args...)
}

// Info 格式化输出 InfoLevel 级别日志
func (l *Logger) Info(msg string, args ...any) {
	l.log(context.Background(), InfoLevel, msg, args...)
}

// Warn 格式化输出 WarnLevel 级别日志
func (l *Logger) Warn(msg string, args ...any) {
	l.log(context.Background(), WarnLevel, msg, args...)
}

// Error 格式化输出 ErrorLevel 级别日志
func (l *Logger) Error(msg string, args ...any) {
	l.log(context.Background(), ErrorLevel, msg, args...)
}

// Panic 格式化输出 PanicLevel 级别日志
func (l *Logger) Panic(msg string, args ...any) {
	l.log(context.Background(), PanicLevel, msg, args...)
}

// Fatal 格式化输出 FatalLevel 级别日志
func (l *Logger) Fatal(msg string, args ...any) {
	l.log(context.Background(), FatalLevel, msg, args...)
}

// TraceContext 格式化输出 TraceLevel 级别日志
func (l *Logger) TraceContext(ctx context.Context, msg string, args ...any) {
	l.log(ctx, TraceLevel, msg, args...)
}

// DebugContext 格式化输出 DebugLevel 级别日志
func (l *Logger) DebugContext(ctx context.Context, msg string, args ...any) {
	l.log(ctx, DebugLevel, msg, args...)
}

// InfoContext 格式化输出 InfoLevel 级别日志
func (l *Logger) InfoContext(ctx context.Context, msg string, args ...any) {
	l.log(ctx, InfoLevel, msg, args...)
}

// WarnContext 格式化输出 WarnLevel 级别日志
func (l *Logger) WarnContext(ctx context.Context, msg string, args ...any) {
	l.log(ctx, WarnLevel, msg, args...)
}

// ErrorContext 格式化输出 ErrorLevel 级别日志
func (l *Logger) ErrorContext(ctx context.Context, msg string, args ...any) {
	l.log(ctx, ErrorLevel, msg, args...)
}

// PanicContext 格式化输出 PanicLevel 级别日志
func (l *Logger) PanicContext(ctx context.Context, msg string, args ...any) {
	l.log(ctx, PanicLevel, msg, args...)
}

// FatalContext 格式化输出 FatalLevel 级别日志
func (l *Logger) FatalContext(ctx context.Context, msg string, args ...any) {
	l.log(ctx, FatalLevel, msg, args...)
}

func (l *Logger) log(ctx context.Context, level LogLevel, msg string, args ...any) {
}
