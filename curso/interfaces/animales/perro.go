package animales

import "fmt"

type Perro struct {
	Nombre string
	Edad   int8
	Color  string
}

//Comunica al perro
func (p Perro) Comunicarse() { //para implementar una interface no es necesario agregar nada, ella reconoce los nombre de los metodos en automatico
	fmt.Println("Woouff soy un Perro")
}
