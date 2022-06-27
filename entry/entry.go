package main

import (
	"go-learn/tree"
)

func main() {
	var root = tree.TreeNode{Value: 1}
	root.Left = &tree.TreeNode{Value: 2}
	root.Right = &tree.TreeNode{Value: 3}
	root.Right.Left = new(tree.TreeNode)
	// fmt.Println(root.Left)
	root.Left.SetValue(11)
	root.Left.Print()

}
