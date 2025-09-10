package main

import (
	"fmt"
	"sort"
	"strings"
)

func grupAnagram(arr []string) map[string][]string {
	anagramGroups := make(map[string][]string)

	for _, word := range arr {
		// Transforma a palavra em uma slice de caracteres
		letters := strings.Split(word, "")
		sort.Strings(letters)                   // Ordena as letras
		sortedWord := strings.Join(letters, "") // Junta de volta

		// Usa a palavra ordenada como chave
		anagramGroups[sortedWord] = append(anagramGroups[sortedWord], word)
	}

	return anagramGroups
}

func main() {
	arr := []string{"bat", "tab", "tap", "pat"}
	fmt.Println(grupAnagram(arr))
}
