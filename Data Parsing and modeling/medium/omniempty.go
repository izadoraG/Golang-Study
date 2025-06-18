package main

import (
	"encoding/json"
	"fmt"
)

type Produto struct {
	Nome     string   `json:"name"`
	Price    float64  `json:"price"`
	Disconto *float64 `json:"discount,omitempty"`
}

func main() {
	var produtos []Produto

	jsonData := `[
  { "name": "Laptop", "price": 1000 },
  { "name": "Mouse", "price": 50, "discount": 5 }
]
`

	json.Unmarshal([]byte(jsonData), &produtos)

	for _, p := range produtos {
		if p.Disconto != nil {
			fmt.Printf("Produto: %s | Preco: %.2f  | Desconto: %.2f\n", p.Nome, p.Price, *p.Disconto)
		} else {
			fmt.Printf("Produto: %s | Preco: %.2f  | Sem desconto\n", p.Nome, p.Price)
		}
	}

}
