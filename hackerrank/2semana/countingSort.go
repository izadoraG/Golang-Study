package main

import "fmt"

func countingSort(arr []int32) []int32 {
	result := make([]int32, 100) // vetor de contagem com 100 posições

	for _, num := range arr {
		result[num]++ // incrementa o índice correspondente
	}

	return result
}

func main() {
	input := []int32{1, 1, 3, 2, 1}
	frequencias := countingSort(input)
	fmt.Println(frequencias)
}
