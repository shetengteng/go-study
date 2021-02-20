package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 解析成64位精度的float
	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f)

	// 0 表示从字符串中推出基数，64 表示数据大小64bit
	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println(i)

	// 自动识别16进制的字符串
	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)

	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println(u)

	// ParseInt(s, 10, 0) 10进制数字字符串转为数字
	k, _ := strconv.Atoi("135")
	fmt.Println(k)

	// 非数字会报错
	_, e := strconv.Atoi("wat")
	fmt.Println(e)
}
