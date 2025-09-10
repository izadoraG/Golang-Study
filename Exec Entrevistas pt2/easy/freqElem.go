package main

import "fmt"

func freqElem(list []int) []int {
	frqcount := make(map[int]int)
	for _, v := range list {
		frqcount[v]++
	}
	for num, count := range frqcount {
		fmt.Println(num, count)
	}
	return nil
}

func main() {
	list := []int{1, 2, 3, 3, 4, 5, 5, 5}
	freqElem(list)
}
