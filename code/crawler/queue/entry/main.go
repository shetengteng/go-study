package main

import (
	"learngo/queue"
	"fmt"
)


type Person struct {
	Name string
	Age int
}

func (p Person) Say() string{
	return p.Name
}

func (p *Person) GetAge() int{
	return p.Age
}

type P interface{
	Say() string
}

type M interface {
	GetAge() int
}

func main() {
	q := queue.Queue{1}
	q.Push(2).Push(3)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())


	var p P
	p = Person{Name:"ss"}
	fmt.Println(p.Say())
	fmt.Printf("%T %v \n",p,p)

	var m M
	m = &Person{Name:"tt",Age:22}
	fmt.Println(m.GetAge())
	fmt.Printf("%T %v \n",m,m)

	// 转换
	pp := p.(Person)
	fmt.Println(pp)
	mp := m.(*Person)
	fmt.Println(mp)
	
	
}
