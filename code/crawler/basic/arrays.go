package main

import "fmt"

func main() {
	var arr [5]int // [0 0 0 0 0]
	arr2 := [3]int{1,2,3}
	arr3 := [...]int{1,2,3}
	fmt.Println(arr,arr2,arr3)

	var grid [4][5]int // [[0 0 0 0 0] [0 0 0 0 0] [0 0 0 0 0] [0 0 0 0 0]]
	fmt.Println(grid)

	for i:= 0;i<len(arr3);i++{
		fmt.Println(arr3[i])
	}

	for i := range arr3 {
		fmt.Println(arr3[i])
	}

	for i,v := range arr3{
		fmt.Println(i,v)
	}

	for _,v := range arr3{
		fmt.Println(v)
	}
}


