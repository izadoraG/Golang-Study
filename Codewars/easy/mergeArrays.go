package main

import (
	"fmt"
	"sort"
)

func MergeArrays(arr1, arr2 []int) []int {
	if len(arr1) < 1 || len(arr2) < 1 {
		return []int{}
	}
	unique := make(map[int]bool)

	for _, num := range arr1 {
		unique[num] = true
	}
	for _, num := range arr2 {
		unique[num] = true
	}

	var result []int
	for num := range unique {
		result = append(result, num)
	}

	sort.Ints(result)
	return result
}

func main() {
	arr1 := []int{1, 2, 3, 4, 5}
	arr2 := []int{6, 7, 6, 9, 2}
	fmt.Println(MergeArrays(arr1, arr2))
}
