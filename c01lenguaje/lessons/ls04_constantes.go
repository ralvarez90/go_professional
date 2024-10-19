package lessons

import "fmt"

// CONSTANTES
//
// Se definene en tiempo de compilación y pueden tener scope de función o
// globales en un archivo.
//
// Empleamos la palabra resvervada const para definir una constante.
func LS04Constantes() {
	fmt.Println("LS04 - CONSTANTES")

	// define constant string
	const USERNAME = "RAlvarez90"
	fmt.Println("USERNAME:", USERNAME)

	// accedemos a constantes globales
	fmt.Println("UNO:", UNO)
	fmt.Println("UNO:", DOS)
	fmt.Println("UNO:", TRES)
	fmt.Println()
}

// Constanes definidas en una sola expresión
const (
	UNO  = 1
	DOS  = 2
	TRES = 3
)
