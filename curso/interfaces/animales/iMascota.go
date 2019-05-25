package animales

import "fmt"

//Una interface se declara como una structura y en vez de contener propiedades contiene metodos
type IMascota interface {
	Comunicarse() //cualquier clase que este en el mismo paquete y contenga el mismo nombre metodo la interface lo adopta
}

type Gorilla struct {
	Nombre string
}

func (gorilla Gorilla) Comunicarse() {
	fmt.Println("Haaaaa Akkkk Uuuuukkk soy un Gorilla")
}

type Vaca struct {
	Nombre      string
	LecheListro int8
}

func (v Vaca) Comunicarse() {
	fmt.Println("Muuuu soy una Vaca")
}
