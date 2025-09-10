package main

import "fmt"

func productExceptSelf(nums []int) []int {
	var result []int

	for i := 0; i < len(nums); i++ {
		product := 1
		for j := 0; j < len(nums); j++ {
			if j != i {
				product *= nums[j]
			}
		}
		result = append(result, product)
	}
	return result
}

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(productExceptSelf(nums))
}
