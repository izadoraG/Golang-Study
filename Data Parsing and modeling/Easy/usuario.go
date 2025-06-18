package main

import (
	"encoding/json"
	"fmt"
)

type Pessoas struct {
	Id   int    `json:"id"`
	Nome string `json:"name"`
}

func main() {

	var pessoas []Pessoas
	jsonData := `[
  { "id": 1, "name": "Alice" },
  { "id": 2, "name": "Bob" },
  { "id": 3, "name": "Carol" }
]`
	json.Unmarshal([]byte(jsonData), &pessoas)
	for _, pessoa := range pessoas {
		fmt.Println("Nome:", pessoa.Nome)
	}
}
