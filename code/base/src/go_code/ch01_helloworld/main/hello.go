// 需求：打印 "hello world"
package main

import "fmt"

func main() {

	fmt.Println("hello go")

	fmt.Println(hell("%s %s", "xx", "yy"))

}

func hell(format string, value ...interface{}) string {
	// 将value数组变成一个个元素
	sprintf := fmt.Sprintf(format, value...)
	return sprintf
}
