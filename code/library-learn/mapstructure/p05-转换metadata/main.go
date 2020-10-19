package main

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
)

type Person struct {
	Name string
	Age  int
}

func main() {

	data := map[string]interface{}{
		"name": "ss",
		"age":  11,
		"job":  "yy",
	}
	// 声明一个metadata
	var metadata mapstructure.Metadata
	var p Person
	mapstructure.DecodeMetadata(&data, &p, &metadata)

	fmt.Println("keys", metadata.Keys)
	fmt.Println("unused", metadata.Unused)
}
