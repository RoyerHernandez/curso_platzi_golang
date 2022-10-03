package main

import (
	//"fmt"
	"fmt"
	"log"
	"strconv"
)

func esPar(par string) int{
	value, error := strconv.Atoi(par)
	result  := 0
	if error!= nil{
		log.Fatal(error)
	}else{
		module := value % 2
		if module!=0{
		result = 1
		}else{
		result = 2
		}		
	} 
	return  result
}

func validaPassword(user string, pass int) int{
	usuario:= user
	contraseña:= pass
	respuesta := 0

	if(usuario == "Israfel2112" && contraseña == 123){
		respuesta = 1
	}
	return respuesta
}


func main() {	
	//Condicionales
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

	valor, error:= strconv.Atoi("89")

	if error != nil{
		log.Fatal(error)
	}else{
		fmt.Println("Value: ", valor)
	}

	esPar:= esPar("2")

	if esPar  == 1{
		fmt.Println("El número ingresado no es par", )
	}else{
		fmt.Println("El número ingresado es par")
	}
	
	validaPassword:= validaPassword("Israfel2112", 123)

	if validaPassword != 1{
		fmt.Println("El usuario no es valido")
	}else{
		fmt.Println("El usuario es valido")
	}
	// 

}
