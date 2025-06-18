package main

import (
	"fmt"
	"sort"
)

func maxDif(arr []int32) int32 {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	dif := arr[len(arr)+1] - arr[0]
	return dif
}

func main() {
	arr := []int32{1, 2, 3, 4, 5}
	fmt.Println(maxDif(arr))
}
