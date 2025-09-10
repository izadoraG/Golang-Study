package main

import "fmt"

func countingValleys(steps int32, path string) int32 {

	var nivel, vales int32

	for _, passo := range path {
		if passo == 'U' {
			nivel++
			if nivel == 0 {
				vales++
			}

		} else if passo == 'D' {
			nivel--
		}

	}
	return vales
}

func main() {
	valley := "UDDDUDUU"
	steps := int32(8)
	fmt.Println(countingValleys(steps, valley))
}
