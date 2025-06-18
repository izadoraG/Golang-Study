package main

import "fmt"

func fibonatti(num int) int {
	if num == 0 {
		return 0
	}

	if num == 1 || num == 2 {
		return 1
	}

	return fibonatti(num-1) + fibonatti(num-2)
}

func main() {
	num := 8
	fmt.Println(fibonatti(num))

	for i := 0; i <= 8; i++ {
		fmt.Println(i, fibonatti(i))
	}
}
