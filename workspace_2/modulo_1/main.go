package main

import "fmt"

// Vamos a ver y usar un poco sobre lo que es la recursividad
// Veremos el caso de la recursividad con fibonacci
func main() {
	var resultado int = factorial(4)
	fmt.Println(resultado)

	var resultado_2 int = fibonacci(7)
	fmt.Println(resultado_2)
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}
func fibonacci(n int) int {
	if n == 0 {
		return 1
	} else if n == 1 {
		return 1
	} else {
		return fibonacci(n-1) + fibonacci(n-2)
	}
}
