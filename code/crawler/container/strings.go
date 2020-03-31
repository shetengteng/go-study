package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "go语言" // utf-8 可变长编码
	fmt.Println(len(s)) // s的byte的大小，英文1个byte，中文3个byte
	for _,b := range []byte(s) {
		fmt.Printf("%X ",b)
	}
	fmt.Println()
	for i,ch := range s {
		fmt.Printf("(%d %X)",i,ch)
	}
	fmt.Println()

	fmt.Println("rune count:",utf8.RuneCountInString(s))

	for i,ch := range []rune(s){ // 转换后rune占用的字节是4个字节
		fmt.Printf("(%d %c)",i,ch)
	}
	fmt.Println()
}
