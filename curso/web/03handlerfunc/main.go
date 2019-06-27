package main

import (
	"fmt"
	"log"
	"net/http"
)

//para implementar un handler sin necesidad de una structura solo se de de crar una funcion que reciba los
//parametros de un handler (w http.ResponseWriter, r *http.Request)
func miHandlerFunc(w http.ResponseWriter, r *http.Request) {
	msg := "<h1>HOLA DESDE HANDLER FUNC</h1>"
	fmt.Fprintf(w, msg)
}

func miHandlerFunc2(w http.ResponseWriter, r *http.Request) {
	msg := "<h1>HOLA DESDE HANDLER FUNC 2 CON GO.</h1>"
	fmt.Fprintf(w, msg)
}

//funcion que retorna un handler y recibe un mensaje
func messageFuncHandler(message string) http.Handler { //retonar un Handler
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, message)
		},
	)
}

func main() {

	//declaramos el servermux
	mux := http.NewServeMux()

	//poner la ruta del handler pero ahora usando handlerfunc
	mux.HandleFunc("/micontrolador", miHandlerFunc)                   //le pasamos la ruta y la funcion del handler
	mux.HandleFunc("/mic2", miHandlerFunc2)                           //le pasamos la ruta y la funcion del handler
	mux.Handle("/mic3", messageFuncHandler("<h1>MI HANDLER C3</h1>")) //como no es un haldlerfunc definido se consume con la otra funcion
	log.Println("Lanzar servidor local")
	log.Println(http.ListenAndServe(":8080", mux)) //lanzamos el servidor
}
