package main

import "fmt"

func invertString(s string) string {

	aux := ""

	for _, v := range s {
		aux = string(v) + aux
	}
	return aux
}

func main() {
	s := "SAP Labs"
	fmt.Println(invertString(s))
}
