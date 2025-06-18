package main

import (
	"fmt"
	"math"
	"sort"
)

func maxMin(k int32, arr []int32) int32 {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	minDif := int32(math.MaxInt32)

	for i := 0; i <= len(arr)-int(k); i++ {
		dif := arr[i+int(k)-1] - arr[i]
		if dif < minDif {
			minDif = dif
		}
	}
	return minDif

}

func main() {
	arr := []int32{10, 1, 2, 3, 4, 5, 7, 8}
	k := int32(4)
	fmt.Println(maxMin(k, arr))
}
