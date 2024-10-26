package lessons

import (
	"c01lenguaje/tools"
	"fmt"
	"math"
)

// INTERFACES
//
// En Go, una interfaz es un tipo que define un conjunto de métodos. Es como un contrato que
// especifica qué comportamiento debe tener un tipo para cumplir con esa interfaz.
//
// No se define cómo se implementan esos métodos, solo se declaran sus firmas. Es común
// encontrar funciones que reciban interfaces vacías o genéricas, donde ya en el interior
// de dichas funciones se hace el manejo concreto.
//
// Estas nos permiten implementar polimorfismo en nuestras estructuras. Se definen igual
// con type y la palanra reservada interafacae.
func LS16Interfaces() {
	fmt.Println(tools.MakeTitle("INTERFACES"))

	// ejemplo 1
	{
		auto1, _ := NewVehicle(CarVehicle, 10)
		fmt.Println("\tDistancia recorrida:", auto1.Distance(), "Km ->", "Car")

		moto1, _ := NewVehicle(MotorcycleVehicle, 20)
		fmt.Println("\tDistancia recorrida:", moto1.Distance(), "Km ->", "Moto")

		truck1, _ := NewVehicle(TruckVehicle, 30)
		fmt.Println("\tDistancia recorrida:", truck1.Distance(), "Km ->", "Truck")
	}

	// ejemplo 2
	{
		t1 := Triangle{Base: 12, Altura: 23}
		c1 := Circle{Radius: 100}
		MostrarInfoFigura(t1)
		MostrarInfoFigura(c1)
	}

	// end application
	fmt.Println()
}

const (
	CarVehicle        = "CAR"
	MotorcycleVehicle = "MOTORCYCLE"
	TruckVehicle      = "TRUCK"
)

type Vehicle interface {
	Distance() float64
}

func NewVehicle(v string, time int) (Vehicle, error) {
	switch v {
	case CarVehicle:
		return &Automovil{Time: time}, nil
	case MotorcycleVehicle:
		return &Motorcycle{Time: time}, nil
	case TruckVehicle:
		return &Truck{Time: time}, nil
	}

	return nil, fmt.Errorf("vehicle: '%s' not exists", v)
}

type Automovil struct {
	Time int
}

func (a *Automovil) Distance() float64 {
	return 100 * (float64(a.Time) / 60)
}

type Motorcycle struct {
	Time int
}

func (m *Motorcycle) Distance() float64 {
	return 120 * (float64(m.Time) / 60)
}

type Truck struct {
	Time int
}

func (t *Truck) Distance() float64 {
	return 80 * (float64(t.Time) / 60)
}

type Geometric interface {
	Area() float64
	Perimetro() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimetro() float64 {
	return 2 * math.Pi * c.Radius
}

type Triangle struct {
	Base   float64
	Altura float64
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Altura) / 2
}

func (t Triangle) Perimetro() float64 {
	return t.Base + t.Altura + math.Hypot(t.Base, t.Altura)
}

func MostrarInfoFigura(f Geometric) {
	fmt.Printf("\tArea     : %.2f\n", f.Area())
	fmt.Printf("\tPerímetro: %.2f\n", f.Perimetro())
}
