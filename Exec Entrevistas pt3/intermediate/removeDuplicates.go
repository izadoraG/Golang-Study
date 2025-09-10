package main

import "fmt"

func removeDuplicates(arr []int) []int {
	count := make(map[int]int)
	var result []int

	for _, v := range arr {
		count[v]++
	}
	for x, _ := range count {
		result = append(result, x)
	}
	return result

}

func main() {
	arr := []int{1, 2, 2, 3, 4, 4, 5}
	fmt.Println(removeDuplicates(arr))
}
