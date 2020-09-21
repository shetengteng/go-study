package stt

import (
	"stt/lru"
	"sync"
)

// 并发控制
type cache struct {
	mu       sync.Mutex // 此处是排他锁，而非读写锁的原因是在get 和 add 都涉及到lru中的数据的移动
	lru      *lru.Cache
	capacity int64 // 缓存大小
}

func (c *cache) add(key string, value ByteView) {
	c.mu.Lock()
	if c.lru == nil {
		c.lru = lru.New(c.capacity, nil)
	}
	c.lru.Add(key, value)
	c.mu.Unlock()
}

func (c *cache) get(key string) (value ByteView, ok bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.lru == nil {
		return
	}
	if v, ok := c.lru.Get(key); ok {
		return v.(ByteView), ok
	}
	return
}
