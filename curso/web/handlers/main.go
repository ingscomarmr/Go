package main

import (
	"fmt"
	"log"
	"net/http"
)

//Estructura que sirvira para implementar un handler
type MsgHandler struct { //vamos a declarar una struct que se usara para heredar un handler
	message string
}

//Request = peticion del usuario, ResponseWriter envia la informacion
func (m MsgHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) { //imprmementamos el metodo de la interfaz
	fmt.Fprintf(w, m.message)
}

func main() {
	log.Println("Iniciar servidor web con handlers")
	mux := http.NewServeMux()

	//mi handler
	m1 := &MsgHandler{message: "<h1>Hola este es un mensaje</h1>"} //declaramos los controladores
	m2 := &MsgHandler{message: "<h1>Handler 2</h1>"}
	m3 := &MsgHandler{message: "<h1>Handler 3</h1>"}

	mux.Handle("/han1", m1) //los agregamos al multiplexor para que responsa con ellos
	mux.Handle("/han2", m2)
	mux.Handle("/han3", m3)

	log.Println("Lanzar server en local host")
	log.Println(http.ListenAndServe(":8080", mux)) //lanzamos el servidor
}
