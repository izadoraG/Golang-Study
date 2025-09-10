package main

import (
	"fmt"
	"sort"
)

func maximumPerimeterTriangle(sticks []int32) []int32 {

	sort.Slice(sticks, func(i, j int) bool {
		return sticks[i] < sticks[j]
	})

	for i := len(sticks) - 1; i >= 2; i-- {
		a, b, c := sticks[i-2], sticks[i-1], sticks[i]
		if a+b > c {
			return []int32{a, b, c}
		}
	}

	return []int32{-1}
}

func main() {
	sticks := []int32{1, 2, 3, 4, 5, 10}
	fmt.Println(maximumPerimeterTriangle(sticks))
}
