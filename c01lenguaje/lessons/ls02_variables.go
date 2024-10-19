package lessons

import "fmt"

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
// Notas:
// - Se usa byte para representar caracteres ascii y rune para caracteres unicode.
// - Si usamos fmt.Printf empleamos el especificador %p para obtener la dirección
// de memporia de una variable.
func LS02Variables() {
	// string ejemplo
	fmt.Println("LS02 VARIABLES")
	var welcomeMsg string = "Welcome to Golang Language Programming!"
	fmt.Println(welcomeMsg)
	fmt.Println()
}
