package interfaces

import (
	"fmt"
)

type Figuras2D interface {
	Area() float64
}

func CalcularArea (f Figuras2D){
	fmt.Println("Area: ", f.Area())
}
