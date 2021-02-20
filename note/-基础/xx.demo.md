https://gobyexample.com/

# hello world

- 编写hello-world.go

```go
package main

import "fmt"

func main() {
    fmt.Println("hello world")
}
```

- 在go目录下执行go run执行该文件

```go
D:\Personal Files\note\go-study\code\base\demo>go run P01-hello-world.go
hello world
```

- 使用go build 可以编译成二进制文件

```go
go build P01-hello-world.go
```



# value

- 基本数据类型

- - string

- - - 可以使用+进行连接

- - integer
  - float
  - boolean

```go
package main

import "fmt"

func main() {
    fmt.Println("go" + "lang")
    fmt.Println("1+1=", 1+1)
    fmt.Println("7.0/3.0=", 7.0/3.0)
    fmt.Println(true && false)
    fmt.Println(true || false)
    fmt.Println(!true)
}

// 结果
golang
1+1= 2
7.0/3.0= 2.3333333333333335
false
true
false
```



# variable

- 变量的声明

- - 使用关键字var ，go编译器通过类型推断出变量的具体类型
  - 可以同时什么多个变量
  - 使用:= 可以完成声明与赋值操作，省去var 关键字
  - 没有赋值的赋上初始值

```go
package main

import "fmt"

func main() {
    var a = "initial"
    fmt.Println(a)

    var b, c int = 1, 2
    fmt.Println(b, c)

    var d = true
    fmt.Println(d)

    var e int
    fmt.Println(e)

    f := "apple"
    fmt.Println(f)
}

// result
initial
1 2
true
0
apple
```



# constant

- 常量

- - 可以出现任何var 出现的位置
  - 常量表达式以任意的精度执行算术
  - 数字常量只有在指定类型---如通过显式转换时 才具有类型
  - 通过在需要类型的上下文中使用数字，如在变量赋值或者函数调用，可以给数字指定类型，如math.Sin期望是一个float64类型
  - 小结：在上下文中使用时指定类型

```go
package main

import (
    "fmt"
    "math"
)

const s string = "constant"

func main() {
    fmt.Println(s)

    const n = 50000000
    const d = 3e20 / n
    fmt.Println(d)

    fmt.Println(int64(d))

    fmt.Println(math.Sin(n))
}

// result
constant
6e+12
6000000000000
0.8256467432733234
```



# for

- 循环

- - 使用break跳出循环
  - 使用return跳出方法
  - 使用continue执行循环的下一个迭代



```go
package main

import "fmt"

func main() {

    i := 1
    for i <= 3 {
        fmt.Println(i)
        i = i + 1
    }

    for j := 7; j <= 9; j++ {
        fmt.Println(j)
    }

    for {
        fmt.Println("loop")
        break
    }

    for n := 0; n <= 5; n++ {
        if n%2 == 0 {
            continue
        }
        fmt.Println(n)
    }
}
// result
1
2
3
7
8
9
loop
1
3
5
```



# if/else

- 没有三元运算符

```go
package main

import "fmt"

func main() {
    if 7%2 == 0 {
        fmt.Println("7 is even")
    } else {
        fmt.Println("7 is odd")
    }
    if 8%4 == 0 {
        fmt.Println("8 is divisible by 4")
    }

    if num := 9; num < 0 {
        fmt.Println(num, "is negative")
    } else if num < 10 {
        fmt.Println(num, "has 1 digit")
    } else {
        fmt.Println(num, "has multiple digits")
    }
}

// result
7 is odd
8 is divisible by 4
9 has 1 digit
```



# switch

- 在同一个case中，可以使用逗号分隔多个表达式
- 可以使用default默认分支
- 不带表达式的switch是if else的另一种写法
- switch可以对接口进行类型的匹配

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    i := 2
    fmt.Println("write ", i, " as ")
    switch i {
    case 1:
        fmt.Println("one")
    case 2:
        fmt.Println("two")
    case 3:
        fmt.Println("three")
    }

    switch time.Now().Weekday() {
    case time.Saturday, time.Sunday:
        fmt.Println("It is the weekend")
    default:
        fmt.Println("It is a weekday")
    }

    t := time.Now()
    switch {
    case t.Hour() < 12:
        fmt.Println("It is before noon")
    default:
        fmt.Println("It is after noon")
    }

    whatAmI := func(i interface{}) {
        switch t := i.(type) {
        case bool:
            fmt.Println("I am a bool")
        case int:
            fmt.Println("I am an int")
        default:
            fmt.Printf("Don't know type %T\n", t)
        }
    }

    whatAmI(true)
    whatAmI(1)
    whatAmI("hi")
}

// result
write  2  as 
two
It is a weekday
It is before noon
I am a bool
I am an int
Don't know type string
```



# array

- 指定个数
- 没有初始化则含有默认值
- 没有slice使用的广泛
- fmt.Println打印array则会全部将元素全部打印出

```go
package main

import "fmt"

func main() {

    var a [5]int
    fmt.Println("emp:", a)

    a[4] = 100
    fmt.Println("set:", a)
    fmt.Println("get:", a[4])
    fmt.Println("len:", len(a))

    b := [5]int{1, 2, 3, 4, 5}
    fmt.Println("dcl:", b)

    var twoD [2][3]int
    for i := 0; i < 2; i++ {
        for j := 0; j < 3; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d", twoD)
}

// result
emp: [0 0 0 0 0]
set: [0 0 0 0 100]
get: 100
len: 5
dcl: [1 2 3 4 5]
2d [[0 1 2] [1 2 3]]
```



# slice

- 切片是go的关键类型
- 切片操作类似与数组
- 使用make关键字进行初始化
- 切片可以通过append关键字进行添加操作，返回一个新的切片
- 使用copy关键字可以复制切片
- 使用slice[low:hight]进行切片的截取，返回一个新的切片

- - 范围是[low,hight)

- make append copy 都属于builtin包下的方法

- - builtin包还定义了基本的数据类型

```go
package main

import "fmt"

func main() {

    s := make([]string, 3)
    fmt.Println("emp:", s) // empty

    s[0] = "a"
    s[1] = "b"
    s[2] = "c"

    fmt.Println("set:", s)
    fmt.Println("get:", s[2])
    fmt.Println("len:", len(s))

    s = append(s, "d")
    s = append(s, "e", "f")
    fmt.Println("apd:", s)

    c := make([]string, len(s))
    copy(c, s)
    fmt.Println("cpy:", c)

    l := s[2:5]
    fmt.Println("sl1:", l)

    l = s[:5]
    fmt.Println("sl2:", l)

    l = s[2:]
    fmt.Println("sl3:", l)

    t := []string{"g", "h", "l"}
    fmt.Println("dcl:", t)

    twoD := make([][]int, 3)
    for i := 0; i < 3; i++ {
        innerLen := i + 1
        twoD[i] = make([]int, innerLen)
        for j := 0; j < innerLen; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d:", twoD)
}

// result
emp: [  ]
set: [a b c]
get: c
len: 3
apd: [a b c d e f]
cpy: [a b c d e f]
sl1: [c d e]
sl2: [a b c d e]
sl3: [c d e f]
dcl: [g h l]
2d: [[0] [1 2] [2 3 4]]
```



# map

- 使用builtin的make关键字
- 使用name[key]获取元素

- - 返回一个是值
  - 返回2个一个是值，一个是 是否存在的bool，第一个元素可以用 _ 替换

- 使用len获取键值对个数
- 使用delete删除元素（map专用）

```go
package main

import "fmt"

func main() {
    m := make(map[string]int)
    m["k1"] = 7
    m["k2"] = 13
    fmt.Println("map:", m)

    v1 := m["k1"]
    fmt.Println("v1:", v1)
    fmt.Println("len:", len(m))

    delete(m, "k2")
    fmt.Println("map:", m)

    _, isExists := m["k2"]
    fmt.Println("prs:", isExists)

    n := map[string]int{"foo": 1, "bar": 2}
    fmt.Println("map:", n)
}
// result
map: map[k1:7 k2:13]
v1: 7
len: 2
map: map[k1:7]
prs: false
map: map[bar:2 foo:1]
```



# range

- 遍历各种数据结构中的元素
- range array or slice 返回 index 和 value
- range map 返回 key 和 value，也可以只用key
- range on strings遍历Unicode码位。第一个值是符文的起始字节索引，第二个值是符文本身

```go
package main

import "fmt"

func main() {

    nums := []int{2, 3, 4}
    sum := 0
    for _, num := range nums {
        sum += num
    }
    fmt.Println("sum:", sum)

    for i, num := range nums {
        if num == 3 {
            fmt.Println("index:", i)
        }
    }

    kvs := map[string]string{"a": "apple", "b": "banana"}
    for k, v := range kvs {
        fmt.Printf("%s -> %s \n", k, v)
    }
    for k := range kvs {
        fmt.Println("key:", k)
    }
    for i, c := range "go" {
        fmt.Println(i, c)
    }
}
// result
sum: 9
index: 1
a -> apple 
b -> banana 
key: a
key: b
0 103
1 111
```



# function

- 函数声明
- 如果参数类型一样，可以声明一次

```go
package main

import "fmt"

func plus(a int, b int) int {
    return a + b
}

func plusPlus(a, b, c int) int {
    return a + b + c
}

func main() {
    res := plus(1, 2)
    fmt.Println("1+2=", res)

    res = plusPlus(1, 2, 3)
    fmt.Println("1+2+3=", res)
}
```



# multiple return values

- 函数支持返回多个值
- 如果要忽略其中一个，使用 _ 

```go
package main

import "fmt"

func values() (int, int) {
    return 3, 7
}

func main() {
    a, b := values()
    fmt.Println(a)
    fmt.Println(b)

    _, c := values()
    fmt.Println(c)
}
```



# variadic fuction

- 可变入参函数
- fmt.Println就是一个可变入参的函数
- 如果使用一个可变入参的函数，同时传入的是一个切片对象slice_arg，那么使用slice_arg...

```go
package main

import "fmt"

func sum(nums ...int) {
    fmt.Print(nums, " ")
    total := 0
    for _, num := range nums {
        total += num
    }
    fmt.Println(total)
}

func main() {
    sum(1, 2)
    sum(1, 2, 3)

    nums := []int{1, 2, 3, 4}
    sum(nums...)
}
// result
[1 2] 3
[1 2 3] 6
[1 2 3 4] 10
```



# closure

- go 支持匿名函数，可以形成闭包

```go
package main

import "fmt"

func intSeq() func() int {
    i := 0
    return func() int {
        i++
        return i
    }
}

func main() {
    nextInt := intSeq()
    fmt.Println(nextInt())
    fmt.Println(nextInt())
    fmt.Println(nextInt())

    // 闭包的变量只在特定声明的方法中，对于创建的新的方法对象，原先的变量值不会继续生效
    nextInt2 := intSeq()
    fmt.Println(nextInt2())
}
// result
1
2
3
1
```



# recursion

- 递归

```go
package main

import "fmt"

func fact(n int) int {
    if n == 0 {
        return 1
    }
    return n * fact(n-1)
}

func main() {
    fmt.Println(fact(7))
}
// result
5040
```



# pointer

- 指针

- - 传递指针会修改对应的值
  - 传递对象不会修改，会拷贝一份

```go
package main

import "fmt"

func zeroval(ival int) {
    ival = 0
}

func zeroptr(iptr *int) {
    *iptr = 0
}

func main() {
    i := 1
    fmt.Println("initial:", i)

    zeroval(i)
    fmt.Println("zeroval:", i)

    zeroptr(&i)
    fmt.Println("zeroptr:", i)

    fmt.Println("pointer:", &i)
}
// result
initial: 1
zeroval: 1
zeroptr: 0
pointer: 0xc00000a0b8
```



# struct

- 结构体

- - 类似于java class
  - 初始化时可以指定相应的字段赋值
  - 省略的字段为默认值
  - 结构体是可变的
  - 可以使用结构体的指针对象进行字段的获取，会自动解引用

```go
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
// result
{bob 11}
{alice 23}
{fred 0}
&{ann 40}
&{jon 12}
ss
99
44
```



# method

- 使用方法，对结构体对象进行调用
- 如果使用指针方式，则可以修改对象内部字段，否则是值拷贝处理
- go会自动在指针对象和值对象之间进行调用的转换

```go
package main

import "fmt"

type rect struct {
    width, height int
}

func (r *rect) area() int {
    return r.width * r.height
}

func (r rect) perim() int {
    return 2*r.width + 2*r.height
}

func main() {
    r := rect{width: 10, height: 5}
    fmt.Println("area:", r.area())
    fmt.Println("perim:", r.perim())

    rp := &r
    fmt.Println("area", rp.area())
    fmt.Println("perim", rp.perim())
}
// result
area: 50
perim: 30
area 50
perim 30
```



# interface

- 接口，实现继承
- 弱继承，只要实现了对应的方法，就实现了继承
- https://jordanorelli.com/post/32665860244/how-to-use-interfaces-in-go#_=_

```go
package main

import (
    "fmt"
    "math"
)

type geometry interface {
    area() float64
    perim() float64
}

type rect2 struct {
    width, height float64
}

type circle struct {
    radius float64
}

func (r rect2) area() float64 {
    return r.width * r.height
}

func (r rect2) perim() float64 {
    return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
    return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
    return 2 * math.Pi * c.radius
}

func measure(g geometry) {
    fmt.Println(g)
    fmt.Println(g.area())
    fmt.Println(g.perim())
}

func main() {
    r := rect2{width: 3, height: 4}
    c := circle{radius: 5}

    measure(r)
    measure(c)
}
// result
{3 4}
12
14
{5}
78.53981633974483
31.41592653589793
```



# error

- https://blog.golang.org/error-handling-and-go
- go语言中，对错误的处理是显式的
- 返回nil表示没有错误
- 自定义错误类型，需要实现Error() string方法
- 使用内联错误检查，是go的常见语法
- 如果希望判断错误类型，需要通过类型断言的方式判断

```go
package main

import (
    "errors"
    "fmt"
)

func f1(arg int) (int, error) {
    if arg == 42 {
        return -1, errors.New("can't work with 42")
    }
    return arg + 3, nil
}

type argError struct {
    arg  int
    prob string
}

// implement builtin error interface
func (e *argError) Error() string {
    return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

// error 是一个接口
// 这里argError是一个实现了error接口的struct
// 直接返回argError表示的是argError这个对象，需要返回&argError的引用，才能用error接口对象接收
func f2(arg int) (int, error) {
    if arg == 42 {
        return -1, &argError{arg, "cannot work with it"}
    }
    return arg + 3, nil
}

func main() {

    for _, i := range []int{7, 42} {
        if r, e := f1(i); e != nil {
            fmt.Println("f1 failed ", e)
        } else {
            fmt.Println("f1 worked ", r)
        }
    }

    for _, i := range []int{7, 42} {
        if r, e := f2(i); e != nil {
            fmt.Println("f2 failed ", e)
        } else {
            fmt.Println("f2 worked ", r)
        }
    }

    _, e := f2(42)
    if ae, ok := e.(*argError); ok {
        fmt.Println(ae.arg)
        fmt.Println(ae.prob)
    }

}
// result
f1 worked  10
f1 failed  can't work with 42
f2 worked  10
f2 failed  42 - cannot work with it
42
cannot work with it
```



# goroutine

- 轻量级线程执行器
- 使用go f(s) 开启一个协程，也可也开启一个匿名函数执行

```go
package main

import (
    "fmt"
    "time"
)

func f(from string) {
    for i := 0; i < 4; i++ {
        fmt.Println(from, ":", i)
    }
}

func main() {
    f("direct")
    go f("goroutine")

    // 匿名函数的方式
    go func(msg string) {
        fmt.Println(msg)
    }("going")

    time.Sleep(time.Second)
    fmt.Println("done")

}
// result
direct : 0
direct : 1
direct : 2
direct : 3
going // 此处执行顺序有变化，说明是并行的
goroutine : 0
goroutine : 1
goroutine : 2
goroutine : 3
done
```



# channel

- channel是连接goroutine的管道
- 一个goroutine可以向channel中发送数据，另一个goroutine负责接收数据
- 使用make定义指定类型的channel
- 默认情况下，发送方和接收方都会阻塞，直到双方都准备好

- - 注意：默认情况下长度是0，要求双方都准备好，否则报死锁，或者直接走default（select情况下）

```go
package main

import "fmt"

func main() {
    messages := make(chan string)
    go func() { messages <- "ping" }()

    msg := <-messages
    fmt.Println(msg)
}
// result
ping
```



# channel buffering

- 默认情况下 channel 是没有缓冲的，必须有发送方和接收方
- 缓冲的channel接收一定的值，而这些值可以没有接收方

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    messages := make(chan string, 2)
    messages <- "buffered"
    messages <- "channel"

    go func() {
        fmt.Println(<-messages)
        fmt.Println(<-messages)
        fmt.Println(<-messages) // 会阻塞
        
    }()

    time.Sleep(time.Second)
    fmt.Println("done")
}
// result
buffered
channel
done
```



# channel  synchronization

- 使用channel进行协程之间的同步操作
- 一个线程可以等待另一个线程执行完成后继续执行

```go
package main

import (
    "fmt"
    "time"
)

func worker(done chan bool) {
    fmt.Println("working")
    time.Sleep(time.Second)
    fmt.Println("done")
    done <- true
}

func main() {

    done := make(chan bool, 1)
    go worker(done)

    // 如果没有这句，那么main线程会直接停止
    <-done

}
// result
working
done
```



# channel  directions

- 指定channel的方向，如只能读取，和只能写入

```go
package main

import "fmt"

func ping(pings chan<- string, msg string) {
    pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
    msg := <-pings
    pongs <- msg
}

func main() {
    pings := make(chan string, 1)
    pongs := make(chan string, 1)
    ping(pings, "passed msg")
    pong(pings, pongs)
    fmt.Println(<-pongs)
}
// result
passed msg
```



# select

- 用于等待多channel操作
- 结合goroutine 和 channel使用
- select是Go中的一个控制结构，类似于switch语句，用于处理异步IO操作。select会监听case语句中channel的读写操作，当case中channel读写操作为非阻塞状态（即能读写）时，将会触发相应的动作
- select中的case语句必须是一个channel操作，select中的default子句总是可运行的
- 如果有多个case都可以运行，select会随机公平地选出一个执行，其他不会执行
- 如果没有可运行的case语句，且有default语句，那么就会执行default的动作
- 如果没有可运行的case语句，且没有default语句，select将阻塞，直到某个case通信可以运行

```go
package main

import (
    "fmt"
    "time"
)

// 模拟接收rpc消息
func main() {

    c1 := make(chan string)
    c2 := make(chan string)

    go func() {
        time.Sleep(time.Second)
        c1 <- "one"
    }()

    go func() {
        time.Sleep(2 * time.Second)
        c2 <- "two"
    }()

    for i := 0; i < 2; i++ {
        // 会在此处阻塞
        select {
        case msg1 := <-c1:
            fmt.Println("received:", msg1)
        case msg2 := <-c2:
            fmt.Println("received:", msg2)
        }
    }

}
// result
received: one
received: two
```



# timeout

- time.sleep 延迟一段时间
- time.After 返回一个channel，在指定超时之后返回一个标识

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    c1 := make(chan string, 1)
    go func() {
        time.Sleep(2 * time.Second)
        c1 <- "result 1"
    }()

    select {
    case res := <-c1:
        fmt.Println(res)
    case <-time.After(1 * time.Second):
        fmt.Println("timeout 1")
    }

    c2 := make(chan string, 1)
    go func() {
        time.Sleep(2 * time.Second)
        c2 <- "result 2"
    }()

    select {
    case res := <-c2:
        fmt.Println(res)
    case <-time.After(3 * time.Second):
        fmt.Println("timeout 2")
    }
}
// result
timeout 1
result 2
```



# non-blocking channel operation

- 通常是阻塞的channel
- 通过select 和 default可以实现非阻塞的发送，接收，以及多路 select
- 当没有消息的时候，会走向default分支
- 关于打印结果的疑虑

- - https://segmentfault.com/q/1010000011312721

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    // 无缓存channel，接收双方必须同时准备好
    messages := make(chan string)
    signals := make(chan bool)

    // 由于此时没有发送者，则直接进入default分支
    select {
    case msg := <-messages:
        fmt.Println("received message", msg)
    default: // 没有消息直接走到default处理
        fmt.Println("no message received")
    }

    msg := "hi"
    select {
    // 此时没有接收者，也直接进入到default
    case messages <- msg:
        fmt.Println("send message", msg)
    // 如果去除default，那么fatal error: all goroutines are asleep - deadlock!
    default:
        fmt.Println("no message send")
    }

    select {
    case msg := <-messages:
        fmt.Println("received message", msg)
    case sig := <-signals:
        fmt.Println("received signal", sig)
    default:
        fmt.Println("no activity")
    }

    time.Sleep(time.Second)
}

// result
no message received
no message send
no activity
```

- 可以执行的情况

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    // 无缓存channel，接收双方必须同时准备好
    // 如果有缓存，则不需要，都准备好，直到缓存满阻塞
    messages := make(chan string, 1)
    //
    go func() {
        for {
            select {
            case msg := <-messages:
                fmt.Println("received message", msg)
            default:
                //fmt.Println("no message received")
            }
        }
    }()

    msg := "hi"
    go func() {
        for {
            select {
            // 当接收方没有准备好，则直接执行default
            case messages <- msg:
                fmt.Println("send message", msg)
            default:
                //fmt.Println("no message send")
            }
        }
    }()

    time.Sleep(1 * time.Second)
}
```



# closing channel

- 使用builtin 中的 close方法关闭channel，同时从channel中获取数据，第二个参数表明是否存在值

```go
package main

import "fmt"

func main() {
    jobs := make(chan int, 5)
    done := make(chan bool)

    go func() {
        for {
            j, more := <-jobs
            // more 在close方法后，设置成false
            if more {
                fmt.Println("received job", j)
            } else {
                fmt.Println("received all jobs")
                done <- true
                return
            }
        }
    }()

    for j := 1; j <= 3; j++ {
        jobs <- j
        fmt.Println("sent job", j)
    }
    close(jobs)
    fmt.Println("sent all jobs")
    <-done
}
// result
sent job 1
sent job 2
sent job 3
sent all jobs
received job 1
received job 2
received job 3
received all jobs
```



# range over channel

- 使用range迭代循环channel

```go
package main

import "fmt"

func main() {
    queue := make(chan string, 2)
    queue <- "one"
    queue <- "two"
    close(queue) // 如果注释掉，那么for range 会报错 fatal error: all goroutines are asleep - deadlock!

    for elem := range queue {
        fmt.Println(elem)
    }
}
// result
one
two
```

- 如果没有close，同时for range 又是一个go func执行，那么该goroutine会阻塞
- for range 等价于 for { msg:=<-channel_xx }



# timer

- built-in go 内置 对象
- 用于处理延时执行

- - 使用timer的Stop方法可以停止

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    timer1 := time.NewTimer(2 * time.Second)
    fmt.Println("start ", time.Now())
    time.Sleep(7 * time.Second)
    c := <-timer1.C // 2s后 timer1自动会放置一个标识给C，而与之前的sleep无关
    fmt.Println("timer1 fired ", c)

    timer2 := time.NewTimer(time.Second)
    go func() {
        <-timer2.C
        fmt.Println("timer2 fired")
    }()
    stop2 := timer2.Stop()
    if stop2 {
        fmt.Println("timer2 stopped")
    }
    time.Sleep(3 * time.Second)
}
// result
start  2021-02-17 23:15:04.8662214 +0800 CST m=+0.008738401
timer1 fired  2021-02-17 23:15:06.8669218 +0800 CST m=+2.009438801
timer2 stopped
```



# ticker

- built-in 内置，周期性执行
- 也使用Stop进行停止

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    ticker := time.NewTicker(500 * time.Millisecond)
    done := make(chan bool)

    go func() {
        for {
            select {
            case <-done:
                return
            case t := <-ticker.C:
                fmt.Println("tick at ", t)

            }
        }
    }()

    time.Sleep(1600 * time.Millisecond)
    ticker.Stop()
    done <- true
    fmt.Println("ticker stopped")
}
// result
tick at  2021-02-17 23:20:51.3318632 +0800 CST m=+0.509640201
tick at  2021-02-17 23:20:51.831537 +0800 CST m=+1.009314001
tick at  2021-02-17 23:20:52.3322258 +0800 CST m=+1.510002801
ticker stopped
```



# worker pool

- 通过2个channel实现一个worker pool

```go
package main

import (
    "fmt"
    "time"
)

func workerOp(id int, jobs <-chan int, results chan<- int) {
    // 如果没有jobs,则会阻塞，直到close(jobs)
    for j := range jobs {
        fmt.Println("worker", id, "started job", j)
        time.Sleep(time.Second)
        fmt.Println("worker", id, "finished job", j)
        results <- j * 2
    }
}

func main() {

    const numJobs = 5
    jobs := make(chan int, numJobs)
    results := make(chan int, numJobs)

    // 创建worker
    for w := 1; w <= 3; w++ {
        go workerOp(w, jobs, results)
    }
    // 创建任务
    for j := 1; j <= numJobs; j++ {
        jobs <- j
    }
    close(jobs)

    for a := 1; a <= numJobs; a++ {
        <-results
    }
}
// result
worker 2 started job 2
worker 1 started job 1
worker 3 started job 3
worker 3 finished job 3
worker 2 finished job 2
worker 1 finished job 1
worker 3 started job 4
worker 2 started job 5
worker 3 finished job 4
worker 2 finished job 5
```



# waitGroup

- 等待多个goroutine关闭
- sync.WatiGroup传参必须是指针类型

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

func workerWg(id int, wg *sync.WaitGroup) {
    defer wg.Done()
    fmt.Printf("Worker %d starting\n", id)
    time.Sleep(time.Second)
    fmt.Printf("Worker %d done\n", id)
}

func main() {

    var wg sync.WaitGroup

    for i := 1; i <= 5; i++ {
        // waitGroup 增加 1 ,wg.Done 减少1
        wg.Add(1)
        // 传参需要引用对象地址
        go workerWg(i, &wg)
    }

    // 阻塞，等待wg内计数是0
    wg.Wait()
}
// result
Worker 5 starting
Worker 3 starting
Worker 4 starting
Worker 1 starting
Worker 2 starting
Worker 2 done
Worker 5 done
Worker 4 done
Worker 3 done
Worker 1 done
```



# rate limiting

- 速率控制机制，支持go的goroutines 和 channel 以及tickers
- 使用time.Tick产生固定速度进行消息

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    // 同时产生5个请求
    requests := make(chan int, 5)
    for i := 1; i <= 5; i++ {
        requests <- i
    }
    close(requests)

    // Tick 是 time.NewTicker的封装
    limiter := time.Tick(200 * time.Millisecond)

    for req := range requests {
        // 固定速率从requests中获取信息，进行消费处理
        <-limiter
        fmt.Println("request", req, time.Now())
    }

    fmt.Println("---------")

    // 示例2 上个例子是串行执行，这个例子允许同时处理3个请求
    burstyLimiter := make(chan time.Time, 3)

    for i := 0; i < 3; i++ {
        burstyLimiter <- time.Now()
    }
    // 开启一个协程，固定速率在limiter中存放请求执行标识
    go func() {
        for t := range time.Tick(200 * time.Millisecond) {
            // 如果 chan 满，则阻塞
            burstyLimiter <- t
        }
    }()

    burstyRequests := make(chan int, 5)
    for i := 1; i <= 5; i++ {
        burstyRequests <- i
    }
    close(burstyRequests)

    for req := range burstyRequests {
        <-burstyLimiter
        fmt.Println("request", req, time.Now())
    }
}
// result
// 如下基本都在200ms串行执行
request 1 2021-02-18 15:34:25.3032719 +0800 CST m=+0.212492601
request 2 2021-02-18 15:34:25.5015196 +0800 CST m=+0.410740301
request 3 2021-02-18 15:34:25.6999848 +0800 CST m=+0.609205501
request 4 2021-02-18 15:34:25.9004833 +0800 CST m=+0.809704001
request 5 2021-02-18 15:34:26.0999467 +0800 CST m=+1.009167401
--------- 
// 如下前3个是并发执行，然后并发标识消费完，变成了串行执行，等于在空闲的时候存储并发标识，等请求来了可以突击处理一批
request 1 2021-02-18 15:34:26.1000173 +0800 CST m=+1.009238001
request 2 2021-02-18 15:34:26.1000173 +0800 CST m=+1.009238001
request 3 2021-02-18 15:34:26.1000173 +0800 CST m=+1.009238001
request 4 2021-02-18 15:34:26.3005277 +0800 CST m=+1.209748401
request 5 2021-02-18 15:34:26.5008614 +0800 CST m=+1.410082101
```



# atomic counter

- 使用sync/atomic 实现多goroutine的交互
- 使用atomic.AddUint64实现原子性操作

```go
package main

import (
    "fmt"
    "sync"
    "sync/atomic"
)

func main() {

    var ops uint64
    var wg sync.WaitGroup
    for i := 0; i < 50; i++ {
        wg.Add(1)
        go func() {
            for c := 0; c < 1000; c++ {
                atomic.AddUint64(&ops, 1)
            }
            wg.Done()
        }()
    }
    wg.Wait()
    fmt.Println("ops:", ops)
}
// result
ops: 50000
```



# mutex

- 使用锁进行状态的读写

```
package main

import (
    "fmt"
    "math/rand"
    "sync"
    "sync/atomic"
    "time"
)

func main() {
    var state = make(map[int]int)
    var mutex = &sync.Mutex{}

    var readOps uint64
    var writeOps uint64

    for r := 0; r < 100; r++ {
        go func() {
            total := 0
            for {
                key := rand.Intn(5)
                mutex.Lock()
                total += state[key]
                mutex.Unlock()
                atomic.AddUint64(&readOps, 1)
                time.Sleep(time.Millisecond)
            }
        }()
    }

    for w := 0; w < 10; w++ {
        go func() {
            for {
                key := rand.Intn(5)
                val := rand.Intn(100)
                mutex.Lock()
                state[key] = val
                mutex.Unlock()
                atomic.AddUint64(&writeOps, 1)
                time.Sleep(time.Millisecond)
            }
        }()
    }
    time.Sleep(time.Second)

    // 读取执行次数
    readOpsFinal := atomic.LoadUint64(&readOps)
    fmt.Println("readOps:",readOpsFinal)

    writeOpsFinal := atomic.LoadUint64(&writeOps)
    fmt.Println("writeOps:",writeOpsFinal)

    mutex.Lock()
    fmt.Println("state:",state)
    mutex.Unlock()
}
// result
readOps: 56980
writeOps: 5692
state: map[0:47 1:97 2:31 3:30 4:73]
```



# stateful goroutine

- 使用channel达到上面mutex的共享内存变量的效果
- 思路是每个数据又一个channel，然后这个channel同时只属于一个goroutine
- 该例子比上面的例子复杂，但是逻辑上更容易理解，使用互斥锁会减少代码的鲁棒性，尽量少用互斥锁

```go
package main

import (
    "fmt"
    "math/rand"
    "sync/atomic"
    "time"
)

type readOp struct {
    key  int
    resp chan int
}

type writeOp struct {
    key  int
    val  int
    resp chan bool
}

func main() {

    var readOps uint64
    var writeOps uint64

    reads := make(chan readOp)
    writes := make(chan writeOp)

    // 集中处理读写，解耦
    go func() {
        var state = make(map[int]int)
        for {
            select {
            case read := <-reads:
                read.resp <- state[read.key]
            case write := <-writes:
                state[write.key] = write.val
                write.resp <- true
            }
        }
    }()

    // 读操作
    for r := 0; r < 100; r++ {
        go func() {
            for {
                read := readOp{
                    key:  rand.Intn(5),
                    resp: make(chan int),
                }

                // 阻塞
                reads <- read
                // 阻塞
                <-read.resp
                atomic.AddUint64(&readOps, 1)
                time.Sleep(time.Millisecond)
            }
        }()
    }

    // 写操作
    for w := 0; w < 10; w++ {
        go func() {
            for {
                write := writeOp{
                    key:  rand.Intn(5),
                    val:  rand.Intn(100),
                    resp: make(chan bool),
                }
                writes <- write
                <-write.resp
                atomic.AddUint64(&writeOps, 1)
                time.Sleep(time.Millisecond)
            }
        }()
    }

    time.Sleep(time.Second)
    readOpsFinal := atomic.LoadUint64(&readOps)
    fmt.Println("readOps:", readOpsFinal)

    writeOpsFinal := atomic.LoadUint64(&writeOps)
    fmt.Println("writeOps:", writeOpsFinal)
}
// result
readOps: 56505
writeOps: 5664
```



# sorting

- 排序
- 使用sort.go中的排序方法

```go
package main

import (
    "fmt"
    "sort"
)

func main() {
    strs := []string{"c", "a", "b"}
    // 对string进行排序
    sort.Strings(strs)
    fmt.Println("strings:", strs)

    ints := []int{7, 2, 4}
    // 对 int 进行排序
    sort.Ints(ints)
    fmt.Println("Ints: ", ints)

    // 判断是否已经排序
    s := sort.IntsAreSorted(ints)
    fmt.Println("sorted:", s)
}
// result
strings: [a b c]
Ints:  [2 4 7]
sorted: true
```



# sorting by function

- 自定义排序，需要实现sort.Interface Len Swap Less方法

```go
package main

import (
    "fmt"
    "sort"
)

// 给 []string 的类型别名
type byLength []string

// 继承 sort的接口
func (s byLength) Len() int {
    return len(s)
}

func (s byLength) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}

func (s byLength) Less(i, j int) bool {
    return len(s[i]) < len(s[j])
}

func main() {
    fruits := []string{"peach", "banana", "kiwi"}
    // sort 接收 排序的interface
    // byLength这里是强转，转换为byLength对象
    sort.Sort(byLength(fruits))
    fmt.Println(fruits)
}
// result
[kiwi peach banana]
```



# panic

- 快速处理正常操作中不该发生的错误
- 如果对一个异常不知道如何处理，或者非期望的异常，则抛出
- 注意，与某些使用异常来处理许多错误的语言不同，Go中的习惯用法是尽可能使用指示错误的返回值

```go
package main

import "os"

func main() {
    panic("a problem")

    _, err := os.Create("/tmp/file")
    if err != nil {
        panic(err)
    }
}
// result
panic: a problem

goroutine 1 [running]:
main.main()
    D:/Personal Files/note/go-study/code/base/demo/P42-panic.go:6 +0x40
```



# defer

- 延迟用于确保函数调用在程序执行的后期执行，通常用于清除。延迟通常用于确保和最终将在其他语言中使用

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    f := createFile("/tmp/defer.txt")
    // 在main结束后执行
    defer closeFile(f)
    writeFile(f)
}

func createFile(p string) *os.File {
    fmt.Println("creating")
    f, err := os.Create(p)
    if err != nil {
        panic(err)
    }
    return f
}

func writeFile(f *os.File) {
    fmt.Println("writing")
    fmt.Fprintln(f, "data")
}

func closeFile(f *os.File) {
    fmt.Println("closing")
    err := f.Close()

    if err != nil {
        fmt.Fprintln(os.Stderr, "error: %v\n", err)
        os.Exit(1)
    }
}
```



# collection function

- go不支持泛型，有些集合的函数系统提供，或者自定义实现

```go
package main

import (
    "fmt"
    "strings"
)

func Index(vs []string, t string) int {
    for i, v := range vs {
        if v == t {
            return i
        }
    }
    return -1
}

func Include(vs []string, t string) bool {
    return Index(vs, t) >= 0
}

// 满足f的判断函数，则返回true
// 只要有一个满足就true
func Any(vs []string, f func(string) bool) bool {
    for _, v := range vs {
        if f(v) {
            return true
        }
    }
    return false
}

// 有一个不满足，就false
func All(vs []string, f func(string) bool) bool {
    for _, v := range vs {
        if !f(v) {
            return false
        }
    }
    return true
}

func Filter(vs []string, f func(string) bool) []string {
    vsf := make([]string, 0)
    for _, v := range vs {
        if f(v) {
            vsf = append(vsf, v)
        }
    }
    return vsf
}

// 映射，从一个[]string映射到另一个[]string
func Map(vs []string, f func(string) string) []string {
    vsm := make([]string, len(vs))
    for i, v := range vs {
        vsm[i] = f(v)
    }
    return vsm
}

func main() {
    var strs = []string{"peach","apple","pear","plum"}
    fmt.Println(Index(strs,"pear"))
    fmt.Println(Include(strs,"grape"))
    fmt.Println(Any(strs, func(v string) bool {
        return strings.HasPrefix(v,"p")
    }))
    fmt.Println(All(strs, func(v string) bool {
        return strings.HasPrefix(v,"p")
    }))
    fmt.Println(Filter(strs, func(v string) bool {
        return strings.Contains(v,"e")
    }))
    fmt.Println(Map(strs,strings.ToUpper))
}
// result
2
false
true
false
[peach apple pear]
[PEACH APPLE PEAR PLUM]
```



# string function

- strings包的函数

```go
package main

import (
    "fmt"
    s "strings"
)

var p = fmt.Println

func main() {
    p("contains: ", s.Contains("test", "es"))
    p("count: ", s.Count("test", "t"))
    p("hasPerfix: ", s.HasPrefix("test", "te"))
    p("hasSuffix: ", s.HasSuffix("test", "st"))
    p("Index: ", s.Index("test", "e"))
    p("join: ", s.Join([]string{"a", "b"}, "-"))
    p("repeat: ", s.Repeat("a", 5))
    p("replace:", s.Replace("foo", "o", "0", -1)) // 全部
    p("replace:", s.Replace("foo", "o", "0", 1)) // 替换1个
    p("split:", s.Split("a-b-c-d-e", "-"))
    p("toLower", s.ToLower("TEST"))
    p("toUpper", s.ToUpper("test"))
    p()

    p("len:", len("hello"))
    p("char:", "hello"[1])
}
// result
contains:  true
count:  2
hasPerfix:  true
hasSuffix:  true
Index:  1
join:  a-b
repeat:  aaaaa
replace: f00
replace: f0o
split: [a b c d e]
toLower test
toUpper TEST

len: 5
char: 101
```



# string formatting

- go 提供string转换的支持，通过printf进行转换

```go
package main

import (
    "fmt"
    "os"
)

type point struct {
    x, y int
}

func main() {

    p := point{1, 2}
    fmt.Printf("%v\n", p)  // 打印出point实例的值
    fmt.Printf("%+v\n", p) // 打印出 point 实例的值，并包含 属性的名称
    fmt.Printf("%#v\n", p) // 打印出 point 的包名等描述信息，同时含有值

    fmt.Printf("%T\n", p)  // 打印出point的类型
    fmt.Printf("%p\n", &p) // 打印出point的地址值

    fmt.Printf("%t\n", true) // 格式化boolean类型

    fmt.Printf("%d\n", 124) // 转换为10进制输出
    fmt.Printf("%b\n", 14)  // 转换为2进制输出
    fmt.Printf("%x\n", 456) // 转换为对应的16进制

    fmt.Printf("%c\n", 33)   // 转换为对应的字符值
    fmt.Printf("%f\n", 78.9) // 对float类型，使用%f进行输出

    fmt.Printf("%e\n", 12340000.0) // 科学计数法 e
    fmt.Printf("%E\n", 12340000.0) // 科学计数法 E

    fmt.Printf("%s\n", "\"string\"") // 基本的字符串打印，会将 \ 进行转义
    fmt.Printf("%q\n", "\"string\"") // 原样输出字符串，不会将 \ 进行转义，会打印出 \

    fmt.Printf("%x\n", "hex this") // 对每个字符进行16进制输出

    fmt.Printf("|%6d|%6d|\n", 12, 345)       // 格式化输出数字，默认右对齐，不足部分用空格填充
    fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45) // 格式化输出浮点类型数字，6表示总数字数目，.2表示保留小数个数
    fmt.Printf("|%6s|%6s|\n", "foo", "b")   // 格式化输出 字符
    fmt.Printf("|%-6s|%-6s|\n", "foo", "b") // 格式化输出，左对齐，使用 -

    s := fmt.Sprintf("a %s", "string") // 返回一个格式化的字符串，而非打印出来
    fmt.Println(s)

    fmt.Fprintf(os.Stderr, "an %s\n", "error") // 输出到流中
}
// result
{1 2}
{x:1 y:2}
main.point{x:1, y:2}
main.point
0xc00000a0d0
true
124
1110
1c8
!
78.900000
1.234000e+07
1.234000E+07
"string"
"\"string\""
6865782074686973
|    12|   345|
|  1.20|  3.45|
|   foo|     b|
|foo   |b     |
a string
an error
```



# regular expression

- 正则表达式

```go
package main

import (
    "bytes"
    "fmt"
    "regexp"
)

func main() {

    // 直接正则检验
    match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
    fmt.Println(match) // true

    // 声明一个正则表达式对象
    r, _ := regexp.Compile("p([a-z]+)ch")
    // 使用该对象进行正则处理
    fmt.Println(r.MatchString("peach")) // true
    // 找到第一个子串
    fmt.Println(r.FindString("peach punch")) // peach
    // 找到第一个子串 并返回对应的位置
    fmt.Println(r.FindStringIndex("peach punch")) // [0 5]
    // 子匹配变量包括关于整个模式匹配和这些匹配中的子匹配的信息。例如，这将返回p([a-z]+)ch和([a-z]+)的信息。
    fmt.Println(r.FindStringSubmatch("peach punch"))
    fmt.Println(r.FindStringSubmatchIndex("peach punch"))
    //
    fmt.Println(r.FindAllStringSubmatchIndex("peach punch pinch", -1))
    fmt.Println(r.FindAllString("peach punch pinch", 2))

    fmt.Println(r.Match([]byte("peach")))

    r = regexp.MustCompile("p([a-z]+)ch")
    fmt.Println(r)

    fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))
    in := []byte("a peach")
    out := r.ReplaceAllFunc(in, bytes.ToUpper)
    fmt.Println(string(out))
}
// result
true
true
peach
[0 5]
[peach ea]
[0 5 1 3]
[[0 5 1 3] [6 11 7 9] [12 17 13 15]]
[peach punch]
true
p([a-z]+)ch
a <fruit>
a PEACH
```



# json

- go built-in内置 json的转换
- http://blog.golang.org/2011/01/json-and-go.html

```go
package main

import (
    "encoding/json"
    "fmt"
    "os"
)

type response1 struct {
    Page   int
    Fruits []string
}

type response2 struct {
    Page   int      `json:"page"`
    Fruits []string `json:"fruits"`
}

func main() {

    // 基本类型的转换为json str
    boolB, _ := json.Marshal(true)
    fmt.Println(string(boolB))

    intB, _ := json.Marshal(1)
    fmt.Println(string(intB))

    fltB, _ := json.Marshal(2.34)
    fmt.Println(string(fltB))

    strB, _ := json.Marshal("gopher")
    fmt.Println(string(strB))

    slcD := []string{"apple", "peach", "pear"}
    slcB, _ := json.Marshal(slcD)
    fmt.Println(string(slcB))

    mapD := map[string]int{"apple": 5, "lettuce": 7}
    mapB, _ := json.Marshal(mapD)
    fmt.Println(string(mapB))

    // 复杂类型
    res1D := &response1{
        Page:   1,
        Fruits: []string{"apple", "peach", "pear"},
    }
    res1B, _ := json.Marshal(res1D)
    fmt.Println(string(res1B))

    res2D := &response2{
        Page:   1,
        Fruits: []string{"apple", "peach", "pear"},
    }
    res2B, _ := json.Marshal(res2D)
    fmt.Println(string(res2B))

    // json str转map
    byt := []byte(`{"num":6.13,"strs":["a","b"]}`)
    var dat map[string]interface{}
    if err := json.Unmarshal(byt, &dat); err != nil {
        panic(err)
    }
    fmt.Println(dat)

    num := dat["num"].(float64)
    fmt.Println(num)

    strs := dat["strs"].([]interface{})
    str1 := strs[0].(string)
    fmt.Println(str1)

    str := `{"page":1,"fruits":["apple","peach"]}`
    res := response2{}
    json.Unmarshal([]byte(str), &res)
    fmt.Println(res)
    fmt.Println(res.Fruits[0])

    // 指定输出流进行输出操作
    enc := json.NewEncoder(os.Stdout)
    d := map[string]int{"apple": 5, "lettuce": 7}
    enc.Encode(d)
}
// result
true
1
2.34
"gopher"
["apple","peach","pear"]
{"apple":5,"lettuce":7}
{"Page":1,"Fruits":["apple","peach","pear"]}
{"page":1,"fruits":["apple","peach","pear"]}
map[num:6.13 strs:[a b]]
6.13
a
{1 [apple peach]}
apple
{"apple":5,"lettuce":7}
```



# xml

```go
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
// result
 <plant>
  <id>12</id>
  <name>coffee</name>
  <origin>Ethiopia</origin>
  <origin>Brazil</origin>
 </plant>
<?xml version="1.0" encoding="UTF-8"?>
 <plant>
  <id>12</id>
  <name>coffee</name>
  <origin>Ethiopia</origin>
  <origin>Brazil</origin>
 </plant>
Plant id=12 name=coffee origin=[Ethiopia Brazil]
 <nesting>
  <parent>
   <child>
    <plant>
     <id>12</id>
     <name>coffee</name>
     <origin>Ethiopia</origin>
     <origin>Brazil</origin>
    </plant>
    <plant>
     <id>11</id>
     <name>tomato</name>
     <origin>Mexico</origin>
     <origin>California</origin>
    </plant>
   </child>
  </parent>
 </nesting>
```



# time

- 时间计算

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    p := fmt.Println

    now := time.Now()
    p(now)

    then := time.Date(2021, 2, 19, 13, 32, 22, 651387237, time.UTC)
    p(then)
    p(then.Year())
    p(then.Month())
    p(then.Day())
    p(then.Hour())
    p(then.Minute())
    p(then.Second())
    p(then.Nanosecond())
    p(then.Location())
    p(then.Weekday())

    // then时间和now进行比较，比较精度到ns
    p(then.Before(now))
    p(then.After(now))
    p(then.Equal(now))

    diff := now.Sub(then)
    p(diff)
    p(diff.Hours()) // diff使用hour进行表示
    p(diff.Minutes()) // diff使用minute 进行表示
    p(diff.Seconds()) // diff 使用second进行表示
    p(diff.Nanoseconds())

    // 时间上加法
    p(then.Add(diff))
    // 时间上减法
    p(then.Add(-diff))
}
// result
2021-02-19 13:38:05.489651 +0800 CST m=+0.010963501
2021-02-19 13:32:22.651387237 +0000 UTC
2021
February
19
13
32
22
651387237
UTC
Friday
false
true
false
-7h54m17.161736237s
-7.904767148954722
-474.28602893728333
-28457.161736237
-28457161736237
2021-02-19 05:38:05.489651 +0000 UTC
2021-02-19 21:26:39.813123474 +0000 UTC
```



# epoch timestamp

- 时间戳

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    now := time.Now()
    fmt.Println(now)
    // 获取时间戳
    secs := now.Unix()
    fmt.Println(secs)

    nanos := now.UnixNano()
    fmt.Println(nanos)

    // 没有ms的时间戳，需要通过ns进行转换
    millis := nanos / 1000000
    fmt.Println(millis)

    // 将时间戳转换为时间
    // 如果没有ns，则可以是0
    fmt.Println(time.Unix(secs, 0))
    // 如果只有s，那么s 是0
    fmt.Println(time.Unix(0, nanos))
}
// result
2021-02-19 13:49:47.4231111 +0800 CST m=+0.011367501
1613713787
1613713787423111100
1613713787423
2021-02-19 13:49:47 +0800 CST
2021-02-19 13:49:47.4231111 +0800 CST
```



# time formatting/parsing

- 时间转换

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    p := fmt.Println

    t := time.Now()
    // time -> time_str
    p(t.Format(time.RFC3339))

    // time_str -> time
    t1, e := time.Parse(
        time.RFC3339,
        "2012-11-01T22:08:41+00:00",
    )
    p(t1)

    p(t.Format("3:04PM"))
    p(t.Format("Mon Jan _2 15:04:05 2006"))
    p(t.Format("2006-01-02T15:04:05.999999-07:00"))

    // 注意时间格式
    form := "3 04 PM"
    t2,e := time.Parse(form,"8 41 PM")
    p(t2)

    fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
        t.Year(), t.Month(), t.Day(),
        t.Hour(), t.Minute(), t.Second())

    ansic := "Mon Jan _2 15:04:05 2006"
    _, e = time.Parse(ansic, "8:41PM")

    p(e)

}
// result
2021-02-19T13:56:24+08:00
2012-11-01 22:08:41 +0000 +0000
1:56PM
Fri Feb 19 13:56:24 2021
2021-02-19T13:56:24.299161+08:00
0000-01-01 20:41:00 +0000 UTC
2021-02-19T13:56:24-00:00
parsing time "8:41PM" as "Mon Jan _2 15:04:05 2006": cannot parse "8:41PM" as "Mon"
```



# random number

- 随机数

```go
package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    // 每次执行的结果都是 81 87
    fmt.Println(rand.Intn(100),",",rand.Intn(100))
    // 获取 float 的随机数 0.0-1.0
    fmt.Println(rand.Float64())
    // 范围 5.0-10.0
    fmt.Println(rand.Float64()*5+5,",",rand.Float64()*5+5)

    // 设置不同的种子，产生不同的随机数
    s1 := rand.NewSource(time.Now().UnixNano())
    r1 := rand.New(s1)

    fmt.Println(r1.Intn(100),",",r1.Intn(100))

    // 一旦固定了种子，那么每次启动main执行的结果都一样
    s2 := rand.NewSource(42)
    r2 := rand.New(s2)
    fmt.Println(r2.Intn(100),",",r2.Intn(100))

    s3 := rand.NewSource(42)
    r3 := rand.New(s3)
    fmt.Println(r3.Intn(100),",",r3.Intn(100))
}
// result
81 , 87
0.6645600532184904
7.1885709359349015 , 7.123187485356329
37 , 68
5 , 87
5 , 87
```



# number parsing

- 数字转换
- 使用strconv 进行转换

```go
package main

import (
    "fmt"
    "strconv"
)

func main() {
    // 解析成64位精度的float
    f, _ := strconv.ParseFloat("1.234", 64)
    fmt.Println(f)

    // 0 表示从字符串中推出基数，64 表示数据大小64bit
    i, _ := strconv.ParseInt("123", 0, 64)
    fmt.Println(i)

    // 自动识别16进制的字符串
    d, _ := strconv.ParseInt("0x1c8", 0, 64)
    fmt.Println(d)

    u, _ := strconv.ParseUint("789", 0, 64)
    fmt.Println(u)

    // ParseInt(s, 10, 0) 10进制数字字符串转为数字
    k, _ := strconv.Atoi("135")
    fmt.Println(k)

    // 非数字会报错
    _, e := strconv.Atoi("wat")
    fmt.Println(e)
}
// result
1.234
123
456
789
135
strconv.Atoi: parsing "wat": invalid syntax
```



# url parsing

```go
package main

import (
    "fmt"
    "net"
    "net/url"
)

func main() {
    s := "postgres://user:pass@host.com:5423/path?k=v#f"

    u, err := url.Parse(s)
    if err != nil {
        panic(err)
    }

    fmt.Println(u.Scheme)          // postgres
    fmt.Println(u.User)            // user:pass
    fmt.Println(u.User.Username()) // user
    p, _ := u.User.Password()
    fmt.Println(p) // pass

    fmt.Println(u.Host) // host.com:5423
    // 分割 host.com 和 5423
    host, port, _ := net.SplitHostPort(u.Host)
    fmt.Println(host) // host.com
    fmt.Println(port) // 5423

    fmt.Println(u.Path)     // /path
    fmt.Println(u.Fragment) // f

    fmt.Println(u.RawQuery) // k=v
    m, _ := url.ParseQuery(u.RawQuery)
    fmt.Println(m)         // map[k:[v]]
    fmt.Println(m["k"][0]) // v
}

// result
postgres
user:pass
user
pass
host.com:5423
host.com
5423
/path
f
k=v
map[k:[v]]
v
```



# sha1 hash

- SHA1哈希经常用于计算二进制或文本blob的短标识。例如，git版本控制系统广泛地使用SHA1s来识别版本化的文件和目录。下面是如何在Go中计算SHA1哈希值
- 使用sha1.New，s.Sum计算
- 结果经常打印为16进制
- 使用其他方式计算hash值，如md5.New

```go
package main

import (
    "crypto/sha1"
    "fmt"
)

func main() {

    s := "sha1 this string"

    h := sha1.New()

    h.Write([]byte(s))
    // sha1求和
    bs := h.Sum(nil)
    fmt.Println(s)
    fmt.Printf("%x\n", bs)
}
// result
sha1 this string
cf23df2207d99a74fbe169e3eba035e633b65d94
```



# base64 encoding

- go支持标准base64以及url兼容的base64

```go
package main

import (
    "encoding/base64"
    "fmt"
)

func main() {

    data := "abc123!?$*&()'-=@~"

    // 标准编码
    sEnc := base64.StdEncoding.EncodeToString([]byte(data))
    fmt.Println(sEnc)

    // 标准解码
    sDec, _ := base64.StdEncoding.DecodeString(sEnc)
    fmt.Println(string(sDec))
    fmt.Println()

    // url base64编码
    uEnc := base64.URLEncoding.EncodeToString([]byte(data))
    fmt.Println(uEnc)

    // url base64解码
    uDec, _ := base64.URLEncoding.DecodeString(uEnc)
    fmt.Println(string(uDec))
}
// result
YWJjMTIzIT8kKiYoKSctPUB+
abc123!?$*&()'-=@~
// 2个编码格式上有所差异，在最后一个字符处
YWJjMTIzIT8kKiYoKSctPUB-
abc123!?$*&()'-=@~
```



# reading file

```go
package main

import (
    "bufio"
    "fmt"
    "io"
    "io/ioutil"
    "os"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {

    // 读取全部文件到内存中
    dat, err := ioutil.ReadFile("tmp/dat")
    check(err)
    fmt.Println(string(dat))

    // 如果是要读取文件的一部分，则先打开文件
    f, err := os.Open("tmp/dat")
    check(err)

    // 读取开始部分
    b1 := make([]byte, 5)
    // 返回的n1表示读取的字节数目
    n1, err := f.Read(b1)
    check(err)
    fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))

    // 读取中间部分，先跳转到该位置
    o2, err := f.Seek(6, 0)
    check(err)

    b2 := make([]byte, 2)
    n2, err := f.Read(b2)
    check(err)
    // 从o2的位置开始读取 n2个字节
    fmt.Printf("%d bytes @ %d", n2, o2)
    fmt.Printf("%v\n", string(b2[:n2]))

    o3, err := f.Seek(6, 0)
    check(err)
    b3 := make([]byte, 2)
    // 如果读的数目没有满足最小的要求的数目的数据，则抛出异常ErrUnexpectedEOF
    n3, err := io.ReadAtLeast(f, b3, 2)
    check(err)
    fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

    // 回退到开头
    _, err = f.Seek(0, 0)
    check(err)

    // 缓冲流
    r4 := bufio.NewReader(f)
    b4, err := r4.Peek(5)
    check(err)
    fmt.Printf("5 bytes: %s\n", string(b4))

    f.Close()
}
    
$ echo "hello" > /tmp/dat
$ echo "go" >>   /tmp/dat
$ go run reading-files.go
hello
go
5 bytes: hello
2 bytes @ 6: go
2 bytes @ 6: go
5 bytes: hello
```



# writing file

```go
package main

import (
    "bufio"
    "fmt"
    "io/ioutil"
    "os"
)

func check2(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    // 直接创建并写入内容到文件
    // 注意：运行前先创建tmp文件夹
    d1 := []byte("hello\ngo\n")
    err := ioutil.WriteFile("d:/tmp/dat1", d1, 0644)
    check2(err)

    // 打开文件，然后开始写
    f, err := os.Create("d:/tmp/dat2")
    check2(err)

    defer f.Close()

    d2 := []byte{115, 111, 109, 101, 10}
    n2, err := f.Write(d2)
    check2(err)
    fmt.Printf("wrote %d bytes\n", n2)

    n3, err := f.WriteString("writes\n")
    check2(err)
    fmt.Printf("wrote %d bytes\n", n3)

    // flush 数据到本地存储中
    f.Sync()

    w := bufio.NewWriter(f)
    n4, err := w.WriteString("buffered\n")
    check2(err)
    fmt.Printf("wrote %d bytes \n", n4)

    w.Flush()
}
// result
wrote 5 bytes
wrote 7 bytes
wrote 9 bytes 
```



# line filter

- 对输入数据进行行过滤
- *从**stdin**读取，输出到**stdout**上*

```go
package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    // 使用缓冲流包裹非缓冲流os.Stdin
    scanner := bufio.NewScanner(os.Stdin)

    for scanner.Scan() {
        ucl := strings.ToUpper(scanner.Text())
        fmt.Println(ucl)
    }

    // 最后检查是否有错误
    if err := scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "error:", err)
        os.Exit(1)
    }

}
// result
hello
HELLO
ss
SS
```

# file path

- 文件路径的解析

```go
package main

import (
    "fmt"
    "path/filepath"
    "strings"
)

func main() {
    // 使用filepath包的工具进行路径的生成，不同系统的路径格式不同
    p := filepath.Join("dir1", "dir2", "filename")
    fmt.Println("p", p)

    // 不建议手动连接路径，使用filepath.join，会去除冗余格式，规范化路径
    fmt.Println(filepath.Join("dir1//", "filename"))
    fmt.Println(filepath.Join("dir1/../dir1", "filename"))

    // dir 和 base 方法会自动分离出路径和文件
    fmt.Println("Dir(p):",filepath.Dir(p))
    fmt.Println("Base(p):",filepath.Base(p))

    filename := "config.json"
    // 分离出扩展名，含有.
    ext := filepath.Ext(filename)
    fmt.Println(ext)

    // 可以分离出文件名，不带扩展名
    fmt.Println(strings.TrimSuffix(filename,ext))

    // 基于base 和 target 之间查找相对路径
    // 如果没有相对路径，则抛出异常
    rel, err := filepath.Rel("a/b", "a/b/t/file")
    if err != nil {
        panic(err)
    }
    fmt.Println(rel)

    rel, err = filepath.Rel("a/b", "a/c/t/file")
    if err != nil {
        panic(err)
    }
    fmt.Println(rel)
}

// result
p dir1\dir2\filename
dir1\filename
dir1\filename
Dir(p): dir1\dir2
Base(p): filename
.json
config
t\file
..\c\t\file
```



# directory

```go
package main

import (
    "fmt"
    "io/ioutil"
    "os"
    "path/filepath"
)

func check3(err error) {
    if err != nil {
        panic(err)
    }
}

func main() {
    // 创建相对文件夹，临时文件夹，在项目的根路径下
    err := os.Mkdir("subdir", 0755)
    check3(err)

    // 删除该文件夹下所有文件，类似rm -rf
    defer os.RemoveAll("subdir")

    createEmptyFile := func(name string) {
        d := []byte("")
        check3(ioutil.WriteFile(name, d, 0644))
    }
    createEmptyFile("subdir/file1")

    // 递归创建文件夹
    err = os.MkdirAll("subdir/parent/child", 0755)
    check3(err)

    createEmptyFile("subdir/parent/file2")
    createEmptyFile("subdir/parent/file3")
    createEmptyFile("subdir/parent/child/file4")

    c, err := ioutil.ReadDir("subdir/parent")
    check3(err)

    fmt.Println("listing subdir/parent")
    for _, entry := range c {
        fmt.Println(" ", entry.Name(), entry.IsDir())
    }

    // 切换到指定文件夹下
    err = os.Chdir("subdir/parent/child")
    check3(err)

    // 读取当前文件夹下的所有文件与文件夹
    c, err = ioutil.ReadDir(".")
    check3(err)

    fmt.Println("Listing subdir/parent/child")
    for _, entry := range c {
        fmt.Println(" ", entry.Name(), entry.IsDir())
    }

    // 回退
    err = os.Chdir("../../..")
    check3(err)

    fmt.Println("Visiting subdir")
    // 传入一个回调方法，在遍历每个文件与文件夹的时候做出相应的动作
    err = filepath.Walk("subdir", visit)

    //fmt.Println(1/0)

}

func visit(p string, info os.FileInfo, err error) error {
    if err != nil {
        return err
    }
    fmt.Println(" ", p, info.IsDir())
    return nil
}
// result
listing subdir/parent
  child true
  file2 false
  file3 false
Listing subdir/parent/child
  file4 false
Visiting subdir
  subdir true
  subdir\file1 false
  subdir\parent true
  subdir\parent\child true
  subdir\parent\child\file4 false
  subdir\parent\file2 false
  subdir\parent\file3 false
```



# temporary files and directories

- go 提供临时文件和文件夹的创建
- 在系统运行完成后销毁，不污染系统

```go
package main

import (
    "fmt"
    "io/ioutil"
    "os"
    "path/filepath"
)


func check4(err error) {
    if err != nil {
        panic(err)
    }
}

func main() {

    // 如果第一个参数为空，则路径为os.TempDir()
    // 第二个参数是文件名称的前缀，每次创建的文件名不同，防止并发情况下冲突
    f, err := ioutil.TempFile("", "sample")
    check4(err)
    // 在linux中，临时文件一般在tmp文件夹下
    fmt.Println("temp file name:",f.Name())
    defer os.RemoveAll(f.Name())

    _, err = f.Write([]byte{1, 2, 3})
    check4(err)

    // 创建临时文件夹
    dname, err := ioutil.TempDir("", "sampledir")
    check4(err)
    fmt.Println("Temp dir name:",dname)

    defer os.RemoveAll(dname)

    fname := filepath.Join(dname, "file1")
    err = ioutil.WriteFile(fname, []byte{1, 2}, 0666)
    check4(err)
}
// result
temp file name: C:\Users\TERREL~2\AppData\Local\Temp\sample136618215
Temp dir name: C:\Users\TERREL~2\AppData\Local\Temp\sampledir649306394
```



# testing

```go
package main

import (
    "fmt"
    "testing"
)

func IntMin(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func TestIntMinBasic(t *testing.T) {
    ans := IntMin(2, -2)
    if ans != -2 {

        t.Errorf("IntMin(2, -2) = %d; want -2", ans)
    }
}

func TestIntMinTableDriven(t *testing.T) {
    var tests = []struct {
        a, b int
        want int
    }{
        {0, 1, 0},
        {1, 0, 0},
        {2, -2, -2},
        {0, -1, -1},
        {-1, 0, -1},
    }

    for _, tt := range tests {

        testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
        t.Run(testname, func(t *testing.T) {
            ans := IntMin(tt.a, tt.b)
            if ans != tt.want {
                t.Errorf("got %d, want %d", ans, tt.want)
            }
        })
    }
}

$ go test -v
== RUN   TestIntMinBasic
--- PASS: TestIntMinBasic (0.00s)
=== RUN   TestIntMinTableDriven
=== RUN   TestIntMinTableDriven/0,1
=== RUN   TestIntMinTableDriven/1,0
=== RUN   TestIntMinTableDriven/2,-2
=== RUN   TestIntMinTableDriven/0,-1
=== RUN   TestIntMinTableDriven/-1,0
--- PASS: TestIntMinTableDriven (0.00s)
    --- PASS: TestIntMinTableDriven/0,1 (0.00s)
    --- PASS: TestIntMinTableDriven/1,0 (0.00s)
    --- PASS: TestIntMinTableDriven/2,-2 (0.00s)
    --- PASS: TestIntMinTableDriven/0,-1 (0.00s)
    --- PASS: TestIntMinTableDriven/-1,0 (0.00s)
PASS
ok      examples/testing    0.023s
```



# cmd-line argument

- 执行程序时读取启动参数

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    // 读取全部的参数
    argsWithProg := os.Args
    // 去除第一个参数，第一个参数是main方法的go文件名
    argsWithoutProg := os.Args[1:]

    arg:=os.Args[3]

    fmt.Println(argsWithProg)
    fmt.Println(argsWithoutProg)
    fmt.Println(arg)
}
// result
D:\Personal Files\note\go-study\code\base\demo>go run P65-command-line-argument.go a b c d
[C:\Users\TERREL~2\AppData\Local\Temp\go-build226346174\b001\exe\P65-command-line-argument.exe a b c d]
[a b c d]
c
```



# cmd-line flag

- 读取命令行中的指定选项，如 wc -l，则-l表示命令行的flag
- go 中的flag库，自带help

```go
package main

import (
    "flag"
    "fmt"
)

func main() {

    // 返回的是一个字符串指针，foo是默认值 usage 是使用描述
    wordPtr := flag.String("word", "foo", "a string")
    numbPtr := flag.Int("numb", 22, "an int")
    boolPtr := flag.Bool("fork", false, "a bool")

    // 可以将flag的值传递给一个变量，前提通过该变量的指针传递
    var svar string
    flag.StringVar(&svar, "svar", "bar", "a string var")

    // 调用parse方法解析flag中的值
    flag.Parse()

    // 输出
    fmt.Println("word:", *wordPtr)
    fmt.Println("numb:",*numbPtr)
    fmt.Println("fork:",*boolPtr)
    fmt.Println("svar:",svar)
    // 解析其他没有定义的flag信息
    fmt.Println("tail:",flag.Args())
}
// result
D:\Personal Files\note\go-study\code\base\demo>go build P66-command-line-flag.go

P66-command-line-flag.exe -h
Usage of P66-command-line-flag.exe:
  -fork
        a bool
  -numb int
        an int (default 22)
  -svar string
        a string var (default "bar")
  -word string
        a string (default "foo")
P66-command-line-flag.exe -word=opt -numb=7 -fork -svar=flag
word: opt
numb: 7
fork: true
svar: flag
tail: []

P66-command-line-flag.exe -word=opt
word: opt
numb: 22
fork: false
svar: bar
tail: []

// 参数要放在最后，否则从参数开始，即使后面是flag，也视为参数
// 所有的标识flag要在arg之前
P66-command-line-flag.exe -word=opt a1 a2 a3 -numb=7
word: opt
numb: 22
fork: false
svar: bar
tail: [a1 a2 a3 -numb=7]
               
// 如果提供一个不存在的flag 则不会正常执行 
P66-command-line-flag.exe -w
flag provided but not defined: -w
Usage of P66-command-line-flag.exe:
...
```



# cmd-line subcommand

- 定义子命令

- - 如go build 与 go get 是不同的go命令

```go
package main

import (
    "flag"
    "fmt"
    "os"
)

func main() {

    // 声明一个命令子集合
    fooCmd := flag.NewFlagSet("foo", flag.ExitOnError)
    // 在该子命令下声明flag
    fooEnable := fooCmd.Bool("enable",false,"enable")
    fooName := fooCmd.String("name","","name")

    barCmd := flag.NewFlagSet("bar",flag.ExitOnError)
    barLevel := barCmd.Int("level",0,"level")

    // 必须要有子命令
    if len(os.Args) < 2 {
        fmt.Println("expected 'foo' or 'bar' subcommands")
        os.Exit(1)
    }

    switch os.Args[1] {
    case "foo":
        fooCmd.Parse(os.Args[2:])
        fmt.Println("subcommand 'foo'")
        fmt.Println(" enable:",*fooEnable)
        fmt.Println(" name:",*fooName)
    case "bar":
        barCmd.Parse(os.Args[2:])
        fmt.Println("subcommand 'bar'")
        fmt.Println(" level:",*barLevel)
        fmt.Println(" tail:",barCmd.Args())
    default:
        fmt.Println("expected 'foo' or 'bar' subcommands")
        os.Exit(1)
    }

}
// result
go build P67-command-line-subcommand.go
P67-command-line-subcommand.exe foo -name=sss -enable
subcommand 'foo'
 enable: true
 name: sss

P67-command-line-subcommand.exe foo -h
Usage of foo:
  -enable
        enable
  -name string
        name

// flag赋值的时候可以使用空格，最好使用=
P67-command-line-subcommand.exe bar -level 8 a1
subcommand 'bar'
 level: 8
 tail: [a1]
```



# environment variable

- 设置环境变量

```go
package main

import (
    "fmt"
    "os"
    "strings"
)

func main() {
    // 环境变量key-value都是string类型
    os.Setenv("FOO", "1")
    fmt.Println("FOO:", os.Getenv("FOO"))
    fmt.Println("BAR:", os.Getenv("BAR"))

    fmt.Println()
    // 遍历所有的环境变量
    for _, e := range os.Environ() {
        pair := strings.SplitN(e, "=", 2)
        //fmt.Println(pair[0]," ",pair[1])
        fmt.Println(pair[0])
    }
}
// result
$ go run environment-variables.go
FOO: 1
BAR: 

TERM_PROGRAM
PATH
SHELL
...

// 环境变量的外部赋值在go命令之前
$ BAR=2 go run environment-variables.go
FOO: 1
BAR: 2
...
```



# HTTP client

```go
package main

import (
    "bufio"
    "fmt"
    "net/http"
)

func main() {
    resp, err := http.Get("http://gobyexample.com")
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    fmt.Println("Response status:", resp.Status)

    scanner := bufio.NewScanner(resp.Body)
    // 读取前5行
    for i := 0; scanner.Scan() && i < 5; i++ {
        fmt.Println(scanner.Text())
    }
    if err := scanner.Err(); err != nil {
        panic(err)
    }

}
// result
Response status: 200 OK
<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8">
    <title>Go by Example</title>
```



# HTTP server

```go
package main

import (
    "fmt"
    "net/http"
)

// 实现一个http.Handler接口，用来处理请求
func hello(w http.ResponseWriter, req *http.Request) {
    // 向w中输出hello字符串
    fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {
    for name, headers := range req.Header {
        for _, h := range headers {
            // 输出header头部的信息
            fmt.Fprintf(w, "%v : %v\n", name, h)
        }
    }
}

func main() {
    // 对get请求进行解析
    http.HandleFunc("/hello", hello)
    http.HandleFunc("/headers", headers)

    http.ListenAndServe(":8080", nil)
}
```



# context

- https://www.jb51.net/article/136970.htm
- context.Context被每个请求创建一次
- 可以用于请求之间的保存上下文信息
- context在Go1.7之后就进入标准库中了。它主要的用处如果用一句话来说，是在于控制goroutine的生命周期。当一个计算任务被goroutine承接了之后，由于某种原因（超时，或者强制退出）我们希望中止这个goroutine的计算任务，那么就用得到这个Context了

```go
package main

import (
    "fmt"
    "net/http"
    "time"
)

func hello2(w http.ResponseWriter, req *http.Request) {

    ctx := req.Context()
    fmt.Println("server: hello handler started")
    defer fmt.Println("server: hello handler ended")

    select {
    case <-time.After(10 * time.Second):
        fmt.Fprintf(w, "hello\n")
    case <-ctx.Done():

        err := ctx.Err()
        fmt.Println("server:", err)
        internalError := http.StatusInternalServerError
        http.Error(w, err.Error(), internalError)
    }
}

func main() {

    http.HandleFunc("/hello", hello2)
    http.ListenAndServe(":8090", nil)
}
```



# spwaning process

- 执行系统命令，在go进程内调用另一个进程的命令
- exec.Command

```go
package main

import (
    "fmt"
    "io/ioutil"
    "os/exec"
)

func main() {

    // 系统命令执行
    dateCmd := exec.Command("date")

    dateOut, err := dateCmd.Output()
    if err != nil {
        panic(err)
    }
    fmt.Println("> date")
    fmt.Println(string(dateOut))

    grepCmd := exec.Command("grep", "hello")

    grepIn, _ := grepCmd.StdinPipe()
    grepOut, _ := grepCmd.StdoutPipe()
    grepCmd.Start()
    grepIn.Write([]byte("hello grep\ngoodbye grep"))
    grepIn.Close()
    grepBytes, _ := ioutil.ReadAll(grepOut)
    grepCmd.Wait()

    fmt.Println("> grep hello")
    fmt.Println(string(grepBytes))

    lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
    lsOut, err := lsCmd.Output()
    if err != nil {
        panic(err)
    }
    fmt.Println("> ls -a -l -h")
    fmt.Println(string(lsOut))
}

// result
$ go run spawning-processes.go 
> date
Wed Oct 10 09:53:11 PDT 2012
> grep hello
hello grep
> ls -a -l -h
drwxr-xr-x  4 mark 136B Oct 3 16:29 .
drwxr-xr-x 91 mark 3.0K Oct 3 12:50 ..
-rw-r--r--  1 mark 1.3K Oct 3 16:28 spawning-processes.go
```



# Exec-ing process

- https://www.jianshu.com/p/e1de8fc52718
- syscall.Exec会执行参数指定的命令，但是并不创建新的进程，只在当前进程空间内执行，即替换当前进程的执行内容，他们重用同一个进程号PID

```go
package main

import (
    "os"
    "os/exec"
    "syscall"
)

func main() {

    // 切换到执行路径 如 ls路径下
    path, err := exec.LookPath("ls")
    if err != nil {
        panic(err)
    }

    // 命令语句，第一个是命令，后面的是该命令的参数
    args := []string{"ls", "-a", "-l", "-h"}
    // 获取环境参数
    env := os.Environ()
    // 执行对应的命令，执行必须放在main语句最后面
    // 执行成功后，当前的进程id会给执行的命令的进程使用
    execErr := syscall.Exec(path, args, env)
    if execErr != nil {
        panic(execErr)
    }
}
```



# signal

- go 语言中接收系统的信号，并进行响应

```go
package main

import (
    "fmt"
    "os"
    "os/signal"
    "syscall"
)

func main() {

    // 声明一个通道用于接收信号
    sigs := make(chan os.Signal, 1)
    done := make(chan bool, 1)

    // 注册通道用于接收消息，后面的2个参数是接收系统信号的类型
    signal.Notify(sigs,syscall.SIGINT,syscall.SIGTERM)

    go func() {
        // 接收到信号后关闭
        sig := <-sigs
        fmt.Println()
        fmt.Println(sig)
        done <- true
    }()

    fmt.Println("awaiting signal")
    <-done
    fmt.Println("exiting")
}
// result
go build P74-signal.go
P74-signal.exe

awaiting signal

interrupt
exiting
```



# exit

```go
package main

import (
    "fmt"
    "os"
)

func main() {

    // 这句答应不会执行
    defer fmt.Println("xx")
    // 关闭程序，返回一个关闭码
    os.Exit(3)
}
```