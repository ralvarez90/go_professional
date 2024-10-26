package lessons

import (
	"c01lenguaje/tools"
	"fmt"
	"reflect"
)

// BUCLE FOR
//
// Go tiene como único bucle el for, que se puede usar como el for
// tradicional, como bucle while y do-while.
//
// El for nos sirve de utilizad para recorrer colecciones de manera
// similar a python.
func LS12BucleFor() {
	fmt.Println(tools.MakeTitle("Bucle For"))

	// for como while
	{
		fmt.Println("For como while")
		for {
			fmt.Println("\tDentro de for...")
			break
		}
	}

	// for tradicional
	{
		fmt.Println("For tradicional")
		sum := 0
		for i := 1; i <= 10_000; i++ {
			sum += i
		}
		fmt.Println("\t1+2+...+10000 =", sum)
	}

	// for para iterar arrays/slices
	{
		fmt.Println("For para iterar arrays/slices")
		someNumbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		fmt.Println("\tsomeNumbers:", someNumbers, ", with type:", reflect.TypeOf(someNumbers))
		for i, v := range someNumbers {
			fmt.Printf("\tsomeNumbers[%d] -> %d\n", i, v)
		}
	}

	// for para iterar maps
	{
		fmt.Println("For para iterar mapas")
		someMap := map[int]string{
			1: "uno",
			2: "dos",
			3: "tres",
		}

		fmt.Println("\tsomeMap:", someMap)
		for k, v := range someMap {
			fmt.Printf("\tsomeMap[%v] -> %v\n", k, v)
		}

		otherMap := map[string][]int{
			"uno":  nil,
			"dos":  {2, 24, 24, 53, 2},
			"tres": {1, 2, 4, 55},
		}

		fmt.Println("\ttotherMap:", otherMap, ", with type:", reflect.TypeOf(otherMap))
		for _, arr := range otherMap {
			fmt.Printf("\t%v\n", arr)
		}
	}

	// new line
	fmt.Println()
}
