package main

import "fmt"

func fibonnati(n int) int {

	if n == 0 {
		return 0
	}

	if n == 1 || n == 2 {
		return 1
	}

	return fibonnati(n-1) + fibonnati(n-2)
}

func main() {
	n := 6
	fmt.Println(fibonnati(n))
	for i := 0; i <= 6; i++ {
		fmt.Println(i, fibonnati(i))
	}
}
