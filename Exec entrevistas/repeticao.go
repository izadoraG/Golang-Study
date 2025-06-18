package main

import "fmt"

func main() {

	for i := 0; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println("Numero:", i, "Fizz")
		} else if i%5 == 0 {
			fmt.Println("Numero:", i, "Buzz")
		} else if i%3 == 0 && i%5 == 0 {
			fmt.Println("Numero:", i, "FizzBuzz")
		}
	}
}
