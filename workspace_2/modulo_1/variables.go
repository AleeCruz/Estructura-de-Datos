package main

import "fmt"

/*Vamos a estudiar sobre las variables en go desde cero*/
func main() {
	// se usar the kayword var + nombre + el tipo de dato con el cual estamos t
	//estamos trabajando
	var numero int
	numero = 25
	fmt.Println(numero)
	numero = 40

	fmt.Println(numero)

	//Vamos a declarar y asignar un dato de tipo string a la variable nombre
	//Veamos como hacerlo
	var nombre string
	nombre = "Alexander"
	//Esta es otra forma de declarar y asignar un string a una vrible de otra
	//forma un tanto especial
	fmt.Println(nombre)
	apellido := "Cruz Apaza"
	fmt.Println(apellido)
	/*Podemos realizr una declaracion y asignacio multiple veamos como */

	name, surname, age := "Alexander", "Cruz", 26

	fmt.Println(name, surname, age)

	//ALgo interesante que podemos hacer es un swap en go que nos permite
	//intercambiar los valores de las variables

	name_1, name_2 := "Maria", "Antonieta"
	fmt.Println(name_1, name_2)

	name_2, name_1 = name_1, name_2
	fmt.Println(name_1)
	fmt.Println(name_2)

	//Importante todas las variables que uno esta creando debe de inicializarle

	//Algo importante sobre  go es la siguiente linea
	/*Es importante notar que aunque nosotros no estemos asignando un valor
	a la variable de numero_dni esta por defecto esta siendo asignada con
	el valor cero por default*/
	var numero_dni int
	fmt.Println(numero_dni)
}
