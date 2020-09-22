package consistenthash

import (
	"hash/crc32"
	"sort"
	"strconv"
)

// hash 算法，将data转换为uint32数字
type Hash func(data []byte) uint32

type Map struct {
	hash     Hash // hash 算法
	replicas int  // 虚拟节点倍数，用于防止数据倾斜，多少个虚拟节点

	keys    []int          // hash环
	hashMap map[int]string // 存储虚拟节点和真实节点的映射表，key 是虚拟节点的hash值，value是真实节点的名称
}

// 构造器
func New(replicas int, fn Hash) *Map {
	m := &Map{
		replicas: replicas,
		hash:     fn,
		hashMap:  make(map[int]string),
	}
	if m.hash == nil {
		m.hash = crc32.ChecksumIEEE
	}
	return m
}

// 添加虚拟节点
func (m *Map) Add(keys ...string) {
	for _, key := range keys {
		for i := 0; i < m.replicas; i++ {
			// 使用Map中的hash算法计算出hashKey
			hashKey := int(m.hash([]byte(strconv.Itoa(i) + key)))
			m.keys = append(m.keys, hashKey)
			m.hashMap[hashKey] = key
		}
	}
	// 进行排序
	sort.Ints(m.keys)
}

// 选择节点，通过key找到最近的节点，返回节点的名称
func (m *Map) Get(key string) string {
	if len(m.keys) == 0 {
		return ""
	}
	hashKey := int(m.hash([]byte(key)))
	// 通过二分查找法获取最近的点
	idx := sort.Search(len(m.keys), func(i int) bool {
		return m.keys[i] >= hashKey
	})

	//顺时针找到第一个匹配的虚拟节点的下标 idx，从 m.keys 中获取到对应的哈希值。如果 idx == len(m.keys)，说明应选择 m.keys[0]，因为 m.keys 是一个环状结构，所以用取余数的方式来处理这种情况
	return m.hashMap[m.keys[idx%len(m.keys)]]
}
