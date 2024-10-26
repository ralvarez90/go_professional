package lessons

import (
	"fmt"
	"reflect"
	"unsafe"
)

// SLICES
//
// Son arreglos pero de tamaño dinámico. Por obvias razonas son menos eficientes
// en memoria. En la práctica se suelen utilizar los slices en lugar de los
// arrays.
//
// Podemos obtener sus slices a partir de otros slices. O slices a apartir de
// arreglos.
//
// Agregamos valores empleando la funcion append. Por otro lado, en el caso de que
// queramos obtener una porción de un array.
//
// Podemos crear slices de manera directa empleando la palabra función make.
func LS10Slices() {

	// uso de slices
	{
		var someIntegers []int
		fmt.Println("someIntegers:", someIntegers)
		fmt.Printf("Type of: %T, len: %d, capacity: %d\n", someIntegers, len(someIntegers), cap(someIntegers))
		for i := 1; i <= 10_000; i++ {
			if i%2 == 0 {
				someIntegers = append(someIntegers, i)
			}
		}

		fmt.Println("Adding elements...")
		fmt.Println("\tlen(someIntegers):", len(someIntegers))
		fmt.Println("\tcap(someIntegers):", cap(someIntegers))
		fmt.Println("\tSize in bytes    :", unsafe.Sizeof(someIntegers))
	}

	// obteniendo porcion de arreglo
	{
		// array
		arrIntegers := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		fmt.Println(arrIntegers)
		fmt.Println("arrIntegers address:", &arrIntegers)

		// slice
		mySlices := arrIntegers[2:4]
		fmt.Println("mySlices:", mySlices, "with type:", reflect.TypeOf(mySlices))
		fmt.Println("mySlices[0] address   :", &mySlices[0])
		fmt.Println("arrIntegers[2] address:", &arrIntegers[2])

		fmt.Println("mySlices:", mySlices)
		mySlices[0] = -1
		mySlices[1] = -1
		fmt.Println("mySlices:", mySlices)
		fmt.Println("arrIntegers:", arrIntegers)
	}

	// más sobre slicing
	{
		arrIntegers := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		fmt.Println("arrIntegers:", arrIntegers)
		fmt.Println("arrIntegers:", arrIntegers[1:])
	}

	// uso de make
	{
		someSlice := make([]int, 10)
		fmt.Println("someSlice:", someSlice)
	}

	// print separator
	fmt.Println()
}
