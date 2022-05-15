package main

import (
	"fmt"
)

func update(arr []int) {
	arr[1] = 100
}
func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s := arr[2:6] //左闭右开
	s1 := arr[:6]
	s2 := arr[2:]
	s3 := arr[:]

	fmt.Println("arr[2:6] =", s)
	/*	fmt.Println("arr[2:6] =", s)
		fmt.Println("arr[:6] =", s1)
		fmt.Println("arr[2:] =", s2)
		fmt.Println("arr[:] =", s3)

		fmt.Println("-----------")
		fmt.Println("s1: ", s1)
		fmt.Println("After update")
		update(s1)
		fmt.Println("s1: ", s1)
		fmt.Println("arr：", arr)

		fmt.Println("Reslice")
		fmt.Println(s2)
		s2 = s2[:5]
		fmt.Println(s2)
		s2 = s2[2:]
		fmt.Println(s2)*/

	fmt.Println("Extends")
	fmt.Println(arr)
	s1 = arr[2:6]
	s2 = s1[3:5]
	fmt.Println("s1= ", s1)
	fmt.Println("s2= ", s2)

	s3 = append(s2, 10)
	s4 := append(s3, 11)
	s5 := append(s4, 12)
	fmt.Println(s3, s4, s5)
	fmt.Println(arr)
}
