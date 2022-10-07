package main

import "fmt"

type figuras2D interface{
	area() float64
}

type cuadrado struct {
	base float64
}

type rectcangulo struct{
	base float64
	altura float64
}

type triangulo struct{
	lado float64
}

func (c cuadrado) area() float64{
	return c.base * c.base
}

func (r rectcangulo) area() float64{
	return r.base * r.altura
}

func (t triangulo) area() float64{
	return t.lado * t.lado * t.lado
}

func calcular(f figuras2D){
	fmt.Println("Area :", f.area())
}

func main() {
	miCuadrado := cuadrado{base: 2}
	miRectangulo := rectcangulo{base: 4, altura: 2}
	miTriangulo := triangulo{lado: 3}

	// Lista de interfaces
	miInterface := []interface{}{"Hola", 23, 32.3, "Hello"}

	calcular(miCuadrado)
	calcular(miRectangulo)
	calcular(miTriangulo)

	fmt.Println(miInterface...)
}
