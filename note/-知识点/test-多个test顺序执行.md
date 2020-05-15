- 当碰到多个test执行，并且要安装顺序执行的时候
  - 针对测试有上下依赖的情况
- 使用t.Run()顺序执行

```go
package main

import (
	"testing"
    "fmt"
)

func TestPrint(t *testing.T){
    t.Run("a1",func(t *testing.T){
        fmt.Println("a1")
    })
    t.Run("a2",func(t *testing.T){
        fmt.Println("a2")
    })
}
```

- 示例

```go
package main

import (
	"testing"
    "fmt"
)

// 小写不能被go test 命令执行
func testPrint1(t *testing.T) {
    res:= Print20()
    fmt.Println("test1")
    if res != 20 {
        t.Errorf("wrong for testPrint1")
    }
}

func testPrint2(t *testing.T) {
    res:= Print20()
    fmt.Println("test2")
    if res != 20 {
        t.Errorf("wrong for testPrint2")
    }
}

// 通过TestAll 执行所有流程
func TestAll(t *testing.T) {
    t.Run("test1",testPrint1)
    t.Run("test2",testPrint2)
}

// 如果有test main 则先执行，如果有m.Run()则执行剩下的Test函数
func TestMain(m *testing.M){
    fmt.Println("test start...")
    m.Run()
}
```

