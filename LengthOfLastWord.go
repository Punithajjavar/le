package main

import "fmt"

//type Solution struct{}

func  lengthOfLastWord(s string) int {
	count := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			count++
		} else {
			if count > 0 {
				return count
			}
		}
	}
	return count
}

func main() {
	//solution := Solution{}
	input := "Hello World "
	// 5
	length := lengthOfLastWord(input)
	fmt.Println("Length of the last word: ", length)
}