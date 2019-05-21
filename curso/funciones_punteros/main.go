package main

/*
* Pasar un valor por referencia muy importante en las funciones
* En Go el paso de parametos por referencia se realiza por medio de punteros
*
 */

import "fmt"

func main() {
	var a int32 = 4
	fmt.Println("El valor de la variable se encuenta en la direccion de memoria ", &a)
	fmt.Println("El valor original de A es ", a)
	pasoPorValor(a)                                                 //lo que se pasa es el valor de la variable y no la direccion de memoria y por lo tanto se le pasa una copia del valor
	fmt.Println("El valor de A depues de func pasoPorValor es ", a) //el valor sigue siendo el mismo pues no se afecto la direecion de memoria
	pasoPorRef(&a)                                                  //se le pasa el puntero, la direccion de memoria
	fmt.Println("El valor de A depues de func pasoPorRef es ", a)   //el valor sera el que tenga la direccion de memoria y como se modifico ese imprime
}

func pasoPorValor(a int32) {
	a *= 2
	fmt.Println("El valor de A en la funcion(pasoPorValor) es ", a)
}

func pasoPorRef(a *int32) { //para esperar una variable puntero se declara con *[tipo]
	*a = (*a) * 2                                                    //para manejar el valor de la variable y no el puntero se declara *[nombre_variable]
	fmt.Println("El valor de A en la funcion(pasoPorValor) es ", *a) //el valor cambiado se mantinene durante todo el programa porque se cambio directamente en dir de memoria
}
