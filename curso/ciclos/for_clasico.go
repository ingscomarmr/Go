package main

import (
	"fmt"
)

func main() {
	fmt.Println(" ########## CICLO FOR CLASICO  ##########")

	for i := 1; i <= 5; i++ {
		fmt.Printf("%v,", i)
	}
	fmt.Println()

	matriz2D := [2][3]int{}

	var valor int
	//llenar matriz
	for c := 0; c < len(matriz2D); c++ {
		for r := 0; r < len(matriz2D[c]); r++ {
			valor++
			matriz2D[c][r] = valor
		}
	}

	//imprime matriz
	for c := 0; c < len(matriz2D); c++ {
		for r := 0; r < len(matriz2D[c]); r++ {
			fmt.Printf("Matriz[%v][%v] = %v\n", c, r, matriz2D[c][r])
		}
	}

	fmt.Println()
	//fmt.Println(matriz2D)

}
