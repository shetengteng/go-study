package main

import "fmt"

// 结构体
type Student struct {
	name string
	age  int
}

func (s *Student) hello(person string) string {
	return fmt.Sprintf("hello %s i am %s", person, s.name)
}

// 接口
type Person interface {
	getName() string
}

//Student 实现了 Person接口
func (s *Student) getName() string {
	return s.name
}

type Worker struct {
	name string
	age  int
}

func (w *Worker) getName() string {
	return w.name
}

func main() {

	stu := &Student{
		name: "tom",
	}
	msg := stu.hello("xx")
	fmt.Println(msg)

	// 强制类型转换为Person
	var p Person = &Student{
		name: "ss",
		age:  11,
	}
	fmt.Println(p.getName())

	// 如果没有此处的声明，那么删除func (w *Worker) getName() 则不会报错
	// 这里声明的意义在于确保编辑器知道Worker需要实现Person接口，没有实现的话则在编译器报错
	// 将nil转化为*Worker类型，再强转为Person类型，如果转换失败则报错
	var _ Person = (*Worker)(nil)

}
