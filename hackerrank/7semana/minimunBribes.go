package main

import "fmt"

func minimumBribes(q []int32) {
	bribes := 0 // Contador de subornos (bribes)

	// Percorre a fila de trás para frente
	for i := len(q) - 1; i >= 0; i-- {

		// 1. Verifica se a pessoa está mais de 2 posições à frente do que deveria
		// A pessoa com número q[i] deveria estar na posição q[i] - 1
		if q[i]-(int32(i)+1) > 2 {
			fmt.Println("Too chaotic") // Violou a regra de no máximo 2 subornos
			return
		}

		// 2. Conta quantas pessoas maiores que q[i] estão na frente dela
		// Começa de no máximo duas posições atrás da posição original de q[i]
		// Isso cobre os possíveis subornos que a pessoa pode ter recebido

		// Ex: se q[i] = 5, então a pessoa poderia ter começado no índice 4
		// e ido no máximo para a frente (posição i). Só olhamos de (5-2)=3 até i
		for j := max(0, int(q[i])-2); j < i; j++ {
			if q[j] > q[i] {
				bribes++ // Essa pessoa passou na frente de q[i], ou seja, deu um suborno
			}
		}
	}

	fmt.Println(bribes) // Total de subornos identificados
}

func main() {
	q := []int32{2, 5, 1, 3, 4}
	minimumBribes(q)
}
