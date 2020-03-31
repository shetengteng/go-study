package main

import "fmt"

func printSlice(s []int){
	fmt.Printf("%v, len = %d cap = %d\n",s, len(s),cap(s))
}

func main() {
	var s []int // zero value for slice is nil

	for i := 0;i<100;i++{
		//printSlice(s)
		s = append(s,2*i+1)
	}
	fmt.Println(s)

	// 关于创建
	s1 := []int{2,3,4,5}
	printSlice(s1)
	s2 := make([]int,16)
	printSlice(s2)
	s3 := make([]int,10,32)
	printSlice(s3)

	fmt.Println("copying slice")
	copy(s2,s1) // 将s1拷贝到s2
	printSlice(s2)

	fmt.Println("deleting elements from slice")
	// 删除下标为3的
	s2 = append(s2[:3],s2[4:]...)
	printSlice(s2)

	// 删除头尾
	fmt.Println("popping from front")
	front := s2[0]
	s2 = s2[1:]
	fmt.Println(front)
	printSlice(s2)

	fmt.Println("Popping from back")
	tail := s2[len(s2)-1]
	s2 = s2[:len(s2)-1]
	fmt.Println(tail)
	printSlice(s2)


}
