package tools

import (
	"fmt"
	"math/rand"
	"strings"
)

// Muetra mensaje en pantalla y espera una entrada por teclado.
func SystemPause() {
	fmt.Print("\nPRESS ANY KEY TO CONTINUE . . . ")
	fmt.Scanln()
}

// Gnera un aleatorio entero dentro del intervalo cerrado [min, max]
func RandInt(min, max int) int {
	return rand.Intn(max-min+1) + min
}

// Retorna un string con un formato específico. Se usa para generar un encabezado de una
// lección con el formato:
//
// ------------------------------ Nombre Lección ------------------------------
func MakeTitle(s string) string {
	return fmt.Sprintf("%s %s %s", strings.Repeat("-", 30), strings.TrimSpace(s), strings.Repeat("-", 30))
}

// Retorna un string con un formato específico. Se usa para generar un encabezado de una
// lección con el formato usando mayúsculas:
//
// ------------------------------ NOMBRE LECCIÓN ------------------------------
func MakeTitleUpper(s string) string {
	return fmt.Sprintf("%s %s %s", strings.Repeat("-", 30), strings.ToUpper(strings.TrimSpace(s)), strings.Repeat("-", 30))
}
