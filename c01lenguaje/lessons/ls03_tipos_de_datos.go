package lessons

import "fmt"

// TIPOS DE DATOS
func LS03TiposDeDatos() {

	fmt.Println("LS03 - TIPOS DE DATOS")
	tiposEnteros := "int, int8, int16, int32, int64"
	tiposEnterosSinSigno := "uint, uint8, uint16, uint32, uint64"
	tiposFlotantes := "float32, float64"
	tiposComplejos := "complex64, complex128"
	tiposBooleanos := "true, false"
	otrosTipos := "byte (alias de uint8), rune (alias de int32)"

	fmt.Println(tiposEnteros)
	fmt.Println(tiposEnterosSinSigno)
	fmt.Println(tiposFlotantes)
	fmt.Println(tiposComplejos)
	fmt.Println(tiposBooleanos)
	fmt.Println(otrosTipos)
	fmt.Println()
}
