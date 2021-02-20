package main

import (
	"errors"
	"fmt"
)

func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("can't work with 42")
	}
	return arg + 3, nil
}

type argError struct {
	arg  int
	prob string
}

// implement builtin error interface
func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

// error 是一个接口
// 这里argError是一个实现了error接口的struct
// 直接返回argError表示的是argError这个对象，需要返回&argError的引用，才能用error接口对象接收
func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "cannot work with it"}
	}
	return arg + 3, nil
}

func main() {

	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed ", e)
		} else {
			fmt.Println("f1 worked ", r)
		}
	}

	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed ", e)
		} else {
			fmt.Println("f2 worked ", r)
		}
	}

	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}

}
