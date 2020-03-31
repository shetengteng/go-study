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
	w := createWorker(0)
	// 非阻塞式的获取,谁先到先取谁
	for {
		select {
		case n := <-c1:
			w <- n
		case n := <-c2:
			w <- n
		}
	}
}
