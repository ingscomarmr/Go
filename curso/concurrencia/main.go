package main

import (
	"fmt"
	"time"
)

/*
* LA CONCURRENCIA EN GO ES TAN SIMPLE DE LANZAR QUE SOLO SE UTILIZA go
* EN ESTE CASO TODAS LA FUNCIONES CON go SE EJECUTAN AL MISMO TIEMPO QUE LAS DEMAS LINIAS DEL PROGRAMA
* NO SON THREADS ES CONCURRENCIA YA QUE LAS ROTINIAS SE EJECUTAN DE MANERA INDEPENDIETE PERO NO ES UN PROCESO PARALELO
* LOS THREADS SE EJCUTAN EN BLOQUES INDEPENDIENTES, LAS GO RUTINAS PUDIERAN ASEMEJARSE A LAS FUNCIONES ASINCRONAS
* NO BLOEAN EL DEMAS CODIGO SI NO QUE SE PUEDE EJCUTAR LAS SIGUIENTES INSTRUCCIONES
 */

func main() {
	go mostrarNumeros() //le indicamos a go que queremos ejecutar esta funcion como una rutina

	/*var input string
	fmt.Print("Introduce un texto:")
	fmt.Scan(&input)
	fmt.Println("Introduciste esto:", input)*/
}

func mostrarNumeros() {
	for i := 1; i <= 10; i++ {
		fmt.Println("Numero:", i)
		time.Sleep(time.Second)
	}
}
