package main

import (
	"fmt"
	"mooc_Go_study/tree"
)

type myTreeNode struct {
	*tree.Node //Embedding
}

func (node *myTreeNode) Traverse() {
	fmt.Println("This is shadowed.")
}

func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.Node == nil {
		return
	}

	left := myTreeNode{myNode.Left}
	left.postOrder()
	right := myTreeNode{myNode.Right}
	right.postOrder()
	myNode.Print()
}

func main() {

	root := myTreeNode{&tree.Node{Value: 3}}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreatNode(2)

	root.Right.Left.SetValue1(4)

	/*	var proot *treeNode
		proot.setValue1(300)
		proot = &root
		proot.setValue1(500)
		proot.print()*/
	fmt.Println("In-order traversal")
	fmt.Print("root.Traverse():  ")
	root.Traverse()
	fmt.Print("root.Node.Traverse():  ")
	root.Node.Traverse()
	fmt.Println()
	fmt.Println("post-order traversal")
	root.postOrder()

}
