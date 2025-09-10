package main

import (
	"fmt"
	"sort"
)

func miniMaxSum(arr []int32) {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	sum := 0
	sum2 := 0
	for i := 1; i < len(arr); i++ {
		sum += int(arr[i])
	}
	for i := 0; i < len(arr)-1; i++ {
		sum2 += int(arr[i])
	}
	fmt.Println(sum2, sum)

}

func main() {
	arr := []int32{1, 2, 3, 4, 5, 6}
	miniMaxSum(arr)
}
