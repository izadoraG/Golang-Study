package main

import "fmt"

func sumSubarray(arr []int, target int) int {

	soma := 0
	start := 0
	for i := 0; i < len(arr); i++ {
		soma += arr[i]

		for soma > target && start <= i {
			soma -= arr[start]
			start++
		}

		if soma == target {
			return start
		}
	}
	return 0
}

func main() {
	arr := []int{1, 9, 3, 7, 5}
	target := 12
	fmt.Println(sumSubarray(arr, target))
}
