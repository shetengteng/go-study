package main

import (
	"fmt"
	"math"
)

const s1 string = "constant"

func main() {
	fmt.Println(s1)

	const n = 50000000
	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
}
