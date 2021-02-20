package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// 从stdin读取，输出到stdout上
	// 使用缓冲流包裹非缓冲流os.Stdin
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		ucl := strings.ToUpper(scanner.Text())
		fmt.Println(ucl)
	}

	// 最后检查是否有错误
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}

}
