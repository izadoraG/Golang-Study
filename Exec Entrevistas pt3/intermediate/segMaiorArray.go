package main

import (
	"fmt"
	"sort"
)

func segMaiorArray(arr []int) int {

	sort.Slice(arr, func(i, j int) bool {
		return arr[i] > arr[j]
	})

	return arr[1]
}

func main() {
	arr := []int{10, 20, 4, 45, 99}
	fmt.Println(segMaiorArray(arr))
}
