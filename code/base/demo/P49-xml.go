package main

import (
	"encoding/xml"
	"fmt"
)

type Plant struct {
	XMLName xml.Name `xml:"plant"`
	Id      int      `xml:"id,atrr"`
	Name    string   `xml:"name"`
	Origin  []string `xml:"origin"`
}

func (p Plant) String() string {
	return fmt.Sprintf("Plant id=%v name=%v origin=%v", p.Id, p.Name, p.Origin)
}

func main() {

	coffee := &Plant{Id: 12, Name: "coffee"}
	coffee.Origin = []string{"Ethiopia","Brazil"}
	// 转换为xml
	out,_:= xml.MarshalIndent(coffee," "," ")
	fmt.Println(string(out))
	// 添加xml头部
	fmt.Println(xml.Header + string(out))


	// 转换为 对象
	var p Plant
	if err := xml.Unmarshal(out,&p);err != nil {
		panic(err)
	}
	fmt.Println(p)

	tomato := &Plant{Id: 11,Name: "tomato"}
	tomato.Origin = []string{"Mexico","California"}

	// 内嵌xml
	type Nesting struct {
		XMLName xml.Name `xml:"nesting"`
		Plants []*Plant `xml:"parent>child>plant"`
	}

	nesting := &Nesting{}
	nesting.Plants = []*Plant{coffee,tomato}

	out,_ = xml.MarshalIndent(nesting," "," ")
	fmt.Println(string(out))
}
