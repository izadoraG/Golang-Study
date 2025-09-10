package main

import "fmt"

func TwoSum(arr []int, target int) []int {

	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i]+arr[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	target := 9
	fmt.Println(TwoSum(arr, target))
}
