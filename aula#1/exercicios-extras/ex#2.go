package main

import (
	"fmt"
	"time"
)

func main() {
	var option int
	var year int

	fmt.Println("Deseja inserir o ano de nascimento?")
	fmt.Println("1 - Inserir ano \n2 - Realizar cálculo automático")

	fmt.Scan(&option)

	if option == 1 {
		fmt.Scan(&year)
	} else {
		year = 1997
	}

	currentYear := time.Now()
	age := currentYear.Year() - year

	fmt.Printf("A idade é: %v", age)
}