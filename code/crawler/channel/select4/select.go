package main

import (
	"fmt"
	"math/rand"
	"time"
)

func work(id int, c chan int) {
	for n := range c {
		time.Sleep(time.Second)
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
			time.Sleep(time.Duration(rand.Intn(3000)) * time.Millisecond)
			out <- i
			i++
		}
	}()
	return out
}

func main() {
	var c1, c2 = generator(), generator()

	var worker = createWorker(0)
	// 进行缓存，不同的消耗速度
	var values []int
	// 定时器
	tm := time.After(10 * time.Second) // 10s后结束，是一个channel，10后接收到一个数据
	tick := time.Tick(time.Second)
	i := 0
	for {
		var activeWorker chan int // nil channel 会阻塞
		var activeValue int
		if len(values) > 0 {
			activeWorker = worker
			activeValue = values[0]
		}

		fmt.Println("-----", i)
		i++

		select {
		case n := <-c1:
			values = append(values, n)
		case n := <-c2:
			values = append(values, n)
		case activeWorker <- activeValue: // 如果是nil channel则不作任何操作，阻塞跳过case
			values = values[1:]
		case <-time.After(800 * time.Millisecond): // 每次select进行一个超时统计，800ms,相邻2个时间的超时
			fmt.Println("timeout")
		case <-tick: // 每隔1s给一个值
			fmt.Println("queue len = ", len(values)) // 每隔1s统计len
		case <-tm:
			fmt.Println("finish") // 从程序开始总的时间
			return
		}
	}
}
