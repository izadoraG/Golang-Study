package main

import "fmt"

func matchingStrings(strings []string, queries []string) []int32 {

	charCount := make(map[string]int)

	for _, palavra := range strings {
		charCount[palavra]++

	}

	var result []int32
	for _, querie := range queries {
		result = append(result, int32(charCount[querie]))
	}

	return result
}

func main() {
	strings := []string{"aba", "baba", "aba", "xzxb"}
	queries := []string{"aba", "xzxb", "ab"}
	fmt.Println(matchingStrings(strings, queries))
}
