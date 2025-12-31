package main

import "fmt"

func main() {
	/*Vamos a hablar sobre las impresiones por pantalla
	en nuestra consolas tenemos 2 grandes candidatos el primero es
	el fmt.printf() y despues tenemos el fmt.Println()*/
	algo := "Algoritmos y Programacion"

	fmt.Println("Esto", "es", algo, 2)

	//Tambien tenemos el fmt.Printf
	fmt.Printf("Esto es %s %d\n", algo, 2)
}
