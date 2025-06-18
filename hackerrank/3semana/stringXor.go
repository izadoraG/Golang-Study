package main

import (
	"fmt"
)

func xorStrings(a, b string) string {
	n := len(a)
	result := ""

	for i := 0; i < n; i++ {
		if a[i] == b[i] {
			result += "0"
		} else {
			result += "1"
		}
	}

	return result
}

func main() {
	a := "10101"
	b := "00101"

	xor := xorStrings(a, b)
	fmt.Println(xor) // Saída: 10000
}
