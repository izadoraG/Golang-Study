package main

import "fmt"

func rotateLeft(d int32, arr []int32) []int32 {

	if len(arr) >= int(d) {
		prefix := arr[:d]
		arr = append(arr[d:], prefix...)
	}
	return arr

}

func main() {
	arr := []int32{1, 2, 3, 4, 5}
	d := int32(4)
	fmt.Println(rotateLeft(d, arr))
}
