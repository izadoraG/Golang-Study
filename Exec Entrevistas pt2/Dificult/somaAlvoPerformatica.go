package main

import "fmt"

func somaAlvoPerformatica(arr []int, k int) [][]int {
	seen := make(map[int]bool)
	var result [][]int

	for _, num := range arr {
		complement := k - num
		if seen[complement] {
			result = append(result, []int{complement, num})
		}
		seen[num] = true
	}
	return result
}

func main() {
	arr := []int{1, 2, 3, 2}
	k := 4
	fmt.Println(somaAlvoPerformatica(arr, k)) // Exemplo: [[2 2], [1 3]]
}
