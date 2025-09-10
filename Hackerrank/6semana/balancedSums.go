package main

import "fmt"

func balancedSums(arr []int32) string {
	total := int32(0)
	for _, num := range arr {
		total += num
	}

	leftSum := int32(0)
	for _, num := range arr {
		if leftSum == total-leftSum-num {
			return "YES"
		}
		leftSum += num
	}

	return "NO"
}

func main() {
	arr := []int32{1, 2, 3}
	fmt.Println(balancedSums(arr))
}
