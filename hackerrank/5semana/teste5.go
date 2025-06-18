package main

import "fmt"

func fibonacciModified(t1 int32, t2 int32, n int32) int32 {
	if n == 0 {
		return 0
	}

	if n == 1 || n == 2 {
		return 1
	}

	return fibonacciModified(n-1) + fibonacciModified(n-2)
}

func main() {
	t1 := int32(0)
	t2 := int32(1)
	n := int32(5)

	for i := 0; i <= 8; i++ {
		fmt.Println(fibonatti(i[5]))
	}
}
