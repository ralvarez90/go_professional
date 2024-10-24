package homework

import (
	"bufio"
	"c01lenguaje/tools"
	"fmt"
	"os"
	"strings"
)

// TAREA 04
//
// Se van ingresando claves-valores y se almacenan en un mapa. El tipo de las claves
// es entero (guardados como string) y un string de descripción.
func Tarea04() {
	fmt.Println(tools.MakeTitle("Tarea 04"))

	catalog := map[string]string{
		"10": "notebook",
		"15": "tv",
		"21": "heladera",
		"27": "monitor",
		"35": "camara",
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Código: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "0" {
			fmt.Println("Adios!")
			break
		}

		value, ok := catalog[input]
		if ok {
			fmt.Println(">>>", value)
		} else {
			fmt.Println("No encontrado")
		}
	}

	// new line
	fmt.Println()
}
