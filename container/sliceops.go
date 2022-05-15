package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("%v, len=%d, cap=%d\n", s, len(s), cap(s))
}
func main() {
	var s []int
	s1 := []int{2, 4, 6, 8}
	s2 := make([]int, 16)
	s3 := make([]int, 16, 32) //长度为16，容量为32的切片

	fmt.Println(s, s1, s2, s3)

	fmt.Println("Coping elements from slices")
	copy(s2, s1)
	printSlice(s2)

	fmt.Println("Deleting elements from slices")
	s2 = append(s2[:3], s2[4:]...)
	printSlice(s2)

	fmt.Println("Poping from front")
	front := s2[0]
	s2 = s2[1:]
	fmt.Println(front)
	printSlice(s2)

	fmt.Println("Poping from tail")
	tail := s2[len(s2)-1]
	s2 = s2[:len(s2)-1]
	fmt.Println(tail)
	printSlice(s2)
}
