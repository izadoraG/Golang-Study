package main

import "fmt"

func Anagramas(str1, str2 string) bool {

	if len(str1) != len(str2) {
		return false
	}

	charCount := make(map[rune]int)
	for _, char := range str1 {
		charCount[char]++
	}

	for _, char := range str2 {
		charCount[char]--
		if charCount[char] < 0 {
			return false
		}
	}
	return true

}

func main() {
	str1 := "listen"
	str2 := "silent"
	fmt.Println(Anagramas(str1, str2))
}
