package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Println("working")
	time.Sleep(time.Second)
	fmt.Println("done")
	done <- true
}

func main() {

	done := make(chan bool, 1)
	go worker(done)

	// 如果没有这句，那么main线程会直接停止
	<-done

}
