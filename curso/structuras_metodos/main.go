package main

import "fmt"

//mi estructura

type Persona struct { //los metodos en las struct se definen afuera
	nombre string
	edad   int8
}

//defino un metodo que sea independiente de la estructura
func (p Persona) saludar() { //para indicar que es un metodo de Persona se antepone func ([nombre] [Tipo_Declarado]) [nombre_funcion]
	fmt.Println("Hola Soy metodo independiente pero de la estructura Persona")
}

//para realizar un siempre set o get es necesario pasar un puntero
func (p *Persona) setEdad(e int8) {
	if e >= 0 {
		p.edad = e
	} else {
		fmt.Println("Edad incorrecta")
	}
}

func (p *Persona) getEdad() int8 {
	return p.edad
}

func main() {
	var per Persona
	per.saludar()
	//per.setEdad(35)
	per.setEdad(-35) //para usar un metodo basta con acceder a el desde la instancia
	fmt.Printf("Edad de persona:%v\n", per.getEdad())

	var carro Carro
	carro.color = "Rojo"
	carro.alto = 1.5
	carro.largo = 4.35
	carro.getMedidas()
}

type Carro struct {
	color string
	largo float32
	alto  float32
}

func (c *Carro) getMedidas() {
	fmt.Printf("Carro de medidas:%vm X %vm\n", c.largo, c.alto)
}
