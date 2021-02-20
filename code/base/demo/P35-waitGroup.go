package main

import (
	"fmt"
	"sync"
	"time"
)

func workerWg(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {

	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		// waitGroup
		wg.Add(1)
		// 传参需要引用对象地址
		go workerWg(i, &wg)
	}

	wg.Wait()
}
