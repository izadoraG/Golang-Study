package main

import (
	"fmt"
)

func countingSort(arr []int32) []int32 {

	counting := make(map[int32]int)

	var result []int32
	for _, v := range arr {
		counting[v]++
	}

	for _, count := range counting {
		result = append(result, int32(count))
	}
	return result

}

func main() {
	arr := []int32{1, 2, 1, 3, 3}
	fmt.Println(countingSort(arr))
}
