package main

import "fmt"

/*********************************************************************************************************
*											DEFER
* Se utiliza para marcar las funciones que se ejecutaran al final de la ejecución de una instruccion
* estas funciones se ejecutan en el orden inverso en que se agregaron, sindo la primera la ultima
* que se ejecute.
*
* Otra cosa buena de defer es que se puede utilizar en cualquier funcion, osea otra función pude tener
* mas defer y ejecutarse una vez que la función termina de ejecutar la ultima instruccion,
* realizar un return o pasa a un caso de panico.
*********************************************************************************************************/

func main() {
	fmt.Println("Conectar a la base de datos...")
	defer closeConnection() //esta funcion se ejecuta al final de todas las demas porque fue la primera

	//si corru un caso de panico y no llega hasta esta instruccioón solo se ejecutan los defer antes declarados
	fmt.Println("Consultar datos de DataSet")
	defer closeDataSet() //esta funcion se ejecuta solo si no ocurrio algo antes y llego a entrar hasta esta instrucción
}

func closeConnection() {
	fmt.Println("	* Cerrar Conexion BD")
}

func closeDataSet() {
	fmt.Println("	* Cerrar DataSet")
}
