package main

import (
	"fmt"
	"strings"
)

func isPalindromo(texto string){

	minText := strings.ToLower(texto)

	textReverse:= ""

	for i:= len(minText) - 1; i >=0; i--{
		textReverse += string(minText[i])
	}

	if(textReverse == minText){
		fmt.Println("Es un Palindromo")
	}else{
		fmt.Println("No es un Palindromo")
	}

}

func ToLower(texto string) {
	panic("unimplemented")
}

func main() {

	isPalindromo("aMor a roma")

	slice:= []string{"Hola", "como", "estas"}

	for i, valor := range slice{
		fmt.Println(i,valor)
	}

}
