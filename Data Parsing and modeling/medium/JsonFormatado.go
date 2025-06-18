package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Livro struct {
	Titulo string `json: "titulo"`
	Autor  string `json: "autor"`
	Ano    int    `json: "ano"`
}

func main() {
	livros := []Livro{
		{Titulo: "Assim que acaba", Autor: "Colleen hover", Ano: 2010},
		{Titulo: "Assim que comeca", Autor: "Colleen hover", Ano: 2020},
		{Titulo: "O lado feio do amor", Autor: "Colleen hover", Ano: 2022},
	}

	jsondata, err := json.MarshalIndent(livros, " ", " ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(jsondata))

}
