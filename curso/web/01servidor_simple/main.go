package main

import (
	"fmt"
	"net/http"
)

func main() {

	dirPublic := "www"
	dirHost := "localhost:8080"

	//forma tradicional pero limitada de lanzar un servidor
	//http.Handle("/", http.FileServer(http.Dir(dirPublic))) //agregar una carpeta que sera publica
	//fmt.Printf("Server lanzado:%s carpeta:%s\n", dirHost, dirPublic)
	//fmt.Println(http.ListenAndServe(dirHost, nil)) //lansar el servidor para que este escuchando

	//lanzar un servidor web por medio de un multiplexor, es mas dinamico y robusto que un servidor simple
	mux := http.NewServeMux()                  //declaramos el multiplexo
	fs := http.FileServer(http.Dir(dirPublic)) //declaramos la ruta de donde servira
	mux.Handle("/", fs)                        //indicamos si contiene alguna ruta web especifica
	fmt.Println("Iniciar servicio web en ", dirPublic)
	fmt.Println(http.ListenAndServe(dirHost, mux)) //iniciamos el server

}
