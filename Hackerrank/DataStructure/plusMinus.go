package main

import "fmt"

func plusMinus(arr []int32) {

	positive := 0.0
	negative := 0.0
	zero := 0.0
	resPositive := 0.0
	resNegative := 0.0
	resZero := 0.0
	for _, item := range arr {
		if item > 0 {
			positive++
		} else if item == 0 {
			zero++
		} else if item < 0 {
			negative++
		}
	}
	resPositive = positive / float64(len(arr))
	resNegative = negative / float64(len(arr))
	resZero = zero / float64(len(arr))
	fmt.Printf("%.6f\n%.6f\n%.6f\n", resPositive, resNegative, resZero)

}

func main() {
	arr := []int32{-1, 2, 3, 0}
	plusMinus(arr)
}
