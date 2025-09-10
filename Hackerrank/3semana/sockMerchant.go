package main

import (
	"fmt"
)

func sockMerchant(n int32, ar []int32) int32 {
	count := make(map[int32]int32)
	for _, sock := range ar {
		count[sock]++
	}

	var result int32
	for _, v := range count {
		result += v / 2
	}
	return result
}

func main() {
	n := int32(10)
	ar := []int32{1, 1, 2, 2, 3}
	fmt.Println(sockMerchant(n, ar))
}
