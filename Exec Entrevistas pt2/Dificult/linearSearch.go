package main

import "fmt"

func binarySearch(arr []int, alvo int) int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == alvo {
			return i
		}
	}
	return -1
}

func main() {
	arr := []int{4, 5, 6, 7, 0, 1, 2}
	alvo := 1
	fmt.Println(binarySearch(arr, alvo))
}
