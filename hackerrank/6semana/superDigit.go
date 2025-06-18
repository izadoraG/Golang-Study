package main

import (
	"fmt"
)

func superDigit(n string, k int32) int32 {
	// Função auxiliar interna (closure) para calcular o super dígito de um número inteiro
	var calc func(int64) int32
	calc = func(num int64) int32 {
		if num < 10 {
			return int32(num)
		}
		var sum int64
		for num > 0 {
			sum += num % 10
			num /= 10
		}
		return calc(sum)
	}

	// Soma os dígitos da string n
	var sum int64
	for _, ch := range n {
		sum += int64(ch - '0')
	}

	// Multiplica pela quantidade de repetições
	total := sum * int64(k)

	// Chama o cálculo recursivo
	return calc(total)
}

func main() {
	n := "9875"
	k := int32(4)
	fmt.Println(superDigit(n, k))
}
