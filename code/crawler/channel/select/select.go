package main

import (
	"fmt"
	"time"
	"math/rand"
)

func generator() chan int{
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500))*time.Millisecond)
			out <- i
			i ++
		}
	}()
	return out
}

func main() {

	//var c1, c2 chan int // c1 and c2 = nil

	var c1,c2 = generator(),generator()

	// 非阻塞式的获取,谁先到先取谁
	for {
		select {
		case n := <-c1:
			fmt.Println("received from c1:", n)
		case n := <-c2:
			fmt.Println("received from c2:", n)
		//default:
		//	fmt.Println("no received")
		}
	}
}
