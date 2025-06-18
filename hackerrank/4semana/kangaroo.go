package main

import "fmt"

func kangaroo(x1 int32, v1 int32, x2 int32, v2 int32) string {

	if (v1 > v2) && (x1 < x2) && (x2-x1)%(v1-v2) == 0 {
		return "YES"
	}

	if (v2 > v1) && (x2 < x1) && (x1-x2)%(v2-v1) == 0 {
		return "YES"
	}

	return "NO"
}

func main() {
	x1 := int32(0)
	v1 := int32(2)
	x2 := int32(5)
	v2 := int32(3)
	fmt.Println(kangaroo(x1, v1, x2, v2))
}

//x1 + n*v1 = x2 + n*v2
//n*v1 - n*v2 = x2 - x1
//n * (v1 - v2) = x2 - x1
//n = (x2 - x1) / (v1 - v2)
