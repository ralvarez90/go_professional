package lessons

import (
	"c01lenguaje/tools"
	"errors"
	"fmt"
)

// MANEJO DE ERRORES
//
// Permiten controlar errores en tiempo de ejecución. El tipo de dato de los errores
// es error. Disponemos del package errors.
//
// Por otro lado, la función fmt.Errorf nos permite generar errores con mensajes
// epecíficos.
//
// En Go, defer es una palabra clave que se utiliza para programar la ejecución de
// una función o método justo antes de que la función que lo contiene finalice,
// sin importar si la función termina normalmente o debido a un error (panic).
//
// La función recover se emplea para recuperar el control de una goroutine que
// ha entrado en panic.
func LS18Errores() {
	fmt.Println(tools.MakeTitle("MANEJO DE ERRORES"))

	// función diferida
	defer func() {
		fmt.Println("\tEsto siempre se ejecuta, i.e. cuando LS18Errores finalize: exitosamente o no.")
		if r := recover(); r != nil {
			fmt.Printf("\tRecovered: %v\n", r)
		}
	}()

	// creacion de error
	{
		err := errors.New("my new error")
		fmt.Printf("\t%v\n", err.Error())
	}

	// error desde package fmt
	{
		err := fmt.Errorf("%s", "my format error")
		fmt.Printf("\t%v\n", err.Error())
	}

	{
		x := 4
		y := 0
		z := Div(float64(x), float64(y))
		fmt.Printf("\tResult x/y: %v\n", z)
	}

	// end lesson
	fmt.Println()
}

func Div(x, y float64) float64 {
	if y == 0 {
		panic("divisor can not be zero")
	}

	return x / y
}
