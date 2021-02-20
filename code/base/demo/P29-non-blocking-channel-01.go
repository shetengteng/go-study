package main

import (
	"fmt"
	"time"
)

func main() {
	// 无缓存channel，接收双方必须同时准备好
	messages := make(chan string)
	signals := make(chan bool)

	// 由于此时没有发送者，则直接进入default分支
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default: // 没有消息直接走到default处理
		fmt.Println("no message received")
	}

	msg := "hi"
	select {
	// 此时没有接收者，也直接进入到default
	case messages <- msg:
		fmt.Println("send message", msg)
	// 如果去除default，那么fatal error: all goroutines are asleep - deadlock!
	default:
		fmt.Println("no message send")
	}

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}

	time.Sleep(time.Second)
}
