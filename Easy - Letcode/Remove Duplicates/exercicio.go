package main

import "fmt"

func removeDuplicates(nums []int) int {

	index := 1 // Posição para gravar o próximo valor único

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[index] = nums[i]
			index++
		}
	}

	return index
}

func main() {
	nums := []int{1, 1, 2, 2, 3}
	fmt.Println(removeDuplicates(nums))
}
