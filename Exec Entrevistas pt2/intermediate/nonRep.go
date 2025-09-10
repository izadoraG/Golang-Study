package main

import "fmt"

func nonRep(arr string) string {

	countingChar := make(map[string]int)

	for _, v := range arr {
		countingChar[string(v)]++
	}

	for k, v := range countingChar {
		if v <= 1 {
			return k
		}
	}
	return ""
}

func main() {
	arr := "swiss"
	fmt.Println(nonRep(arr))
}
