package main

import "fmt"

func climbingLeaderboard(ranked []int32, player []int32) []int32 {

	// Gerar os ranks
	rank := make([]int32, len(ranked))
	rank[0] = 1

	for i := 1; i < len(ranked); i++ {
		if ranked[i] == ranked[i-1] {
			rank[i] = rank[i-1]
		} else {
			rank[i] = rank[i-1] + 1
		}
	}

	var res []int32
	index := len(ranked) - 1

	for _, p := range player {
		for index >= 0 && p >= ranked[index] {
			index--
		}
		if index == -1 {
			res = append(res, 1)
		} else {
			res = append(res, rank[index]+1)
		}
	}
	return res
}

func main() {
	ranked := []int32{100, 100, 50, 40, 40, 20, 10}
	player := []int32{5, 25, 50, 120}
	fmt.Println(climbingLeaderboard(ranked, player))
}
