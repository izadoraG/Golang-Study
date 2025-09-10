package main

import "fmt"

func binarySearch(x []int, target int) int {

	for i := 0; i < len(x); i++ {
		if x[i] == target {
			return i
		}
	}
	return -1
}

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	target := 7
	fmt.Println(binarySearch(x, target))
}
