package main

import (
	"fmt"
	"sort"
)

func missingNumbers(arr []int32, brr []int32) []int32 {

	countarr := make(map[int32]int)
	countbrr := make(map[int32]int)

	for _, num := range arr {
		countarr[num]++
	}

	for _, num := range brr {
		countbrr[num]++
	}

	var result []int32

	seen := make(map[int32]bool)
	for num, count := range countbrr {
		if countarr[num] < count && !seen[num] {
			result = append(result, num)
			seen[num] = true
		}
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i] < result[j]
	})
	return result

}

func main() {
	arr := []int32{201, 202, 203, 204}
	brr := []int32{201, 202, 203, 204, 205}
	fmt.Println(missingNumbers(arr, brr))
}
