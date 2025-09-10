package main

import (
	"fmt"
	"strings"
)

func reverseVowels(s string) string {

	vowels := "aeiouAEIOU"
	str := []rune(s)
	left, right := 0, len(s)-1

	for left < right {
		for left < right && !strings.ContainsRune(vowels, str[left]) {
			left++
		}
		for left < right && !strings.ContainsRune(vowels, str[right]) {
			right--
		}
		if left < right {
			str[left], str[right] = str[right], str[left]
			left++
			right--
		}
	}
	return string(str)
}

func main() {
	s := "IceCreAm"
	fmt.Println(reverseVowels(s))
}
