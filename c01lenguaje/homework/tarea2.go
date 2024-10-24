package homework

import (
	"c01lenguaje/tools"
	"fmt"
)

// TAREA 2
//
// Recorrer el arreglo definido e incrementar cada uno de sus
// elementos en 20.
func Tarea02() {
	fmt.Println(tools.MakeTitle("Tarea 02"))
	array := [5]int{4, 2, 5, 6, 7}
	fmt.Println("\tInitial array:", array)

	// update values
	for i, _ := range array {
		array[i] += 20
	}

	// results
	fmt.Println("\tFinal array:", array)
	fmt.Println()
}
