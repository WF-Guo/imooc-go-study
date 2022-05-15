package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var aa = 3     //作用于为包内，go没有全局变量的概念
var ss = "kkk" //函数外面也可以定义变量，但是必须使用var，不可以使用冒号等号

var ( //批量定义方式
	aaa = 3
	bbb = 4
	ccc = true
)

func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s) //%q可以将字符串的引号打出来
}

func variableInitValue() {
	var a, b int = 3, 4
	var s string = "abc"

	fmt.Println(a, b, s) //定义的变量一定要用到，不然会报错
}

func variableTypeDeduction() { //自动推断变量类型
	var a, b, c, s = 3, 4, true, "def"
	fmt.Println(a, b, c, s)
}

func variableShorter() {
	a, b, c, s := 3, 4, true, "def" //推荐定义方式
	b = 5                           //再次赋值不可以用冒号
	fmt.Println(a, b, c, s)
}

func euler() {
	c := 3 + 4i
	fmt.Println(cmplx.Abs(c))

	fmt.Println(cmplx.Pow(math.E, 1i*math.Pi) + 1)

	fmt.Printf("%.3f\n", cmplx.Exp(1i*math.Pi)+1)
}

func triangle() {
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

func consts() {
	const filename = "abc.txt"
	const a, b = 3, 4
	var c int
	c = int(math.Sqrt(a*a + b*b))

	fmt.Println(filename, c)
}

func enums() {
	const (
		cpp = iota
		_
		goland
		python
		javascript
	)

	//b, kb, mb, gb, tb, pb
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)

	fmt.Println(cpp, javascript, goland, python) //0 1 2 3
	fmt.Println(b, kb, mb, gb, tb, pb)
}
func main() {
	fmt.Println("hello  world")
	variableZeroValue()
	variableInitValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(aaa, bbb, ccc)

	euler()
	triangle()
	consts()
	enums()
}
