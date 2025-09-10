package main

import (
	"fmt"
	"strconv"
)

//func compress(chars []byte) int {
//	var result []string
//
//	count := 1
//	for i := 0; i < len(chars); i++ {
//		if i == len(chars)-1 || chars[i] != chars[i+1] {
//			result = append(result, string(chars[i]))
//			if count > 1 {
//				result = append(result, strconv.Itoa(count))
//			}
//			count = 1
//		} else {
//			count++
//		}
//	}
//
//	return len(result)
//}

func compress(chars []byte) int {
	write := 0 // posição onde vamos escrever no array
	count := 1

	for i := 0; i < len(chars); i++ {
		// se é o último ou diferente do próximo
		if i == len(chars)-1 || chars[i] != chars[i+1] {
			chars[write] = chars[i]
			write++

			if count > 1 {
				// converte número para string e escreve cada dígito
				for _, d := range strconv.Itoa(count) {
					chars[write] = byte(d)
					write++
				}
			}
			count = 1
		} else {
			count++
		}
	}

	fmt.Printf("%q\n", chars[:write]) // imprime ["a" "2" "b" "2" "c" "3"]
	return write
}

func main() {
	chars := []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}
	fmt.Println(compress(chars))
}
