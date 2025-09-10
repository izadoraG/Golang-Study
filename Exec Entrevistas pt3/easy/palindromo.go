package main

import "fmt"

func palindromo(s string) bool {

	aux := ""

	for _, v := range s {
		aux = string(v) + aux
	}

	if aux == s {
		return true
	}
	return false
}

func main() {
	s := "PAP"
	fmt.Println(palindromo(s))
}
