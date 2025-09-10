package main

import "fmt"

func rotacionarArray(arr []int, target int) []int {

	var result []int
	for i := 0; i < len(arr); i++ {
		result = append(arr[target:], arr[:target]...)
	}
	return result

}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	target := 3
	fmt.Println(rotacionarArray(arr, target))
}
