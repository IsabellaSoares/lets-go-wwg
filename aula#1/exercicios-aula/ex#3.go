package main

import (
	"fmt"
)

func main() {
	var chocolate, cafe, sorvete float64 = 4.50, 10.0, 19.90
	var totalDaCompra = (chocolate * 2) + (cafe * 3) + (sorvete * 2)

	fmt.Printf("o valor da compra deu: \n%v", totalDaCompra)
}