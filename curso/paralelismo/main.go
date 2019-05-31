package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	var nPro int
	fmt.Print("Cuantos procesadores quieres usar:")
	fmt.Scan(&nPro)
	//si queremos aumentar mas aun la potencia podemos decirle al programa que use mas de un procesador
	runtime.GOMAXPROCS(nPro)

	fmt.Println()

	//para esperar a las go rutinas agregamos este wait group
	var wg sync.WaitGroup

	t := time.Now()
	fmt.Println("Inicia el proceso:", t)

	numbers := []uint32{1111222232, 7888122, 99912, 111, 44433312, 889981231, 34341212, 244, 71233, 899912, 331222, 888812, 999999, 11111, 66666666, 888812812, 9991222, 888812, 99912221, 888123311, 111222}
	wg.Add(len(numbers)) //como vamos a lanzar una rutina por cada numero vamos poner que espere el mismo numero de rutinas que se lanzan

	/* 1.- NORMALMENTE SE ARIA ESTE PROCESO EN UN SOLO PROGRAMA LINIAL
	for _, v := range numbers {
		fmt.Printf("El numero %v tiene %v primos \n", v, cuantosPrimosTengo(v))
	}
	*/
	//para mejorarlo con paralelismo vamos a implementar rutinas

	for _, v := range numbers {

		//agregamos una funcion para implementar una rutina
		go func(a uint32) { //con la palabra "go" le indicamos que es una go rutina
			defer wg.Done() //con defer le decimos que cuando termine la funcion cuantosPrimosTengo se ejecute el Done de la tunia el cual inidca que ya termino la rutina
			fmt.Printf("El numero %v tiene %v primos \n", a, cuantosPrimosTengo(a))
		}(v) //le debemos pasar los parametros aquí
		//para que las go runitas se ejecuten y no termine el proceso si no que las espere vamos agregar un WaitGroup
	}

	wg.Wait() // le vamos a decir al programa que espere las go rutinas

	t = time.Now()
	fmt.Println("Finiliza el proceso:", t)
}

//obtener el numero de primos que se encuentran en ese numero
func cuantosPrimosTengo(n uint32) uint32 {
	var c uint32
	var i uint32
	for i = 1; i <= n; i++ {
		if i%2 != 0 {
			c = c + 1
		}
	}
	return c
}
