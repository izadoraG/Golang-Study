package main

import "fmt"

func CalculateYears(years int) (result [3]int) {

	catYears := 0
	dogYears := 0

	if years == 1 {
		catYears += 15
	} else if years == 2 {
		catYears += 15 + 9
	} else if years > 2 {
		catYears = 15 + 9 + (years-2)*4
	}

	if years == 1 {
		dogYears += 15
	} else if years == 2 {
		dogYears += 15 + 9
	} else if years > 2 {
		dogYears = 15 + 9 + (years-2)*5
	}

	return [3]int{catYears, dogYears, years}

}

func main() {
	years := 10
	fmt.Println(CalculateYears(years))
}
