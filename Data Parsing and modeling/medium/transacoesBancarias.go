package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Cliente struct {
	Nome       string      `json:"nome"`
	Transacoes []Transacao `json:"transacoes"`
}

type Transacao struct {
	Descricao string  `json:"descricao"`
	Tipo      string  `json:"tipo"`
	Valor     float64 `json:"valor"`
}

func main() {
	var clientes []Cliente
	jsonData := `
	[
		{
			"nome": "João",
			"transacoes": [
				{ "descricao": "Salário", "tipo": "entrada", "valor": 3000 },
				{ "descricao": "Aluguel", "tipo": "saida", "valor": 1200 },
				{ "descricao": "Mercado", "tipo": "saida", "valor": 500 }
			]
		}
	]`

	err := json.Unmarshal([]byte(jsonData), &clientes)
	if err != nil {
		log.Fatal(err)
	}

	for _, cliente := range clientes {
		fmt.Println("Cliente:", cliente.Nome)

		totalEntrada := 0.0
		totalSaida := 0.0

		fmt.Println("Transacoes:")
		for _, transacao := range cliente.Transacoes {
			fmt.Printf("- %s (%s): R$%.2f\n", transacao.Descricao, transacao.Tipo, transacao.Valor)

			if transacao.Tipo == "entrada" {
				totalEntrada += transacao.Valor
			} else if transacao.Tipo == "saida" {
				totalSaida += transacao.Valor
			}

		}
		resto := totalEntrada - totalSaida
		fmt.Printf("Total de Entradas: R$%.2f\n", totalEntrada)
		fmt.Printf("Total de Saídas: R$%.2f\n", totalSaida)
		fmt.Println(resto)
	}

}
