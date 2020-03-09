package main
import (
	"fmt"
	"d:/study/go-study/code/base/src/go_code/demo/utils" 
	// 如果代码不在GOPATH中，那么引入的自定义的包需要全路径
)
func main(){
	var re = utils.Cal(1,2,'+')
	fmt.Println("结果是=",re)
}
