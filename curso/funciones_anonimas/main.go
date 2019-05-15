package main

import "fmt"

func main() {
	//Al igual que en javascript go tambien soporta funciones anonimas, y se pueden delcarar y ejcutar de la siguiente forma
	func() {
		fmt.Println("		FUNCIONES ANONIMAS")
	}()

	//Una func anomina contenida en una variable, me pueden madar a ejecutar en la variable misma
	printSeparador := func() {
		fmt.Println("---------------------------------------------------")
	}

	printSeparador()

	ope1 := operacion() //la funcion permanece viva durante la vigencia de la variable
	fmt.Println(ope1()) //la variable x permanecera viva todo el tiempo y aumentara
	fmt.Println(ope1())
	fmt.Println(ope1())

	printSeparador()

	ope2 := operacion() //Si declaramos otra instancia sera completamente diferente a la primera
	fmt.Println(ope2()) //su vida de la funcion ahora depende esta variable ope2
	fmt.Println(ope2())
	fmt.Println(ope2())
	fmt.Println(ope2())
	fmt.Println(ope2())

}

//esta es una funcion que retorna una funcion anonima la cual suma un valor ala variable x
func operacion() func() int32 {
	var x int32
	return func() int32 {
		x++
		return x
	}
}
