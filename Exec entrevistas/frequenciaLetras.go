package main

import "fmt"

func Frequencia(str []string) map[rune]int {

	charCount := make(map[rune]int)

	for _, palavra := range str {
		for _, char := range palavra {
			charCount[char]++
		}
	}
	return charCount

}

func main() {
	str := []string{"banana", "amora"}
	result := Frequencia(str)

	for char, qtd := range result {
		fmt.Printf("%c: %d\n", char, qtd)
	}

}
