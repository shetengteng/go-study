package main

import (
	"fmt"
	"sync"
)

type atomicInt struct {
	value int
	lock  sync.Mutex
}

func (a *atomicInt) increment() {
	a.lock.Lock()
	defer a.lock.Unlock()
	a.value++
}

// 如果要将锁控制在代码块中，需要一个匿名函数
func (a *atomicInt) increment2() {
	func() {
		a.lock.Lock()
		defer a.lock.Unlock()
		a.value++
	}()
}

func (a *atomicInt) get() int {
	a.lock.Lock()
	defer a.lock.Unlock()
	return int(*a)
}

func main() {

	var a atomicInt
	a.increment()
	go func() {
		a.increment()
	}()

	fmt.Println(a.get())
}

// 使用go run -race atomic.go 查看风险
