package bufferPool

import "gslog/pool"

var (
	// 全局变量
	_pool = pool.NewBufferPool()
	//Get 仅暴露Get方法获取缓存对象
	Get = _pool.Get()
)
