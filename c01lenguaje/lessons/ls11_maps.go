package lessons

import (
	"fmt"
	"reflect"
)

// MAPS
//
// Son colecciones de pares clave valor, equivalentes a los diccionarios
// en Python.
//
// Para declarar un map usamos el tipo de dato map de la forma:
// map[tk]tv donde tk es el tipo de dato de los keys y tv es el
// tipo de dato de los valores.
//
// Para instanciarlo se emplea la palabra reservada make. Para eliminar
// elementos en su interior empleamos la función delete.
func LS11Maps() {

	// declaración/definición forma 1
	{
		m1 := make(map[int]string)
		m1[1] = "One"
		m1[2] = "Two"
		m1[3] = "Three"
		fmt.Println("m1:", m1, "with type:", reflect.TypeOf(m1))
	}

	// declaración/definición forma 2
	{
		var m2 map[int]string = map[int]string{
			1: "one",
			2: "two",
		}

		// update map
		m2[5] = "five"
		fmt.Println("m2 map:", m2)

		// delete if 1 key ok
		value, ok := m2[1]
		if ok {
			fmt.Printf("Existe m2[%d] -> %s\n", 1, value)
			fmt.Println("Deleting...")
			delete(m2, 1)
		}
	}

	// new line
	fmt.Println()
}
