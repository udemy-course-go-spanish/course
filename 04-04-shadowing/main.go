package main

import "fmt"

var nombre string

func main() {
	nombre := 35
	fmt.Println(nombre)
	// varias líneas después
	// var nombre int
	nombre = 3
	fmt.Println(nombre)
	// otras líneas después
	if true {
		var nombre bool
		nombre = true
		fmt.Println(nombre)
	}
	fmt.Println(nombre)
}
