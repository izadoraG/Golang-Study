package main

import "fmt"

func gamingArray(arr []int32) string {
	var count int32 = 0
	var currentMax int32 = 0

	for _, num := range arr {
		if num > currentMax {
			currentMax = num
			count++
		}
	}

	if count%2 == 0 {
		return "ANDY"
	}
	return "BOB"
}

func main() {
	arr := []int32{3, 2, 1}
	fmt.Println(gamingArray(arr))
}
