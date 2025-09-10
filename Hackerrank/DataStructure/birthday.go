package main

import "fmt"

func birthday(arr []int32, d int32, m int32) int32 {
	count := int32(0)
	for i := 0; int(i) <= len(arr)-int(m); i++ {
		sum := int32(0)
		for j := 0; j < int(m); j++ {
			sum += arr[i+j]
		}
		if sum == d {
			count++
		}
	}
	return count
}

func main() {
	arr := []int32{2, 2, 1, 3, 2}
	d := int32(4)
	m := int32(2)
	fmt.Println(birthday(arr, d, m))
}
