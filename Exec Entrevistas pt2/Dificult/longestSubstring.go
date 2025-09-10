package main

import "fmt"

func longestSubstring(s string) string {

	charCount := make(map[rune]bool)
	seen := make(map[rune]bool)

	for _, v := range s {
		charCount[v] = true
	}
	var newString string
	for _, char := range s {
		if charCount[char] && !seen[char] {
			newString += string(char)
			seen[char] = true

		}
	}
	return newString

}

func main() {
	s := "abcabcbb"
	fmt.Println(longestSubstring(s))
}
