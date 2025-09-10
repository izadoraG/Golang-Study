package main

import (
	"fmt"
)

func gcdOfStrings(str1 string, str2 string) string {
	if str2+str1 != str1+str2 {
		return ""
	}
	gcdLength := gcd(len(str1), len(str2))
	return str1[:gcdLength]
}
func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

//achar o minimo divisor comum

func main() {
	str1 := "ABABAB"
	str2 := "ABAB"
	fmt.Println(gcdOfStrings(str1, str2))
}
