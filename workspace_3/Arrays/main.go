package main

import "fmt"

func modificar(arreglo *[4]int) {
	arreglo[0] = 0
	arreglo[3] = 100
}

func main() {

	fmt.Println("----------Arrays--------")

	//Primero vamos a declarar una varibale de tipo array de enteros

	/*var arreglo [4]int = [4]int{1, 2, 3, 4}

	fmt.Println(arreglo)

	//Podemos acceder a cada uno de los elementos de un arreglos del siguiente modo

	fmt.Println(arreglo[0])
	fmt.Println(arreglo[2])

	fmt.Println(arreglo[3])

	//Ahora veremos otra manera  de declarar  arreglos es la siguiente

	arreglo_nuevo := [8]int{5, 6, 8}
	fmt.Println(arreglo_nuevo)

	fmt.Println("\n\nEl tamaño del arreglo es de: ")
	fmt.Println(len(arreglo_nuevo))

	//hay que tener cuidado con las declaraciones y asignanciones de arreglos
	/*Veremos ejemplos concretos de lo que no se debe de realizar por
	cierto motivos:

	var  arreglo [4]int = [5]int {1,2,3,4} go no te permite realizar esto

	veremos otro ejemplo de lo que podria llegar a pasar

	arreglo := [4]int {1,2,3,4}
	arreglo[4] = 5

	Este ultimo ejemplo tambien nos dara error porque el arreglo solo tiene posiciones
	del 0 1 2 3.
	La pi¿osicion 4 no existe asi que seria un error de sintaxis


	*/

	arreglo := [4]int{5, 6, 9, 2}
	fmt.Println(arreglo)
	modificar(&arreglo)
	fmt.Println(arreglo)

}
