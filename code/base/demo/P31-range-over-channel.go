package main

import (
	"fmt"
	"time"
)

func main() {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	//close(queue)

	go func() {
		for elem := range queue {
			fmt.Println(elem)
		}
		fmt.Println("--finish--")
	}()

	time.Sleep(time.Second * 10)
}
