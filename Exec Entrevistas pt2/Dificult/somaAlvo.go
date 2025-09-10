package main

import "fmt"

func somaAlvo(arr []int, k int) []int {

	var result []int
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i]+arr[j] == k {
				result = append(result, arr[i], arr[j])
			}
		}
	}
	return result
}

func main() {
	arr := []int{1, 2, 3, 2}
	k := 4
	fmt.Println(somaAlvo(arr, k))
}
