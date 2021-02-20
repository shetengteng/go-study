package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	// 声明一个通道用于接收信号
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	// 注册通道用于接收消息，后面的2个参数是接收系统信号的类型
	// SIGTERM 接收关闭信号
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		// 接收到信号后关闭
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")
}
