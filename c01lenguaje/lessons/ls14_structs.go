package lessons

import (
	"c01lenguaje/tools"
	"encoding/json"
	"fmt"
)

// STRUCTS
//
// Permiten crear tipos de datos propios. Se definen empleando la palabra
// reservada struct y en su interior sus campos.
//
// Básicamaente son colecciones de campos asociados a un nombre. La sintaxis
// es:
//
//	type NombreStruct struct {
//			field1 dataType1
//			field2 dataType2
//			...
//	}
//
// Los structs pueden tener como campos otras estructuras, arrays, slices,
// y cualquier otro tipo de dato.
//
// En go, los campos dentro de un struct puede contener una etiqueta, que
// son metadatos adjuntos que proporcionan información adicional sobre
// como se deben de usar esos campos.
//
// El uso más frecuente de las etiquetas en campos es para controlar la
// serialización y deserialización en formatos como json o xml.
//
// Con el api de reflection podemos crearn nuestros propios tags personalizados,
// así como la etiqueta json.
//
// Para convertir un struct en json usamos la funciín Marshall. Esta función
// recibe un objeto de tipo any (interface vacía)
func LS14Structs() {
	fmt.Println(tools.MakeTitle("Structs"))

	// ejemplo de estructura Product
	{
		type Type struct {
			ID          uint
			Code        string
			Description string
		}

		type Product struct {
			ID    uint
			Name  string
			Type  Type
			Price float32
			Tags  []string
		}

		// declaración forma 1
		var p1 Product
		p1.ID = 1
		p1.Name = "Samsung S24+"
		p1.Price = 123.34
		p1.Type = Type{ID: 1, Code: "T1C", Description: "Telefonía"}
		fmt.Printf("\t%v\n", p1)

		// declaración forma 2
		p2 := Product{}
		p2.ID = 2
		p2.Name = "Huawei P40"
		p2.Price = 6600
		p2.Type = Type{ID: 1, Code: "T1C", Description: "Telefonía Movil"}
		fmt.Printf("\t%v\n", p2)

		// declaración forma 3
		p3 := Product{3, "Motorola Neo 4", Type{ID: 1, Code: "T1D", Description: "Telefonía IP"}, 4500, []string{"Electrónica", "Celulares"}}
		fmt.Printf("\t%v\n", p3)

		// declaración forma 4
		p4 := Product{ID: 4, Price: 8000, Name: "Iphone 12", Type: Type{ID: 1, Code: "T5", Description: "Telefonía"}}
		p4.Tags = []string{"Electrónica", "Celulares", "Apple"}
		fmt.Printf("\t%v\n", p4)
	}

	// struct con metadatos
	{
		type Address struct {
			Calle   string `json:"calle"`
			Numero  uint   `json:"numero"`
			Colonia string `json:"colonia"`
		}

		type Person struct {
			Name    string  `json:"name"`
			Age     uint    `json:"age"`
			Address Address `json:"address"`
		}

		person1 := Person{Name: "John Wick", Age: 45}
		person1.Address.Calle = "Fray Servando T. de Mier"
		person1.Address.Numero = 787
		person1.Address.Colonia = "Jardín Balbuena"
		fmt.Printf("\t%v\n", person1)

		jsonBytes, err := json.Marshal(person1)
		if err != nil {
			panic("\terror al convertir json:")
		}

		// fmt.Printf("\t%v\n", jsonBytes)
		fmt.Printf("\t%v\n", string(jsonBytes))
	}

	// end lesson
	fmt.Println()
}
