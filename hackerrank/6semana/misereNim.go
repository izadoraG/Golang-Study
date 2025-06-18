package main

import "fmt"

func misereNim(s []int32) string {

	allOnes := true
	for _, num := range s {
		if num != 1 {
			allOnes = false
			break
		}
	}
	if allOnes {
		if len(s)%2 != 0 {
			return "Second"
		} else {
			return "First"
		}
	}

	result := int32(0)
	for _, num := range s {
		result ^= num
	}

	if !allOnes {
		if result == 0 {
			return "Second"
		} else {
			return "First"
		}
	}
	return ""

}

func main() {
	s := []int32{1, 1, 1}
	fmt.Println(misereNim(s))
}
