package main

import "fmt"

func increasingTriplet(nums []int) bool {
	first, second := 1000, 1000 // ambos começam com o maior int possível

	for _, n := range nums {
		if n <= first {
			first = n
		} else if n <= second {
			second = n
		} else {
			return true
		}
	}
	return false
}

func main() {
	nums := []int{20, 100, 10, 12, 5, 13}
	fmt.Println(increasingTriplet(nums))
}
