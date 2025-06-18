package main

import "fmt"

func plusMinus(arr []int32) {
	positivos := 0
	negativos := 0
	zeros := 0

	for _, n := range arr {
		if n > 0 {
			positivos++
		} else if n == 0 {
			zeros++
		} else if n < 0 {
			negativos++
		}
	}
	total := float64(len(arr))
	fmt.Printf("%.6f\n", float64(positivos)/total)
	fmt.Printf("%.6f\n", float64(negativos)/total)
	fmt.Printf("%.6f\n", float64(zeros)/total)

}

func main() {
	arr := []int32{1, -3, 4, -5, 0, 6, -7}
	plusMinus(arr)
}
