package main

import "fmt"

//func sumXor(n int64) int64 {
//
//	count := int64(0)
//	for i := 0; i <= int(n)/2; i++ {
//		if int(n)+i == int(n)^i {
//			count++
//		}
//	}
//	return count
//
//}

func sumXor(n int64) int64 {
	var count int64 = 0
	for n > 0 {
		if n&1 == 0 {
			count++
		}
		n >>= 1
	}
	return 1 << count
}

func main() {
	n := int64(10)
	fmt.Println(sumXor(n))
}
