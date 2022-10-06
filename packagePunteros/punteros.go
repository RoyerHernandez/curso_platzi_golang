package punteros

import "fmt"

//Pc crea un struct para pc
type Pc struct {
	Ram   int
	Disk  int
	Brand string
}


// Imprime el brand y el pong
func (myPc Pc) Ping() {
	fmt.Println(myPc.Brand, "Poong")
}

func (myPc *Pc) DuplicateRam(){
	myPc.Disk = myPc.Disk * 2
	myPc.Ram = myPc.Ram * 2
}