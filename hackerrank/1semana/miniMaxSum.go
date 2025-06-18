package main

import (
	"fmt"
	"sort"
)

func miniMaxSum(arr []int32) {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	var minSum int64 = 0
	var maxSum int64 = 0

	// Soma os 4 menores
	for i := 0; i < 4; i++ {
		minSum += int64(arr[i])
	}

	// Soma os 4 maiores
	for i := 1; i < 5; i++ {
		maxSum += int64(arr[i])
	}

	fmt.Printf("%d %d\n", minSum, maxSum)
}

func main() {
	arr := []int32{1, 2, 3, 4, 5}
	miniMaxSum(arr)
}
