package tree

import "fmt"

type Node struct {
	Value int
	Left, Right *Node
}

// 含有接收者，此处是传值
func (node Node) Print(){
	fmt.Print(node)
}
// 注意指针接收者
func (node *Node) SetValue(value int){
	node.Value = value
}


// 使用工厂方法代替构造函数
func CreateNode(value int) *Node {
	return &Node{Value:value}
}