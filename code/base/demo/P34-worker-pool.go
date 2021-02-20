package main

import (
	"fmt"
	"time"
)

func workerOp(id int, jobs <-chan int, results chan<- int) {
	// 如果没有jobs,则会阻塞，直到close(jobs)
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {

	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// 创建worker
	for w := 1; w <= 3; w++ {
		go workerOp(w, jobs, results)
	}
	// 创建任务
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= numJobs; a++ {
		<-results
	}

}
