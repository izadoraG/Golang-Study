package main

import "fmt"

func migratoryBirds(arr []int32) int32 {
	birdCount := make(map[int32]int32)

	for _, bird := range arr {
		birdCount[bird]++
	}

	var maxCount int32
	var birdType int32 = 6

	for id, count := range birdCount {
		if count > maxCount || (count == maxCount && id < birdType) {
			maxCount = count
			birdType = id
		}
	}

	return birdType
}

func main() {
	birds := []int32{1, 4, 4, 4, 5, 3, 3, 3}
	fmt.Println(migratoryBirds(birds))
}
