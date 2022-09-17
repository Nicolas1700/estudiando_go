package mypackage

import "fmt"

//CarPublic, acceso publico
type CarPublic struct {
	Brand string
	Year  int
}

// PrintMessage, para imprimir un mensaje
func PrintMessage(text string) {
	fmt.Println(text)
}
