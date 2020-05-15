- benchmark函数一般以Benchmark开头
- benchmark的case一般会执行b.N次
- 在执行过程中依据实际case的执行时间是否稳定会增加b.N的次数

```go
package main

import (
	"testing"
)

func BenchmarkAll(b *testing.B) {
	for n := 0; n < b.N; n++ { // b.N会变化，做调整，知道测试的方法返回时间趋于稳定之后
		mytest(n) // 方法要有稳定的状态，不要有非稳定的状态，否则跑不完
	}
}
// go test -bench . 进行测试
```

- 稳态的测试

```go
// 一定能跑完，每次返回的时间差不多
func mytest(n int) int {
    return n
}
```

- 非稳态的测试

```go
func mytest(n int) int {
    for n >0 {
        n --
    }
    return n
}
```

