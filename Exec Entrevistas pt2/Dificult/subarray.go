package main

import "fmt"

func subarraySum(nums []int, k int) int {
	count := 0
	sum := 0
	prefixSums := map[int]int{0: 1} // Começamos com 0:1 para lidar com somas desde o índice 0

	for _, num := range nums {
		sum += num // Atualiza a soma acumulada

		// Verifica se já vimos uma soma anterior que faz a subtração sum - k ser válida
		if freq, exists := prefixSums[sum-k]; exists {
			count += freq // Adiciona o número de vezes que essa soma apareceu
		}

		// Atualiza o mapa de prefix sums com a nova soma atual
		prefixSums[sum]++
	}

	return count
}

func main() {
	arr := []int{1, 1, 1, 3, 4, 1, 1}
	k := 2
	fmt.Println("Total de subarrays:", subarraySum(arr, k)) // Saída esperada: 2
}
