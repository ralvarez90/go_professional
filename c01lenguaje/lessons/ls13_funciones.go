package lessons

import (
	"c01lenguaje/tools"
	"errors"
	"fmt"
)

// FUNCIONES
//
// Son bloques de código reutilizables. Pueden o no regresar valores, y pueden o no
// recibir parámetros.
//
// La sintaxis es:
//
//	func NombreFuncion(p1 t1, p2, t2, ..., pn tn) (t1, t2, ..., tm) {
//			cuerpo
//	}
//
// donde p1, ..., pn son los parámetros con sus respectivo tipo tipo cada uno y t1, ..., tn
// son los tipos de retorno en caso de que retorno dichos elementos.
//
// En go las funciones son ciudadanos de primera clase, por lo se pueden tratar como
// cualquier otra variable.
//
// Go soporta el retorno múltiple.
//
// Notas:
// - Go maneja los errores retornarn un valor de tipo error.
func LS13Funciones() {
	fmt.Println(tools.MakeTitle("Funciones"))

	// fcs funcion
	{
		obtenerSuma := func(a, b int) int {
			return a + b
		}

		fmt.Println("\t1 + 2 =", obtenerSuma(1, 2))
	}

	// función con retorno múltiple
	{
		obtenerProducto := func(x, y int) (int, bool) {
			result := x * y
			return result, true
		}

		result, ok := obtenerProducto(5, 6)
		if ok {
			fmt.Println("\t5 + 6 =", result)
		}
	}

	// función con retorno múltiple y manejo de errores.
	{
		type Operation int
		const (
			SUMA Operation = iota
			RESTA
			MULTIPLICACION
			DIVISION
		)

		// función que recibe una operación, y dos flotantes
		calc := func(op Operation, x float64, y float64) (float64, error) {
			switch op {
			case SUMA:
				return x + y, nil
			case RESTA:
				return x - y, nil
			case DIVISION:
				if y == 0 {
					return 0, errors.New("y cannot be 0")
				} else {
					return x / y, nil
				}
			case MULTIPLICACION:
				return x * y, nil
			}

			return 0, errors.New("error al efectuar operación")
		}

		// invocamos función
		result, err := calc(SUMA, 1, 2)
		fmt.Println("\t1 + 2 =", result, ", err:", err)

		result, err = calc(RESTA, 1, 2)
		fmt.Println("\t1 - 2 =", result, ", err:", err)

		result, err = calc(MULTIPLICACION, 1, 2)
		fmt.Println("\t1 x 2 =", result, ", err:", err)

		result, err = calc(DIVISION, 1, 2)
		fmt.Println("\t1 / 2 =", result, ", err:", err)

		result, err = calc(DIVISION, 1, 0)
		fmt.Println("\t1 / 0 =", result, ", err:", err)

		_, err = calc(-1, 1, 0)
		if err != nil {
			fmt.Printf("\t%s\n", err)
		}
	}

	// Retorno de valores nombrados
	{
		split := func(n int) (x, y int) {
			x = 4*9 + n
			y = -x
			return
		}

		x, y := split(5)
		fmt.Println("\tx:", x)
		fmt.Println("\ty:", y)
	}

	// variadic functions
	{
		getSum := func(integers ...int) int {
			total := 0
			for _, value := range integers {
				total += value
			}
			return total
		}

		fmt.Println("\t1 + 2 + ... + 10 =", getSum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
	}

	// factory (función que regresa otra función)
	{
		type Operation int
		const (
			SUMA Operation = iota
			RESTA
			PRODUCTO
			DIV
		)

		factoryOperation := func(op Operation) func(float64, float64) float64 {
			switch op {
			case SUMA:
				return func(x, y float64) float64 { return x + y }
			case RESTA:
				return func(x, y float64) float64 { return x - y }
			case PRODUCTO:
				return func(x, y float64) float64 { return x * y }
			case DIV:
				return func(x, y float64) float64 {
					if y == 0 {
						panic("error: y cannot be 0")
					} else {
						return x / y
					}
				}
			}
			return nil
		}

		otherSum := factoryOperation(SUMA)
		fmt.Printf("\t%v\n", otherSum(1, 2))
	}

	// new line
	fmt.Println()
}
