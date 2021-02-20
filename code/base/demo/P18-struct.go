package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 12
	return &p
}

func main() {
	fmt.Println(person{"bob", 11})
	fmt.Println(person{name: "alice", age: 23})
	fmt.Println(person{name: "fred"})
	fmt.Println(&person{name: "ann", age: 40})
	fmt.Println(newPerson("jon"))

	s := person{name: "ss", age: 99}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 44
	fmt.Println(sp.age)

}
