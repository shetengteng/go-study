package main

import (
	"encoding/json"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"log"
)

type Person struct {
	Name   string
	Age    int
	Job    string
	Salary int `mapstructure:"gongzi"` // 使用别名，从map中查找对应的属性
}

type Cat struct {
	Name  string
	Age   int
	Breed string
	Size  int
}

// mapstructure 忽略大小写 没有匹配上的用默认值代替
func main() {

	// 准备数据
	datas := []string{
		`{"type":"person","name":"dd","age":18,"job":"pp","gongzi":100000}`,
		`{"type":"cat","name":"ddx","age":182,"breed":"xx"}`,
	}
	for _, data := range datas {
		// 将每个字符转换为map[string]interface{}
		var m map[string]interface{}
		err := json.Unmarshal([]byte(data), &m)
		if err != nil {
			log.Fatal(err)
		}

		switch m["type"].(string) {
		// 判断是person类型，然后转换为person对象
		case "person":
			var p = new(Person)
			err := mapstructure.Decode(m, p)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("person", p)
		case "cat":
			var c = new(Cat)
			err := mapstructure.Decode(m, c)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("cat", c)
		}
	}
}
