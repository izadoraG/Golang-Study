package main

import "fmt"

func Palindromo(str string) (result bool) {

	aux := ""

	for _, v := range str {
		aux = string(v) + aux
	}

	if aux != str {
		return false
	}
	return true

}

func main() {
	str := "racecare"
	fmt.Println(Palindromo(str))
}
