package main

import "fmt"

func elemFreq(list []int) int {
	frqcount := make(map[int]int)
	for _, v := range list {
		frqcount[v]++
	}
	maiorFreq := 0
	for _, count := range frqcount {
		if count > maiorFreq {
			maiorFreq = count
		}
	}
	return maiorFreq
}

func main() {
	list := []int{1, 2, 3, 3, 4, 5, 5, 5}
	fmt.Println(elemFreq(list))
}
