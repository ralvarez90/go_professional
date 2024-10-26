package homework

import (
	"c01lenguaje/tools"
	"fmt"
)

// TAREA 05
//
// Crear una función que se llame New que reciba una serie de valores
// de punto flotante
func Tarea05() {
	fmt.Println(tools.MakeTitle("TAREA 05"))

	// test 1
	m := NewMatrix([]float64{2, 7, 8}, []float64{4, 4, 7}, []float64{5, 7, 8})
	m.Print()
}

type Dimenson struct {
	Alto  uint
	Ancho uint
}

type Matrix struct {
	Dim      Dimenson
	Data     [][]float64
	IsSquare bool
}

func NewMatrix(rows ...[]float64) Matrix {
	if len(rows) == 0 {
		return Matrix{Data: [][]float64{}, Dim: Dimenson{Alto: 0, Ancho: 0}}
	}

	// matriz a regresar
	m := Matrix{Data: rows, IsSquare: true}
	firstSize := len(rows[0])

	for _, r := range rows {
		if len(r) != firstSize {
			m.IsSquare = false
		}
	}

	if m.IsSquare {
		m.Dim = Dimenson{Alto: uint(firstSize), Ancho: uint(firstSize)}
	}

	return m
}

func (m Matrix) Print() {
	for _, r := range m.Data {
		fmt.Println(r)
	}

	if m.IsSquare {
		fmt.Printf("%d X %d\n", m.Dim.Ancho, m.Dim.Alto)
	}

	fmt.Println("Cuadratic:", m.IsSquare)
}
