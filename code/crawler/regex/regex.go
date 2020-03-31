package main

import (
	"fmt"
	"regexp"
)

func main() {

	// 测试邮件
	const text = `my email is ccs@gmail.com xx@abc.com
	ffdsf sss@qq.com
	add@abc.com.cn
	`

	// 写法1
	//compile := regexp.MustCompile(".+@.+\\..+")

	// 写法2
	compile := regexp.MustCompile(`[a-zA-z0-9]+@[a-zA-z0-9.]+\.[a-zA-z0-9]+`)
	//match := compile.FindString(text) // 得到第一匹配的
	match := compile.FindAllString(text, -1) // 找到所有的 -1 表示所有
	fmt.Println(match)

	// 写法3，提取

	compile = regexp.MustCompile(`([a-zA-z0-9]+)@([a-zA-z0-9]+)(\.[a-zA-z0-9.]+)`)
	submatch := compile.FindAllStringSubmatch(text, -1)
	for _,m := range submatch {
		fmt.Println(m)
	}


}
