package main

import (
	"fmt"
	"math"
	"sort"
)

func closestNumbers(arr []int32) []int32 {

	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	menorDif := int32(math.MaxInt32)
	var result []int32

	for i := 0; i < len(arr)-1; i++ {
		diferenca := arr[i+1] - arr[i]
		if diferenca < menorDif {
			menorDif = diferenca
			result = []int32{arr[i], arr[i+1]}
		} else if diferenca == menorDif {
			result = append(result, arr[i], arr[i+1])
		}
	}
	return result

}

func main() {
	arr := []int32{4, 2, 1, 3}
	fmt.Println(closestNumbers(arr))
}
