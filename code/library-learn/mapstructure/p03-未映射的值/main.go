package main

import (
	"encoding/json"
	"fmt"
	"github.com/mitchellh/mapstructure"
)

type Person struct {
	Name  string
	Age   int
	Job   string
	Other map[string]interface{} `mapstructure:",remain"` // 保留没有匹配的元素
}

func main() {
	data := `{"name":"ss","age":1,"job":"y","height":11,"handsome":true}`
	var m map[string]interface{}
	json.Unmarshal([]byte(data), &m)
	var p = new(Person)
	mapstructure.Decode(m, p)
	fmt.Println(p)
}
