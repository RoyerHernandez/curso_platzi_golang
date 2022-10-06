package main

import (
	pk "curso_platzi_golang/src/mypackage"
	"fmt"
)

func main() {

	var myCar = pk.CarPublic{}
	myCar.Brand = "Ferrari"
	myCar.Year = 2022
	fmt.Println(myCar)

	pk.PrintMessage("Hola Mundo")

}
