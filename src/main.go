package main

import (
	"fmt"
	"log"
	"strconv"
)

/*
func funNormal() {
	fmt.Println("Hello word")
}

func funUnParametro(message string) {
	fmt.Println(message)
}

func funVariosParametros(a, b int, c string) {
	fmt.Println(a, b, c)
}

func retunrValue(a int) int {
	return a * 2
}

func doubleValue(a int) (b, c int) {
	return a, a * 2
}

// Funciones sobre encontrar el area de un rectangulo, tapecio y circulo
// Rectangulo
func areaRectangulo(baseRectangulo, alturaRectangulo float64) float64 {
	return baseRectangulo * alturaRectangulo
}

// Trapecio
func areaTrapecio(baseMenor, baseMayor, alturaTrapecio float64) float64 {
	return (baseMenor + baseMayor) * alturaTrapecio / 2
}

// Circulo
func areaCirculo(radio float64) float64 {
	const pi float64 = 3.141592
	return radio * radio * pi
}
*/

// Reto crear funciones
// Crear una fun que reciba un numero y diga si es par o impar
func numParImpar (num int) string{
	if num % 2 == 0  {
		return "Es par"
	}else{
		return "Es impar"
	}
}
// Crear fun que reciba un usu y pass, y diga si son validos o no para ing a la plataforma
func valUsuPass (usuName, pass string) bool {
	if usuName == "Nicolas" && pass == "1313u" {
		return true
	}else{
		return false
	}
}


func main() {

	/*
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
	*/

	// Uso paquete fmt
	/*
		helloMenssage := "Hello"
		worldMenssage := "world"

		//Println
		fmt.Println(helloMenssage, worldMenssage)
		fmt.Println(helloMenssage, worldMenssage)

		//Printf
		nombre := "Nicolas"
		edad := 18

		fmt.Printf("%s tiene %d años \n", nombre, edad)
		fmt.Printf("%v tiene %v años \n", nombre, edad)

		//Spintf
		menssage := fmt.Sprintf("%s tiene %d años.", nombre, edad)
		fmt.Println(menssage)

		// Conocer tipo de dato
		fmt.Printf("Tipo de nombre: %T \n", nombre)
		fmt.Printf("Tipo de edad: %T \n", edad)
	*/

	//Uso de funciones
	/*
		//Una funcion sin parametros
		funNormal()
		//funcion con un parametro
		funUnParametro("Hello world")
		//Funcion con varios parametros
		funVariosParametros(1, 2, "Hello")
		// Funcion con return
		value := retunrValue(2)
		fmt.Println(value)
		// Funcion con varios return
		value1, _ := doubleValue(2)
		fmt.Println("Value 1: ", value1)

		//Reto funciones de areas de figuras geometricas
		//Rectangulo
		areaRectangulo := areaRectangulo(5, 8)
		fmt.Println("Area Rectangulo: ", areaRectangulo)
		//Trapecio
		areaTrapecio := areaTrapecio(5, 6, 9)
		fmt.Println("Area Trapecio: ", areaTrapecio)
		//Ciculo
		areaCirculo := areaCirculo(5)
		fmt.Println("Area circulo: ", areaCirculo)
	*/

	//Ciclo for
	/*
			// For condicional
			for i := 0; i <= 10; i++{
				fmt.Println(i)
			}

			fmt.Printf("\n")

			// For while
			counter := 0
			for counter <= 10 {
				fmt.Println(counter)
				counter++
			}

			// For forever
		/*	counterForever := 0
			for {
				fmt.Println(counterForever)
				counterForever++
			}
	*/
	/*
		// For arreglo
		// For arreglo segun (cantidad del array)
		fmt.Printf("\n")
		nombresArray := [5]string{"Ana", "José", "Daniel", "María", "Carlos"}
		for i := 0; i < len(nombresArray); i++ {
			fmt.Println(nombresArray[i])
		}
		fmt.Printf("\n")

		//For Range (para obtener indice y valor)
		arreglo := []int{0, 1, 4, 6, 10, 9}
		fmt.Println("Arreglo:", arreglo)

		fmt.Println("Primer ejemplo")
		for i, j := range arreglo {
			fmt.Printf("indice i: %d tiene como valor #%d\n", i, j)
		}

	*/

	valor1 := 1
	valor2 := 2

	if valor1 == 1 {
		fmt.Println("Es 1")
	} else {
		fmt.Println("No es 1")
	}

	if valor1 == 1 && valor2 == 2 {
		fmt.Println("Se cumple")
	}

	if valor1 == 0 || valor2 == 2 {
		fmt.Println("Al menos 1 es verdadero")
	}

	// Convertir texto a número
	value, err := strconv.Atoi("58")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Value: ", value)

	// Uso funciones
	// Numero es par o impar
	numero := numParImpar(8)
	fmt.Println(numero)

	//Validad usuario con password
	ingreso := valUsuPass("Nicolas","1313u")
	fmt.Println(ingreso)


}
