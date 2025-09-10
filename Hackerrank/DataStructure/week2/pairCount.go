package main

import "fmt"

func sockMerchant(n int32, ar []int32) int32 {
	numCount := make(map[int32]int32)

	for _, v := range ar {
		numCount[v]++
	}
	var result int32
	for _, v := range numCount {
		result += v / 2
	}
	return result

}
func main() {
	n := int32(7)
	ar := []int32{1, 2, 1, 2, 1, 1, 3}
	fmt.Println(sockMerchant(n, ar))
}
