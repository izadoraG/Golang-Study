package main

import "fmt"

func boyerMoore(arr []int) int {
	metadeArr := len(arr) / 2
	numCount := make(map[int]int)

	for _, v := range arr {
		numCount[v]++
	}

	for index, num := range numCount {
		if num > metadeArr {
			return index
		}
	}
	return 0
}

func main() {
	arr := []int{2, 1, 1, 1, 2}
	fmt.Println(boyerMoore(arr))
}
