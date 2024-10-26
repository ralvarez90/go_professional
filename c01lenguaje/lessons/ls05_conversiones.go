package lessons

import (
	"fmt"
	"reflect"
	"strconv"
)

// CONVERSIONES
//
// Podemos realizar conversiones entre distintos tipos de datos
// de diversa manera.  Desde usar el tipo de dato como función, el
// paquete strings y strconv.
func LS05Conversiones() {
	fmt.Println("LS05 - CONVERSIONES ENTRE TIPOS DE DATOS")

	// conversion entero a flotante
	{
		var someInt int32 = 123
		fmt.Printf("someInt: %d with type '%T'", someInt, someInt)
		someIntFloatVersion := float64(someInt)
		fmt.Println(someIntFloatVersion, reflect.TypeOf(someIntFloatVersion))
	}

	// numero a flotante
	{
		var someFloat = 123.123
		var strFloat string = fmt.Sprintf("%.2f", someFloat)
		fmt.Println("strFloat:", strFloat)
	}

	// string a numero, uso de paquete strconv
	{
		var strFloat = "123.123"
		someFloat, _ := strconv.ParseFloat(strFloat, 64)
		fmt.Printf("someFloat: %.2f, with type: %T\n", someFloat, someFloat)
		someInt, _ := strconv.ParseInt("123", 0, 64)
		fmt.Println("someInt:", someInt, "with type:", reflect.TypeOf(someInt))
	}

	// new line
	fmt.Println()
}
