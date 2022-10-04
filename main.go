package main

import "fmt"

func main() {

	// Switch
	

	switch modulo := 4 % 2; modulo {
	case 0:
		fmt.Println("es par")
	default:
		fmt.Println("es impar")
	}

	// Sin Condicion

	value:= 300

	switch {
	case value > 200:
		fmt.Println("Es mayor a 200")
	case value < 0:
		fmt.Println("ES menor a cero") 
	default:
		fmt.Println("No hay condicion")
	}
}
