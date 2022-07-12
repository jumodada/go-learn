package main

import (
	"go-learn/tree"
)

type myTreeNode struct {
	node *tree.Node
}

func (myNode myTreeNode) postOrder() {
	if myNode.node == nil {
		return
	}
	myTreeNode{myNode.node.Left}.postOrder()
	myTreeNode{myNode.node.Right}.postOrder()
	myNode.node.Print()
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
