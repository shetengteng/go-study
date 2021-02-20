package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

func main() {

	p := point{1, 2}
	fmt.Printf("%v\n", p)  // 打印出point实例的值
	fmt.Printf("%+v\n", p) // 打印出 point 实例的值，并包含 属性的名称
	fmt.Printf("%#v\n", p) // 打印出 point 的包名等描述信息，同时含有值

	fmt.Printf("%T\n", p)  // 打印出point的类型
	fmt.Printf("%p\n", &p) // 打印出point的地址值

	fmt.Printf("%t\n", true) // 格式化boolean类型

	fmt.Printf("%d\n", 124) // 转换为10进制输出
	fmt.Printf("%b\n", 14)  // 转换为2进制输出
	fmt.Printf("%x\n", 456) // 转换为对应的16进制

	fmt.Printf("%c\n", 33)   // 转换为对应的字符值
	fmt.Printf("%f\n", 78.9) // 对float类型，使用%f进行输出

	fmt.Printf("%e\n", 12340000.0) // 科学计数法 e
	fmt.Printf("%E\n", 12340000.0) // 科学计数法 E

	fmt.Printf("%s\n", "\"string\"") // 基本的字符串打印，会将 \ 进行转义
	fmt.Printf("%q\n", "\"string\"") // 原样输出字符串，不会将 \ 进行转义，会打印出 \

	fmt.Printf("%x\n", "hex this") // 对每个字符进行16进制输出

	fmt.Printf("|%6d|%6d|\n", 12, 345)       // 格式化输出数字，默认右对齐，不足部分用空格填充
	fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45) // 格式化输出浮点类型数字，6表示总数字数目，.2表示保留小数个数
	fmt.Printf("|%6s|%6s|\n", "foo", "b")   // 格式化输出 字符
	fmt.Printf("|%-6s|%-6s|\n", "foo", "b") // 格式化输出，左对齐，使用 -

	s := fmt.Sprintf("a %s", "string") // 返回一个格式化的字符串，而非打印出来
	fmt.Println(s)

	fmt.Fprintf(os.Stderr, "an %s\n", "error") // 输出到流中
}
