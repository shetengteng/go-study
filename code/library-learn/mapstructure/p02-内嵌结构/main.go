package main

import (
	"encoding/json"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"log"
)

// 内嵌结构，2种方式构成
type Person struct {
	Name string
}

type Friend struct {
	Person
}

// 将Person内的属性扁平放置在Friend中
type Friend2 struct {
	Person `mapstructure:",squash"`
}

func main() {

	var datas = []string{
		`{"type":"friend","person":{"name":"ss"}}`,
		`{"type":"friend2","name":"ss2"}`,
	}

	// 将datas转换为map[string]interface{}对象
	for _, data := range datas {
		var m map[string]interface{}
		err := json.Unmarshal([]byte(data), &m)
		if err != nil {
			log.Fatal(err)
		}

		switch m["type"].(string) {
		case "friend":
			var f = new(Friend)
			err := mapstructure.Decode(m, f)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("f1", f)
			fmt.Println("f1", f.Person.Name)
		case "friend2":
			var f = new(Friend2)
			err := mapstructure.Decode(m, f)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("f2", f)
		}
	}
}
