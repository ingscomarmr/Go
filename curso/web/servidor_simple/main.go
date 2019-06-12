package main

import (
	"fmt"
	"net/http"
)

func main() {

	dirPublic := "www"
	dirHost := "localhost:8080"
	http.Handle("/", http.FileServer(http.Dir(dirPublic)))
	fmt.Printf("Server lanzado:%s carpeta:%s\n", dirHost, dirPublic)
	fmt.Println(http.ListenAndServe(dirHost, nil))

}
