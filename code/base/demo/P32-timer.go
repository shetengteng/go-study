package main

import (
	"fmt"
	"time"
)

func main() {
	timer1 := time.NewTimer(2 * time.Second)
	fmt.Println("start ", time.Now())
	time.Sleep(7 * time.Second)
	c := <-timer1.C
	fmt.Println("timer1 fired ", c)

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("timer2 fired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("timer2 stopped")
	}
	time.Sleep(3 * time.Second)
}
