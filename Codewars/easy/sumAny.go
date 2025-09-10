package main

import (
	"fmt"
	"strconv"
)

func SumMix(arr []any) int {

	count := 0
	for i := 0; i < len(arr); i++ {
		num, err := strconv.ParseInt(fmt.Sprint(arr[i]), 10, 32)
		if err != nil {
			fmt.Println(err)
		}
		count += int(num)
	}
	return count
}

func main() {
	arr := []any{1, "7", 2, "3", 4, "9", 5}
	fmt.Println(SumMix(arr))
}
