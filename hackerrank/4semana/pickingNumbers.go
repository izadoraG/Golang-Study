package main

import (
	"fmt"
)

func pickingNumbers(a []int32) int32 {
	freq := make(map[int32]int32)
	var maxCount int32

	for _, num := range a {
		freq[num]++
	}

	for num, count := range freq {
		sum := count + freq[num+1]
		if sum > maxCount {
			maxCount = sum
		}
	}

	return maxCount
}

func main() {
	a := []int32{1, 2, 2, 2, 3, 3, 5, 5, 5}
	fmt.Println(pickingNumbers(a))
}
