package lessons

import (
	"c01lenguaje/tools"
	"fmt"
)

// TIPOS DE DATOS
//
// Go tiene un sistema estático de tipos, es decir, una vez que se establece
// un tipo a una variable este se mantiende durante la ejecución de todo el
// programa.
//
// Podemos definir tipos de datos mediante la palabra reservada type a partir
// de tipos de datos ya existente.
func LS03TiposDeDatos() {
	fmt.Println(tools.MakeTitle("Tipos de Datos"))

	tiposEnteros := "\tint, int8, int16, int32, int64"
	tiposEnterosSinSigno := "\tuint, uint8, uint16, uint32, uint64"
	tiposFlotantes := "\tfloat32, float64"
	tiposComplejos := "\tcomplex64, complex128"
	tiposBooleanos := "\ttrue, false"
	otrosTipos := "\tbyte (alias de uint8), rune (alias de int32)"

	fmt.Println(tiposEnteros)
	fmt.Println(tiposEnterosSinSigno)
	fmt.Println(tiposFlotantes)
	fmt.Println(tiposComplejos)
	fmt.Println(tiposBooleanos)
	fmt.Println(otrosTipos)
	fmt.Println()
}
