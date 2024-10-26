package lessons

import (
	"fmt"
	"math"
)

// TIPO DE DATO BYTE
//
// Este es un alias del tipo de dato uint8. Los enteros que podemos a abarcar en 8 bits son
// del 0 al 255.
//
// A una variable byte podemos asignarle un caracter.
func LS06Byte() {

	// uso de byte (alias uint8)
	{
		var someByteNum byte = 0
		fmt.Println("someByteNum:", someByteNum)

		// iteramos rango 0 - 255
		fmt.Println("Byte type data max value:", math.MaxUint8)
		for {
			someByteNum += 1
			if someByteNum == math.MaxUint8 {
				fmt.Println("someByteNum:", someByteNum)
				someByteNum++
				break
			}
		}

		fmt.Println("final someByteNum valiue:", someByteNum)
	}

	// use de byte usando %c
	{
		var counter byte = 0
		var lowerBound byte = 65
		var upperBound byte = 122
		for counter = lowerBound; counter < upperBound; counter++ {
			fmt.Printf("Character %c is in ascci code: %d\n", counter, counter)
		}

		var aCharacter byte = 'A'
		var bCharacter byte = 'B'
		var cCharacter byte = 66
		fmt.Println("aCharacter:", aCharacter, string(aCharacter))
		fmt.Println("bCharacter:", bCharacter, string(bCharacter))
		fmt.Println("cCharacter:", cCharacter, string(cCharacter))
	}

	// creando arreglo de byte
	{
		vector := []byte{65, 97, 82, 115}
		fmt.Println("vector:", vector)
		fmt.Println("string(vector):", string(vector))
	}

	fmt.Println()
}
