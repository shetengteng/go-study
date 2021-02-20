package main

import (
	"fmt"
	"strings"
)

func Index(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

func Include(vs []string, t string) bool {
	return Index(vs, t) >= 0
}

// 满足f的判断函数，则返回true
// 只要有一个满足就true
func Any(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

// 有一个不满足，就false
func All(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

func Filter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

// 映射，从一个[]string映射到另一个[]string
func Map(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

func main() {
	var strs = []string{"peach","apple","pear","plum"}
	fmt.Println(Index(strs,"pear"))
	fmt.Println(Include(strs,"grape"))
	fmt.Println(Any(strs, func(v string) bool {
		return strings.HasPrefix(v,"p")
	}))
	fmt.Println(All(strs, func(v string) bool {
		return strings.HasPrefix(v,"p")
	}))
	fmt.Println(Filter(strs, func(v string) bool {
		return strings.Contains(v,"e")
	}))
	fmt.Println(Map(strs,strings.ToUpper))
}
