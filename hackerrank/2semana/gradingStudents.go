package main

import "fmt"

func gradingStudents(grades []int32) []int32 {
	var result []int32

	for _, grade := range grades {
		if grade < 38 {
			result = append(result, grade)
		} else {
			proximoMultiploDe5 := ((grade / 5) + 1) * 5
			if proximoMultiploDe5-grade < 3 {
				result = append(result, proximoMultiploDe5)
			} else {
				result = append(result, grade)
			}
		}
	}
	return result
}

func main() {
	grades := []int32{84, 27, 57}
	fmt.Println(gradingStudents(grades))
}
