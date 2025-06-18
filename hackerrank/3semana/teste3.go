package main

import (
	"fmt"
)

func getTotalX(a []int32, b []int32) int32 {

	var result int32

	maxA := a[0]
	for _, val := range a {
		if val > maxA {
			maxA = val
		}
	}

	minB := b[0]
	for _, val := range b {
		if val < minB {
			minB = val
		}
	}

	for num := maxA; num <= minB; num++ {
		valid := true

		for _, valA := range a {
			if num%valA != 0 {
				valid = false
				break
			}
		}
		if valid {
			for _, valB := range b {
				if valB%num != 0 {
					valid = false
					break

				}
			}

		}
		if valid {
			result++
		}
	}
	return result

}

func main() {
	a := []int32{2, 4}
	b := []int32{16, 32, 96}
	fmt.Println(getTotalX(a, b))
}
