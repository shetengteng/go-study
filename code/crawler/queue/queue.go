package queue

// QUEUE FIFO
type Queue []interface{}

func (q *Queue) Push(v int) *Queue {
	// 指针接收者，接收的q和返回的q是不同的，与其他面向对象的语言不一样
	*q = append(*q, v)
	return q
}

func (q *Queue) Pop() int {
	head := (*q)[0]
	*q = (*q)[1:]
	return head.(int)
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
