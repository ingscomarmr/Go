package main

import "../interfaces/animales"

func main() {
	var p animales.Perro
	var g animales.Gato

	p.Nombre = "Negra"
	p.Edad = 3
	p.Color = "Negro"

	g.Nombre = "Galleta"
	g.Color = "Pardo"
	g.Edad = 1

	ComunicaMascota(p)
	ComunicaMascota(g)

	var v animales.Vaca
	v.Nombre = "Lola"

	var gor animales.Gorilla
	gor.Nombre = "Monky"

	ComunicaMascota(gor)
	ComunicaMascota(v)
}

func ComunicaMascota(m animales.IMascota) {
	m.Comunicarse()
}
