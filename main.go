package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {	
	// Condicionales
	Valor1:= 1
	Valor2:= 2

	if(Valor1 == Valor2){
		fmt.Println("Son iguales")
	}else{
		fmt.Println("Son diferentes")
	}

	// Puerta logica and &&
	
	if Valor1 == 1 && Valor2 ==3{
		fmt.Println("Es correcto")
	}else{
		fmt.Println("Es incorrecto")
	}

	// Puerta logica or ||
	
	if Valor1 == 1 || Valor2 ==3{
		fmt.Println("Es correcto")
	}else{
		fmt.Println("Es incorrecto")
	}

	// Convertir texto a número

	value, error:= strconv.Atoi("ldjsnkjfdnsf")

	if error != nil{
		log.Fatal(error)
	}else{
		fmt.Println("Value: ", value)
	}

}
