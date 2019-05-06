package main

import (
	"fmt"
)

func main() {
	//for continuo o for tipo while
	fmt.Println(" ########## CICLO FOR CONTINUO  ##########")
	a := 10 //tine una variable que serive como flag

	for a > 0 { //el bucle que valida el ciblo

		a--                               //el cambiador de flag
		fmt.Println("Valor de a:%v\n", a) //las instrucciones que se repiten
	}

	fmt.Println("Fin del ciclo")
}
