package main

import (
	"fmt"
)

func main() {
	var odd = []int{1, 3, 5, 7, 9, 11}
	var even = []int{2, 4, 6, 8, 10, 12}
	var union []int

	union = append(odd, even...)

	fmt.Println(odd)
	fmt.Println(even)
	fmt.Println(union)
}
