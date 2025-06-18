package main

import "fmt"

func anagram(s string) int32 {
	runeSlice := []rune(s)
	mid := len(runeSlice) / 2
	firstString := string(runeSlice[:mid])
	lastString := string(runeSlice[mid:])

	count := int32(0)

	if len(runeSlice)%2 != 0 {
		return -1

	}

	charCount := make(map[rune]int)
	for _, char := range firstString {
		charCount[char]++
	}

	for _, char := range lastString {
		charCount[char]--
		if charCount[char] < 0 {
			count++
		}
	}
	return count

}

func main() {
	s := "abc"
	fmt.Println(anagram(s))
}
