//ceheck if N is power of 2
//iff not, power of 2 until < N and reduce this number of N. N = resut os sub
//if yes, divide by 2 and pass the rest

package main

import (
	"fmt"
	"math/bits"
)

func counterGame(n int64) string {

	count := 0

	for n > 1 {
		if n%2 == 0 {
			nextNumber := n / 2
			n = nextNumber
			count++
		} else {
			n -= 1 << (bits.Len64(uint64(n)) - 1)
			count++
		}
	}

	if count%2 == 0 {
		return "Richard"
	}
	return "Louise"

}

func main() {
	n := int64(132)
	fmt.Println(counterGame(int64(n)))
}
