package main

import "fmt"

func insertionSort(A []int) {
	// Empezamos desde el segundo elemento (índice 1)
	// En el pseudocódigo es j = 2
	for j := 1; j < len(A); j++ {
		key := A[j] // El elemento actual a insertar

		// Empezamos a comparar con los elementos a la izquierda
		i := j - 1

		// Mientras el elemento a la izquierda sea mayor que la llave,
		// lo desplazamos a la derecha
		for i >= 0 && A[i] > key {
			A[i+1] = A[i]
			i = i - 1
		}

		// Colocamos la llave en su posición correcta
		A[i+1] = key
	}
}

func main() {

	arreglo := []int{4, 1, 5, 7, 8, 2, 1, 4}
	fmt.Println("Original:", arreglo)

	insertionSort(arreglo)

	fmt.Println("Ordenada:", arreglo)
}
