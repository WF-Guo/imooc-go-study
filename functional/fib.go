package main

import (
	"bufio"
	"fmt"
	"io"
	"mooc_Go_study/functional/Fib"
	"strings"
)

type intGen func() int

func (g intGen) Read(p []byte) (n int, err error) {
	next := g()
	if next > 10000 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)
	return strings.NewReader(s).Read(p)
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() { //相当于while
		fmt.Println(scanner.Text())
	}
}
func main() {
	var f intGen = Fib.Fibonacci()

	printFileContents(f)

}
