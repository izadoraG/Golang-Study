package main

import (
	"fmt"
	"math"
)

func diagonalDifference(arr [][]int32) int32 {
	size := len(arr)
	sumPrimary := int32(0)
	sumSecondary := int32(0)

	for i := 0; i < size; i++ {
		sumPrimary += arr[i][i]
		sumSecondary += arr[i][size-1-i]
	}

	diference := int32(math.Abs(float64(sumPrimary - sumSecondary)))
	return diference
}

func main() {
	matrix := [][]int32{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println(diagonalDifference(matrix))
}
