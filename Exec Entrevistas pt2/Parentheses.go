package main

import "fmt"

func isValid(s string) bool {
	pairs := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}

	stack := []rune{}

	for _, c := range s {
		if closing, ok := pairs[c]; ok {
			// É um caractere de abertura → empilha o fechamento esperado
			stack = append(stack, closing)
		} else {
			// É um fechamento → verifica com o topo da pilha
			if len(stack) == 0 || stack[len(stack)-1] != c {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}

func main() {
	s := "()[]{"
	fmt.Println(isValid(s))
}
