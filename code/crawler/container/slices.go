package main

import "fmt"

func updateSlice(s []int){
	s[0] = 100
}

func printArr(a []int){
	fmt.Println(a)
}

func main() {
	arr:=[...]int{1,2,3,4,5,6,7,8,9,10}
	fmt.Println(arr[2:6])
	fmt.Println(arr[:6])
	fmt.Println(arr[2:])
	fmt.Println(arr[:])

	fmt.Println("------------")
	s := arr[:5]
	updateSlice(s)
	fmt.Println(s)
	fmt.Println(arr)
	printArr(arr[:])

	fmt.Println("reslice")
	fmt.Println(s)
	s = s[:3]
	printArr(s)
	fmt.Printf("s=%v len(s)=%d cap(s)=%d \n",s,len(s),cap(s))

	arr2:=[...]int{1,2,3,4,5,6,7,8,9,10}
	s2 := arr2[2:3]
	s3 := append(s2,10)
	s4 := append(s3,11)
	fmt.Println("s2,s3,s4",s2,s3,s4)
	fmt.Println("arr",arr2)
}

