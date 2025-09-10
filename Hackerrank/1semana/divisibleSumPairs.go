package main

import "fmt"

func divisibleSumPairs(n int32, k int32, ar []int32) int32 {

	var result int32 = 0

	for i := int32(0); i < n; i++ {
		for j := i; j < n; j++ {
			if i < j && (ar[i]+ar[j])%k == 0 {
				result++
			}
		}

	}
	return result

}

func main() {
	ar := []int32{1, 3, 2, 6, 1, 2}
	var k int32 = 3
	var n int32 = 6
	fmt.Println(divisibleSumPairs(n, k, ar))
}
