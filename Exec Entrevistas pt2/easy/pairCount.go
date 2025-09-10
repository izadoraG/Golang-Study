package main

import "fmt"

func pairCount(a []int) int {
	count := 0
	for _, num := range a {
		if num%2 != 0 {
			count++
		}
	}
	return count
}

func main() {
	a := []int{1, 2, 3, 4, 5}
	fmt.Println(pairCount(a))
}
