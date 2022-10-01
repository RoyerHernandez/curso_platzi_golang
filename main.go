package main

import "fmt"

func normalFunction(message string){
	fmt.Println(message)
}

func tripleArcument(a, b int, c string){
	fmt.Println(a,b,c)
}

func returnValue (a int) int{
	return a * 2
}

func returnDoblevalue (a int) (c, d int){
	return a, a * 2
}

func areaRectangulo(base, altura int) int{
	return base * altura
}

func areaTrapecio(base, Base, altura int) int{
	return ((base + Base)*altura/2)
}

func areaCirculo(radio float64) float64{
	const pi float64 = 3.14159
	return pi * (radio*radio)
}

func main() {	
	// Funciones
	normalFunction("Hello World")
	tripleArcument(2, 4, "Hello")
	value:= returnValue(2)
	fmt.Println("Value: ",value)

	value1, value2 := returnDoblevalue(2)
	fmt.Println("value1:", value1)
	fmt.Println("value2:", value2)

	value3, _ := returnDoblevalue(3)
	fmt.Println("value3:", value3)

	_ , value4 := returnDoblevalue(4)
	fmt.Println("value3:", value4)

	// Reto

	// Área Rectangulo
	areaRectangulo := areaRectangulo(30, 10)
	fmt.Println("Área Rectangulo:", areaRectangulo)

	// Área Trapecio
	areaTrapecio := areaTrapecio(30, 50, 40)
	fmt.Println("Área Trapecio:", areaTrapecio)

	// Área Circulo
	areaCirculo := areaCirculo(6)
	fmt.Println("Área Circulo:", areaCirculo)
}
