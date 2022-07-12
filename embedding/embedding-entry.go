package main

import (
	"fmt"
	"go-learn/tree"
	"strconv"
)

type myTreeNode struct {
	*tree.Node
}

func (myNode myTreeNode) postOrder() {
	if myNode.Node == nil {
		return
	}
	myTreeNode{myNode.Node.Left}.postOrder()
	myTreeNode{myNode.Node.Right}.postOrder()
	myNode.Node.Print()
}

func main() {
	//var root = tree.Node{Value: 1}
	////root.Left = &tree.Node{Value: 2}
	////root.Right = &tree.Node{Value: 3}
	////root.Right.Left = new(tree.Node)
	////root.Left.SetValue(11)
	////root.Left.Print()
	//myTreeNode{&root}.postOrder()
	i, _ := strconv.ParseInt("1542366690027773952", 10, 64)
	fmt.Println(i)
}
