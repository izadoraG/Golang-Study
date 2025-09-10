package main

import "fmt"

func matchingStrings(strings []string, queries []string) []int32 {

	var result []int32
	charCount := make(map[string]int)
	for _, str := range strings {
		charCount[str]++
	}
	for _, querie := range queries {
		result = append(result, int32(charCount[querie]))
	}
	return result

}

func main() {
	strings := []string{"ab", "ab", "abc"}
	queries := []string{"ab", "abc", "cd"}
	fmt.Println(matchingStrings(strings, queries))
}
