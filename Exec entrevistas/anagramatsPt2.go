package main

import (
	"fmt"
	"sort"
)

func Anagramas2(words []string) map[string][]string {
	anag := make(map[string][]string)

	for _, word := range words {
		// Converte a palavra para slice de runes
		chars := []rune(word)
		// Ordena os caracteres
		sort.Slice(chars, func(i, j int) bool {
			return chars[i] < chars[j]
		})
		// Converte de volta para string
		sortedWord := string(chars)

		// Adiciona a palavra original no grupo do anagrama
		anag[sortedWord] = append(anag[sortedWord], word)
	}

	return anag
}

func main() {
	words := []string{"listen", "silent", "enlist", "hello", "below", "elbow"}
	resultado := Anagramas2(words)

	// Imprime os grupos de anagramas
	for chave, grupo := range resultado {
		if len(grupo) > 1 {
			fmt.Printf("Chave: %s -> Anagramas: %v\n", chave, grupo)
		}
	}
}
