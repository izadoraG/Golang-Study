package main

import "fmt"

func marsExploration(s string) int32 {
	numPalavras := len(s) / 3

	expected := ""
	for i := 0; i < numPalavras; i++ {
		expected += "SOS"
	}

	var count int32
	for i := range s {
		if s[i] != expected[i] {
			count++
		}
	}
	return count
}

func main() {
	s := "SOSSOT"
	fmt.Println(marsExploration(s))
}
