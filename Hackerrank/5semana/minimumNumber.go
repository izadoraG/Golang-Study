package main

import (
	"fmt"
	"strings"
)

func minimumNumber(n int32, password string) int32 {

	numbers := "0123456789"
	lower_case := "abcdefghijklmnopqrstuvwxyz"
	upper_case := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	special_characters := "!@#$%^&*()-+"

	hasLower := false
	hasUpper := false
	hasSpecial := false
	hasNumber := false

	for _, char := range password {
		if strings.Contains(lower_case, string(char)) {
			hasLower = true
		} else if strings.Contains(upper_case, string(char)) {
			hasUpper = true
		} else if strings.Contains(special_characters, string(char)) {
			hasSpecial = true
		} else if strings.Contains(numbers, string(char)) {
			hasNumber = true
		}
	}
	count := int32(0)
	if !hasLower {
		count++
	}
	if !hasUpper {
		count++
	}
	if !hasNumber {
		count++
	}
	if !hasSpecial {
		count++
	}

	if n+count < 6 {
		return 6 - n
	}
	return count

}

func main() {
	password := "2bb#AA"
	n := int32(6)
	fmt.Println(minimumNumber(n, password))
}
