package main

import "fmt"

func intersecaoArray(arr, arr2 []int) []int {
	charCount := make(map[int]bool)
	seen := make(map[int]bool)

	for _, v := range arr {
		charCount[v] = true
	}

	var result []int
	for _, char := range arr2 {
		if charCount[char] && !seen[char] {
			result = append(result, char)
			seen[char] = true // Marca que já adicionou esse número
		}
	}

	return result
}

func main() {
	arr := []int{1, 2, 3}
	arr2 := []int{1, 1}
	fmt.Println(intersecaoArray(arr, arr2))
}
