package main

import "fmt"

func twoSum(arr []int) []int {

	target := 9

	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i]+arr[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil // nenhum par encontrado
}

func main() {
	arr := []int{2, 7, 11, 15}
	fmt.Println(twoSum(arr)) // Saída esperada: [0 1]
}
