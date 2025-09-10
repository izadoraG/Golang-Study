package main

import "fmt"

func soma(num []int) int {

	result := 0

	for _, v := range num {
		result += v
	}
	return result
}

func main() {
	num := []int{1, 2, 3, 4, 5}
	fmt.Println(soma(num))
}
