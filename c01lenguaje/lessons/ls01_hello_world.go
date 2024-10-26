package lessons

import (
	"c01lenguaje/tools"
	"fmt"
)

// Hello World
//
// Usamos la función fmt.Println para mostrar mensajes en la salida
// standard.
func LS01HelloWorld() {
	fmt.Println(tools.MakeTitle("HELLO WORLD"))
	fmt.Println("\tHello World in Go!")
	fmt.Println()
}
