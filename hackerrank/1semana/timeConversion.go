package main

import (
	"fmt"
	"strconv"
)

func timeConversion(s string) string {
	hora := s[:2]
	minuto := s[3:5]
	segundo := s[6:8]
	periodo := s[8:] // AM ou PM

	horaInt, _ := strconv.Atoi(hora)

	if periodo == "AM" {
		if horaInt == 12 {
			horaInt = 0
		}
	} else if periodo == "PM" {
		if horaInt != 12 {
			horaInt += 12
		}
	}

	return fmt.Sprintf("%02d:%s:%s", horaInt, minuto, segundo)
}

func main() {
	input := "12:05:45AM"
	result := timeConversion(input)
	fmt.Println(result)
}
