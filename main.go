package main

import (
	pk "curso_platzi_golang/src/packagePunteros"
	"fmt"
)

type pc struct {
	ram   int
	disk  int
	brand string
}

func (myPc pc) ping() {
	fmt.Println(myPc.brand, "pong")
}

func (myPc *pc) duplicateRam() {
	myPc.ram = myPc.ram * 2
}

func main() {

	a := 50
	b := &a

	fmt.Println(b)
	fmt.Println(*b)

	*b = 100
	fmt.Println(a)

	var myPc pk.Pc
	myPc.Ram = 16
	myPc.Disk = 500
	myPc.Brand = "MSI"

	myPc.Ping()

	fmt.Println(myPc.String())

	fmt.Println(myPc)
	myPc.DuplicateRam()

	fmt.Println(myPc)
	myPc.DuplicateRam()

	fmt.Println(myPc)
	myPc.DuplicateRam()

	fmt.Println(myPc)
	myPc.DuplicateRam()
}
