package tree

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}

//结构体方法
func (node Node) Print() {
	fmt.Print(node.Value, " ")
}
func (node Node) SetValue(value int) {
	node.Value = value
}
func (node *Node) SetValue1(value int) {
	if node == nil {
		fmt.Println(" Setting value to nil. Ignored.")
		return
	}
	node.Value = value
}

func CreatNode(value int) *Node {
	return &Node{Value: value}
}
