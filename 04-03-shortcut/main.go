package main

import "fmt"

func main() {
	// :=
	// nombredelavariable := valor
	// nombre := "Alexys"
	// nombre, apellido, edad := "Alexys", "Lozada", 37
	a, b := 5, 10
	a, b = b, a
	// tenemos más líneas de código
	c, a := 3, 4
	fmt.Println(a, b, c)
}
