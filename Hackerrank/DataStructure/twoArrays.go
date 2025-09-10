package main

import (
	"fmt"
	"sort"
)

func twoArrays(k int32, A []int32, B []int32) string {
	sort.Slice(A, func(i, j int) bool {
		return A[i] < A[j]
	})
	sort.Slice(B, func(i, j int) bool {
		return B[i] > B[j]
	})

	for i := 0; i < len(A); i++ {
		if A[i]+B[i] < k {
			return "No"
		}
	}
	return "Yes"

}

func main() {
	k := int32(10)
	A := []int32{2, 1, 3}
	B := []int32{7, 8, 9}
	fmt.Println(twoArrays(k, A, B))
}
