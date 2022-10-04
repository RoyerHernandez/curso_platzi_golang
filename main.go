package main

import "fmt"

func isPalindromo(texto string){

	textReverse:= ""

	for i:= len(texto) - 1; i >=0; i--{
		textReverse += string(texto[i])
	}

	if(textReverse == texto){
		fmt.Println("Es un Palindromo")
	}else{
		fmt.Println("No es un Palindromo")
	}

}

func main() {

	isPalindromo("amor al roma")

	slice:= []string{"Hola", "como", "estas"}

	for i, valor := range slice{
		fmt.Println(i,valor)
	}

}
