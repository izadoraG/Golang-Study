package main

import "fmt"

func birthday(s []int32, d int32, m int32) int32 {
	var result int32

	for i := 0; i <= len(s)-int(m); i++ {
		var sum int32
		for j := i; j < i+int(m); j++ {
			sum += s[j]
		}
		if sum == d {
			result++
		}
	}

	return result
}

func main() {
	s := []int32{2, 2, 1, 3, 2, 2}
	d := int32(4)
	m := int32(2)
	fmt.Println(birthday(s, d, m))
}
