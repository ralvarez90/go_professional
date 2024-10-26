package lessons

import (
	"c01lenguaje/tools"
	"fmt"
	"math"
	"unsafe"
)

// VARIABLES
//
// Son identificadores de espacios de memoria donde se almacenarán datos. Cada
// variable tiene un tipo  de dato fijo durante la ejecución de un programa.
//
// Existen en Go diferentes tipos de datos que una variable puede tener, enteros,
// enteros sin signo, flotantes, strings, etc.
//
// Go puede inferir el tipo de dato el definirso o se puede especificar de manera
// explícita.
//
// Todas las variables tienen un valor default asignado. Para los punteros y las interfaces
// el valor default siempre será null.
//
// Como en C, empleamos el operador & para obtener la dirección de memoria de una
// variable.
//
// Podemos definir variables sin especificar el tipo y sin usar la palabra reservada
// var usando el operador de asignación :=.
//
// Notas:
// - Se usa byte para representar caracteres ascii y rune para caracteres unicode.
func LS02Variables() {
	fmt.Println(tools.MakeTitle("Variables"))

	var welcomeMsg string = "\tWelcome to Golang Language Programming!"
	fmt.Println(welcomeMsg, "wiht address:", &welcomeMsg)

	// show size some int64
	showIntegersSize()

	// other var
	myAge := 34
	myAge++
	fmt.Println("\tEl siguiente año cumplo", myAge, "años.")
	fmt.Println()
}

func showIntegersSize() {
	var numInt64 int64 = 0
	fmt.Printf("\tnumInt64: %d, with type %T, and size (bytes) %d\n", numInt64, numInt64, unsafe.Sizeof(numInt64))

	var maxUint uint64 = math.MaxUint64
	fmt.Println("\tmath.MaxUint64:", maxUint)
}
