package mypc

import "fmt"

// Struct Pc
type Pc struct {
	Ram   int
	Disk  int
	Brand string
}

// Fun/met struct ping
func (myPc Pc) Ping(){
	fmt.Println(myPc.Ram, "Pong")
}

// fun/met duplicar ram
func (myPc *Pc) DuplicateRam (){
	myPc.Ram = myPc.Ram * 2
}