package main

import "fmt"

type treeNode struct {
	value       int
	left, right *treeNode
}

func createNode(value int) *treeNode {
	return &treeNode{
		value: value,
	}
}

func (node treeNode) print() {
	fmt.Println(node.value)
}

func (node *treeNode) setValue(value int) {
	node.value = value
}

func main() {
	var root = treeNode{value: 1}
	root.left = &treeNode{value: 2}
	root.right = &treeNode{value: 3}
	root.right.left = new(treeNode)
	//fmt.Println(root)
	root.left.setValue(11)
	root.left.print()
}
