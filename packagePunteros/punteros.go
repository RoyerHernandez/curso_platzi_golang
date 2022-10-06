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

func (myPC Pc) String() string{
	return fmt.Sprintf("Tengo %d GB de Ram, %d de disco y mi soy de tipo %s", myPC.Ram, myPC.Disk, myPC.Brand)
}