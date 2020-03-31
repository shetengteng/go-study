package main

import (
	"fmt"
	"learngo/tree"
)

type myTreeNode struct {
	node *tree.Node
}

func (myNode *myTreeNode) postOrder(){
	if myNode == nil || myNode.node == nil {
		return
	}

	// 左右节点必须要提取出来，否则报错
	//  cannot call pointer method on myTreeNode literal
	//  cannot take the address of myTreeNode literal
	left := myTreeNode{myNode.node.Left}
	right := myTreeNode{myNode.node.Right}
	left.postOrder()
	right.postOrder()
	myNode.node.Print()

}

func main() {

	var root tree.Node
	fmt.Println(root)

	root = tree.Node{Value:3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5,nil,nil}
	root.Right.Left = new(tree.Node)
	root.Right.Right = tree.CreateNode(2)
	root.Print()
	root.SetValue(56)
	root.Traverse()
	fmt.Println()
	//
	//nodes := []tree.Node{
	//	{Value:3},
	//	{},
	//	{2,nil,&root},
	//}
	//fmt.Println(nodes)

	var myNode = myTreeNode{&root}
	myNode.postOrder()

}