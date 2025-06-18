package main

import "fmt"

func invertFirstRow(matrix [][]int) int {

	for i, j := 0, len(matrix)-1; i < j; i, j = i+1, j-1 {
		matrix[i][2], matrix[j][2] = matrix[j][2], matrix[i][2]
	}

	firstRow := matrix[0]
	for i, j := 0, len(firstRow)-1; i < j; i, j = i+1, j-1 {
		firstRow[i], firstRow[j] = firstRow[j], firstRow[i]
	}

	rows := len(matrix)
	rowHalf := rows / 2

	sum := 0
	for i := 0; i < rowHalf; i++ {
		for j := 0; j < rowHalf; j++ {
			sum += matrix[i][j]
		}
	}
	return sum

}

func main() {
	matrix := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
	}

	fmt.Println("Original Matrix:")
	for _, row := range matrix {
		fmt.Println(row)
	}

	sum := invertFirstRow(matrix)

	fmt.Println("\nMatrix with inverted first row:")
	for _, row := range matrix {
		fmt.Println(row)

	}

	fmt.Println(sum)
}
