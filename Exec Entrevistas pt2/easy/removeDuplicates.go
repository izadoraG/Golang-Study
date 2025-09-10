package main

import (
	"fmt"
	"slices"
)

func removeDuplicates(strList []int) []int {
	var list []int
	for _, item := range strList {
		if slices.Contains(list, item) == false {
			list = append(list, item)
		}
	}
	return list
}

func main() {
	strList := []int{1, 2, 3, 4, 5, 1, 2, 3}
	fmt.Println(removeDuplicates(strList))
}
