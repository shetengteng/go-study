package main

import (
	"fmt"
	"math/rand"
	"time"
)

func work(id int, c chan int) {
	for n := range c {
		fmt.Printf("worker %d received %d \n", id, n)
	}
}
func createWorker(id int) chan int {
	c := make(chan int)
	go work(id, c)
	return c
}

func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			out <- i
			i++
		}
	}()
	return out
}

func main() {
	var c1, c2 = generator(), generator()

	var worker = createWorker(0)

	n := 0
	hasValue := false
	for {
		var activeWorker chan int // nil channel 会阻塞
		if hasValue {
			activeWorker = worker
		}

		select {
		case n = <-c1: // 如果生成的速率过快，消耗的速率慢，n收到了1,2,3,那么最后只有3生效
			hasValue = true
		case n = <-c2:
			hasValue = true
		case activeWorker <- n: // 如果是nil channel则不作任何操作，阻塞住，跳过case
			hasValue = false
		}
	}
}
