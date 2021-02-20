package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)
	// 获取时间戳
	secs := now.Unix()
	fmt.Println(secs)

	nanos := now.UnixNano()
	fmt.Println(nanos)

	// 没有ms的时间戳，需要通过ns进行转换
	millis := nanos / 1000000
	fmt.Println(millis)

	// 将时间戳转换为时间
	// 如果没有ns，则可以是0
	fmt.Println(time.Unix(secs, 0))
	// 如果只有s，那么s 是0
	fmt.Println(time.Unix(0, nanos))
}
