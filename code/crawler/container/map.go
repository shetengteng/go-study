package main

import "fmt"

func main() {
	// key在map中是无序的
	m := map[string]string {
		"name":"sss",
		"course":"golang",
		"site":"sss",
	}
	fmt.Println(m)

	m2 := make(map[string]string) // m2 == empty map
	fmt.Println(m2,m2 == nil)

	var m3 map[string]int // m3 == nil
	fmt.Println(m3 == nil)

	// go 语言中的nil 是可以参与运算的，和empty map是可以混用的
	// 遍历
	fmt.Println("traversing map")

	for k,v := range m {
		fmt.Println(k,v)
	}
	for k := range m {
		fmt.Println(k)
	}
	for _,v := range m {
		fmt.Println(v)
	}

	fmt.Println("getting values")
	courseName := m["course"]
	fmt.Println(courseName)

	// 如果key不存在
	k,ok:= m["k"]
	fmt.Printf("k=%q ok=%v \n",k,ok)

	fmt.Println("deleting values")
	delete(m,"name")
	name, ok := m["name"]
	fmt.Println(name,ok)
}
