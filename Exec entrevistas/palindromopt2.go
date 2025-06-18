package main

import "fmt"

func Palindromo(str []string) (result bool) {

	for _, palavra := range str {
		aux := ""

		for _, v := range palavra {
			aux = string(v) + aux
		}

		if aux != palavra {
			return false
		}

	}
	return true
}

func main() {
	str := []string{"racecar", "ama"}
	fmt.Println(Palindromo(str))
}
