package main

import (
	"errors"
	"fmt"
)

func main() {
	var dividendo, divisor float32

	dividendo = 100
	divisor = 0

	fmt.Printf("Divide %v entre %v\n", dividendo, divisor)

	r, err := division(dividendo, divisor)

	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	fmt.Println("Resultado:", r)

}

func division(dividendo, divisor float32) (resultado float32, err error) { //error es un tipo de dato que se declara como tal

	if divisor == 0 {
		err = errors.New("No se puede dividir por 0") //con errors podemos declarar el error
	} else {
		resultado = dividendo / divisor
	}
	return //retornara el valor por defecto de err(nil) y de resultado(0) si no se le asigno nada a la variable
}
