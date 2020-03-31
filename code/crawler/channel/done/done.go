package main

import (
	"fmt"
)

type worker struct {
	in   chan int
	done chan bool
}

func doWork(id int, c chan int, done chan bool) {
	for n := range c {
		fmt.Printf("worker %d received %c \n", id, n)
		go func() { done <- true }() // 需要异步处理，由于要<-done多次，发送该数据后阻塞

	}
}

func createWorker(id int) worker {
	w := worker{
		in:   make(chan int),
		done: make(chan bool),
	}
	go doWork(id, w.in, w.done)
	return w
}

func chanDemo() {

	var workers [10]worker

	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i)
	}
	for i, worker := range workers {
		worker.in <- 'a' + i
		//<-workers[i].done
	}
	for i, worker := range workers {
		worker.in <- 'A' + i
		//<-workers[i].done
	}

	for _, worker := range workers {
		<-worker.done // 取出后才可以继续放入值，否则阻塞
		<-worker.done
	}

}

func main() {
	chanDemo()
}
