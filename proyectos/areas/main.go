package main

import (
	"fmt"
	//Importaremos el paquete de areas de la siguiente manera
	"utils/areas"
)

func main() {
	per := areas.CalcularPerimetroRectangulo(3, 4)
	fmt.Println("Hola Mundo!")
	fmt.Println(per)

}
