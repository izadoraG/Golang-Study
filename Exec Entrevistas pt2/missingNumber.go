package main

import (
	"fmt"
	"sort"
)

func missingnumber(x []int) int {

	sort.Slice(x, func(i, j int) bool {
		return x[i] < x[j]
	})

	for i := 1; i < len(x); i++ {
		if x[i-1] != x[i]-1 {
			return x[i] - 1
		}
	}
	return 0
}

func main() {
	x := []int{3, 0, 2}
	fmt.Println(missingnumber(x))
}
