package main

import "fmt"

func lonelyInteger(nums []int32) int32 {
	result := int32(0)
	for _, num := range nums {
		result ^= num
	}
	return result
}

func main() {
	arr := []int32{1, 2, 3, 4, 3, 2, 1}
	lonely := lonelyInteger(arr)
	fmt.Println("The lonely integer is:", lonely) // Output: 4
}
