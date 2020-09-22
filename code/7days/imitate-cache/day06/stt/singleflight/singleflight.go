package singleflight

import "sync"

// 主要功能：防止重复查询

// call 代表正在进行中，或已经结束的请求。使用 sync.WaitGroup 锁避免重入
type call struct {
	wg  sync.WaitGroup
	val interface{}
	err error
}

// Group 是 singleflight 的主数据结构，管理不同 key 的请求(call)。
type Group struct {
	mu sync.Mutex // 对map进行锁保护
	m  map[string]*call
}

// 多个线程等待一个线程执行完成后返回结果
// 对同一个key,fn只执行一次
func (g *Group) Do(key string, fn func() (interface{}, error)) (interface{}, error) {

	g.mu.Lock()
	if g.m == nil {
		g.m = make(map[string]*call)
	}
	if c, ok := g.m[key]; ok {
		g.mu.Unlock()
		// 该线程进行等待，说明该key已经有线程在处理中
		c.wg.Wait()
		// 等线程处理完成，返回结果，每个线程都返回一个结果
		return c.val, c.err
	}
	//如果key不存在，则将key放入map

	c := new(call)
	// 添加一个线程计数放入waitGroup中，当所有线程都done的时候(这里就1个)，处于wait位置的线程开始执行
	c.wg.Add(1)
	g.m[key] = c
	g.mu.Unlock()

	// 线程继续执行
	c.val, c.err = fn()
	// 线程执行完成，通知wait位置的线程继续执行
	c.wg.Done()

	g.mu.Lock()
	delete(g.m, key)
	g.mu.Unlock()

	return c.val, c.err
}
