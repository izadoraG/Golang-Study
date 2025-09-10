package main

import (
	"fmt"
)

func sansaXor(arr []int32) int32 {

	result := int32(0)

	for i := 0; i < len(arr); i++ {
		count := (i + 1) * (len(arr) - i)

		if count%2 == 1 {
			result ^= arr[i]
		}
	}

	return result
}

func main() {
	arr := []int32{1, 2, 3}
	fmt.Println(sansaXor(arr))
}
