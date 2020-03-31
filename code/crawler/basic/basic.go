package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

// 变量初始值
func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d %q\n",a,s)
}

// 变量赋值
func variableInitialValue() {
	var a,b int = 3,4
	var s string = "abc"
	fmt.Println(a,b,s)
}

// 类型推断
func variableTypeDeduction(){
	var a,b,c,s = 2,3,true,"df"
	fmt.Println(a,b,c,s)
}

// 第一次用变量可以使用 := 推荐
func variableShorter(){
	a,b,c,s := 2,3,true,"df"
	fmt.Println(a,b,c,s)
}

// 外部变量赋值，包内部变量
var ss = "kk"
//aa:=11 error 不允许

// 使用() 省略var的书写
var (
	aa = 3
	bb = true
)

func consts() {
	const filename string = "xx.txt"
	const a,b = 3,4 // 不指定int 那么默认是float类型
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename,c)

	// 写法2
	const (
		e = 1
		f = 2
	)
}


func enums(){
	const(
		cpp = 0
		java = 1
		python = 2
		golang = 3
	)
	fmt.Println(cpp,java,python,golang)
}

// 使用iota做简化，表示自增
func enums2(){
	const(
		cpp2 = iota
		_ // 表示省略，此处有+1效果
		java2
		python2
		golang2
	)

	const(
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(cpp2,java2,python2,golang2)
	fmt.Println(b,kb,mb,gb,tb,pb)
}


func euler(){
	pow := cmplx.Pow(math.E, 1i*math.Pi)
	fmt.Println(pow)
	pow = cmplx.Exp(1i*math.Pi)+1
	fmt.Println(pow)
}

func triangle(){
	var a,b int = 3,4
	var c int
	c = int(math.Sqrt(float64(a*a+b*b)))
	fmt.Println(c)
}

func main() {
	fmt.Println("hello world")
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
	consts()
	enums()
	enums2()
	euler()
	triangle()
}