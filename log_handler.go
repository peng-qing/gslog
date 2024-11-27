package gslog

import (
	"context"
	"io"
	"sync"
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
	// Closer 需要实现 io.Closer 接口
	io.Closer
}

// commonHandler 应该是 LogHandler 实现的一个基类
type commonHandler struct {
	mutex       sync.Mutex
	writeSyncer io.WriteCloser
	options     *LogOptions
}

// newCommonHandler 实例化commonHandler方法 基类 不对外
func newCommonHandler(writeSyncer io.WriteCloser, opts ...Options) *commonHandler {
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
