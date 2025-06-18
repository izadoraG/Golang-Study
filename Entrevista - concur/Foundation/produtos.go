//Ordenar Produtos por Popularidade e Preço
//Enunciado:
//Você possui uma lista de produtos com nome, preço e comentários (reviews). Cada comentário é uma string separada por "|".
//Ordene os produtos pelo maior número de comentários.
//Depois pelo menor preço.
//Exporte o resultado como JSON.

//Um novo campo ratings com uma lista de notas (de 1 a 5).
//A lógica esperada pode passar a considerar média de avaliações além de preço e quantidade de comentários
//Ordene os produtos primeiro pela maior média das avaliações (ratings), depois pelo maior número de comentários (comments) e, por fim, pelo menor preço.
//Retorne os dados em formato JSON sem alterar os nomes dos campos.

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"sort"
	"strings"
)

type Produtos struct {
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Comments string  `json:"comments"`
	Rating   []int   `json:"rating"`
}

func countComments(comments string) int {
	if comments == "" {
		return 0
	}
	return len(strings.Split(comments, "|"))
}

func main() {

	var produtos []Produtos
	jsonData := `[
  {
    "name": "Notebook",
    "price": 4500.00,
    "comments": "Fast|Reliable|Great battery",
    "ratings": [5, 5, 4, 5, 4]
  },
  {
    "name": "Smartphone",
    "price": 2300.00,
    "comments": "Sleek design|Good camera",
    "ratings": [4, 4, 3, 5]
  },
  {
    "name": "Mouse",
    "price": 120.00,
    "comments": "Comfortable|Cheap|Durable|Wireless",
    "ratings": [5, 5, 4, 4, 5]
  },
  {
    "name": "Monitor",
    "price": 980.00,
    "comments": "Sharp image",
    "ratings": [3, 4, 4]
  },
  {
    "name": "Headphones",
    "price": 299.99,
    "comments": "Noise-cancelling|Bass|Comfy",
    "ratings": [5, 4, 5, 5, 5]
  }
]`

	json.Unmarshal([]byte(jsonData), &produtos)

	sort.Slice(produtos, func(i, j int) bool {
		if countComments(produtos[i].Comments) == countComments(produtos[j].Comments) {
			return produtos[i].Price < produtos[j].Price
		}
		return produtos[i].Comments > produtos[j].Comments
	})

	result, err := json.MarshalIndent(produtos, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(result))

}
