package main

import (
	"encoding/json"
	"fmt"
	"github.com/mitchellh/mapstructure"
)

type Person struct {
	Name string
	Age  int
	Job  string `mapstructure:"job,omitempty"` // 转换到map中的时候，如果该属性不存在，则忽略，使用该omitempty，前面需要一个名称表示该属性在map中的名称
}

func main() {

	p := &Person{
		Name: "ss",
		Age:  11,
		Job:  "x", // 如果不存在，则不会放入到map中
	}
	var m map[string]interface{}
	// 入参是引用
	mapstructure.Decode(p, &m)
	fmt.Println(m)

	// 转换为json
	data, _ := json.Marshal(&m)
	fmt.Println(string(data))

}
