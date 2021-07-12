package main

import (
	"fmt"
)

func main() {
	var slice = make([]string, 12)

	slice[0] = "Aquário"
	slice[1] = "Peixes"
	slice[2] = "Áries"
	slice[3] = "Touro"
	slice[4] = "Gêmeos"
	slice[5] = "Câncer"
	slice[6] = "Leão"
	slice[7] = "Virgem"
	slice[8] = "Libra"
	slice[9] = "Escorpião"
	slice[10] = "Sagitário"
	slice[11] = "Capricórnio"

	fmt.Println("Todos os signos do Zodíaco:", slice)
	fmt.Println("Alguns os signos do Zodíaco:", slice[2:8])
}
