package main

func lengthOfLongestSubstring(s string) int {
	lastSeen := make(map[rune]int)
	maxLen := 0
	start := 0

	for i, char := range s {
		if lastI, found := lastSeen[char]; found && lastI >= start {
			// caractere repetido dentro da janela atual
			start = lastI + 1
		}

		lastSeen[char] = i
		maxLen = max(maxLen, i-start+1)
	}

	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
