package main

import "fmt"

func productExceptSelf(nums []int) []int {
	ret := make([]int, len(nums)) // cria array de saída
	prefix := 1

	for i, n := range nums {
		ret[i] = prefix
		prefix *= n
	}

	suffix := 1
	for i := len(nums) - 1; i >= 0; i-- {
		ret[i] *= suffix
		suffix *= nums[i]
	}

	return ret
}

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(productExceptSelf(nums))
}
