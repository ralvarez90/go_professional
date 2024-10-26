package lessons

import (
	"c01lenguaje/tools"
	"fmt"
)

// CONSTANTES
//
// Se definene en tiempo de compilación y pueden tener scope de función o
// globales en un archivo.
//
// Empleamos la palabra resvervada const para definir una constante.
func LS04Constantes() {
	fmt.Println(tools.MakeTitle("Constantes"))

	// define constant string
	const USERNAME string = "RAlvarez90"
	const PASSWORD = "Leslie123asd"
	fmt.Println("USERNAME:", USERNAME)
	fmt.Println("PASSWORD:", PASSWORD)

	// accedemos a constantes globales
	fmt.Println("\tUNO:", UNO)
	fmt.Println("\tUNO:", DOS)
	fmt.Println("\tUNO:", TRES)
	fmt.Println()
}

// Constanes definidas en una sola expresión
const (
	UNO  = 1
	DOS  = 2
	TRES = 3
)
