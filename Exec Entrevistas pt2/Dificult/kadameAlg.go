package main

import "fmt"

func kadameAlg(arr []int) int {

	current_max := 0
	max_so_far := 0

	for i := 0; i < len(arr); i++ {
		max_so_far += arr[i]
		if max_so_far < 0 {
			max_so_far = 0
		} else if max_so_far > current_max {
			current_max = max_so_far
		}
	}
	return current_max
}

func main() {
	arr := []int{-2, 1, -3, 4, -1, 2, 2, -5, 4}
	fmt.Println(kadameAlg(arr))
}
