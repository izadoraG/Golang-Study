package main

import "fmt"

func sum(list []int) int {
	sum := 0
	for _, value := range list {
		sum += value
	}
	return sum
}

func main() {
	list := []int{1, 2, 3, 4, 5}
	fmt.Println(sum(list))
}
