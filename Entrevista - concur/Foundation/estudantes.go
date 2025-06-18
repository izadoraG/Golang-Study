//Você tem uma lista de estudantes com nome, nota final e número de faltas. Ordene os dados:
//Primeiro pela maior nota.
//Depois pelo menor número de faltas.
//Imprima a lista ordenada em formato JSON.

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"sort"
)

type Estudantes struct {
	Name     string  `json:"name"`
	Grade    float64 `json:"grade"`
	Absences int     `json:"absences"`
}

func main() {
	var estudantes []Estudantes
	jsonData := `[
	{
	"name": "Alice",
	"grade": 9.5,
	"absences": 2
	},
	{
	"name": "Bruno",
	"grade": 8.0,
	"absences": 0
	},
	{
	"name": "Carlos",
	"grade": 9.5,
	"absences": 1
	},
	{
	"name": "Diana",
	"grade": 7.5,
	"absences": 3
	},
	{
	"name": "Eduardo",
	"grade": 8.0,
	"absences": 1
	}
]`

	json.Unmarshal([]byte(jsonData), &estudantes)

	sort.Slice(estudantes, func(i, j int) bool {
		if estudantes[i].Grade == estudantes[j].Grade {
			return estudantes[i].Absences > estudantes[j].Absences
		}
		return estudantes[i].Grade > estudantes[j].Grade
	})

	result, err := json.MarshalIndent(estudantes, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(result))

}
