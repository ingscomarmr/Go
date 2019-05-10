package main

import (
	"fmt"
)

func main() {
	//diHola()
	//saludar("Omar", 31) //pasar parametros a una funcion

	nums := []int{100, 500, -9, 56, 24}
	max, min := maxYmin(nums)
	fmt.Printf("Numero Maximo:%d, Minimo:%d\n", max, min)

	max, min = maxYmin2(nums)
	fmt.Printf("Numero Maximo:%d, Minimo:%d\n", max, min)
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

//funcion de retorno multiple
func maxYmin(numeros []int) (int, int) { //indicamos los tipos que va retornar (int, int)
	var max, min int

	for _, v := range numeros {
		if v > max {
			max = v
		}

		if v < min {
			min = v
		}
	}

	return max, min //indicamos el orden en el que se debe retornar
}

//funcion de retorno multiple version mas corta y mas definida
func maxYmin2(numeros []int) (max int, min int) { //indicamos los tipos y nombre que va retornar (int, int)
	for _, v := range numeros {
		if v > max {
			max = v
		}

		if v < min {
			min = v
		}
	}

	return //ya no necesita return porque ya le indicamos como va retornar y que orden
}
