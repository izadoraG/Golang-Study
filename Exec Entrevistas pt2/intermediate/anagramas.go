package main

import "fmt"

func anagramas(str, str2 string) bool {
	charCount := make(map[rune]int)
	for _, char := range str {
		charCount[char]++
	}
	for _, char := range str2 {
		charCount[char]--
		if charCount[char] > 0 {
			return true
		}
	}
	return false

}

func main() {
	str := "silent"
	str2 := "listen"
	fmt.Println(anagramas(str, str2))
}
