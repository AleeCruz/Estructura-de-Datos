package main

import (
	"fmt"
	mate "modulo_4/matematicas"
	"utils/areas"
)

func main() {
	per := areas.CalcularPerimetroDeUnRectangulo(3, 4)

	fmt.Println("\n\nHola Mundo")
	fmt.Println(per)

	fmt.Println(mate.AreaCirculo(3))
	fmt.Println(mate.Cuadrado(3))
	fmt.Println("Adios Mundo\n\n")

}
