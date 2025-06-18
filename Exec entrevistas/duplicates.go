package main

import "fmt"

func removeDuplicates(arr []int) []int {

	var result []int
	result = append(result, arr[0])

	for i := 1; i < len(arr); i++ {
		if arr[i] != arr[i-1] {
			result = append(result, arr[i])
		}
	}
	return result

}

func main() {
	arr := []int{1, 1, 2, 2, 3, 4, 5, 5}
	fmt.Println(removeDuplicates(arr))
}
