package mypackage

import "fmt"

// CarPublic car con acceso public
type CarPublic struct{
	Brand string
	Year int
}

type carPrivate struct{
	brand string
	year int
}

// PrintMessage, imprime un mensaje
func PrintMessage(text string) {
	fmt.Println(text)
}