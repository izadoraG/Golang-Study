package main

import (
	"fmt"
	"math"
	"sort"
)

func minimumAbsoluteDifference(arr []int32) int32 {

	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	menorDif := int32(math.MaxInt32)

	for i := 0; i < len(arr)-1; i++ {
		diferenca := arr[i+1] - arr[i]
		diferencas := int32(math.Abs(float64(diferenca)))
		if diferencas < menorDif {
			menorDif = diferencas

		}

	}
	return menorDif
}

func main() {
	arr := []int32{3, -7, 0}
	fmt.Println(minimumAbsoluteDifference(arr))
}
