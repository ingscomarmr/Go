package animales

import "fmt"

//Structura de tipo que representa a un gato
type Gato struct {
	Nombre string
	Edad   int8
	Color  string
}

//Metodo que comunica al gato y que aparte se une con la interface IMascota, esto debido a que son del mismo paquete
func (g Gato) Comunicarse() { //para implementar una interface no es necesario agregar nada, ella reconoce los nombre de los metodos en automatico
	fmt.Println("Miauuuu soy un Gato.")
}

/* Metodos con punteros no son tomados por las interfaces
func (g *Gato) Comunicarse() {
	fmt.Println("Miauuuu soy un Gato:", g.Nombre)
}

*/
