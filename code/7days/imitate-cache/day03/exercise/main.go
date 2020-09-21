package main

import (
	"fmt"
	"sync"
	"time"
)

// 结构体对象可以直接声明创建一个对象
var m sync.Mutex

// 开启100个 goroutine打印一次
var set = make(map[int]bool, 0)

func main() {

	for i := 0; i < 100; i++ {
		// 如果没有上锁，则会有多个 99 ,同时有异常的产生 fatal error: concurrent map writes
		go printOnce(99)
	}

	time.Sleep(time.Second)
}

func printOnce(num int) {
	// 添加锁
	m.Lock()
	if _, ok := set[num]; ok {
		return
	}
	fmt.Println(num)
	set[num] = true
	m.Unlock()
}
