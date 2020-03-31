# 工厂模式

- Golang的结构体没有构造函数，通常可以使用工厂模式来解决这个问题

```go
package model

type student struct {
    Name string
    Score float64
}

func NewStudent(n string,s float64) *student {
    return &student{
        Name: n,
        Score: s,
    }
}
```

- main.go

```go
package main
import{
    "fmt"
    "demo/model"
}

func main(){
    var stu = model.NewStudent("ss",22)
    fmt.Println(*stu)
    fmt.Println("name=",stu.Name," score=",stu.Score)
}
```

