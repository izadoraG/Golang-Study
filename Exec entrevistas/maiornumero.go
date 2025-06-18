package main

import "fmt"

func maiorNumero(arr []int) int {

	maior := 0

	for _, a := range arr {
		if a > maior {
			maior = a
		}
	}
	return maior
}

func main() {
	arr := []int{1, 3, 6, 4, 70, 8, 9, 4}
	fmt.Println(maiorNumero(arr))
}
