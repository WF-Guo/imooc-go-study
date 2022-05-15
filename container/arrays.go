package main

import "fmt"

func printArray(arr [5]int) {
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func printArray2(arr *[5]int) {
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func printArray3(arr []int) {
	arr[0] = 100
	for _, v := range arr {
		fmt.Print(" ", v)
	}
}
func main() {
	var arr1 [5]int
	arr2 := [3]int{1, 3, 4}
	arr3 := [...]int{2, 4, 6, 8, 10} //让编译器自动判断有多少数据，需要使用...
	var grid [4][5]int
	fmt.Println(arr1, arr2, arr3)
	fmt.Println(grid)

	/*	for i := 0; i < len(arr3); i++ {
			fmt.Println(arr3[i])
		}
		for i, v := range arr3 {
			fmt.Println(i, v)
		}
		for _, v := range arr3 {
			fmt.Println(v)
		}
	*/
	/*	fmt.Println("==============")
		fmt.Println("Modified arr3:")
		printArray(arr3)
		fmt.Println("arr3:")
		fmt.Println(arr3)*/

	/*	fmt.Println("==============")
		fmt.Println("Modified arr3:")
		printArray2(&arr3)
		fmt.Println("arr3:")
		fmt.Println(arr3)*/

	for v := range arr3 {
		fmt.Print(" ", v)
	}
	fmt.Println()
	printArray3(arr3[:])
}
