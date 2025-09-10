package main

import (
	"fmt"
	"sort"
)

func mergeSortedLists(list1 []int, list2 []int) []int {

	var result []int

	for _, num := range list1 {
		result = append(result, num)
	}

	for _, num := range list2 {
		result = append(result, num)
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i] < result[j]
	})

	return result
}

func main() {
	list1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	list2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(mergeSortedLists(list1, list2))
}
