package main

import "fmt"

func pageCount(n int32, p int32) int32 {
	// Mínimo entre virar da frente (começo) ou virar de trás (fim)
	fromFront := p / 2
	fromBack := (n / 2) - (p / 2)

	if fromFront < fromBack {
		return fromFront
	}
	return fromBack
}

func main() {
	n := int32(5)
	p := int32(4)
	fmt.Println(pageCount(n, p)) // Saída esperada: 0
}
