package main

import (
	"fmt"
	"os"
)

func main() {

	// 这句答应不会执行
	defer fmt.Println("xx")
	// 关闭程序，返回一个关闭码
	os.Exit(3)
}
