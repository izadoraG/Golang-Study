package main

import (
	"fmt"
	"strconv"
)

func flippingBits(n int64) int64 {

	binario := fmt.Sprintf("%032b", n)

	invertido := ""

	for _, bin := range binario {
		if bin == '0' {
			invertido += "1"
		} else {
			invertido += "0"
		}
	}

	resultado, err := strconv.ParseInt(invertido, 2, 64)
	if err != nil {
		panic(err)
	}

	return resultado

}

func main() {
	decimalNumber := 9
	fmt.Println(flippingBits(int64(decimalNumber)))
}
