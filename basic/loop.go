package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

//十进制转二进制：除以2反序取余
func convertToBinary(n int) string {
	if n == 0 {
		return strconv.Itoa(0)
	}
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2                        //取余
		result = strconv.Itoa(lsb) + result //反序
	}
	return result
}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	printFileContents(file)

}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() { //相当于while
		fmt.Println(scanner.Text())
	}
}

func main() {
	fmt.Println(
		convertToBinary(5),  //101
		convertToBinary(13), //1101
		convertToBinary(0),
	)

	printFile("basic/abc.txt")
	s := `abc"d"
			fsdkla

fsadf
fdasdf
gggg
`
	printFileContents(strings.NewReader(s))

}
