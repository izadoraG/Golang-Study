package main

import "fmt"

func Reverse(str string) (result string) {

	aux := ""

	for _, a := range str {
		aux = string(a) + aux
	}

	return aux

}

func main() {
	str := "izadora"

	fmt.Println(Reverse(str))
}
