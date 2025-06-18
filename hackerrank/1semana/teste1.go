package main

import (
	"fmt"
	"sort"
)

func findMedian(arr []int32) int32 {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	n := len(arr)
	middle := n / 2

	if n%2 == 0 {

		return int32(arr[middle-1]+arr[middle]) / 2.0
	}
	return int32(arr[middle])
}

func main() {
	arr := []int32{0, 1, 2, 4, 6, 5, 3}
	fmt.Println(findMedian(arr))
}
