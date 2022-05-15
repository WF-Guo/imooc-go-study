package main

import "fmt"

/*func lengthOfNonRepeatingSubstr(s string) int {
	lastOccurred := make(map[byte]int)
	maxLength := 0
	start := 0
	for i, ch := range []byte(s) {
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength
}*/

func lengthOfNonRepeatingSubstr(s string) int {
	lastOccurred := make(map[rune]int)
	maxLength := 0
	start := 0
	for i, ch := range []rune(s) {
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength
}

func main() {
	fmt.Println(lengthOfNonRepeatingSubstr("asssbbbs"))
	fmt.Println(lengthOfNonRepeatingSubstr("abcdefg"))
	fmt.Println(lengthOfNonRepeatingSubstr("aaaa"))
	fmt.Println(lengthOfNonRepeatingSubstr("我爱慕课网"))
	fmt.Println(lengthOfNonRepeatingSubstr("我爱爱爱爱你"))
}
