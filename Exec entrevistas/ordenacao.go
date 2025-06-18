package main

import "fmt"

func ordenacaoArray(arr *[]int) {

	n := len(*arr)

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if (*arr)[j] > (*arr)[j+1] {
				(*arr)[j], (*arr)[j+1] = (*arr)[j+1], (*arr)[j]

			}
		}
	}
}

func main() {
	input := []int{5, 2, 9, 1, 5, 6}
	ordenacaoArray(&input) // Passamos o ponteiro do slice
	fmt.Println(input)     // Saída esperada: [1 2 5 5 6 9]
}
