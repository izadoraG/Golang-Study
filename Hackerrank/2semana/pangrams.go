package main

import (
	"fmt"
	"strings"
)

func pangrams(s string) string {

	letrasEncontradas := make(map[rune]bool)
	fraseMin := strings.ToLower(s)

	for _, letra := range fraseMin {
		if letra >= 'a' && letra <= 'z' {
			letrasEncontradas[letra] = true
		}
	}

	if len(letrasEncontradas) != 26 {
		return fmt.Sprintf("not pangram")
	}

	return fmt.Sprintf("pangram")
}

func main() {
	fraseTeste := "We promptly judged antique ivory buckles for the prize"
	fmt.Println(pangrams(fraseTeste))
}
