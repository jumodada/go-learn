package main

import (
	"go-learn/tree"
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
	var root = tree.Node{Value: 1}
	//root.Left = &tree.Node{Value: 2}
	//root.Right = &tree.Node{Value: 3}
	//root.Right.Left = new(tree.Node)
	//root.Left.SetValue(11)
	//root.Left.Print()
	myTreeNode{&root}.postOrder()

}
