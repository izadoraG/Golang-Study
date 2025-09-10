package main

import (
	"fmt"
	"sort"
)

func findMedian(arr []int32) int32 {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	meio := len(arr) / 2

	for i := 1; i < len(arr); i++ {
		if len(arr)%2 == 0 {
			mid1 := arr[meio]
			mid2 := arr[meio-1]
			return (mid1 + mid2) / 2
		}
	}
	return arr[meio]

}

func main() {
	arr := []int32{0, 1, 2, 3, 4, 5}
	fmt.Println(findMedian(arr))
}
