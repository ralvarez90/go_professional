package homework

import (
	"c01lenguaje/tools"
	"fmt"
)

// TAREA 1
//
// Se tendrán 2 variables definidas:
// - licence: indica si la persona tiene licencia
// - age: edad de la persona
func Tarea01() {
	fmt.Println(tools.MakeTitle("Tarea 01"))
	licence, age := true, 19
	if licence && age > 15 {
		fmt.Println("\tPuede seguir avanzando")
	} else if age <= 15 || !licence {
		fmt.Println("\tNo puede seguir circulando")
	}

	fmt.Println()
}
