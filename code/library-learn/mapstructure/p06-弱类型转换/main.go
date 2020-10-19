package main

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
)

type Person struct {
	Name   string
	Age    int
	Emails []string
}

func main() {
	m := map[string]interface{}{
		"name":   123,
		"age":    "18",
		"emails": []int{1, 2, 3},
	}
	var p Person
	mapstructure.WeakDecode(&m, &p)
	fmt.Println(p)

}
