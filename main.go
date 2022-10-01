package main

import "fmt"

func main() {	
	// Area de un Cuadrado

	const baseCuadrado int = 10
    areaCuadrado:= baseCuadrado * baseCuadrado

	fmt.Println("Área Cuadraro:", areaCuadrado)

	// Suma

	x:= 10
	y:= 20
	result:= x+y

	fmt.Println("Suma:", result)

	// Resta

	result =  y - x
	fmt.Println("Resta:", result)

	// Multiplicacion

	result = x * y
	fmt.Println("Multiplicación:", result)

	// División

	result = y/x
	fmt.Println("División:", result)

	//Modulo

	result = y % x
	fmt.Println("Modulo:", result)

	// Incremental

	x++
	fmt.Println("Incremental:", x)

	// Decremental

	x--
	fmt.Println("Decremental:", x)

	// Retos

	// Rectangulo base * altura

	baseRectangulo:= 10
	alturaRectangulo:= 40
	areaRectangulo:= baseRectangulo * alturaRectangulo

	fmt.Println("Área Rectangulo:", areaRectangulo)

	// Trapecio (Base + base) * altura / 2
	
	base:= 10
	Base:= 40
	altura:= 30

	areaTrapecio:= ((base+Base) * altura) / 2

	fmt.Println("Área Trapecio:", areaTrapecio)

	// Circulo pi * radio 2

	var radio float64= 6
	const pi float64 = 3.14159

	areaCirculo:= pi * (radio*radio)

	fmt.Println("Área Circulo:", areaCirculo)

}
