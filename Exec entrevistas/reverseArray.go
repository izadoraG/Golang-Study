package main

import "fmt"

func reverseArray(a []int32) []int32 {
	n := len(a)
	for i := 0; i < n/2; i++ {
		a[i], a[n-1-i] = a[n-1-i], a[i]
	}
	return a
}

func main() {
	numeros := []int32{10, 20, 30, 40, 50}
	fmt.Println("Antes:", numeros)
	reverseArray(numeros)
	fmt.Println("Depois:", numeros)
}
