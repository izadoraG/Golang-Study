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
			return "NO"
		}
	}
	return "YES"

}

func main() {
	k := int32(5)

	A := []int32{1, 2, 2, 1}

	B := []int32{3, 3, 3, 4}

	fmt.Println(twoArrays(k, A, B))

}
