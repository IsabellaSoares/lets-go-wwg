package main

import (
	"fmt"
)

func main() {
	var ano = make(map[int]string)

	ano[0] = "Janeiro"
	ano[1] = "Fevereiro"
	ano[2] = "Março"
	ano[3] = "Abril"
	ano[4] = "Maio"
	ano[5] = "Junho"
	ano[6] = "Julho"
	ano[7] = "Agosto"
	ano[8] = "Setembro"
	ano[9] = "Outubro"
	ano[10] = "Novembro"
	ano[11] = "Dezembro"

	fmt.Println("Primeiro mês do ano:", ano[0])
	fmt.Println("Último mês do ano:", ano[11])
	fmt.Println(len(ano))
}
