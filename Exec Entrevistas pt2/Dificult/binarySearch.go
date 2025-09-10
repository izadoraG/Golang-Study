package main

import "fmt"

func buscaRotacionado(nums []int, alvo int) int {
	inicio := 0
	fim := len(nums) - 1

	for inicio <= fim {
		meio := (inicio + fim) / 2

		if nums[meio] == alvo {
			return meio
		}

		// Verifica se o lado esquerdo está ordenado
		if nums[inicio] <= nums[meio] {
			if nums[inicio] <= alvo && alvo < nums[meio] {
				fim = meio - 1
			} else {
				inicio = meio + 1
			}
		} else {
			// Caso contrário, o lado direito está ordenado
			if nums[meio] < alvo && alvo <= nums[fim] {
				inicio = meio + 1
			} else {
				fim = meio - 1
			}
		}
	}

	return -1 // não encontrado
}

func main() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	alvo := 0
	fmt.Println(buscaRotacionado(nums, alvo)) // Saída: 4
}
