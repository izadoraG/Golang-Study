package main

import (
	"fmt"
	"sort"
)

func isValid(s string) string {

	noRep := make(map[rune]int)

	for _, char := range s {
		noRep[char]++
	}

	var frequency []int
	for _, count := range noRep {
		frequency = append(frequency, count)
	}
	sort.Ints(frequency)

	if len(frequency) == 1 {
		return "YES"
	}

	first := frequency[0]
	second := frequency[1]
	last := frequency[len(frequency)-1]
	secondLast := frequency[len(frequency)-2]

	if first == last {
		return "YES"
	} else if first == 1 && last == second {
		return "YES"
	} else if first == second && second == secondLast && secondLast == last-1 {
		return "YES"
	}
	return "NO"

}

func main() {
	s := "abccc"
	fmt.Println(isValid(s))
}
