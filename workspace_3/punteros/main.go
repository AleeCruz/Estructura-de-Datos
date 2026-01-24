package main

import "fmt"

func pasaUnDia(valor *int) {
	*valor = *valor * 2
}

func main() {
	fmt.Println("--------Punteros--------")
	//Puntero es una variable que permite almacenar una direccion de memoria

	//Por ejemplo vamos a declarar una variable
	/*var dir *int //Acabamos de crear una variable de tipo puntero

	var numero int = 5

	fmt.Println(numero)

	dir = &numero

	fmt.Println(*dir)
	*/

	//vamos a declarar una variable
	/*var valor int = 5           //Se declaro una variable de tipo int
	var direccion *int = &valor //Se le esta asignando la direccion de memoria a la variable direccion

	//Como usarlas??
	fmt.Println(valor) //Imprime el valor almacenado en la varibale numero

	fmt.Println(direccion) //imprime la direccion en la cual vive la variable numero

	//Vamos a ealizar modificaciones

	//Existen 2 manera
	valor++

	fmt.Println("Se le aumento el valor en uno directamente a valor ", valor)
	//Ahora se le va a aumentar indirectamente el valor en uno

	*direccion++

	fmt.Println("Se le aumento indirectamente el valor de lavariable valor ", valor)

	//Podriamos asignar el valor indirecta a una nueva variable
	valor_2 := *direccion

	fmt.Println("El valor de la nueva variable valor_2 es: ", valor_2)
	*/

	//Ahora vamos a ver los ejemplos de como usarlos en funciones

	var dolar int = 10
	fmt.Println("El dolar ahora vale: ", dolar)
	pasaUnDia(&dolar)
	fmt.Println("El dolar ahora vale: ", dolar)

}
