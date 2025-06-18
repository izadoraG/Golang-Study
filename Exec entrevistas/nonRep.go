package main

import "fmt"

func Repeticao(str string) string {
	nonRepChar := make(map[rune]int)

	for _, char := range str {
		nonRepChar[char]++
	}

	for _, char := range str {
		if nonRepChar[char] == 1 {
			return string(char)
		}
	}

	return ""
}

func main() {
	str := "swiss"
	fmt.Println(Repeticao(str))
}
