package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql" //solo se ejecuta la funcion init que es la requerida para gorm
	"github.com/jinzhu/gorm"
)

func main() {
	//abrimos la conexion
	conDb, err := gorm.Open("mysql", //le indicamos la base de datos, puede ser oracle o etc
		"omarmr:Omar.2019@/escuela?charset=utf8&parseTime=True&loc=Local") //agreamos la cadena de conexion

	if err != nil { //si encontramos un error que mande un panic y que ya no siga
		panic("ERROR EN LA CONECION DE BASE DE DATOS")
	}

	defer conDb.Close() //con el defer cerramos la conexion
	fmt.Println("SE CONECTO A LA BASE DE DATOS")
}
