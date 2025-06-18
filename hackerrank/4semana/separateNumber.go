package main

import (
	"fmt"
	"strconv"
)

func separateNumbers(s string) {

	for i := 0; i <= len(s)/2; i++ {
		firstStr := s[:i]
		firstNum, err := strconv.Atoi(firstStr)
		if err != nil {
			continue
		}
		sequence := firstStr
		nextNum := firstNum

		for len(sequence) < len(s) {
			nextNum++
			sequence += strconv.Itoa(nextNum)
		}

		if sequence == s {
			fmt.Println("YES", firstStr)
			return
		}

	}
	fmt.Println("NO")

}

func main() {
	s := "99100"
	separateNumbers(s)
}
