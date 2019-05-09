package main

import (
	"fmt"
)

func main() {
	diHola()
	saludar("Omar", 31) //pasar parametros a una funcion
}

//una funcion se declara con la palabra reservada func
func diHola() {
	//el contenido de la funcion esta aqui
	fmt.Println("Funcion Hola")
}

//funcion con parametros
func saludar(nombre string, edad uint8) {
	desEdad := edadPersonaDesc(edad)
	fmt.Printf("Hola %s no sabia que ya tienes %d años de edad, eres %s \n", nombre, edad, desEdad)
}

//cuando una funcion retorna algo se pone el tipo que retorna despues de los parentesis
func edadPersonaDesc(edad uint8) string {
	var tipoEdadDes string

	switch {
	case edad < 3:
		tipoEdadDes = "Bebe"
	case edad < 15:
		tipoEdadDes = "Jovencito"
	case edad < 60:
		tipoEdadDes = "Joven"
	case edad < 100:
		tipoEdadDes = "Anciano"
	default:
		tipoEdadDes = "Nada"
	}
	return tipoEdadDes
}

//si los parametros son del mismo tipo se puede acoortar la firma de la funcion
func suma(a, b int) int {
	return a + b
}
