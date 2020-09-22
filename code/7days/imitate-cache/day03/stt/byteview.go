package stt

// 存储的是不可变byte数组
type ByteView struct {
	// 存储真实的缓存值，byte类型可以存储任意类型的值，图片，字符串等
	// 作为lru 中的 Value部分，实现Value的Len接口，返回内存的大小
	// 注意：b 是切片对象，因此ByteView在深拷贝对内存消耗不大
	b []byte
}

func (v ByteView) Len() int {
	return len(v.b)
}

// 值拷贝切片
// 防止缓存值被修改，b是只读的
func (v ByteView) ByteSlice() []byte {
	return cloneBytes(v.b)
}

func (v ByteView) String() string {
	return string(v.b)
}

func cloneBytes(b []byte) []byte {
	c := make([]byte, len(b))
	copy(c, b)
	return c
}
