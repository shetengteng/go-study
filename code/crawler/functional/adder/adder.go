package main

import (
	"fmt"
	"strings"
	"io"
	"bufio"
)

// 写法1：闭包
func adder() func(int) int{
	sum := 0
	return func (i int) int {
		sum += i
		return sum
	}
}

// 写法2：正统函数式编程写法
type iAdder func(int) (int,iAdder)

// 没有变量，只有函数和常数
func adder2(base int) iAdder {
	return func(v int) (int,iAdder){
		return base +v,adder2(base + v)
	}
}

func fibonacci() intGen {
	a,b := 0,1
	return func() int{
		a,b = b,a+b
		return a
	}
}

type intGen func() int
// 函数也可以实现接口
func (g intGen) Read(p []byte)(n int,err error){
	next := g()
	if next > 100 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n",next)
	return strings.NewReader(s).Read(p)

}

func printFileContents(reader io.Reader){
	scanner := bufio.NewScanner(reader)
	for scanner.Scan(){
		fmt.Println(scanner.Text())
	}
}

func main() {
	a := adder()
	a2 := adder2(0)
	for i:=0;i<10 ; i++ {
		fmt.Printf("0+1+...%d = %d \n",i,a(i))
		var r int
		r,a2 = a2(i)
		fmt.Printf("0+1+...%d = %d \n",i,r)
	}

	f := fibonacci()
	fmt.Println(f())

	printFileContents(f)

}
