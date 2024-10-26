package main

import (
	"c01lenguaje/homework"
	"c01lenguaje/lessons"
	"c01lenguaje/tools"
)

func runLessons() {
	// show hello world
	lessons.LS01HelloWorld()

	// variables
	lessons.LS02Variables()

	// tipos de datos
	lessons.LS03TiposDeDatos()

	// constantes
	lessons.LS04Constantes()

	// conversiones de tipos de datos
	lessons.LS05Conversiones()

	// tipo de dato byte
	lessons.LS06Byte()

	// operadores
	lessons.LS07Operadores()

	// condicionales
	lessons.LS08Condicionales()

	// vectores
	lessons.LS09Vectores()

	// rebanadas
	lessons.LS10Slices()

	// maps
	lessons.LS11Maps()

	// bucle for
	lessons.LS12BucleFor()

	// funciones
	lessons.LS13Funciones()

	// structs
	lessons.LS14Structs()

	// métodos (structs)
	lessons.LS15Methods()

	// interfaces
	lessons.LS16Interfaces()

	// pointers
	lessons.LS17Punteros()

	// errores
	lessons.LS18Errores()

	// paquetes internos
	lessons.LS19PaquetesInternos()

	// paquetes internos, parte 2
	lessons.LS20PaquetesInternos()
}

func runHomeworks() {
	// homework.Tarea01()
	// homework.Tarea02()
	// homework.Tarea03()
	// homework.Tarea04()
	homework.Tarea05()
}

// Entry point
func main() {
	runLessons()
	tools.SystemPause()
}
