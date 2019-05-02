package main

import "fmt"

// La estrucutra Persona contiene diferentes actributos, es to es similar a una clase en C#
type Persona struct {
	Nombre    string
	Edad      uint
	Sexo      string
	Atributos []string
}

func main() {

	//de esta forma se declara una instancia de Persona
	var pedro Persona
	pedro.Nombre = "Pedro Perez"
	pedro.Edad = 21
	pedro.Sexo = "M"
	pedro.Atributos = []string{"Moreno", "Con lunar en cuello", "Pelo castaño"}

	fmt.Println(pedro)

	//inicializacion rapida
	luis := Persona{
		Nombre:    "Luis Pablo Munguia",
		Edad:      7,
		Sexo:      "M",
		Atributos: []string{"Blanco", "Sabiendo", "Inquieto"},
	}

	fmt.Println(luis)

	//inicializar mas cortar
	juan := Persona{
		"Juan Munguia",
		24,
		"M",
		[]string{"Blanco", "Sabiendo"},
	}

	fmt.Println(juan)
}
