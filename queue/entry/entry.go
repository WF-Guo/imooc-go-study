package main

import (
	"fmt"
	"mooc_Go_study/queue"
)

func main() {
	q := queue.Queue{1, 2}
	q.Push(3)
	q.Push(4)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
	//初始的q和后面的q不是同一个q，push和pop操作改变了指针

	/*	q.Push("abc")
		fmt.Println(q.Pop())*/
}
