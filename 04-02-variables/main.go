package main

import "fmt"

var (
	nombre string
	edad   int
	peso   float32
)

func main() {
	// var nombredelavariable tipo
	// var nombre string
	// nombre = "Alexys"
	// var nombre string = "Alexys"
	// var nombre = "Alexys"
	// var nombre, apellido string
	// nombre = "Alexys"
	// apellido = "Lozada"
	// var nombre, apellido string = "Alexys", "Lozada"
	// var nombre, apellido = "Alexys", "Lozada"
	// var (
	//	nombre string
	//	edad   int
	//	peso   float32
	// )
	nombre, edad, peso = "Alexys", 37, 80
	fmt.Println(nombre, edad, peso)
}
