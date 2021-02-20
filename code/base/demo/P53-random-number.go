package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 每次执行的结果都是 81 87
	fmt.Println(rand.Intn(100),",",rand.Intn(100))
	// 获取 float 的随机数 0.0-1.0
	fmt.Println(rand.Float64())
	// 范围 5.0-10.0
	fmt.Println(rand.Float64()*5+5,",",rand.Float64()*5+5)

	// 设置不同的种子，产生不同的随机数
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	fmt.Println(r1.Intn(100),",",r1.Intn(100))

	// 一旦固定了种子，那么每次启动main执行的结果都一样
	s2 := rand.NewSource(42)
	r2 := rand.New(s2)
	fmt.Println(r2.Intn(100),",",r2.Intn(100))

	s3 := rand.NewSource(42)
	r3 := rand.New(s3)
	fmt.Println(r3.Intn(100),",",r3.Intn(100))
}
