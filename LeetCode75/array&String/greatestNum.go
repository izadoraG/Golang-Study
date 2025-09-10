package main

import "fmt"

func kidsWithCandies(candies []int, extraCandies int) []bool {
	greatestNum := 0
	var result []bool

	for i := 0; i < len(candies); i++ {
		if candies[i] > greatestNum {
			greatestNum = candies[i]
		}
	}

	for i := 0; i < len(candies); i++ {
		if candies[i]+extraCandies >= greatestNum {
			result = append(result, true)
		} else if candies[i]+extraCandies < greatestNum {
			result = append(result, false)
		}
	}
	return result
}

func main() {
	candies := []int{2, 3, 5, 1, 3}
	extraCandies := 3
	fmt.Println(kidsWithCandies(candies, extraCandies))
}
