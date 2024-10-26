package lessons

import (
	"c01lenguaje/tools"
	"fmt"
)

// PUNTEROS
//
// Son variables que almacenan direcciones de memoria. Pueden almacenar direcciones
// de memoria de otras variables.
//
// Los punteros al ser variables también tienen su propia dirección de memoria. Para
// obtener la dirección de memoria de una variable empleamos el &.
//
// Usamos el caracter *nombrePuntero para acceder al valor al que apunta su dirección
// de memoria almacenada.
//
// Al igual que cualquier variable apuntador, si tengo un apuntador a un struct, por
// medio *ptrTosTruct accedemos a la estructura a la que apunta.
func LS17Punteros() {
	fmt.Println(tools.MakeTitle("PUNTEROS"))

	// apuntadores
	{
		// variable entera
		var someInt int = 123
		fmt.Printf("\tVariable someInt con dirección %p, almacena %v\n", &someInt, someInt)

		// acceso a variable entera desde apuntador
		ptrSomeInt := &someInt
		fmt.Printf("\tVariable someInt con dirección %p, almacena %v\n", ptrSomeInt, *ptrSomeInt)

		// mostramos dirección de memoria de apuntador
		fmt.Printf("\tDirección de apuntador ptrSomeIn: %p\n", &ptrSomeInt)

		// función que recibe apuntador
		increment := func(ptrInt *int) {
			*ptrInt++
		}

		increment(ptrSomeInt)
		fmt.Printf("\tVariable someInt con dirección %p, almacena %v\n", ptrSomeInt, *ptrSomeInt)
		increment(&someInt)
		fmt.Printf("\tVariable someInt con dirección %p, almacena %v\n", &someInt, someInt)
	}

	// ejemplo con slices
	{
		someSlice := []int{1, 2, 3}

		resetSlice := func(slice []int) {
			fmt.Printf("\tslice: %p\n", slice)
			if len(slice) != 0 {
				fmt.Printf("\tslice[0] -> %v, with address: %p\n", slice[0], &slice[0])
			}

			for i := range slice {
				slice[i] = 0
			}
		}

		resetSlice(someSlice)
		fmt.Printf("\t%v\n", someSlice)
	}

	// aputadores a estructuras
	{
		// puntero a estructura
		ptrMyStruct := &MyStruct{ID: 123, Name: "Test"}
		fmt.Printf("\tmyStruct: %v\n", ptrMyStruct)

		// cambio de estado desde apuntador
		ptrMyStruct.ID = 1
		ptrMyStruct.Name = "test"
		fmt.Printf("\t%v\n", ptrMyStruct)
		fmt.Printf("\t%v\n", *ptrMyStruct)

		// invoke metood
		updateMyStruct(ptrMyStruct)

		// otra estructura
		otherStruct := MyStruct{ID: 0}
		otherStruct.Set("John Wick 1")
		fmt.Printf("\totherStruct: %v\n", otherStruct)

		otherStruct.SetP("John Wick 2")
		fmt.Printf("\totherStruct: %v\n", otherStruct)
	}

	// end application
	fmt.Println()
}

type MyStruct struct {
	ID   uint
	Name string
}

func (m MyStruct) Set(name string) {
	fmt.Printf("\tAddress: %p\n", &m)
	m.Name = name
}

func (m *MyStruct) SetP(name string) {
	fmt.Printf("\tAddress: %p\n", m)
	m.Name = name
}

func updateMyStruct(v *MyStruct) {
	fmt.Printf("\tAddress in function: %p\v", v)
	v.ID = 0
	v.Name = "empty"
	fmt.Printf("\t%v\n", *v)
}
