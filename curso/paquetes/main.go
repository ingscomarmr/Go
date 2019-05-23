package main

import (
	"fmt"

	"../paquetes/utils"
)

func main() {
	var nombre string = "---OMAR MUNGUIA RIVERA??????"
	fmt.Println("String Input:", nombre)
	nombre = utils.GetLettersOnly(nombre) //ocupamos la funcion del paquete, solo tengo acceso a la funciones publicas (con letra mayuscula)
	//utils.IsNilOrEmpty(nombre) Esta si es accesible
	//utils.isLetter(32) esta no lo es
	fmt.Println("String Output:", nombre)

}
