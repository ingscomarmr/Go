package main

import (
	"html/template"
	"log"
	"net/http"
)

//Bean para guardar datos de una persona
type Persona struct {
	Nombre string
	Edad   uint8
}

//creamos el controlador o handler con handlerfunc
func renderIndex(w http.ResponseWriter, r *http.Request) {

	miPersona := &Persona{"Omar Munguia", 31}           //objeto de datos
	t, err := template.ParseFiles("./views/index.html") //le pasamos el template que va renderizar
	if err != nil {                                     //si eexiste un error al leer el template entonces mandamos un error
		http.Error(w, err.Error(), http.StatusInternalServerError) //error 500 que es un insertnal error
	}
	t.Execute(w, miPersona) //si todo bien mandamos a renderizar
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", renderIndex) //poniendo el controlador

	log.Println("Lanzar servidor localhost:8080/")
	log.Println(http.ListenAndServe(":8080", mux)) //lanzamos el servidor

}
