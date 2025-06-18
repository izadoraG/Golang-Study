package main

import (
	"encoding/json"
	"fmt"
)

type Pedido struct {
	Order  int     `json:"order_id"`
	Amount float64 `json:"amount"`
	Client Cliente `json:"client"` // <- Aqui estava faltando a tag
}

type Cliente struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	var pedidos []Pedido

	jsonData := `[
	  {
		"order_id": 101,
		"client": { "id": 1, "name": "Alice" },
		"amount": 250.50
	  },
	  {
		"order_id": 102,
		"client": { "id": 2, "name": "Bob" },
		"amount": 100.00
	  },
	  {
		"order_id": 103,
		"client": { "id": 1, "name": "Alice" },
		"amount": 199.99
	  }
	]`

	// Parse do JSON
	err := json.Unmarshal([]byte(jsonData), &pedidos)
	if err != nil {
		fmt.Println("Erro ao decodificar JSON:", err)
		return
	}

	// Mapa para acumular os totais por cliente
	totaisPorCliente := make(map[string]float64)

	for _, pedido := range pedidos {
		totaisPorCliente[pedido.Client.Name] += pedido.Amount
	}

	// Imprimir os totais
	for nome, total := range totaisPorCliente {
		fmt.Printf("Cliente: %s | Total: R$%.2f\n", nome, total)
	}
}
