package test

import "math"

func Triangle(a,b int) int {
	return int(math.Sqrt(float64(a*a+b*b)))
}