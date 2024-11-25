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
	writeSyncer io.Writer
	options     *LogOptions
}

// newCommonHandler 实例化commonHandler方法 不对外
func newCommonHandler(writeSyncer io.Writer, opts ...Options) *commonHandler {
	logOptions := &LogOptions{}

	for _, optFunc := range opts {
		optFunc.apply(logOptions)
	}

	return &commonHandler{
		writeSyncer: writeSyncer,
		options:     logOptions,
	}
}

func (c *commonHandler) Enabled(ctx context.Context, level LogLevel) bool {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.options == nil {
		// 默认日志级别 InfoLevel
		return level >= InfoLevel
	}

	return c.options.Level >= level
}

func (c *commonHandler) LogRecord(ctx context.Context, entry *LogEntry) error {
	//TODO implement me
	panic("implement me")
}

func (c *commonHandler) Close() error {
	panic("implement me")
}
