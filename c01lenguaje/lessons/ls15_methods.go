package lessons

import (
	"c01lenguaje/tools"
	"encoding/json"
	"fmt"
	"strings"
)

// MÉTODOS
//
// Para crear un método en Go, se define una función con un receptor asociado
// a la estructura. El receptor es un parámetro especial que se coloca entre
// la palabra reservada func y el nomnbre del método.
//
// La sintaxis es:
//
//	func (receptor Tipo) nombreMetodo(parametros) retorno {
//		cuerpo...
//	}
//
// En go los métodos setters se concideran mala práctica.
//
// Notas:
// - Los receptores en métodos pueden ser de valor o de puntero.
func LS15Methods() {
	fmt.Println(tools.MakeTitle("MÉTODOS"))

	// ejemplo 01. creando isntancía de rectángulo
	{
		r1 := Rectangle{10, 20}
		fmt.Printf("\t%v\n", r1)
		fmt.Printf("\tRectangle area: %.2f u^2\n", r1.Area())
	}

	// ejemplo 02. creando instancia e invocando método
	{
		p1 := Product{ID: 1, Count: 10, Price: 14}
		p1.SetName("Soda")

		p1.Type = Type{ID: 3, Code: "CC-10", Description: "Soda 600 ml."}
		p1.AddTags("soda", "cola", "fresh", "sugar free")

		jsonByte, _ := json.Marshal(p1)
		fmt.Printf("\tProduct: %v\n", string(jsonByte))
		fmt.Printf("\tTotal: $%.2f\n", p1.TotalPrice())
	}

	// ejemplo 02. Array de productos
	{
		mercadoLibreCarrito := NewCar(123)
		mercadoLibreCarrito.ID = 1
		mercadoLibreCarrito.AddProducts(
			Product{ID: 1, Name: "Filwans Zero1", Type: Type{ID: 2, Code: "123-v", Description: "Audífionos HI-RES-1"}, Count: 10, Price: 760, Tags: []string{"Audífonos", "Technology"}},
			Product{ID: 2, Name: "Filwans Zero2", Type: Type{ID: 3, Code: "123-x", Description: "Audífionos HI-RES-2"}, Count: 20, Price: 780, Tags: []string{"Audífonos", "Technology"}},
			Product{ID: 3, Name: "Filwans Zero3", Type: Type{ID: 4, Code: "123-y", Description: "Audífionos HI-RES-3"}, Count: 30, Price: 790, Tags: []string{"Audífonos", "Technology"}},
			Product{ID: 4, Name: "Filwans Zero4", Type: Type{ID: 5, Code: "123-z", Description: "Audífionos HI-RES-5"}, Count: 40, Price: 800, Tags: []string{"Audífonos", "Technology"}},
		)

		fmt.Println("Productos: ", len(mercadoLibreCarrito.Products))
		mercadoLibreCarrito.ShowProducts()
		fmt.Printf("\tMercado libre: $%.2f\n", mercadoLibreCarrito.GetTotal())
	}

	// end lesson
	fmt.Println()
}

// definimos estructura
type Rectangle struct {
	ancho float64
	alto  float64
}

func (r Rectangle) Area() float64 {
	return r.alto * r.ancho
}

type Type struct {
	ID          uint   `json:"id"`
	Code        string `json:"code"`
	Description string `json:"description"`
}

type Product struct {
	ID    uint     `json:"id"`
	Name  string   `json:"name"`
	Type  Type     `json:"type"`
	Count uint64   `json:"count"`
	Price float64  `json:"price"`
	Tags  []string `json:"tags"`
}

func (p Product) TotalPrice() float64 {
	return p.Price * float64(p.Count)
}

func (p *Product) SetName(name string) {
	p.Name = strings.TrimSpace(name)
}

func (p *Product) AddTags(tags ...string) {
	p.Tags = append(p.Tags, tags...)
}

type Car struct {
	ID       uint      `json:"id"`
	UserID   uint      `json:"user_id"`
	Products []Product `json:"products"`
}

func (c *Car) AddProducts(products ...Product) {
	c.Products = append(c.Products, products...)
}

func (c Car) ShowProducts() {
	for _, product := range c.Products {
		fmt.Printf("\t- %v\n", product)
	}
}

func (c Car) GetTotal() float64 {
	var total float64 = 0
	for _, product := range c.Products {
		total += product.TotalPrice()
	}

	return total
}

func NewCar(userId uint) Car {
	return Car{ID: userId}
}
