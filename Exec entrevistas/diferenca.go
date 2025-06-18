package main

import "fmt"

func maiorDiferenca(arr []int) int {

	maiorValor := arr[0]
	menorValor := arr[0]

	for _, a := range arr {
		if a > maiorValor {
			maiorValor = a
		} else if a < menorValor {
			menorValor = a
		}
	}

	diferenca := maiorValor - menorValor
	return diferenca

}

func main() {
	arr := []int{2, 3, 5, 6, 7, 4, 6}
	fmt.Println(maiorDiferenca(arr))
}
