package main

import (
	"fmt"
)

func main() {
	fmt.Println("####### FOR EVER ########")
	//no tiene condicion y se ejecuta por siempre
	//se usa para sockets u otros procesos que necesitan estar escuchandose por siempre.
	a := 10
	for {
		fmt.Printf("Ejecutando for ever %v \n", a)
		a--
		if a == 0 {
			break //la unica forma de parar un for ever es una instruccion break, de lo contrario sera infinito
		}
	}
}
