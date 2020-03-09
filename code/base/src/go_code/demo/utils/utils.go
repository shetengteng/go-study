package utils

import (
	"fmt"
)

func Cal(n1 float64,n2 float64,operator byte) float64 {
	var re float64
	switch operator {
	case '+':
		re = n1 + n2
	case '-':
		re = n1 - n2
	case '*':
		re = n1 * n2
	case '/':
		re = n1 / n2
	default:
		fmt.Println("操作符号错误")
	}
	return re
}