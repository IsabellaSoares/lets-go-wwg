package main

import (
	"fmt"
)

func main() {
	var quilometros int16
	quilometros = 150

	fmt.Println(quilometros)
}

// não compila porque dá erro de overflow, o tipo de dado int8 não é suficiente para conter o valor 150
// int8 -> 8 bits -> -128 ~ 127
// int16 -> 16 bits -> -2^15 ~ 2^15-1