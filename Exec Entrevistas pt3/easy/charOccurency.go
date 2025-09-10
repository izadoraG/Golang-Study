package main

import "fmt"

func charOcurrency(s string) map[rune]int {

	charCount := make(map[rune]int)

	for _, char := range s {
		charCount[char]++
	}

	for char, count := range charCount {
		fmt.Println(string(char), ":", count, "")
	}
	return charCount
}

func main() {
	s := "success"
	charOcurrency(s)
}
