package main

import "fmt"

func Unique(input string) bool {

	charCount := make(map[rune]bool)

	for _, char := range input {
		if charCount[char] {
			return false
		}
		charCount[char] = true

	}
	return true
}

func main() {
	input := "aba"
	fmt.Println(Unique(input))
}
