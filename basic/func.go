package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		//return a / b
		q, _ := div1(a, b)
		return q, nil
	default:
		return 0, fmt.Errorf("unsupported operation: %s", op)
	}
}

func apply(op func(int, int) int, a, b int) int {
	pointer := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(pointer).Name()
	fmt.Printf("Calling function %s with args "+
		"(%d, %d)\n", opName, a, b)
	return op(a, b)
}
func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

// 13 / 3 =  4.... 1
func div(a, b int) (int, int) {
	return a / b, a % b
}
func div1(a, b int) (q, r int) { //返回值起别名
	q = a / b
	r = a % b
	return
}

func sum(numbers ...int) int {
	sum := 0
	for i := range numbers {
		sum += numbers[i]
	}
	return sum
}

func swap(a, b *int) {
	*b, *a = *a, *b
}
func swap1(a, b int) (int, int) {
	return b, a
}

func main() {
	fmt.Println(eval(3, 4, "X"))
	fmt.Println(div(13, 3))
	q, r := div1(18, 4)
	fmt.Println(q, r)

	fmt.Println(apply(pow, 3, 2))
	fmt.Println(sum(1, 2, 3, 4, 5, 6, 7, 8, 9))

	a, b := 3, 4
	swap(&a, &b)
	fmt.Println(a, b)
	c, d := 5, 6
	c, d = swap1(c, d)
	fmt.Println(c, d)
}
