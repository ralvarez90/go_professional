package lessons

import (
	"fmt"
	"reflect"
	"unsafe"
)

// VECTORES (arrays)
//
// Son colecciones homogeneas de datos. Estos son de tamaño fijo.
func LS09Vectores() {

	// vector de enteros
	{
		// definición de vector
		someInts := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		fmt.Println("someInts:", someInts)
		fmt.Println("Type of someInts:", reflect.TypeOf(someInts))
		fmt.Printf("someInts: %v\nType of someInts: %T\n", someInts, someInts)
		fmt.Println("len(someInts):", len(someInts))
		fmt.Println("cap(someInts):", cap(someInts))
	}

	// declaración
	{
		var someInts [10]int
		fmt.Println("Again someInts:", someInts)
		for i := 0; i < len(someInts); i++ {
			someInts[i] = i + 1
		}
		fmt.Println("Again someInts:", someInts)
		fmt.Printf("someInts: %v, with type: %T, length: %d and byte size: %d\n", someInts, someInts, len(someInts), unsafe.Sizeof(someInts))
	}

	// print separator
	fmt.Println()
}
