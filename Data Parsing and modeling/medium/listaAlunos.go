package main

import "fmt"

type Alunos struct {
	Nome  string
	Notas []float64
}

func main() {
	alunos := []Alunos{
		{
			Nome:  "Ana",
			Notas: []float64{9.5, 7.8, 8.2},
		},
		{
			Nome:  "luiza",
			Notas: []float64{9.2, 7.9, 8.1},
		},
	}

	for _, aluno := range alunos {
		somaNotas := 0.0
		for _, nota := range aluno.Notas {
			somaNotas += nota
		}

		mediaAvaliacao := somaNotas / 3
		fmt.Printf("A media final do aluno %s é %.2f\n", aluno.Nome, mediaAvaliacao)
	}

}
