package lru

import "container/list"

// 当前是线程非安全的，在并发访问时有问题
type Cache struct {
	maxBytes  int64                         // 允许使用的最大内存，如果设置为0表示无限制
	usedBytes int64                         // 已使用的内存
	ll        *list.List                    // 双向链表，存储节点，用于lru操作
	cache     map[string]*list.Element      // map存储每个链表节点信息，key是节点的名称，便于查询
	OnEvicted func(key string, value Value) // 回调函数，当一个key的元素被移除时进行回调处理，可以是nil
}

// 存储的数据单元
type entry struct {
	key   string
	value Value
}

// 存储值的长度，占用多少bytes，此处Value是一个接口，不用特定指定数据类型，通用性高
type Value interface {
	Len() int
}

// 初始化
func New(maxBytes int64, onEvicted func(string, Value)) *Cache {
	return &Cache{
		maxBytes:  maxBytes,
		ll:        list.New(),
		cache:     make(map[string]*list.Element),
		OnEvicted: onEvicted,
	}
}

// 查询
// 从map中查找得到节点，然后在双向列表中放到头部
func (c *Cache) Get(key string) (value Value, ok bool) {
	if element, ok := c.cache[key]; ok {
		// 得到链表的节点，将节点移动
		c.ll.MoveToFront(element)
		// 类型断言，转换为 entry的指针类型
		kv := element.Value.(*entry)
		return kv.value, true
	}
	return
}

// 删除操作
// 先从链表中删除，然后在map中删除
func (c *Cache) RemoveOldest() {
	back := c.ll.Back()
	if back != nil {
		c.ll.Remove(back) // 从字典中删除元素
		kv := back.Value.(*entry)
		delete(c.cache, kv.key) // 删除节点的映射关系
		c.usedBytes -= int64(len(kv.key)) + int64(kv.value.Len())
		if c.OnEvicted != nil {
			c.OnEvicted(kv.key, kv.value)
		}
	}
}

// 新增
func (c *Cache) Add(key string, value Value) {
	// 判断是否存在，存在则更新value
	// 不存在则放入map，然后放到链表的头部
	// 给 userBytes 添加值，判断是否超过maxBytes 超过则删除
	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		c.usedBytes += int64(value.Len()) - int64(kv.value.Len())
		kv.value = value
	} else {
		c.cache[key] = c.ll.PushFront(&entry{key, value})
		c.usedBytes += int64(len(key)) + int64(value.Len())
	}

	for c.maxBytes != 0 && c.maxBytes < c.usedBytes {
		c.RemoveOldest()
	}
}

func (c *Cache) Len() int {
	return c.ll.Len()
}
