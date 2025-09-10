package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Ler input
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		processInput(input)
	}
}

func processInput(input string) {
	// Separar a linha pelo ponto-e-vírgula
	parts := strings.Split(input, ";")
	action := parts[0]
	kind := parts[1]
	text := parts[2]

	if action == "S" {
		// Split
		if kind == "M" {
			text = strings.TrimSuffix(text, "()")
		}
		// Quebrar camelCase
		for i, c := range text {
			if i != 0 && c >= 'A' && c <= 'Z' {
				fmt.Print(" ")
			}
			fmt.Print(strings.ToLower(string(c)))
		}
		fmt.Println()
	} else if action == "C" {
		// Combine
		words := strings.Split(text, " ")
		for i, word := range words {
			if kind == "V" && i == 0 {
				// Variável: primeira palavra minúscula
				fmt.Print(strings.ToLower(word))
			} else {
				// As outras palavras: primeira letra maiúscula
				fmt.Print(strings.Title(word))
			}
		}
		if kind == "M" {
			fmt.Print("()")
		}
		fmt.Println()
	}
}
