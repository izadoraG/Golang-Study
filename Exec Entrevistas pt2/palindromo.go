package main

import "fmt"

func palindromo(x string) bool {

	aux := ""

	for _, v := range x {
		aux = string(v) + aux
	}
	if aux != x {
		return false
	}
	return true
}

func main() {
	x := "racecare"
	fmt.Println(palindromo(x))
}
