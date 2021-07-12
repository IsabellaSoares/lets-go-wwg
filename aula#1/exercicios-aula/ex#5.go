package main

import (
	"fmt"
)

func main() {
	a := 15 >= 15
	b := 100 < 1000
	c := 10 == 10
	d := 5 > 4
	e := 50 <= 40

	fmt.Printf("O tipo de a é %T e seu valor é %v\n", a, a)
	fmt.Printf("O tipo de b é %T e seu valor é %v\n", b, b)
	fmt.Printf("O tipo de c é %T e seu valor é %v\n", c, c)
	fmt.Printf("O tipo de d é %T e seu valor é %v\n", d, d)
	fmt.Printf("O tipo de e é %T e seu valor é %v\n", e, e)
}
