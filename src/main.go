package main

import "fmt"

func main() {

	//Declaracion de constantes
	const pi2 = 3.1415

	fmt.Println("pi2: ", pi2)

	//Declaracion de variables enteras
	base := 12
	var altura int = 10
	var area int

	fmt.Println(base, altura, area)

	// Zero value
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println("a: ", a)
	fmt.Println("b: ", b)
	fmt.Println("c: ", c)
	fmt.Println("d: ", d)

	var areaCuadrado = 10
	baseCuadrado := areaCuadrado * areaCuadrado

	fmt.Println("La base del cuadrado es: ", baseCuadrado)

	x := 10
	y := 50

	// Suma
	result := x + y
	fmt.Println("Suma: ", result)

	// Resta
	result = y - x
	fmt.Println("Resta: ", result)

	// Multiplicacion
	result = x * y
	fmt.Println("Multiplicacion: ", result)

	//División
	result = y / x
	fmt.Println("División: ", result)

	// Division ( modulo, ver el residuo)
	result = y % x
	fmt.Println("Residuo: ", result)

	// Incrementar
	x++
	fmt.Println("Incrementar: ", x)

	// Decrementar
	x--
	fmt.Println("Decrementar: ", x)

	// Reto
	// Calcular el area de un rectangulo, trapecio y circulo

	// Rectangulo
	baseRectangulo := 10
	alturaRectangulo := 50

	areaRectangulo := baseRectangulo * alturaRectangulo

	fmt.Println("Area rectangulo: ", areaRectangulo)

	// Trapecio

	baseMenor := 6
	baseMayor := 8
	alturaTrapecio := 4

	areaTrapecio := ((baseMenor + baseMayor) * alturaTrapecio) / 2

	fmt.Println("Area trapecio: ", areaTrapecio)

	// Circulo
	var radio float64 = 6
	const pi float64 = 3.141592
	areaCirculo := (radio * radio) * pi

	fmt.Println("Area circulo: ", areaCirculo)
}
