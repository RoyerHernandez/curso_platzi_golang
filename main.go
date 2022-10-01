package main

import "fmt"

func main() {
	// Declaración de Constantes

	const pi float64 = 3.1415;
	const pi2 = 3.141516
	
	fmt.Println("pi1:", pi)
	fmt.Println("pi2:", pi2)

	// Declaración de Variables

	base:= 12
	var altura int = 14
	var area int

	fmt.Println(base, altura, area)

	// Zero Values

	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a,b,c,d)

	// Area de un Cuadrado

	const baseCuadrado int = 10
    areaCuadrado:= baseCuadrado * baseCuadrado

	fmt.Println("Área Cuadraro:", areaCuadrado)
	
	fmt.Println("Hello World")
}
