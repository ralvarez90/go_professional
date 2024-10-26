package lessons

import (
	"c01lenguaje/tools"
	"fmt"
)

// CONDICIONALES
//
// Son un mecanismos por el cual podemos tomar desiciones en tiempo de
// ejecución. Son los clásicos bloqus if, else if y else.
//
// Otra alternativa al if, else if, else es el switch, que es muy
// similar a otros lenguajes como java o javascript.
func LS08Condicionales() {

	// bloque if, else if, else
	{
		someRandonInt := tools.RandInt(1, 100)
		fmt.Println("someRandomInt:", someRandonInt)
		if someRandonInt%2 == 0 {
			fmt.Println("El número es par")
		} else {
			fmt.Println("El número es impar")
		}
	}

	// bloque if con retorno de funcion
	{
		if result := fn(); result {
			fmt.Println("Booleano:", result)
		} else {
			fmt.Println("Booleano:", result)
		}

		aleatorio := tools.RandInt(-1, 1)
		if aleatorio < 0 {
			fmt.Println("Aleatorio:", aleatorio, ", es menor que 0")
		} else if aleatorio > 0 {
			fmt.Println("Aleatorio:", aleatorio, ", es mayor que 0")
		} else {
			fmt.Println("El aleatorio", aleatorio, "es 0")
		}
	}

	// switch
	{
		aleatorio := tools.RandInt(1, 10)
		fmt.Println("Aleatorio:", aleatorio)
		switch aleatorio {
		case 1:
			fmt.Println("Cota inferior:", aleatorio)
		case 10:
			fmt.Println("Cota superior:", aleatorio)
		default:
			fmt.Println("Interno:", aleatorio)
		}
	}

	fmt.Println()
}

// Utilidad interna
func fn() bool {
	return tools.RandInt(1, 2)%2 == 0
}
