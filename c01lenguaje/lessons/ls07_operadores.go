package lessons

import "fmt"

// OPERADORES
//
// Son básicamente los mismos que otros lenguajes de programación. Existen los
// aritméticos, relacionales, booleanos, etc.
//
// 1. Relacionales
// Tenemos los clásicos <, <=, >, >=, != y ==.
//
// 2. Lógicos
// Son los clásicos &&, || y el de negación.
//
// 3. Ariméticos
// Son los clásicos +, -, *. /, %, = y las combinaciones de asignación.
func LS07Operadores() {

	// relacionales
	{
		yearsOld := 34
		if yearsOld >= 18 {
			fmt.Println("El usuario es mayor de edad:", yearsOld)
		}
	}

	// lógicos, ||, && y !
	{
		verdadero, falso := true, false
		fmt.Println("Falso || Verdadero    :", falso || verdadero)
		fmt.Println("Verdadero && Verdadero:", true && verdadero)
	}

	fmt.Println()
}
