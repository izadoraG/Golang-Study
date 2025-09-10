package main

import "fmt"

func lonelyinteger(a []int32) int32 {

	numCount := make(map[int32]int32)
	for _, v := range a {
		numCount[v]++
	}

	for num, count := range numCount {
		if count <= 1 {
			return num
		}
	}
	return 0
}

func main() {
	a := []int32{1, 9, 3, 9, 1}
	fmt.Println(lonelyinteger(a))
}
