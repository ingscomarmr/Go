package main

import "fmt"

func main() {
	fmt.Println(" ########## CICLO FOR CONTINUO  ##########")
	nombres := []string{"omar", "evaristo", "rigo"}

	for index, valor := range nombres { // range nos ayuda a sacar el valor y el indice
		fmt.Printf("Index:%v, valor:%v\n", index, valor)
	}

	for _, valor := range nombres { //el indice se puede ignorar para no usarlo
		fmt.Printf("valor:%v\n", valor)
	}

	for index := range nombres { //el valor se puede incluso quitar pero no el indice
		fmt.Printf("Index:%v\n", index)
	}

	//para mapas es similar
	paises := map[string]string{"mx": "Mexico", "us": "USA", "br": "Brasil"}
	for pos, pais := range paises { //el valor se puede incluso quitar pero no el indice
		fmt.Printf("posicion:%v, valor:%v\n", pos, pais)
	}

	//para cadenas es similar
	cadena := "Omar Munguia"
	for i, c := range cadena { //el valor se puede incluso quitar pero no el indice
		fmt.Printf("posicion:%v, valor:%v\n", i, string(c))
	}

	//incluso se puede pasar directo la cadena
	for i, c := range "Esto es mi cadena" { //el valor se puede incluso quitar pero no el indice
		fmt.Printf("posicion:%v, valor:%v\n", i, string(c))
	}
}
