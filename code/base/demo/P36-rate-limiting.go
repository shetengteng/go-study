package main

import (
	"fmt"
	"time"
)

func main() {
	// 同时产生5个请求
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	// Tick 是 time.NewTicker的封装
	limiter := time.Tick(200 * time.Millisecond)

	for req := range requests {
		// 固定速率从requests中获取信息，进行消费处理
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	fmt.Println("---------")

	// 示例2 上个例子是串行执行，这个例子允许同时处理3个请求
	burstyLimiter := make(chan time.Time, 3)

	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}
	// 开启一个协程，固定速率在limiter中存放请求执行标识
	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			// 如果 chan 满，则阻塞
			burstyLimiter <- t
		}
	}()

	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)

	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}
