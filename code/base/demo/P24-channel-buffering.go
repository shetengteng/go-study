package main

import (
	"fmt"
	"time"
)

func main() {
	messages := make(chan string, 2)
	messages <- "buffered"
	messages <- "channel"

	go func() {
		fmt.Println(<-messages)
		fmt.Println(<-messages)
		fmt.Println(<-messages)
		fmt.Println("---")
	}()

	time.Sleep(time.Second)
	fmt.Println("done")
}
