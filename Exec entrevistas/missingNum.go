package main

import "fmt"

func encontrarNumeroFaltando(arr []int) int {
	n := len(arr) + 1 // porque 1 número está faltando

	somaEsperada := n * (n + 1) / 2
	somaReal := 0

	for _, num := range arr {
		somaReal += num
	}

	return somaEsperada - somaReal
}

func main() {
	arr := []int{1, 2, 4, 5, 6}
	fmt.Println("Número faltando:", encontrarNumeroFaltando(arr)) // Saída: 3
}
