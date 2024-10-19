package lessons

import (
	"fmt"
	"unsafe"
)

func LS05Conversiones() {
	fmt.Println("LS05 - CONVERSIONES ENTRE TIPOS DE DATOS")
	conversionesEntreEnteros()
}

func conversionesEntreEnteros() {
	var someInt int32 = 123
	fmt.Println("someInt:", someInt)
	var someFloat = float64(someInt)
	fmt.Printf("someFloat: %.2f, with type %T, and size(bytes): %d\n", someFloat, someFloat, unsafe.Sizeof(someFloat))
}
