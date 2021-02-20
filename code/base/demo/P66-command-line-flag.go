package main

import (
	"flag"
	"fmt"
)

func main() {

	// 返回的是一个字符串指针，foo是默认值 usage 是使用描述
	wordPtr := flag.String("word", "foo", "a string")
	numbPtr := flag.Int("numb", 22, "an int")
	boolPtr := flag.Bool("fork", false, "a bool")

	// 可以将flag的值传递给一个变量，前提通过该变量的指针传递
	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	// 调用parse方法解析flag中的值
	flag.Parse()

	// 输出
	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:",*numbPtr)
	fmt.Println("fork:",*boolPtr)
	fmt.Println("svar:",svar)
	// 解析其他没有定义的flag信息
	fmt.Println("tail:",flag.Args())
}
