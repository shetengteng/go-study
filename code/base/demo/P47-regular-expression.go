package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {

	// 直接正则检验
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match) // true

	// 声明一个正则表达式对象
	r, _ := regexp.Compile("p([a-z]+)ch")
	// 使用该对象进行正则处理
	fmt.Println(r.MatchString("peach")) // true
	// 找到第一个子串
	fmt.Println(r.FindString("peach punch")) // peach
	// 找到第一个子串 并返回对应的位置
	fmt.Println(r.FindStringIndex("peach punch")) // [0 5]
	// 子匹配变量包括关于整个模式匹配和这些匹配中的子匹配的信息。例如，这将返回p([a-z]+)ch和([a-z]+)的信息。
	fmt.Println(r.FindStringSubmatch("peach punch"))
	fmt.Println(r.FindStringSubmatchIndex("peach punch"))
	//
	fmt.Println(r.FindAllStringSubmatchIndex("peach punch pinch", -1))
	fmt.Println(r.FindAllString("peach punch pinch", 2))

	fmt.Println(r.Match([]byte("peach")))

	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println(r)

	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))
	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))
}
