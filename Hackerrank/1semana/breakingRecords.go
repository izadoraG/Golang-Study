package main

import "fmt"

func breakingRecords(scores []int32) []int32 {
	var maxCount int32
	var minCount int32
	max := scores[0]
	min := scores[0]

	for _, score := range scores {
		if score > max {
			max = score
			maxCount++
		}
		if score < min {
			min = score
			minCount++
		}

	}
	return []int32{maxCount, minCount}

}

func main() {
	scores := []int32{10, 5, 20, 20, 4, 5, 2, 25, 1}
	fmt.Println(breakingRecords(scores))
}
