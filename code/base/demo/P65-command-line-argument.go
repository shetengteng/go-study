package main

import (
	"fmt"
	"os"
)

func main() {
	// 读取全部的参数
	argsWithProg := os.Args
	// 去除第一个参数，第一个参数是main方法的go文件名
	argsWithoutProg := os.Args[1:]

	arg:=os.Args[3]

	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)
}
