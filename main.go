package main

import (
pk "curso_platzi_golang/src/packageFiguras"
"fmt"
)
type figuras2D interface{
	Area() float64
	}

func calcular(f figuras2D){
	fmt.Println("El área es:", f.Area())
}

func main() {
	myCuadrado := pk.Cuadrado{Nombre:"Cuadrado", Base: 5}
	myRectangulo := pk.Rectangulo{Nombre: "Rectangulo", Base: 15, Altura: 26}
	myTriangulo := pk.Triangulo{Nombre: "Triangulo",Lado: 3}

	calcular(myCuadrado)
	calcular(myRectangulo)
	calcular(myTriangulo)
}
