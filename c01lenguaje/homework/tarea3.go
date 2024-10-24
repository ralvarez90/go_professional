package homework

import (
	"bufio"
	"c01lenguaje/tools"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// TAREA 03
//
// Se van recibiendo enteros por consola y si se recibe un valor 0,
// entonces el programa finaliza.
//
// Los enteros leidos diferentes de cero se almacenan en un arreglo
// de enteros.
func Tarea03() {
	fmt.Println(tools.MakeTitle("Tarea 3"))

	reader := bufio.NewReader(os.Stdin)
	inputs := []int{}
	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		number, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Error: Ingrese un entero válido")
			continue
		}

		if number == 0 {
			break
		}

		// add elements
		inputs = append(inputs, number)
	}

	// resultados
	fmt.Println("Ingresados:", inputs)

	// new line
	fmt.Println()
}
