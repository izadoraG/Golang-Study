package main

import "fmt"

func reverseOrder(list []int) []int {
	for i, j := 0, len(list)-1; i < j; i, j = i+1, j-1 {
		list[i], list[j] = list[j], list[i]
	}
	return list
}

func main() {
	a := []int{1, 2, 3, 4, 5}
	fmt.Println(reverseOrder(a))

}
