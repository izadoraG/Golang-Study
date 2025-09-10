package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Points(games []string) int {
	countPoints := 0
	for _, game := range games {
		parts := strings.Split(game, ":")

		key, err1 := strconv.Atoi(parts[0])
		value, err2 := strconv.Atoi(parts[1])
		if err1 != nil || err2 != nil {
			continue
		}

		if key > value {
			countPoints += 3
		} else if key < value {
			countPoints += 0
		} else if key == value {
			countPoints += 1
		}
		fmt.Println(countPoints)

	}
	return countPoints

}

func main() {
	games := []string{"3:1", "2:2", "0:1"}
	fmt.Println(Points(games))
}
