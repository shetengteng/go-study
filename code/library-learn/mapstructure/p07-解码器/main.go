package main

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"log"
)

type Person struct {
	Name string
	Age  int
}

// 实现了weakDecode的功能，源码中实现类似，都是先实现一个DecoderConfig，然后通过返回的decoder进行处理
func main() {
	m := map[string]interface{}{
		"name": 123,
		"age":  "18",
		"job":  "programmer",
	}
	var p Person
	var metadata mapstructure.Metadata
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		WeaklyTypedInput: true,
		Result:           &p,
		Metadata:         &metadata,
	})
	if err != nil {
		log.Fatal(err)
	}
	err = decoder.Decode(&m)
	if err == nil {
		fmt.Println(p)
		fmt.Println(metadata.Keys)
		fmt.Println(metadata.Unused)
	} else {
		fmt.Println(err.Error())
	}
}
