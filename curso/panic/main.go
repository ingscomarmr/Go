package main

import (
	"errors"
	"fmt"
)

/*
* Para lanzar errores muy comunmente se ocupa la instruccion "panic" esta practicamente interrumpe el programa
* pero para manejar los "panic" tambien podemos utilizar una intruccion recovery para ejecutar otro tipo de
* instruccion y de esta forma hacer que no pare nuestro programa.
 */

func main() {
	//var input string
	//fmt.Scanf("%s", &input)
	var ope int
	var d1, d2 float32

	for {
		fmt.Println("\n\n####### OPERACIONES ########")
		fmt.Println("1.- Division")
		fmt.Println("2.- Suma")
		fmt.Println("3.- Resta")
		fmt.Println("4.- Multiplica")
		fmt.Println("5.- Salir")
		fmt.Print("Realiza una operacion:")
		fmt.Scan(&ope) //introduce el valor que le metemos desde consola

		if ope == 5 {
			break
		}

		fmt.Print("Introduce el primer operador:")
		fmt.Scanf("%f", &d1)

		fmt.Print("Introduce el segundo operador:")
		fmt.Scanf("%f", &d2)
		resultado, err := operacion(ope, d1, d2)

		if err != nil { //error controlado
			fmt.Println(err)
		} else {
			fmt.Println("El resultado de la operacion es:", resultado)
		}

	}

	fmt.Println("Adiosss!!")

}

func operacion(ope int, d1, d2 float32) (float32, error) {
	//recordemos que defer es una instruccion que indica ala funcion que debe ejecutar al final de todo el codigo o si ocurre un error
	defer func() { //por medio de esta funcion nostros recuperamos el control para que panic no truene todo el programa
		r := recover() //recobramos y trabajamos el error si es que existio
		if r != nil {
			fmt.Println("\n\n!!!Ocurrio un error:", r) //imprimomos el error para que lo vean, pero bien se puede realizar mas cosas
			fmt.Println()
		}
	}()

	if ope == 1 {
		return divide(d1, d2), nil //retornamos los dos valores el resultado y el error
	} else if ope == 2 {
		return suma(d1, d2), nil
	} else if ope == 3 {
		return (d1 - d2), nil
	} else if ope == 4 {
		return (d1 * d2), nil
	} else {
		return 0, errors.New("No existe esta operación!!")
	}

	panic("Operacion invalida")
}

func divide(dividendo, divisor float32) float32 {

	if divisor == 0 {
		panic("No se pude dividir por 0") //aqui pudieramos enviar un error pero decidimos enviar un panic
	}
	return dividendo / divisor
}

func suma(d1, d2 float32) float32 {
	return d1 + d2
}
