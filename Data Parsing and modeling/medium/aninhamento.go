package main

import (
	"encoding/json"
	"fmt"
)

type Task struct {
	Status string `json:"status"`
	Dado   Dados  `json:"data"`
}

type Dados struct {
	Data User `json:"user"`
}

type User struct {
	Id    int      `json:"id"`
	Name  string   `json:"name"`
	Roles []string `json:"roles"`
}

func main() {

	var tasks []Task
	jsonData := `[{
  "status": "success",
  "data": {
    "user": {
      "id": 123,
      "name": "Lucas",
	  "roles": ["admin", "editor"]
    }
  }
}
]`

	json.Unmarshal([]byte(jsonData), &tasks)

	for _, task := range tasks {
		fmt.Println("Nome:", task.Dado.Data.Name)
		fmt.Println("Roles:")
		for _, role := range task.Dado.Data.Roles {
			fmt.Println(role)
		}
	}
}
