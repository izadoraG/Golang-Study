package main

//1) From the pizzas JSON dataset, print to console the following:
//
//- How many "Four Cheese" pizzas were ordered?
//- How many "Caprese" pizzas with "medium" size were ordered?
//- How many unique customer names there are?
//- How many orders sam did?

//Quantas pizzas cada cliente pediu? ✅ médio
//Quantos sabores diferentes cada cliente pediu? ✅ médio
//Qual o sabor mais pedido no total? 🔥 difícil
//Qual o cliente que mais pediu pizzas? 🔥 difícil

import (
	"encoding/json"
	"fmt"
)

type Pizzas struct {
	Name      string `json:"name"`
	Phone     string `json:"phone_number"`
	PizzaName string `json:"pizza_flavor"`
	PizzaSize string `json:"pizza_size"`
}

func main() {

	var pizzas []Pizzas

	jsonData := `[{
"name": "Sam",
"phone_number": "024-135-2468",
"pizza_flavor": "Four Cheese",
"pizza_size": "medium"
},
{
"name": "Lee",
"phone_number": "890-123-4567",
"pizza_flavor": "Chorizo",
"pizza_size": "large"
},
{
"name": "Evan",
"phone_number": "468-579-6802",
"pizza_flavor": "Supreme",
"pizza_size": "medium"
},
{
"name": "Sage",
"phone_number": "579-680-7913",
"pizza_flavor": "Veggie",
"pizza_size": "medium"
},
{
"name": "Dylan",
"phone_number": "357-468-5791",
"pizza_flavor": "Caprese",
"pizza_size": "small"
},
{
"name": "Emery",
"phone_number": "579-680-7913",
"pizza_flavor": "Cheese",
"pizza_size": "large"
},
{
"name": "Aston",
"phone_number": "456-789-0123",
"pizza_flavor": "Steak",
"pizza_size": "medium"
},
{
"name": "Dylan",
"phone_number": "913-024-1357",
"pizza_flavor": "BBQ Chicken",
"pizza_size": "small"
},
{
"name": "Casey",
"phone_number": "246-357-4680",
"pizza_flavor": "Seafood",
"pizza_size": "medium"
},
{
"name": "Gustav",
"phone_number": "456-789-0123",
"pizza_flavor": "Spinach",
"pizza_size": "large"
},
{
"name": "Riley",
"phone_number": "579-680-7913",
"pizza_flavor": "Steak",
"pizza_size": "large"
},
{
"name": "Indiana",
"phone_number": "579-680-7913",
"pizza_flavor": "Roasted Garlic",
"pizza_size": "medium"
},
{
"name": "Chris",
"phone_number": "246-357-4680",
"pizza_flavor": "Spinach and Artichoke",
"pizza_size": "medium"
},
{
"name": "Peter",
"phone_number": "468-579-6802",
"pizza_flavor": "Cheeseburger",
"pizza_size": "medium"
},
{
"name": "Kyle",
"phone_number": "791-802-9135",
"pizza_flavor": "Caprese",
"pizza_size": "medium"
},
{
"name": "Charlie",
"phone_number": "678-901-2345",
"pizza_flavor": "Smoked Salmon",
"pizza_size": "small"
},
{
"name": "Dakota",
"phone_number": "357-468-5791",
"pizza_flavor": "Hawaiian",
"pizza_size": "large"
},
{
"name": "Avery",
"phone_number": "901-234-5678",
"pizza_flavor": "Pesto",
"pizza_size": "large"
},
{
"name": "Kyle",
"phone_number": "913-024-1357",
"pizza_flavor": "Roasted Garlic",
"pizza_size": "large"
},
{
"name": "Dakota",
"phone_number": "024-135-2468",
"pizza_flavor": "Smoked Salmon",
"pizza_size": "medium"
},
{
"name": "Drew",
"phone_number": "012-345-6789",
"pizza_flavor": "Margherita",
"pizza_size": "medium"
},
{
"name": "Pat",
"phone_number": "357-468-5791",
"pizza_flavor": "Roasted Garlic",
"pizza_size": "medium"
},
{
"name": "Avery",
"phone_number": "012-345-6789",
"pizza_flavor": "Caprese",
"pizza_size": "large"
},
{
"name": "Alex",
"phone_number": "579-680-7913",
"pizza_flavor": "Steak",
"pizza_size": "large"
},
{
"name": "Bailey",
"phone_number": "135-246-3579",
"pizza_flavor": "Cheeseburger",
"pizza_size": "medium"
},
{
"name": "Jane",
"phone_number": "567-890-1234",
"pizza_flavor": "Caprese",
"pizza_size": "large"
},
{
"name": "Jordan",
"phone_number": "567-890-1234",
"pizza_flavor": "Pineapple",
"pizza_size": "large"
},
{
"name": "Indiana",
"phone_number": "456-789-0123",
"pizza_flavor": "Steak",
"pizza_size": "small"
},
{
"name": "Indiana",
"phone_number": "579-680-7913",
"pizza_flavor": "Spinach and Artichoke",
"pizza_size": "small"
},
{
"name": "Mandy",
"phone_number": "913-024-1357",
"pizza_flavor": "Roasted Garlic",
"pizza_size": "small"
},
{
"name": "Mandy",
"phone_number": "802-913-0246",
"pizza_flavor": "Mushroom",
"pizza_size": "small"
},
{
"name": "Lee",
"phone_number": "246-357-4680",
"pizza_flavor": "Margherita",
"pizza_size": "large"
},
{
"name": "Lee",
"phone_number": "012-345-6789",
"pizza_flavor": "Buffalo Chicken",
"pizza_size": "small"
},
{
"name": "Dakota",
"phone_number": "234-567-8901",
"pizza_flavor": "Veggie",
"pizza_size": "small"
},
{
"name": "Jordan",
"phone_number": "579-680-7913",
"pizza_flavor": "Pepperoni",
"pizza_size": "small"
},
{
"name": "Morgan",
"phone_number": "234-567-8901",
"pizza_flavor": "Supreme",
"pizza_size": "large"
},
{
"name": "Sam",
"phone_number": "123-456-7890",
"pizza_flavor": "Margherita",
"pizza_size": "medium"
},
{
"name": "Emery",
"phone_number": "791-802-9135",
"pizza_flavor": "Veggie",
"pizza_size": "large"
},
{
"name": "Gustav",
"phone_number": "791-802-9135",
"pizza_flavor": "Cheeseburger",
"pizza_size": "large"
},
{
"name": "John",
"phone_number": "802-913-0246",
"pizza_flavor": "Cheese",
"pizza_size": "medium"
},
{
"name": "Sam",
"phone_number": "579-680-7913",
"pizza_flavor": "Supreme",
"pizza_size": "medium"
},
{
"name": "Gustav",
"phone_number": "468-579-6802",
"pizza_flavor": "Pesto",
"pizza_size": "small"
},
{
"name": "Gustav",
"phone_number": "246-357-4680",
"pizza_flavor": "Pineapple",
"pizza_size": "large"
},
{
"name": "Taylor",
"phone_number": "791-802-9135",
"pizza_flavor": "Mushroom",
"pizza_size": "large"
},
{
"name": "John",
"phone_number": "246-357-4680",
"pizza_flavor": "Pepperoni",
"pizza_size": "medium"
},
{
"name": "Peter",
"phone_number": "579-680-7913",
"pizza_flavor": "Buffalo Chicken",
"pizza_size": "medium"
},
{
"name": "Steve",
"phone_number": "791-802-9135",
"pizza_flavor": "Four Cheese",
"pizza_size": "large"
},
{
"name": "Aidan",
"phone_number": "678-901-2345",
"pizza_flavor": "Spinach",
"pizza_size": "medium"
},
{
"name": "Cameron",
"phone_number": "234-567-8901",
"pizza_flavor": "Roasted Garlic",
"pizza_size": "large"
},
{
"name": "John",
"phone_number": "246-357-4680",
"pizza_flavor": "Chorizo",
"pizza_size": "small"
},
{
"name": "Bob",
"phone_number": "567-890-1234",
"pizza_flavor": "Smoked Salmon",
"pizza_size": "medium"
},
{
"name": "Morgan",
"phone_number": "901-234-5678",
"pizza_flavor": "Hawaiian",
"pizza_size": "large"
},
{
"name": "Skylar",
"phone_number": "135-246-3579",
"pizza_flavor": "BBQ Chicken",
"pizza_size": "large"
},
{
"name": "Dylan",
"phone_number": "345-678-9012",
"pizza_flavor": "Pineapple",
"pizza_size": "large"
},
{
"name": "Kyle",
"phone_number": "024-135-2468",
"pizza_flavor": "Seafood",
"pizza_size": "large"
},
{
"name": "Pat",
"phone_number": "791-802-9135",
"pizza_flavor": "Pepperoni",
"pizza_size": "small"
}]`

	json.Unmarshal([]byte(jsonData), &pizzas)

	var count int
	var count2 int
	var count3 int
	for _, pizza := range pizzas {
		if pizza.PizzaName == "Four Cheese" {
			count++
		} else if pizza.PizzaName == "Caprese" && pizza.PizzaSize == "medium" {
			count2++
		} else if pizza.Name == "Sam" {
			count3++
		}
	}

	//how many pizzas each client order
	charCount := make(map[string]int)
	for _, pizza := range pizzas {
		charCount[pizza.Name]++
	}

	for client, qtd := range charCount {
		fmt.Printf("%s: %d\n", client, qtd)
	}

	//witch client order more pizzas
	mostOrder2 := 0
	clientName := ""
	for client, qtd := range charCount {
		if qtd > mostOrder2 {
			clientName = client
			mostOrder2 = qtd
			fmt.Printf("O cliente %s, pediu %d pizzas\n", clientName, mostOrder2)
		}

	}

	//most favorite flavor
	flavorCount := make(map[string]int)
	for _, pizza := range pizzas {
		flavorCount[pizza.PizzaName]++
	}

	mostOrder := 0
	maxFlavor := ""
	for flavor, qtd := range flavorCount {
		if qtd > mostOrder {
			maxFlavor = flavor
			mostOrder = qtd
			fmt.Printf("pizza mais pedida %s(%d pedidos)\n", maxFlavor, mostOrder)
		}

	}

	fmt.Printf("Pizza coun Four cheese: %d\n", count)
	fmt.Printf("Pizza caprese e medium: %d\n", count2)
	fmt.Printf("Unique users: %d\n", len(charCount))
	fmt.Printf("Sam orders: %d\n", count3)
}
