package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// 环境变量key-value都是string类型
	os.Setenv("FOO", "1")
	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR"))

	fmt.Println()
	// 遍历所有的环境变量
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		//fmt.Println(pair[0]," ",pair[1])
		fmt.Println(pair[0])
	}
}
