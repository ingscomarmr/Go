package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql" //solo se ejecuta la funcion init que es la requerida para gorm
	"github.com/jinzhu/gorm"
)

//Nos ayuda contener el model de la tabla curso
type Alumno struct {
	gorm.Model      //en automatico nos crea las columnas id(ID), created_at(CreatedAt), updated_at(UpdatedAt),deleted_at(DeletedAt)
	Nombre          string
	ApellidoPaterno string
	ApellidoMaterno string
	Edad            uint8
	Correo          string
	Curp            string
}

func main() {
	//abrimos la conexion
	conDb, err := gorm.Open("mysql", //le indicamos la base de datos, puede ser oracle o etc
		"omarmr:Omar.2019@/escuela?charset=utf8&parseTime=True&loc=Local") //agreamos la cadena de conexion

	if err != nil { //si encontramos un error que mande un panic y que ya no siga
		panic("ERROR EN LA CONECION DE BASE DE DATOS")
	}

	defer conDb.Close() //con el defer cerramos la conexion
	fmt.Println("SE CONECTO A LA BASE DE DATOS")

	//conDb.CreateTable(&Alumno{}) //esto crea la tabla en base de datos
	//fmt.Println("SE CREO LA TABLA ALUMNO")

	//consultar un alumno
	var alumno Alumno
	conDb.First(&alumno, 1) //consultamos el registros que tenga id 1

	if alumno.ID != 0 { //si el alumno no existe
		fmt.Println("Alumno ID=1 :", alumno)
		fmt.Println("REALIZAMOS UN UPDATE ")
		alumno.Curp = "MURO880404HPLNVM05"
		conDb.Save(&alumno) //realizamos el update

	} else {
		//pare realizar un insert
		conDb.Create(&Alumno{
			Nombre: "Omar", ApellidoPaterno: "Munguia", ApellidoMaterno: "Rivera",
			Edad: 31, Correo: "omar.munguia@gmail.com", Curp: "MURO880404",
		})
		fmt.Println("Se inserto un alumno")
	}

}
