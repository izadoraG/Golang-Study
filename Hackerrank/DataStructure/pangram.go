package main

import (
	"fmt"
	"strings"
)

func pangram(s string) string {

	alphabet := make(map[rune]bool)
	trimmed := strings.TrimSpace(s)

	for _, char := range trimmed {
		if char >= 'a' && char <= 'z' {
			alphabet[char] = true
		}
	}

	if len(alphabet) != 26 {
		return "not pangram"
	}

	return "pangram"

}

func main() {
	s := "We promptly judged antique ivory buckles for the prize"
	fmt.Println(pangram(s))
}
