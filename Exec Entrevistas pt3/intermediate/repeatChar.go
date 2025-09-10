package main

import (
	"fmt"
)

func longestSubstring(s string) string {
	start := 0
	maxLen := 0
	subStart := 0
	seen := make(map[rune]int)

	for i, ch := range s {
		if prevIndex, found := seen[ch]; found && prevIndex >= start {
			start = prevIndex + 1
		}
		seen[ch] = i

		if i-start+1 > maxLen {
			maxLen = i - start + 1
			subStart = start
		}
	}

	return s[subStart : subStart+maxLen]
}

func main() {
	s := "abcabc"
	fmt.Println(longestSubstring(s))
}
