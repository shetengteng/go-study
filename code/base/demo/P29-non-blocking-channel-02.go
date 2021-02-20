package main

import (
	"fmt"
	"time"
)

func main() {
	// 无缓存channel，接收双方必须同时准备好
	// 如果有缓存，则不需要，都准备好，直到缓存满阻塞
	messages := make(chan string, 1)
	//
	go func() {
		for {
			select {
			case msg := <-messages:
				fmt.Println("received message", msg)
			default:
				//fmt.Println("no message received")
			}
		}
	}()

	msg := "hi"
	go func() {
		for {
			select {
			// 当接收方没有准备好，则直接执行default
			case messages <- msg:
				fmt.Println("send message", msg)
			default:
				//fmt.Println("no message send")
			}
		}
	}()

	time.Sleep(1 * time.Second)
}
