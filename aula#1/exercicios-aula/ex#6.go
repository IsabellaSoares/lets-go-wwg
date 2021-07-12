package main

import (
	"fmt"
)

func main() {
	var array = [5]string{"zero", "um", "dois", "tres", "quatro"}

	fmt.Printf("O tipo do meu array é %T e os valores são %v", array, array)
}
