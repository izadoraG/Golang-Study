package main

import (
	"fmt"
	"strings"
)

func mergeAlternatly(a string, b string) string {

	i := 0
	var result []string

	for i < len(a) || i < len(b) {
		if i < len(a) {
			result = append(result, string(a[i]))
		}
		if i < len(b) {
			result = append(result, string(b[i]))
		}
		i++
	}
	return strings.Join(result, "")

}

func main() {
	a := "abc"
	b := "def"
	fmt.Println(mergeAlternatly(a, b))
}
