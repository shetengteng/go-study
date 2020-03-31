package main

import (
	"fmt"
	"time"
)

func chanDemo() {
	//var c chan int // c == nil
	c := make(chan int)

	// 接收数据
	go func() {
		for {
			n := <-c
			fmt.Println(n)
		}
	}()

	// 发送数据
	go func() {
		for {
			c <- 1
			time.Sleep(time.Millisecond)
		}
	}()

	c <- 1
	c <- 2
}

// 接收一个channel
func work(id int, c chan int) {
	for {
		n, ok := <-c
		if !ok { // 如果收不到消息则退出
			break
		}
		fmt.Printf("worker %d received %d \n", id, n)
	}
	// 方式2
	//for n:= range c {
	//	fmt.Printf("worker %d received %d \n", id, n)
	//}
}

func chanDemo2() {
	var channels [10]chan int
	for i := 0; i < 10; i++ {
		channels[i] = make(chan int)
		go work(i, channels[i])
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}
	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}
	time.Sleep(time.Millisecond)
}

// 返回一个channel
func createWorker(id int) chan int {
	c := make(chan int)
	go func() {
		for {
			fmt.Printf("worker %d received %c \n", id, <-c)
		}
	}()
	return c
}

func chanDemo3() {
	var channels [10]chan int
	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i)
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}
	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}
	time.Sleep(time.Millisecond)
}

// 该chan只能收数据，<-chan表示只能发送数据
func createWorker2(id int) chan<- int {
	c := make(chan int)
	go func() {
		for {
			fmt.Printf("worker %d received %c \n", id, <-c)
		}
	}()
	return c
}

func bufferedChannel() {
	c := make(chan int, 3)
	c <- 1
	c <- 2
	c <- 3
	c <- 4 // 此时产生deadlock
}

func channlClose() {
	c := make(chan int, 3)
	go work(0, c)
	c <- 1
	c <- 2
	c <- 3
	close(c) // 关闭后会一直接收到空串，或者0
	time.Sleep(time.Millisecond)

}

func main() {
	//chanDemo3()
	channlClose()
	time.Sleep(time.Second * 1)
}
