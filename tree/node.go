package main

import "fmt"

type treeNode struct {
	value       int
	left, right *treeNode
}

//结构体方法
func (node treeNode) print() {
	fmt.Print(node.value)
}
func (node treeNode) setValue(value int) {
	node.value = value
}
func (node *treeNode) setValue1(value int) {
	node.value = value
}
func creatNode(value int) *treeNode {
	return &treeNode{value: value}
}

func main() {
	var root treeNode

	root = treeNode{value: 3}
	root.left = &treeNode{}
	root.right = &treeNode{5, nil, nil}
	root.right.left = new(treeNode)
	root.left.right = creatNode(2)

	root.right.left.setValue(4)
	root.right.left.print()
	fmt.Println()
	root.right.left.setValue1(4)
	root.right.left.print()

}
