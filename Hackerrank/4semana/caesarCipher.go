package main

import "fmt"

func caesarCipher(s string, k int32) string {
	k = k % 26
	result := ""

	for _, ch := range s {
		switch {
		case ch >= 'a' && ch <= 'z':
			rotated := ((ch - 'a' + k) % 26) + 'a'
			result += string(rotated)
		case ch >= 'A' && ch <= 'Z':
			rotated := ((ch - 'A' + k) % 26) + 'A'
			result += string(rotated)
		default:
			result += string(ch)
		}
	}
	return result
}
func main() {
	s := "middle-Outz"
	k := int32(2)
	fmt.Println(caesarCipher(s, k))
}
